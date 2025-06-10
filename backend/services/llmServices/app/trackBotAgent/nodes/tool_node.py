from typing import Dict, Any, List
from langchain_core.messages import ToolMessage, AIMessage
from trackBotAgent.state import AgentState
from trackBotAgent.tools.custom_tools import get_available_tools
import logging
import json

logger = logging.getLogger(__name__)

async def tool_node(state: AgentState) -> AgentState:
    """
    Execute tools called by the LLM and return updated AgentState.
    """
    logger.info("Processing tool node")
    
    try:
        # Get the last message (should be from LLM with tool calls)
        last_message = state["messages"][-1]
        
        if not hasattr(last_message, 'tool_calls') or not last_message.tool_calls:
            logger.warning("Tool node called but no tool calls found")
            return {
                **state,
                "next_action": "end"
            }
        
        # Get available tools and execute tool calls
        tool_results = await _execute_tool_calls(last_message.tool_calls)
        
        # Update tools_called in state
        updated_tools_called = _update_tools_called(state.get("tools_called", []), last_message.tool_calls)
        
        # Return updated AgentState
        return {
            **state,
            "messages": state["messages"] + tool_results,
            "tools_called": updated_tools_called,
            "next_action": "llm_call"  # Go back to LLM to process tool results
        }
        
    except Exception as e:
        logger.error(f"Error in tool node: {e}")
        error_message = ToolMessage(
            content=f"Error executing tools: {str(e)}",
            tool_call_id="error"
        )
        return {
            **state,
            "messages": state["messages"] + [error_message],
            "next_action": "end"
        }

async def _execute_tool_calls(tool_calls: List[Dict[str, Any]]) -> List[ToolMessage]:
    """
    Execute a list of tool calls and return their results.
    """
    available_tools = get_available_tools()
    tool_results = []
    
    for tool_call in tool_calls:
        tool_name = tool_call["name"]
        tool_args = tool_call["args"]
        tool_call_id = tool_call["id"]
        
        logger.info(f"Executing tool: {tool_name} with args: {tool_args}")
        
        try:
            if tool_name not in available_tools:
                raise ValueError(f"Tool {tool_name} not available")
            
            tool = available_tools[tool_name]
            result = await tool.ainvoke(tool_args)
            
            tool_message = ToolMessage(
                content=str(result),
                tool_call_id=tool_call_id
            )
            tool_results.append(tool_message)
            logger.info(f"Tool {tool_name} executed successfully")
            
        except Exception as e:
            logger.error(f"Error executing tool {tool_name}: {e}")
            tool_message = ToolMessage(
                content=f"Error executing tool {tool_name}: {str(e)}",
                tool_call_id=tool_call_id
            )
            tool_results.append(tool_message)
    
    return tool_results

def _update_tools_called(existing_tools: List[Dict[str, Any]], new_tool_calls: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
    """
    Update the list of called tools with new tool calls.
    """
    updated_tools = list(existing_tools)
    
    for tool_call in new_tool_calls:
        updated_tools.append({
            "tool_name": tool_call["name"],
            "args": tool_call["args"],
            "call_id": tool_call["id"]
        })
    
    return updated_tools 