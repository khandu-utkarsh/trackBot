# coding: utf-8

from typing import Dict, List  # noqa: F401
import importlib
import pkgutil

from trackbot_api.apis.conversations_api_base import BaseConversationsApi
import trackbot_api.impl

from fastapi import (  # noqa: F401
    APIRouter,
    Body,
    Cookie,
    Depends,
    Form,
    Header,
    HTTPException,
    Path,
    Query,
    Response,
    Security,
    status,
)

from trackbot_api.models.extra_models import TokenModel  # noqa: F401
from pydantic import Field
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.conversation import Conversation
from trackbot_api.models.create_conversation_request import CreateConversationRequest
from trackbot_api.models.create_conversation_response import CreateConversationResponse
from trackbot_api.models.error import Error


router = APIRouter()

ns_pkg = trackbot_api.impl
for _, name, _ in pkgutil.iter_modules(ns_pkg.__path__, ns_pkg.__name__ + "."):
    importlib.import_module(name)


@router.post(
    "/users/{userId}/conversations",
    responses={
        201: {"model": CreateConversationResponse, "description": "Conversation created successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Conversations"],
    summary="Create a new conversation",
    response_model_by_alias=True,
)
async def create_conversation(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
    create_conversation_request: CreateConversationRequest = Body(None, description=""),
) -> CreateConversationResponse:
    """Start a new conversation with the AI assistant"""
    if not BaseConversationsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseConversationsApi.subclasses[0]().create_conversation(userId, create_conversation_request)


@router.delete(
    "/conversations/{conversationId}",
    responses={
        204: {"description": "Operation completed successfully with no content to return"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Conversations"],
    summary="Delete conversation",
    response_model_by_alias=True,
)
async def delete_conversation(
    conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")] = Path(..., description="Conversation ID", ge=1),
) -> None:
    """Delete a conversation and all associated messages"""
    if not BaseConversationsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseConversationsApi.subclasses[0]().delete_conversation(conversationId)


@router.get(
    "/conversations/{conversationId}",
    responses={
        200: {"model": Conversation, "description": "Conversation details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Conversations"],
    summary="Get conversation by ID",
    response_model_by_alias=True,
)
async def get_conversation_by_id(
    conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")] = Path(..., description="Conversation ID", ge=1),
) -> Conversation:
    """Retrieve a specific conversation by its ID"""
    if not BaseConversationsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseConversationsApi.subclasses[0]().get_conversation_by_id(conversationId)


@router.get(
    "/users/{userId}/conversations",
    responses={
        200: {"model": List[Conversation], "description": "List of conversations retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Conversations"],
    summary="List conversations for a user",
    response_model_by_alias=True,
)
async def list_conversations(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
) -> List[Conversation]:
    """Retrieve all conversations for a specific user"""
    if not BaseConversationsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseConversationsApi.subclasses[0]().list_conversations(userId)
