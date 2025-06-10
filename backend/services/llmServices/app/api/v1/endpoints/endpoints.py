from fastapi import APIRouter, HTTPException
from trackbot_client.models import LLMServiceMessageRequest, LLMServiceMessageResponse, Message
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage, BaseMessage
from services.agent_service import AgentService
import logging
from typing import List, Type

router = APIRouter()
logger = logging.getLogger(__name__)

class MessageConverter:
    """Handles conversion between TrackBot and LangChain message formats."""
    
    _message_type_map = {
        "assistant": AIMessage,
        "user": HumanMessage,
        "system": SystemMessage
    }
    
    _reverse_type_map = {
        AIMessage: "assistant",
        HumanMessage: "user",
        SystemMessage: "system"
    }
    
    @classmethod
    def to_langchain(cls, message: Message) -> BaseMessage:
        message_class = cls._message_type_map.get(message.message_type)
        if not message_class:
            raise ValueError(f"Unknown message type: {message.message_type}")
            
        if message_class == AIMessage:
            return message_class(
                content=message.content,
                additional_kwargs=message.additional_kwargs,
                response_metadata=message.response_metadata,
                name=message.name,
                id=message.llm_id,
                tool_calls=message.tool_calls,
                invalid_tool_calls=message.invalid_tool_calls,
                usage_metadata=message.usage_metadata,
                example=message.example
            )
        return message_class(content=message.content)
    
    @classmethod
    def to_trackbot(cls, message: BaseMessage) -> Message:
        message_type = cls._reverse_type_map.get(type(message))
        if not message_type:
            raise ValueError(f"Unknown message type: {type(message)}")
            
        if message_type == "assistant":
            return Message(
                message_type=message_type,
                content=message.content,
                additional_kwargs=message.additional_kwargs,
                response_metadata=message.response_metadata,
                name=message.name,
                llm_id=message.id,
                tool_calls=message.tool_calls,
                invalid_tool_calls=message.invalid_tool_calls,
                usage_metadata=message.usage_metadata,
                example=message.example
            )
        return Message(message_type=message_type, content=message.content)

async def process_with_agent(request: LLMServiceMessageRequest, is_continuation: bool = False) -> LLMServiceMessageResponse:
    """Common processing logic for both conversation endpoints."""
    try:
        langchain_messages = [MessageConverter.to_langchain(m) for m in request.messages]
        agent_service = AgentService(messages=langchain_messages, user_id=request.user_id)
        
        if is_continuation:
            result = await agent_service.continue_from_interruption(request.messages[-1])
        else:
            result = await agent_service.process_messages()
            
        response_message = MessageConverter.to_trackbot(result)
        return LLMServiceMessageResponse(
            message=response_message,
            response_time_ms=None,
            model_name=None
        )
    except ValueError as e:
        raise HTTPException(status_code=404, detail=str(e))
    except Exception as e:
        logger.error(f"Error processing request: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@router.post("/process_conversation", response_model=LLMServiceMessageResponse)
async def process_messages_handler(request: LLMServiceMessageRequest):
    """Process messages through the LangGraph agent workflow."""
    logger.info(f"Received request for user {request.user_id}")
    return await process_with_agent(request)

@router.post("/continue_conversation", response_model=LLMServiceMessageResponse)
async def continue_session_handler(request: LLMServiceMessageRequest):
    """Continue a session after user input."""
    logger.info(f"Continuing session {request.user_id}")
    return await process_with_agent(request, is_continuation=True)