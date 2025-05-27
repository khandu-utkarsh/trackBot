from fastapi import APIRouter

from app.api.v1.endpoints import process_text, chat

api_router = APIRouter()
api_router.include_router(process_text.router, prefix="/process", tags=["process"])
api_router.include_router(chat.router, prefix="/chat", tags=["chat"]) 