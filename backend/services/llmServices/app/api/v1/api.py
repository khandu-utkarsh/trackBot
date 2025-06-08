from fastapi import APIRouter

from api.v1.endpoints import endpoints

api_router = APIRouter()
api_router.include_router(endpoints.router, tags=["endpoints"])