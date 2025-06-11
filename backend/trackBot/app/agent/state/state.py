from typing import Dict, Any, List, Optional
from typing_extensions import TypedDict
from langchain_core.messages import BaseMessage

class AgentState(TypedDict):
    """State for the LangGraph agent workflow."""
    messages: List[BaseMessage]
    user_id: int
    conversation_id: int
    tools_called: List[Dict[str, Any]]
    pending_input_prompt: Optional[str]
    status: str
    next_action: Optional[str]
