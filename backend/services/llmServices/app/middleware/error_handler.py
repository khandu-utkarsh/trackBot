from fastapi import Request, status
from fastapi.responses import JSONResponse
from starlette.middleware.base import BaseHTTPMiddleware
import logging

logger = logging.getLogger(__name__)

class ErrorHandlerMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        try:
            # Log request details
            body = await request.body()
            logger.info(f"Raw request - Method: {request.method}, URL: {request.url}")
            logger.info(f"Request headers: {request.headers}")
            logger.info(f"Request body: {body.decode()}")
            
            return await call_next(request)
        except Exception as e:
            logger.error(f"Error processing request: {str(e)}", exc_info=True)
            return JSONResponse(
                status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
                content={"detail": "Internal server error"}
            ) 