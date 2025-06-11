from fastapi import Request, HTTPException, status
from fastapi.security import HTTPBearer
from app.utils.jwt import verify_token
import logging
from config.settings import get_settings
from typing import Optional
from app.models.models import User

settings = get_settings()
logger = logging.getLogger(__name__)
security = HTTPBearer()

async def verify_token_middleware(request: Request, call_next):
    """
    Middleware to verify JWT token in protected routes
    """
    logger.info("Validating JWT: Middleware Called")

    # Get JWT from cookie
    token = request.cookies.get(settings.COOKIE_NAME)
    if not token:
        logger.info(f"NO {settings.COOKIE_NAME} cookie found")
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Authentication required"
        )

    logger.info(f"FOUND {settings.COOKIE_NAME} cookie")

    try:
        # Verify token and get user context
        payload = verify_token(token)
        request.state.user = User(
            id=payload["user_id"],
            email=payload["email"],
            name=payload["name"],
            picture=payload["picture"]
        )
        
        logger.info(f"JWT validation successful for user: {payload['email']}")
        return await call_next(request)
    except Exception as e:
        logger.error(f"JWT validation failed: {str(e)}")
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid authentication credentials",
            headers={"WWW-Authenticate": "Bearer"},
        )

async def verify_token(request: Request) -> User:
    """
    Dependency to verify token for protected routes.
    Returns the authenticated user.
    """
    logger.info("Validating JWT: Dependency Called")

    # Get JWT from cookie
    token = request.cookies.get(settings.COOKIE_NAME)
    if not token:
        logger.info(f"NO {settings.COOKIE_NAME} cookie found")
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Authentication required"
        )

    logger.info(f"FOUND {settings.COOKIE_NAME} cookie")

    try:
        # Verify token and get user context
        payload = verify_token(token)
        user = User(
            id=payload["user_id"],
            email=payload["email"],
            name=payload["name"],
            picture=payload["picture"]
        )
        
        logger.info(f"JWT validation successful for user: {payload['email']}")
        return user
    except Exception as e:
        logger.error(f"JWT validation failed: {str(e)}")
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid authentication credentials",
            headers={"WWW-Authenticate": "Bearer"},
        )

def get_user_from_request(request: Request) -> Optional[User]:
    """
    Get user context from request
    """
    try:
        user = request.state.user
        if isinstance(user, dict):
            return User(**user)
        return user
    except (AttributeError, ValueError):
        return None 