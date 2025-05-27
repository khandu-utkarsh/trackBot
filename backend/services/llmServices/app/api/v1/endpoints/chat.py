from fastapi import APIRouter, HTTPException
from typing import List, Optional
from pydantic import BaseModel
from app.services import LLMService

router = APIRouter()
llm_service = LLMService()

class ChatMessage(BaseModel):
    role: str  # "user", "assistant", "system"
    content: str

class ChatRequest(BaseModel):
    messages: List[ChatMessage]
    user_id: str
    conversation_id: str
    context: Optional[dict] = None  # For workout context, user preferences, etc.

class ChatResponse(BaseModel):
    message: str
    message_type: str = "assistant"
    metadata: Optional[dict] = None  # For any additional data like suggested workouts, etc.

@router.post("/message", response_model=ChatResponse)
async def process_chat_message(request: ChatRequest):
    """
    Process a chat message and return an AI response.
    This endpoint will be called by the workout app service when a user sends a message.
    """
    try:
        # Extract the latest user message
        user_message = None
        for msg in reversed(request.messages):
            if msg.role == "user":
                user_message = msg.content
                break
        
        if not user_message:
            raise HTTPException(status_code=400, detail="No user message found")
        
        # Process the message with context
        response = await llm_service.process_chat_message(
            user_message=user_message,
            conversation_history=request.messages,
            user_id=request.user_id,
            conversation_id=request.conversation_id,
            context=request.context or {}
        )
        
        return ChatResponse(
            message=response.get("message", "I'm sorry, I couldn't process your request."),
            message_type="assistant",
            metadata=response.get("metadata", {})
        )
        
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Failed to process chat message: {str(e)}")

@router.get("/health")
async def health_check():
    """Health check endpoint for the chat service"""
    return {"status": "healthy", "service": "llm-chat"} 