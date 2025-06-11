from fastapi import Request, status
from fastapi.responses import JSONResponse
import logging
import traceback

logger = logging.getLogger(__name__)

async def recovery_middleware(request: Request, call_next):
    """
    Middleware to handle unhandled exceptions and provide graceful error responses
    """
    try:
        return await call_next(request)
    except Exception as e:
        logger.error(f"Unhandled exception: {str(e)}")
        logger.error(traceback.format_exc())
        
        return JSONResponse(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            content={
                "error": "Internal server error",
                "detail": str(e) if request.app.debug else "An unexpected error occurred"
            }
        ) 