from fastapi import APIRouter, Depends, HTTPException, status, Request
from sqlalchemy.orm import Session
from app.config.database import get_db
from app.utils.jwt import create_access_token, validate_google_jwt
from pydantic import BaseModel
from typing import Optional, List
from app.models.models import GoogleLoginRequest, User
from app.models.database import User as DBUser
from fastapi.responses import JSONResponse
from datetime import timedelta
from middleware.auth import verify_token_middleware

routerLogin = APIRouter()

class GoogleAuthRequest(BaseModel):
    token: str

class UserResponse(BaseModel):
    id: int
    email: str

    class Config:
        from_attributes = True

@routerLogin.post("/google", response_model=User)
async def google_login(request: GoogleLoginRequest, db: Session = Depends(get_db)):
    """
    Handle Google OAuth login
    """
    if not request.googleToken:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Google token is required"
        )

    try:
        # 1. Validate Google JWT token
        google_claims = await validate_google_jwt(request.googleToken)
        
        # 2. Find or create user in database
        user = db.query(DBUser).filter(DBUser.email == google_claims.email).first()
        
        if not user:
            try:
                # Create new user
                user = DBUser(
                    email=google_claims.email,
                    name=google_claims.name,
                    picture=google_claims.picture
                )
                db.add(user)
                db.commit()
                db.refresh(user)
            except Exception as e:
                raise HTTPException(
                    status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
                    detail="Failed to create user"
                )
        else:
            # Update existing user's info
            user.name = google_claims.name
            user.picture = google_claims.picture
            db.commit()
            db.refresh(user)

        # 3. Create JWT token with user information
        try:
            token = create_access_token({
                "user_id": user.id,
                "email": user.email,
                "name": user.name,
                "picture": user.picture,
                "google_sub": google_claims.sub
            })
        except Exception as e:
            raise HTTPException(
                status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
                detail="Failed to create token"
            )

        # 4. Create response with user information
        responseUser = User(
            id=user.id,
            email=user.email,
            name=user.name,
            picture=user.picture,
            created_at=user.created_at,
            updated_at=user.updated_at
        )

        # 5. Set secure HttpOnly cookie
        response = JSONResponse(content=responseUser.model_dump())
        response.set_cookie(
            key="trackbot_auth_token",
            value=token,
            httponly=True,
            secure=True,
            samesite="strict",
            max_age=ACCESS_TOKEN_EXPIRE_MINUTES * 60,  # Convert minutes to seconds
            path="/"
        )
        
        return response

    except HTTPException as e:
        raise e
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=f"Invalid Google token: {str(e)}",
            headers={"WWW-Authenticate": "Bearer"},
        )

router = APIRouter()

@router.middleware("http")
async def protected_middleware(request: Request, call_next):
    return await verify_token_middleware(request, call_next)

@router.post("/logout", response_model=User)
async def logout(request: Request, db: Session = Depends(get_db)):
    """
    Handle user logout
    """
    user = request.state.user
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="User not found"
        )
    
    response = User(
        id=user.id,
        email=user.email,
        name=user.name,
        picture=user.picture,
        created_at=user.created_at,
        updated_at=user.updated_at
    )
    return response

@router.get("/me", response_model=User)
async def get_current_user(request: Request, db: Session = Depends(get_db)):
    """
    Get current user information
    """
    user = request.state.user
    if not user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="User not found"
        )
    
    response = User(
        id=user.id,
        email=user.email,
        name=user.name,
        picture=user.picture,
        created_at=user.created_at,
        updated_at=user.updated_at
    )
    return response 