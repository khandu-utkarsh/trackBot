#This file is the entry point of the LLM Service
# ------------------------------------------------------------------------------------------------
# ------------------------------------------------------------------------------------------------
# ------------------------------------------------------------------------------------------------
from fastapi import FastAPI
from fastapi.middleware.gzip import GZipMiddleware
from config.settings import get_settings
from api.v1.api import api_router
from middleware import ErrorHandlerMiddleware
from middleware.auth import verify_token_middleware
from middleware.recovery import recovery_middleware
from middleware.logging import LoggingMiddleware
from middleware.cors import configure_cors
from middleware.request_id import RequestIDMiddleware
from config.database import init_db
from config.init_langgraph_db import init_langgraph_checkpointing
import logging
from contextlib import asynccontextmanager

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)

settings = get_settings()

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup
    init_db()
    await init_langgraph_checkpointing()
    yield
    # Shutdown
    pass

app = FastAPI(title=settings.PROJECT_NAME, lifespan=lifespan)

# Add middleware in the correct order
app.add_middleware(ErrorHandlerMiddleware)  # Outermost
app.add_middleware(LoggingMiddleware)
app.add_middleware(RequestIDMiddleware)
app.add_middleware(GZipMiddleware, minimum_size=1000)  # Compress responses larger than 1KB
app.middleware("http")(recovery_middleware)
configure_cors(app)  # Innermost

app.include_router(api_router, prefix=settings.API_STR)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:app",
        host= settings.HOST,
        port= settings.PORT,       # From docker-compose
        reload=True
    ) 