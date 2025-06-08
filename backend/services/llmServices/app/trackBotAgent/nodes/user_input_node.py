from typing import Dict, Any
from app.trackBotAgent.state import AgentState
import logging

logger = logging.getLogger(__name__)

async def user_input_node(state: AgentState) -> AgentState:
    """
    Handle user input interruption.
    This node sets up the state for requiring user input and pauses execution.
    """
    logger.info("Processing user input node")
    
    # Get the last message to understand what input is needed
    last_message = state["messages"][-1]
    
    # Set up pending user input requirement
    pending_prompt = last_message.content
    
    logger.info(f"User input required. Prompt: {pending_prompt}")
    
    # Return updated AgentState with all fields preserved
    return {
        **state,  # Preserve all existing state
        "pending_input_prompt": pending_prompt,
        "status": "waiting_for_input",
        "next_action": "pause"  # This will cause the graph to pause and wait for user input
    } 