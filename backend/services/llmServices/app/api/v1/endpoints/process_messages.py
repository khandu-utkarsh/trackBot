from fastapi import APIRouter, HTTPException
from app.models import ProcessMessagesRequest, ProcessMessagesResponse
from app.services import LLMService
import logging
from langchain_core.messages import BaseMessage
router = APIRouter()
llm_service = LLMService()
logger = logging.getLogger(__name__)

@router.post("/process_messages")
async def process_text_handler(request: ProcessMessagesRequest):
    #logger.info(f"Received request body: {request.model_dump()}")
    try:
        print(request.messages)

        langchain_messages = [m.to_langchain() for m in request.messages]
        print(langchain_messages)

        # Convert messages to LangChain format
        llm_result: BaseMessage = await llm_service.process_request(langchain_messages, request.user_id)
        return ProcessMessagesResponse(message=llm_result)
    except Exception as e:
        logger.error(f"Error processing request: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))