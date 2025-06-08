from typing import Dict, Any, List
from langchain_core.messages import BaseMessage, HumanMessage
from app.trackBotAgent.trackBot_agent import BaseAgent
from app.trackBotAgent.state import AgentState
import logging
import uuid

logger = logging.getLogger(__name__)

class AgentService:
    """
    Service class for managing agent interactions.
    """
    
    def __init__(self):
        self.active_sessions: Dict[str, Dict[str, Any]] = {}
    
    async def process_messages(
        self, 
        messages: List[BaseMessage], 
        user_id: int, 
        session_id: str = None
    ) -> Dict[str, Any]:
        """
        Process messages through the agent workflow.
        
        Args:
            messages: List of messages to process
            user_id: ID of the user
            session_id: Session ID for conversation continuity
            
        Returns:
            Processing result including response and metadata
        """
        if not session_id:
            session_id = str(uuid.uuid4())
        
        logger.info(f"Processing messages for user {user_id}, session {session_id}")
        
        try:
            # Create or get agent
            agent = BaseAgent()
            
            # Prepare initial state
            initial_state: AgentState = {
                "messages": messages,
                "user_id": user_id,
                "session_id": session_id,
                "tools_called": [],
                "pending_user_input": None,
                "context": {},
                "next_action": None
            }
            
            # Run the agent
            result = await agent.run(initial_state)
            
            # Store session state for potential continuation
            self.active_sessions[session_id] = {
                "state": result,
                "agent": agent,
                "user_id": user_id
            }
            
            # Prepare response
            response = {
                "messages": result["messages"],
                "session_id": session_id,
                "tools_called": result.get("tools_called", []),
                "needs_user_input": result.get("pending_user_input") is not None,
                "pending_input_prompt": result.get("pending_user_input", {}).get("prompt"),
                "status": "completed" if not result.get("pending_user_input") else "awaiting_input"
            }
            
            logger.info(f"Processing completed for session {session_id}")
            return response
            
        except Exception as e:
            logger.error(f"Error processing messages: {e}")
            raise
    
    async def continue_session(
        self, 
        session_id: str, 
        user_response: str
    ) -> Dict[str, Any]:
        """
        Continue a session after user input.
        
        Args:
            session_id: Session ID to continue
            user_response: User's response to continue with
            
        Returns:
            Continuation result
        """
        logger.info(f"Continuing session {session_id}")
        
        if session_id not in self.active_sessions:
            raise ValueError(f"Session {session_id} not found or expired")
        
        try:
            session_data = self.active_sessions[session_id]
            agent = session_data["agent"]
            current_state = session_data["state"]
            
            # Continue from interruption
            result = await agent.continue_from_interruption(current_state, user_response)
            
            # Update session state
            self.active_sessions[session_id]["state"] = result
            
            # Prepare response
            response = {
                "messages": result["messages"],
                "session_id": session_id,
                "tools_called": result.get("tools_called", []),
                "needs_user_input": result.get("pending_user_input") is not None,
                "pending_input_prompt": result.get("pending_user_input", {}).get("prompt"),
                "status": "completed" if not result.get("pending_user_input") else "awaiting_input"
            }
            
            logger.info(f"Session {session_id} continuation completed")
            return response
            
        except Exception as e:
            logger.error(f"Error continuing session {session_id}: {e}")
            raise
    
    def get_session_status(self, session_id: str) -> Dict[str, Any]:
        """
        Get the status of a session.
        
        Args:
            session_id: Session ID to check
            
        Returns:
            Session status information
        """
        if session_id not in self.active_sessions:
            return {"status": "not_found"}
        
        session_data = self.active_sessions[session_id]
        state = session_data["state"]
        
        return {
            "status": "active",
            "user_id": session_data["user_id"],
            "needs_user_input": state.get("pending_user_input") is not None,
            "pending_input_prompt": state.get("pending_user_input", {}).get("prompt"),
            "tools_called_count": len(state.get("tools_called", []))
        }
    
    def cleanup_session(self, session_id: str) -> bool:
        """
        Clean up a session.
        
        Args:
            session_id: Session ID to clean up
            
        Returns:
            True if session was cleaned up, False if not found
        """
        if session_id in self.active_sessions:
            del self.active_sessions[session_id]
            logger.info(f"Session {session_id} cleaned up")
            return True
        return False 