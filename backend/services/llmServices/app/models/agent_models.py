from pydantic import BaseModel, Field
from typing import List, Dict, Any, Optional
from models.message_models import Message

class ProcessMessagesRequest(BaseModel):
    """Request model for processing messages through the agent."""
    messages: List[Message] = Field(description="List of messages to process")
    user_id: int = Field(description="ID of the user making the request")
    session_id: Optional[str] = Field(default=None, description="Session ID for conversation continuity")

class ProcessMessagesResponse(BaseModel):
    """Response model for agent message processing."""
    message: Message = Field(description="The agent's response message")
    session_id: str = Field(description="Session ID for the conversation")
    tools_called: List[Dict[str, Any]] = Field(default_factory=list, description="List of tools that were called")
    needs_user_input: bool = Field(default=False, description="Whether user input is required")
    pending_input_prompt: Optional[str] = Field(default=None, description="Prompt for required user input")
    status: str = Field(description="Status of the processing (completed/awaiting_input)")

class ContinueSessionRequest(BaseModel):
    """Request model for continuing a session after user input."""
    session_id: str = Field(description="Session ID to continue")
    user_response: str = Field(description="User's response to continue with")

class SessionStatusResponse(BaseModel):
    """Response model for session status."""
    status: str = Field(description="Status of the session")
    user_id: Optional[int] = Field(default=None, description="User ID if session is active")
    needs_user_input: bool = Field(default=False, description="Whether user input is required")
    pending_input_prompt: Optional[str] = Field(default=None, description="Prompt for required user input")
    tools_called_count: int = Field(default=0, description="Number of tools called in this session") 