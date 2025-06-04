from fastapi import APIRouter

from app.api.v1.endpoints import process_messages

api_router = APIRouter()
api_router.include_router(process_messages.router, tags=["process_messages"])