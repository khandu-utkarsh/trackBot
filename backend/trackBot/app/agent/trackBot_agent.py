from typing import Dict, Any
from langchain_core.messages import BaseMessage, HumanMessage
from langgraph.graph import StateGraph, END, START
from langgraph.checkpoint.postgres import PostgresSaver
from agent.state import AgentState
from agent.nodes import llm_node, tool_node, user_input_node
import logging
from config.settings import get_settings

logger = logging.getLogger(__name__)
settings = get_settings()

class TrackBotAgent:
    """
    Agent class that orchestrates the LangGraph workflow for workout logging.
    The workflow follows these steps:
    1. Process user input through LLM
    2. If more information is needed, ask user
    3. If all information is available, execute tools
    4. Process tool results and end conversation
    """
    
    def __init__(self, user_id: int, conversation_id: int, graph: StateGraph):
        self.user_id = user_id
        self.conversation_id = conversation_id
        self.graph = graph
    
    @classmethod
    def create(cls, user_id: int, conversation_id: int) -> "TrackBotAgent":
        agent_builder = StateGraph(AgentState)
        
        # Add nodes
        agent_builder.add_node("llm_call", llm_node)
        agent_builder.add_node("tools", tool_node)
        agent_builder.add_node("user_input", user_input_node)

        # Add edges
        agent_builder.add_edge(START, "llm_call")
        agent_builder.add_edge("tools", "llm_call")

        # Add conditional edges
        agent_builder.add_conditional_edges(
            "llm_call",
            cls._decide_next_step,
            {"tools": "tools", "user_input": "user_input", END: END}
        )

        agent_builder.add_conditional_edges(
            "user_input",
            cls._decide_after_user_input,
            {"pause": END, "continue": "llm_call", END: END}
        )

        # Create the PostgresSaver instance
        checkpointer = PostgresSaver.from_conn_string(settings.DATABASE_URL)
        
        # Compile the graph with the checkpointer
        graph = agent_builder.compile(
            interrupt_before=["user_input"],
        )
        
        return cls(user_id, conversation_id, graph)
    
    @staticmethod
    def _decide_next_step(state: AgentState) -> str:
        next_action = state.get("next_action", "end")
        return {"tools": "tools", "user_input": "user_input", END: END}.get(next_action, END)

    @staticmethod
    def _decide_after_user_input(state: AgentState) -> str:
        next_action = state.get("next_action", "end")
        return {"pause": "pause", "llm_call": "continue", "continue": "continue"}.get(next_action, END)
    
    async def run(self, state: AgentState) -> AgentState:
        """
        Run the agent with the given state.
        """
        config = {
            "configurable": {
                "thread_id": f"{self.user_id}_{self.conversation_id}"
            }
        }
        
        # Run the graph with the state
        result = await self.graph.ainvoke(state, config=config)
        return result