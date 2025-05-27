from fastapi import APIRouter

from fastapi import APIRouter, HTTPException
from app.models import ProcessTextRequest, ProcessTextResponse
from app.services import LLMService

router = APIRouter()
llm_service = LLMService()

@router.post("/process_text")
async def process_text_handler(request: ProcessTextRequest):
    try:
        result: ProcessTextResponse = await llm_service.process_request(request.input)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e)) 