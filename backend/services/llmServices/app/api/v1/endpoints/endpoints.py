from fastapi import APIRouter, HTTPException
from trackbot_client.models import LLMServiceMessageRequest, LLMServiceMessageResponse, Message
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage, BaseMessage

from services.agent_service import AgentService
import logging

router = APIRouter()
logger = logging.getLogger(__name__)


def convert_message_to_langchain(message: Message) -> BaseMessage:
    if message.message_type == "assistant":
        return AIMessage(content=message.content, additional_kwargs=message.additional_kwargs, response_metadata=message.response_metadata, name=message.name, id=message.llm_id, tool_calls=message.tool_calls, invalid_tool_calls=message.invalid_tool_calls, usage_metadata=message.usage_metadata, example=message.example)
    elif message.message_type == "user":
        return HumanMessage(content=message.content)
    elif message.message_type == "system":
        return SystemMessage(content=message.content)
    else:
        raise ValueError(f"Unknown message type: {message.message_type}")

def convert_message_to_trackbot(message: BaseMessage) -> Message:
    if isinstance(message, AIMessage):
        return Message(message_type="assistant", content=message.content, additional_kwargs=message.additional_kwargs, response_metadata=message.response_metadata, name=message.name, llm_id=message.id, tool_calls=message.tool_calls, invalid_tool_calls=message.invalid_tool_calls, usage_metadata=message.usage_metadata, example=message.example)
    elif isinstance(message, HumanMessage):
        return Message(message_type="user", content=message.content)
    elif isinstance(message, SystemMessage):
        return Message(message_type="system", content=message.content)
    else:
        raise ValueError(f"Unknown message type: {type(message)}")


@router.post("/process_conversation", response_model=LLMServiceMessageResponse)
async def process_messages_handler(request: LLMServiceMessageRequest):
    """Process messages through the LangGraph agent workflow."""
    logger.info(f"Received request for user {request.user_id}")
    
    try:
        # Convert request messages to LangChain format
        langchain_messages = [convert_message_to_langchain(m) for m in request.messages]

        print("Calling agent service")
        
        agent_service = AgentService(messages=langchain_messages, user_id=request.user_id)

        print("Agent service called")

        # Process through agent service
        result : AIMessage = await agent_service.process_messages()
        
        print("Result received")
        
        # Convert response message back to our format
        response_message : Message = convert_message_to_trackbot(result)
        
        print("Response message converted")
        
        return LLMServiceMessageResponse( message=response_message, response_time_ms=None, model_name=None)
        
    except Exception as e:
        logger.error(f"Error processing request: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@router.post("/continue_conversation", response_model=LLMServiceMessageResponse)
async def continue_session_handler(request: LLMServiceMessageRequest):
    """Continue a session after user input."""
    logger.info(f"Continuing session {request.user_id}")
    
    try:
        # Convert request messages to LangChain format
        langchain_messages = [convert_message_to_langchain(m) for m in request.messages]
        
        agent_service = AgentService(messages=langchain_messages, user_id=request.user_id)
        result : AIMessage = await agent_service.continue_from_interruption(request.messages[-1])
        
        # Convert response message back to our format
        response_message : Message = convert_message_to_trackbot(result)
        
        return LLMServiceMessageResponse(
            message=response_message,
            response_time_ms=None,
            model_name=None
        )
        
    except ValueError as e:
        raise HTTPException(status_code=404, detail=str(e))
    except Exception as e:
        logger.error(f"Error continuing session: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))