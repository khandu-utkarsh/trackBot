from typing import Dict, Any
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
        
        # Get available tools
        available_tools = get_available_tools()
        tool_results = []
        updated_tools_called = list(state["tools_called"])
        
        # Execute each tool call
        for tool_call in last_message.tool_calls:
            tool_name = tool_call["name"]
            tool_args = tool_call["args"]
            tool_call_id = tool_call["id"]
            
            logger.info(f"Executing tool: {tool_name} with args: {tool_args}")
            
            # Record the tool call
            updated_tools_called.append({
                "tool_name": tool_name,
                "args": tool_args,
                "call_id": tool_call_id
            })
            
            # Execute the tool if available
            if tool_name in available_tools:
                try:
                    tool = available_tools[tool_name]
                    #So, this is responsible for calling the tool.
                    result = await tool.ainvoke(tool_args)
                    
                    # Create tool message
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
            else:
                logger.error(f"Tool {tool_name} not found")
                tool_message = ToolMessage(
                    content=f"Tool {tool_name} not available",
                    tool_call_id=tool_call_id
                )
                tool_results.append(tool_message)
        
        # Return updated AgentState with all fields preserved
        return {
            **state,  # Preserve all existing state
            "messages": state["messages"] + tool_results,
            "tools_called": updated_tools_called,
            "next_action": "llm_call"  # Go back to LLM to process tool results
        }
        
    except Exception as e:
        logger.error(f"Error in tool node: {e}")
        return {
            **state,  # Preserve existing state
            "next_action": "end"
        } 