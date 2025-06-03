package middleware

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserContext represents the authenticated user
type UserContext struct {
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// AuthMiddleware handles authentication
type AuthMiddleware struct{}

// NewAuthMiddleware creates a new auth middleware instance
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// GoogleJWTClaims represents the claims in a Google JWT token
type GoogleJWTClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Iss           string `json:"iss"`
	Aud           string `json:"aud"`
	Iat           int64  `json:"iat"`
	Exp           int64  `json:"exp"`
}

// Your JWT Claims (not Google's)
type TrackBotJWTClaims struct {
	UserID    int    `json:"user_id"` // From your database
	Email     string `json:"email"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	GoogleSub string `json:"google_sub"` // Google's user ID for reference
	jwt.RegisteredClaims
}

// JWT secret key (store in environment variable)
var jwtSecret = []byte(os.Getenv("TRACKBOT_JWT_SECRET_KEY"))

// CreateJWT creates your own JWT token
func (a *AuthMiddleware) CreateJWT(userID int, email, name, picture, googleSub string) (string, error) {
	claims := TrackBotJWTClaims{
		UserID:    userID,
		Email:     email,
		Name:      name,
		Picture:   picture,
		GoogleSub: googleSub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)), // 5 minutes
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "trackbot-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

type contextKey string

const userKey contextKey = "user"

// ValidateJWT middleware - validates YOUR JWT tokens from cookies
func (a *AuthMiddleware) ValidateJWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get JWT from cookie (not Authorization header)
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			// Parse and validate YOUR JWT
			claims, err := a.validateTrackBotJWT(cookie.Value)
			if err != nil {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			// Create user context from YOUR claims
			userCtx := &UserContext{
				UserID:  claims.UserID,
				Email:   claims.Email,
				Name:    claims.Name,
				Picture: claims.Picture,
			}

			ctx := context.WithValue(r.Context(), userKey, userCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// validateTrackBotJWT validates YOUR JWT token
func (a *AuthMiddleware) validateTrackBotJWT(tokenString string) (*TrackBotJWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TrackBotJWTClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TrackBotJWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GooglePublicKey represents a Google public key
type GooglePublicKey struct {
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// GoogleJWKS represents Google's JSON Web Key Set
type GoogleJWKS struct {
	Keys []GooglePublicKey `json:"keys"`
}

// Cache for Google public keys
type GoogleKeysCache struct {
	keys      *GoogleJWKS
	fetchedAt time.Time
	ttl       time.Duration
}

var googleKeysCache = &GoogleKeysCache{
	ttl: 24 * time.Hour, // Cache keys for 24 hours
}

// fetchGooglePublicKeys fetches Google's public keys for JWT verification with caching
func (a *AuthMiddleware) fetchGooglePublicKeys() (*GoogleJWKS, error) {
	// Check if we have cached keys that are still valid
	if googleKeysCache.keys != nil && time.Since(googleKeysCache.fetchedAt) < googleKeysCache.ttl {
		return googleKeysCache.keys, nil
	}

	// Fetch fresh keys
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Google public keys: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch Google public keys: status %d", resp.StatusCode)
	}

	var jwks GoogleJWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode Google public keys: %v", err)
	}

	// Update cache
	googleKeysCache.keys = &jwks
	googleKeysCache.fetchedAt = time.Now()

	return &jwks, nil
}

// getPublicKeyByKid finds a public key by key ID
func (a *AuthMiddleware) getPublicKeyByKid(jwks *GoogleJWKS, kid string) (*rsa.PublicKey, error) {
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			return a.rsaPublicKeyFromJWK(&key)
		}
	}
	return nil, fmt.Errorf("public key not found for kid: %s", kid)
}

// rsaPublicKeyFromJWK converts a JWK to an RSA public key
func (a *AuthMiddleware) rsaPublicKeyFromJWK(jwk *GooglePublicKey) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %v", err)
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %v", err)
	}

	// Convert exponent bytes to int
	var e int
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}

	// Create RSA public key
	publicKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: e,
	}

	return publicKey, nil
}

// validateGoogleJWT validates a Google JWT token with proper signature verification and audience check
func (a *AuthMiddleware) ValidateGoogleJWT(tokenString string) (*GoogleJWTClaims, error) {
	// Parse the token to get the header
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the key ID from the header
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("missing key ID in token header")
		}

		// Fetch Google's public keys
		jwks, err := a.fetchGooglePublicKeys()
		if err != nil {
			return nil, err
		}

		// Get the public key for this kid
		publicKey, err := a.getPublicKeyByKid(jwks, kid)
		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse or verify token: %v", err)
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or claims")
	}

	// Convert map claims to our struct
	googleClaims := &GoogleJWTClaims{}

	if sub, ok := claims["sub"].(string); ok {
		googleClaims.Sub = sub
	}
	if email, ok := claims["email"].(string); ok {
		googleClaims.Email = email
	}
	if emailVerified, ok := claims["email_verified"].(bool); ok {
		googleClaims.EmailVerified = emailVerified
	}
	if name, ok := claims["name"].(string); ok {
		googleClaims.Name = name
	}
	if picture, ok := claims["picture"].(string); ok {
		googleClaims.Picture = picture
	}
	if givenName, ok := claims["given_name"].(string); ok {
		googleClaims.GivenName = givenName
	}
	if familyName, ok := claims["family_name"].(string); ok {
		googleClaims.FamilyName = familyName
	}
	if iss, ok := claims["iss"].(string); ok {
		googleClaims.Iss = iss
	}
	if aud, ok := claims["aud"].(string); ok {
		googleClaims.Aud = aud
	}
	if iat, ok := claims["iat"].(float64); ok {
		googleClaims.Iat = int64(iat)
	}
	if exp, ok := claims["exp"].(float64); ok {
		googleClaims.Exp = int64(exp)
	}

	// Validate issuer
	if googleClaims.Iss != "accounts.google.com" && googleClaims.Iss != "https://accounts.google.com" {
		return nil, errors.New("invalid issuer")
	}

	// Validate audience - check against your Google OAuth client ID
	expectedAudience := os.Getenv("GOOGLE_CLIENT_ID")
	if expectedAudience == "" {
		return nil, errors.New("GOOGLE_CLIENT_ID environment variable not set")
	}
	if googleClaims.Aud != expectedAudience {
		return nil, fmt.Errorf("invalid audience: expected %s, got %s", expectedAudience, googleClaims.Aud)
	}

	// Check if token is expired
	now := time.Now().Unix()
	if googleClaims.Exp > 0 && googleClaims.Exp < now {
		return nil, errors.New("token expired")
	}

	// Check if token is not yet valid
	if googleClaims.Iat > 0 && googleClaims.Iat > now+300 { // Allow 5 minutes clock skew
		return nil, errors.New("token not yet valid")
	}

	// Ensure email is verified
	if !googleClaims.EmailVerified {
		return nil, errors.New("email not verified")
	}

	return googleClaims, nil
}

// GetUserFromContext extracts user context from request context
func GetUserFromContext(ctx context.Context) (*UserContext, bool) {
	user, ok := ctx.Value(userKey).(*UserContext)
	return user, ok
}
