from fastapi import APIRouter
from app.api.v1.conversation import router as conversation_router
from app.api.v1.auth import routerLogin, router
from app.api.v1.health import router as health_router

api_router = APIRouter()

api_router.include_router(health_router, tags=["health"])
api_router.include_router(routerLogin, prefix="/auth", tags=["auth"])
api_router.include_router(router, prefix="/auth", tags=["auth"])
api_router.include_router(conversation_router, prefix="", tags=["conversation"])