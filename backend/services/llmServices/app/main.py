#This file is the entry point of the LLM Service
# ------------------------------------------------------------------------------------------------
# ------------------------------------------------------------------------------------------------
# ------------------------------------------------------------------------------------------------
from fastapi import FastAPI
from app.config.settings import get_settings
from app.api.v1.api import api_router
from app.middleware import ErrorHandlerMiddleware

settings = get_settings()

fastApiApp  = FastAPI(title=settings.PROJECT_NAME)
fastApiApp.add_middleware(ErrorHandlerMiddleware)
fastApiApp.include_router(api_router, prefix=settings.API_V1_STR)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:fastApiApp",
        host="0.0.0.0",  # Docker default
        port=8081,       # From docker-compose
        reload=True
    ) 