from typing import Dict, Any
from langchain_core.messages import BaseMessage, HumanMessage
from langgraph.graph import StateGraph, END, START
from langgraph.checkpoint import MemorySaver
from app.trackBotAgent.state import AgentState
from app.trackBotAgent.nodes import llm_node, tool_node, user_input_node
import logging

logger = logging.getLogger(__name__)

class TrackBotAgent:
    """
    Agent class that orchestrates the LangGraph workflow.
    """
    
    def __init__(self):
        self.graph = self._build_graph()
    
    def _build_graph(self) -> StateGraph:
        """
        Build the LangGraph workflow.
        """
        agent_builder = StateGraph(AgentState)
        
        # Add nodes
        agent_builder.add_node("llm_call", llm_node)
        agent_builder.add_node("tools", tool_node)
        agent_builder.add_node("user_input", user_input_node)

        # Add edges
        agent_builder.add_edge(START, "llm_call")
        agent_builder.add_edge("tools", "llm_call")

        # Add conditional logic from llm_call
        agent_builder.add_conditional_edges(
            "llm_call",
            self._decide_next_step,
            {
                "tools": "tools",
                "user_input": "user_input",
                END: END,
            }
        )

        # Add conditional logic from user_input to handle pause
        agent_builder.add_conditional_edges(
            "user_input",
            self._decide_after_user_input,
            {
                "pause": END,  # End execution to allow for interruption
                "continue": "llm_call",
                END: END,
            }
        )

        # Compile and return the agent with interruption capability
        checkpointer = MemorySaver()
        return agent_builder.compile(
            checkpointer=checkpointer,
            interrupt_before=["user_input"]  # Interrupt before user input for manual handling
        )
    
    def _decide_next_step(self, state: AgentState) -> str:
        """
        Decide the next step based on the current state.
        """
        next_action = state.get("next_action", "end")
        logger.info(f"Next action determined: {next_action}")
        
        if next_action == "tools":
            return "tools"
        elif next_action == "user_input":
            return "user_input"
        else:
            return END

    def _decide_after_user_input(self, state: AgentState) -> str:
        """
        Decide what to do after user input node processing.
        """
        next_action = state.get("next_action", "end")
        logger.info(f"After user input, next action: {next_action}")
        
        if next_action == "pause":
            return "pause"
        elif next_action == "llm_call" or next_action == "continue":
            return "continue"
        else:
            return END
    
    async def run(self, initial_state: AgentState) -> AgentState:
        """
        Run the agent workflow.
        
        Args:
            initial_state: Initial state for the agent
            
        Returns:
            Final state after execution or interruption
        """
        logger.info("Starting agent execution")
        
        try:            
            # Run the graph
            config = {"configurable": {"thread_id": "unique_thread_id"}}
            result = await self.graph.ainvoke(initial_state, config=config)
            
            logger.info("Agent execution completed")
            return result
            
        except Exception as e:
            logger.error(f"Error in agent execution: {e}")
            raise
    
    async def continue_from_interruption(self, state: AgentState, user_response: str) -> AgentState:
        """
        Continue execution after user input interruption.
        
        Args:
            state: Current state when interrupted
            user_response: User's response to continue
            
        Returns:
            Final state after continuation
        """
        logger.info("Continuing from interruption")
        
        try:
            # Add user response to messages
            updated_state = dict(state)
            updated_state["messages"] = state["messages"] + [HumanMessage(content=user_response)]
            updated_state["pending_input_prompt"] = None
            updated_state["status"] = "active"
            updated_state["next_action"] = "continue"  # Signal to continue processing
            
            # Continue execution from the interruption point
            config = {"configurable": {"thread_id": "unique_thread_id"}}
            result = await self.graph.ainvoke(updated_state, config=config)
            
            logger.info("Continuation completed")
            return result
            
        except Exception as e:
            logger.error(f"Error continuing from interruption: {e}")
            raise