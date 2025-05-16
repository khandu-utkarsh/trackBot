from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from services.llm_service import LLMService
from config.settings import get_settings

settings = get_settings()
app = FastAPI(title=settings.PROJECT_NAME)
llm_service = LLMService()

class WorkoutRequest(BaseModel):
    input: str


#Currently this python service, is only going to support the POST request.


@app.post(f"{settings.API_V1_STR}/process")
async def process_workout_request(request: WorkoutRequest):
    try:
        result = await llm_service.process_request(request.input)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:app",
        host=settings.HOST,
        port=settings.PORT,
        reload=True
    ) 