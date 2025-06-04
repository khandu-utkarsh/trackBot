from fastapi import APIRouter, HTTPException
from app.models import ProcessTextRequest, ProcessTextResponse
from app.services import LLMService
import logging

router = APIRouter()
llm_service = LLMService()
logger = logging.getLogger(__name__)

@router.post("/process_messages")
async def process_text_handler(request: ProcessTextRequest):
    logger.info(f"Received request body: {request.model_dump()}")
    try:
        result: ProcessTextResponse = await llm_service.process_request(request.input)
        return result
    except Exception as e:
        logger.error(f"Error processing request: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))