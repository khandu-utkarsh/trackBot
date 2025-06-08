from typing import Dict, Any, List
from langchain.chat_models import init_chat_model
from langchain.prompts import ChatPromptTemplate
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage, BaseMessage

from app.config.settings import get_settings
import logging


settings = get_settings()
logger = logging.getLogger(__name__)
'''
class LLMService:
    def __init__(self):
        self.llm = init_chat_model(
            model_name=settings.MODEL_NAME,
            api_key=settings.OPENAI_API_KEY
        )
        
        # Fitness coaching system prompt
        self.system_prompt = """You are FitBot, an expert AI fitness coach and personal trainer. 
        You help users with:
        - Creating personalized workout plans
        - Providing exercise guidance and form tips
        - Nutrition advice and meal planning
        - Tracking progress and motivation
        - Answering fitness-related questions
        
        Always be encouraging, knowledgeable, and safety-focused. 
        If a user asks about medical conditions or injuries, recommend consulting a healthcare professional.
        
        When suggesting workouts, be specific about:
        - Exercise names
        - Sets and reps
        - Rest periods
        - Proper form cues
        
        You can also help users log their workouts by parsing their descriptions into structured data."""
        
        self.chat_prompt = ChatPromptTemplate.from_messages([
            ("system", self.system_prompt),
            ("human", "{user_message}")
        ])
        
        self.workout_parser_prompt = ChatPromptTemplate.from_messages([
            ("system", """You are a fitness assistant that parses workout descriptions into structured data.
            Parse the input into structured cardio and strength exercises.
            For sets, create separate entries for each set (don't store set count, store individual sets).
            
            Return a JSON object with:
            - exercises: array of exercise objects
            - workout_type: "strength", "cardio", or "mixed"
            - estimated_duration: number in minutes
            """),
            ("human", "{user_input}")
        ])

    async def process_chat_message(
        self, 
        user_message: str, 
        conversation_history: List[Dict[str, str]], 
        user_id: str, 
        conversation_id: str, 
        context: Dict[str, Any]
    ) -> Dict[str, Any]:
        """
        Process a chat message using LangChain/LangGraph for fitness coaching
        """
        try:
            # Convert conversation history to LangChain format
            messages = []
            for msg in conversation_history[-10:]:  # Keep last 10 messages for context
                if msg["role"] == "system":
                    messages.append(SystemMessage(content=msg["content"]))
                elif msg["role"] == "user":
                    messages.append(HumanMessage(content=msg["content"]))
                elif msg["role"] == "assistant":
                    messages.append(AIMessage(content=msg["content"]))
            
            # Add current user message
            messages.append(HumanMessage(content=user_message))
            
            # Check if this looks like a workout logging request
            if self._is_workout_logging_request(user_message):
                return await self._handle_workout_logging(user_message, user_id, context)
            
            # Regular chat response
            response = await self.llm.ainvoke(messages)
            
            return {
                "message": response.content,
                "metadata": {
                    "response_type": "chat",
                    "user_id": user_id,
                    "conversation_id": conversation_id
                }
            }
            
        except Exception as e:
            return {
                "message": "I'm sorry, I encountered an error processing your message. Please try again.",
                "metadata": {
                    "error": str(e),
                    "response_type": "error"
                }
            }

    def _is_workout_logging_request(self, message: str) -> bool:
        """
        Determine if the message is a workout logging request
        """
        workout_keywords = [
            "did", "completed", "finished", "workout", "exercise", "sets", "reps", 
            "ran", "lifted", "bench", "squat", "deadlift", "cardio", "miles", "km"
        ]
        message_lower = message.lower()
        return any(keyword in message_lower for keyword in workout_keywords)

    async def _handle_workout_logging(self, user_message: str, user_id: str, context: Dict[str, Any]) -> Dict[str, Any]:
        """
        Handle workout logging requests by parsing and structuring the data
        """
        try:
            # Parse the workout description
            chain = self.workout_parser_prompt | self.llm
            result = await chain.ainvoke({"user_input": user_message})
            
            # TODO: Here you would integrate with LangGraph for more complex workflow
            # For now, we'll return a structured response
            
            return {
                "message": f"Great job on your workout! I've logged the following exercises: {user_message}. Would you like me to suggest your next workout?",
                "metadata": {
                    "response_type": "workout_logged",
                    "parsed_workout": result.content,
                    "user_id": user_id,
                    "suggested_action": "log_workout"
                }
            }
            
        except Exception as e:
            return {
                "message": "I had trouble parsing your workout. Could you describe it in more detail?",
                "metadata": {
                    "error": str(e),
                    "response_type": "parsing_error"
                }
            }

    async def process_request(self, user_input: str) -> Dict[str, Any]:
        """
        Legacy method for backward compatibility
        """
        # Simple processing for now
        chain = self.chat_prompt | self.llm
        result = await chain.ainvoke({"user_message": user_input})
        
        return {
            "message": result.content,
            "metadata": {
                "response_type": "simple_chat"
            }
        } 

'''

class LLMService:
    def __init__(self):
        self.llm = init_chat_model(
            model=settings.MODEL_NAME,
            api_key=settings.OPENAI_API_KEY
        )
        self.tools_bound = False

    def bind_tools(self, tools: List[Any]) -> None:
        """
        Bind tools to the LLM for tool calling.
        
        Args:
            tools: List of tools to bind
        """
        if tools:
            self.llm = self.llm.bind_tools(tools)
            self.tools_bound = True
            logger.info(f"Bound {len(tools)} tools to LLM")



                 
    async def process_chat_message(
        self, 
        messages: List[BaseMessage],
        user_id: int
    ) -> AIMessage:
        """
        Process a chat message using LangChain/LangGraph for fitness coaching
        """
        logger.info(f"Processing chat message: {messages}")
        logger.info(f"User ID: {user_id}")

        try:
            # Regular chat response
            logger.info(f"Calling LLM with messages: {messages}")
            logger.info(f"User ID: {user_id}")
            response: AIMessage = await self.llm.ainvoke(messages)
            logger.info(f"LLM response: {response.content}")
            return response

        except Exception as e:
            logger.error(f"Error processing chat message: {e}")
            return AIMessage(content="I'm sorry, I encountered an error processing your message. Please try again.")