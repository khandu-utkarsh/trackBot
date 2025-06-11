from typing import Dict, Any
from langchain_core.messages import AIMessage, HumanMessage, SystemMessage
from agent.state import AgentState
import logging
from langchain.chat_models import init_chat_model
from config.settings import get_settings
from app.tools.exercise_tools import get_available_tools
from langchain_core.prompts import ChatPromptTemplate

logger = logging.getLogger(__name__)
settings = get_settings()

systemMessage = """You are an intelligent assistant that helps users log their workout details into a structured format suitable for database storage. Users will describe their workout sessions in natural language.

Your job is to extract structured data and prepare it for submission using a single tool call. Each workout contains one or more exercises.

Responsibilities:

1. For each exercise, extract:
   - **Type**: Determine if the exercise is "cardio" or "strength"
   - **Name**: Name of the exercise (e.g., "RDL", "Run", "Bench Press")
   - If **strength**:
     - `reps` (int): number of repetitions
     - `weight` (int): weight lifted in kilograms
   - If **cardio**:
     - `distance` (int): distance covered, in meters
     - `time` (int): duration of the activity, in seconds

2. If the user mentions **multiple sets with the same reps and weight**, create **multiple entries**, one per set. Do **not** use a `sets` field.

3. **Group all exercises into a single tool call**. Do not make separate tool calls for each exercise.

Workflow:

1. First, determine if you have enough information to generate a valid tool call.
2. If **any required field is missing or ambiguous**, ask the user for clarification.
3. Once all required data is available, make **one tool call** with the full list of exercises.
4. After the tool call, return a success confirmation message and conclude the conversation.

Important Guidelines:

- Do **not** call the tool until all required fields are available for every exercise.
- Always ask for clarification if information is missing or unclear.
- Use the tool **once** per workout to submit **all exercises together**.
- Output must conform strictly to the expected tool input schema.
- Use the *workout_id = 1* and *user_id = 1* for the workout and user.

"""

async def llm_node(state: AgentState) -> AgentState:
    """
    Process messages through LLM and determine next action.
    """
    logger.info("Processing LLM node")
    
    try:
        llm = init_chat_model(
            model=settings.MODEL_NAME,
            api_key=settings.OPENAI_API_KEY
        )

        tools = get_available_tools()
        availableToolsArray = list(tools.values())
        llm_with_tools = llm.bind_tools(availableToolsArray)

        # Build complete message sequence with system prompt
        messages = [SystemMessage(content=systemMessage)] + state["messages"]

        # Call the LLM
        responseFromLLM = await llm_with_tools.ainvoke(messages)
        logger.info("LLM response received")

        aiResponse: AIMessage = responseFromLLM
        
        # Determine the next action based on response content and tool calls
        next_action = _determine_next_action(aiResponse)
        logger.info(f"Next action determined: {next_action}")

        # Return updated AgentState
        return {
            **state,
            "messages": state["messages"] + [aiResponse],
            "next_action": next_action
        }
        
    except Exception as e:
        logger.error(f"Error in LLM node: {e}")
        error_message = AIMessage(content="I encountered an error processing your request. Please try again.")
        return {
            **state,
            "messages": state["messages"] + [error_message],
            "next_action": "end"
        }

def _determine_next_action(response: AIMessage) -> str:
    """
    Determine the next action based on LLM response.
    """
    # If there are tool calls, we should execute them
    if hasattr(response, 'tool_calls') and response.tool_calls:
        logger.info("Tool calls detected, proceeding to tool execution")
        return "tools"
    
    # Check if we need more information from the user
    if _needs_user_input(response.content):
        logger.info("User input required")
        return "user_input"
    
    # If we have a complete response with no tool calls or user input needed
    logger.info("No further action needed, ending conversation")
    return "end"

def _needs_user_input(content: str) -> bool:
    """
    Determine if the response requires user input.
    """
    user_input_indicators = [
        "need more information",
        "can you provide",
        "please specify",
        "which option",
        "what would you prefer",
        "please provide",
        "could you clarify",
        "I need to know",
        "please tell me",
        "could you tell me",
        "need to know",
        "could you tell me",
        "confirm"

    ]
    
    content_lower = content.lower()
    return any(indicator in content_lower for indicator in user_input_indicators) 