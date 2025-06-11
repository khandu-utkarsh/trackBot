from langchain_core.messages import BaseMessage
from agent.trackBot_agent import TrackBotAgent
from agent.state import AgentState
from typing import List
import logging

logger = logging.getLogger(__name__)

class AgentService:
    """
    Service class that provides a simplified interface to the TrackBotAgent.
    This class handles the high-level conversation flow and state management.
    """
    
    def __init__(self, messages: List[BaseMessage], user_id: int):
        self.agent = TrackBotAgent()
        self.state = AgentState(
            messages=messages,
            user_id=user_id,
            tools_called=None,
            pending_input_prompt=None,
            status="started",
            next_action="process_messages"
        )
    
    async def process_messages(self) -> BaseMessage:
        """
        Process the current conversation state through the agent.
        Returns the last message from the agent's response.
        """
        logger.info(f"Processing messages for user {self.state['user_id']}")
        try:
            result_state = await self.agent.run(self.state)
            self.state = result_state
            return result_state["messages"][-1]
        except Exception as e:
            logger.error(f"Error processing messages: {e}")
            raise

    async def continue_from_interruption(self, user_response: BaseMessage) -> BaseMessage:
        """
        Continue the conversation after receiving user input.
        Returns the last message from the agent's response.
        """
        logger.info(f"Continuing conversation for user {self.state['user_id']}")
        try:
            result_state = await self.agent.continue_from_interruption(self.state, user_response)
            self.state = result_state
            return result_state.messages[-1]
        except Exception as e:
            logger.error(f"Error continuing conversation: {e}")
            raise