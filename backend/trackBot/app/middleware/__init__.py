"""
Middleware package for the LLM Service.
"""

from .error_handler import ErrorHandlerMiddleware

__all__ = ["ErrorHandlerMiddleware"]
