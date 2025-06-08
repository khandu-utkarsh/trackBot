from langchain_core.messages import BaseMessage
from app.trackBotAgent.trackBot_agent import TrackBotAgent
from app.trackBotAgent.state import AgentState
from typing import List
import logging

logger = logging.getLogger(__name__)

class AgentService:
    """
    Service class for managing agent interactions.
    """
    
    def __init__(self, messages: List[BaseMessage], user_id: int):
        self.messages = messages
        self.user_id = user_id
        self.agent = TrackBotAgent()
        self.state = AgentState(
            messages=messages,
            user_id=user_id,
            tools_called=None,
            pending_input_prompt=None,
            status="started",
            next_action="process_messages"
        )
    
    async def process_messages(
        self, 
    ) -> BaseMessage:
        """
        Process messages through the agent workflow.
        
        Args:
            
        Returns:
            Processing result including response and metadata
        """        
        logger.info(f"Processing messages for user {self.user_id}")
        
        try:
            # Create or get agent
            agent = TrackBotAgent()
            
            # Prepare initial state
            initial_state: AgentState = self.state
            
            # Run the agent
            self.state : AgentState = await agent.run(initial_state)
            

            #We will be only returning the last message from the agent
            last_message = self.state.messages[-1]            
            logger.info(f"Processing completed for user {self.user_id}")
            return last_message
            
        except Exception as e:
            logger.error(f"Error processing messages: {e}")
            raise

    async def continue_from_interruption(
        self, 
        user_response: BaseMessage, 
    ) -> BaseMessage:
        """
        Continue a session after user input.
        
        Args:
            session_id: Session ID to continue
            user_response: User's response to continue with
            
        Returns:
            Continuation result
        """
        logger.info(f"Continuing from interruption for user {self.user_id}")
        
        try:
            # Continue from interruption
            result : AgentState = await self.agent.continue_from_interruption(self.state, user_response)
            # Update session state
            self.state = result
            
            #We will be only returning the last message from the agent
            last_message = self.state.messages[-1]            
            logger.info(f"Processing completed for user {self.user_id}")
            return last_message
            
        except Exception as e:
            logger.error(f"Error processing messages: {e}")
            raise