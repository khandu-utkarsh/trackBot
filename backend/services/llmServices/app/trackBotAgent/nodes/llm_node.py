from typing import Dict, Any
from langchain_core.messages import AIMessage, HumanMessage, SystemMessage
from trackBotAgent.state import AgentState
import logging
from langchain.chat_models import init_chat_model
from config.settings import get_settings
from trackBotAgent.tools.custom_tools import get_available_tools
from langchain_core.prompts import ChatPromptTemplate

logger = logging.getLogger(__name__)
settings = get_settings()


systemMessage = """You are an intelligent assistant that helps log workout details into a structured database. The user will provide natural language descriptions of their workout sessions.

Your task is to extract structured data about the workout. Each workout consists of one or more exercises, which must be parsed with all relevant parameters.

Your responsibilities:

1. For each exercise, extract:
   - Type: either "Strength" or "Cardio" (must be included and case-sensitive)
   - Name: the name of the exercise (e.g., "RDL", "Run", "Bench Press")
   - For Strength: include "reps" (int) and "weights" (int)
   - For Cardio: include "distance" (int, in meters) and "time" (int, in seconds)

2. Do not omit the "type" field. It is required for every exercise.

3. If a user mentions performing multiple sets with the same reps and weights, generate separate entries for each set. Do not use a "sets" field — the database does not support it.

4. Group all exercises into a single tool call. Do not split into multiple tool invocations — the tool accepts a list of exercises in one input.

Important:
- Only call the tool once per response.
- Once the tool has been called and the data is structured correctly, return a confirmation message and do not call the tool again.
- If any required parameter is missing or unclear, ask the user for clarification instead of guessing.
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
        llm_with_tools = llm.bind_tools(tools)

        # Build complete message sequence with system prompt
        messages = [SystemMessage(content=systemMessage)] + state["messages"]

        # Call the LLM
        responseFromLLM = await llm_with_tools.ainvoke(messages)
        print("Resonse from LLM with tools: ")
        print(responseFromLLM)

        aiResponse :  AIMessage = responseFromLLM
        print("AI Response: ")
        print(aiResponse)
        

        # Determine the next action
        # Determine next action based on response
        next_action = "end"  # Default to ending
        if hasattr(aiResponse, 'tool_calls') and aiResponse.tool_calls:
            next_action = "tools"
            logger.info(f"Tool calls detected: {aiResponse.tool_calls}")
        elif _needs_user_input(aiResponse.content):
            next_action = "user_input"
            logger.info("User input required")
        else:
            next_action = "end"
        # Return updated AgentState with all fields preserved
        return {
            **state,  # Preserve all existing state
            "messages": state["messages"] + [aiResponse],
            "next_action": next_action
        }
        
    except Exception as e:
        logger.error(f"Error in LLM node: {e}")
        error_message = AIMessage(content="I encountered an error processing your request. Please try again.")
        return {
            **state,  # Preserve existing state
            "messages": state["messages"] + [error_message],
            "next_action": "end"
        }

def _needs_user_input(content: str) -> bool:
    """
    Determine if the response requires user input.
    You can customize this logic based on your needs.
    """
    user_input_indicators = [
        "need more information",
        "can you provide",
        "please specify",
        "which option",
        "what would you prefer",
        "please provide"
    ]
    
    content_lower = content.lower()
    return any(indicator in content_lower for indicator in user_input_indicators) 