from datetime import datetime, timedelta
from typing import Optional, Dict, Any
from jose import JWTError, jwt
from fastapi import HTTPException, status
from pydantic import BaseModel
import httpx
from cryptography.hazmat.primitives.asymmetric import rsa
import base64
from config.settings import get_settings
from datetime import UTC

settings = get_settings()

# JWT Configuration
JWT_SECRET_KEY = settings.TRACKBOT_JWT_SECRET_KEY
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 5  # 5 minutes like in Go service

class GoogleJWTClaims(BaseModel):
    sub: str
    email: str
    email_verified: bool
    name: str
    picture: str
    given_name: str
    family_name: str
    iss: str
    aud: str
    iat: int
    exp: int

class TrackBotJWTClaims(BaseModel):
    user_id: int
    email: str
    name: str
    picture: str
    google_sub: str
    exp: datetime
    iat: datetime
    iss: str = "trackbot-app"

class GooglePublicKey(BaseModel):
    kty: str
    alg: str
    use: str
    kid: str
    n: str
    e: str

class GoogleJWKS(BaseModel):
    keys: list[GooglePublicKey]

class GoogleKeysCache:
    def __init__(self, ttl_hours: int = 24):
        self.keys: Optional[GoogleJWKS] = None
        self.fetched_at: Optional[datetime] = None
        self.ttl = timedelta(hours=ttl_hours)

google_keys_cache = GoogleKeysCache()

def create_access_token(data: Dict[str, Any], expires_delta: Optional[timedelta] = None) -> str:
    """
    Create a new JWT access token
    """
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.now(UTC) + expires_delta
    else:
        expire = datetime.now(UTC) + timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES)
    
    to_encode.update({
        "exp": expire,
        "iat": datetime.now(UTC),
        "iss": "trackbot-app"
    })
    
    encoded_jwt = jwt.encode(to_encode, JWT_SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt

def verify_token(token: str) -> Dict[str, Any]:
    """
    Verify and decode a JWT token
    """
    try:
        payload = jwt.decode(token, JWT_SECRET_KEY, algorithms=[ALGORITHM])
        return payload
    except JWTError:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Could not validate credentials",
            headers={"WWW-Authenticate": "Bearer"},
        )

async def fetch_google_public_keys() -> GoogleJWKS:
    """
    Fetch Google's public keys for JWT verification with caching
    """
    if (google_keys_cache.keys is not None and 
        google_keys_cache.fetched_at is not None and 
        datetime.now(UTC) - google_keys_cache.fetched_at < google_keys_cache.ttl):
        return google_keys_cache.keys

    async with httpx.AsyncClient() as client:
        response = await client.get("https://www.googleapis.com/oauth2/v3/certs")
        if response.status_code != 200:
            raise HTTPException(
                status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
                detail="Failed to fetch Google public keys"
            )
        
        jwks = GoogleJWKS(**response.json())
        google_keys_cache.keys = jwks
        google_keys_cache.fetched_at = datetime.now(UTC)
        return jwks

def get_public_key_by_kid(jwks: GoogleJWKS, kid: str) -> Optional[GooglePublicKey]:
    """
    Get a public key by its key ID
    """
    for key in jwks.keys:
        if key.kid == kid:
            return key
    return None

def rsa_public_key_from_jwk(jwk: GooglePublicKey) -> rsa.RSAPublicKey:
    """
    Convert a JWK to an RSA public key
    """
    # Convert base64url to base64
    n = base64.urlsafe_b64decode(jwk.n + '=' * (-len(jwk.n) % 4))
    e = base64.urlsafe_b64decode(jwk.e + '=' * (-len(jwk.e) % 4))
    
    # Convert bytes to integers
    n_int = int.from_bytes(n, 'big')
    e_int = int.from_bytes(e, 'big')
    
    # Create RSA public key
    public_key = rsa.RSAPublicNumbers(e_int, n_int).public_key()
    return public_key

async def validate_google_jwt(token: str) -> GoogleJWTClaims:
    """
    Validate a Google JWT token
    """
    try:
        # Get the key ID from the token header
        unverified_claims = jwt.get_unverified_header(token)
        kid = unverified_claims.get('kid')
        if not kid:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid token: missing key ID"
            )

        # Get the public key
        jwks = await fetch_google_public_keys()
        jwk = get_public_key_by_kid(jwks, kid)
        if not jwk:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid token: key not found"
            )

        # Convert JWK to RSA public key
        public_key = rsa_public_key_from_jwk(jwk)
        
        # Verify the token
        claims = jwt.decode(
            token,
            public_key,
            algorithms=[jwk.alg],
            audience=settings.GOOGLE_CLIENT_ID,
            issuer="https://accounts.google.com"
        )
        
        return GoogleJWTClaims(**claims)
    except JWTError as e:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=f"Invalid token: {str(e)}"
        )