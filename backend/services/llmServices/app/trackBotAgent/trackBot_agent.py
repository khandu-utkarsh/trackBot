from typing import Dict, Any
from langchain_core.messages import BaseMessage, HumanMessage
from langgraph.graph import StateGraph, END, START
from langgraph.checkpoint.memory import MemorySaver
from trackBotAgent.state import AgentState
from trackBotAgent.nodes import llm_node, tool_node, user_input_node
import logging

logger = logging.getLogger(__name__)

class TrackBotAgent:
    """
    Agent class that orchestrates the LangGraph workflow for workout logging.
    The workflow follows these steps:
    1. Process user input through LLM
    2. If more information is needed, ask user
    3. If all information is available, execute tools
    4. Process tool results and end conversation
    """
    
    def __init__(self):
        self.graph = self._build_graph()
    
    def _build_graph(self) -> StateGraph:
        """
        Build the LangGraph workflow with clear state transitions.
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
                "tools": "tools",        # Execute tools if needed
                "user_input": "user_input",  # Ask user for more info
                END: END,               # End if complete
            }
        )

        # Add conditional logic from user_input
        agent_builder.add_conditional_edges(
            "user_input",
            self._decide_after_user_input,
            {
                "pause": END,           # End to wait for user response
                "continue": "llm_call", # Process user response
                END: END,              # End if complete
            }
        )

        # Compile the graph with interruption capability
        checkpointer = MemorySaver()
        return agent_builder.compile(
            checkpointer=checkpointer,
            interrupt_before=["user_input"]  # Allow interruption before user input
        )
    
    def _decide_next_step(self, state: AgentState) -> str:
        """
        Decide the next step based on the current state.
        """
        next_action = state.get("next_action", "end")
        logger.info(f"Next action determined: {next_action}")
        
        if next_action == "tools":
            logger.info("Proceeding to tool execution")
            return "tools"
        elif next_action == "user_input":
            logger.info("Requesting user input")
            return "user_input"
        else:
            logger.info("Ending conversation")
            return END

    def _decide_after_user_input(self, state: AgentState) -> str:
        """
        Decide what to do after user input node processing.
        """
        next_action = state.get("next_action", "end")
        logger.info(f"After user input, next action: {next_action}")
        
        if next_action == "pause":
            logger.info("Pausing for user response")
            return "pause"
        elif next_action in ["llm_call", "continue"]:
            logger.info("Continuing with LLM processing")
            return "continue"
        else:
            logger.info("Ending conversation")
            return END
    
    async def run(self, state: AgentState) -> AgentState:
        """
        Run the agent workflow with the given state.
        
        Args:
            state: Current state for the agent
            
        Returns:
            Updated state after execution
        """
        logger.info(f"Starting agent execution for user {state['user_id']}")
        
        try:            
            # Run the graph with a unique thread ID
            config = {"configurable": {"thread_id": f"thread_{state['user_id']}"}}
            result = await self.graph.ainvoke(state, config=config)
            
            logger.info(f"Agent execution completed for user {state['user_id']}")
            return result
            
        except Exception as e:
            logger.error(f"Error in agent execution: {e}")
            raise
    
    async def continue_from_interruption(self, state: AgentState, user_response: BaseMessage) -> AgentState:
        """
        Continue execution after user input interruption.
        
        Args:
            state: Current state when interrupted
            user_response: User's response to continue
            
        Returns:
            Updated state after continuation
        """
        logger.info(f"Continuing from interruption for user {state['user_id']}")
        
        try:
            # Update state with user response
            updated_state = dict(state)
            updated_state["messages"] = state["messages"] + [user_response]
            updated_state["pending_input_prompt"] = None
            updated_state["status"] = "active"
            updated_state["next_action"] = "continue"
            
            # Continue execution from the interruption point
            config = {"configurable": {"thread_id": f"thread_{state['user_id']}"}}
            result = await self.graph.ainvoke(updated_state, config=config)
            
            logger.info(f"Continuation completed for user {state['user_id']}")
            return result
            
        except Exception as e:
            logger.error(f"Error continuing from interruption: {e}")
            raise