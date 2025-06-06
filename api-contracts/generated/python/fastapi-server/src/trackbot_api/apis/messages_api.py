# coding: utf-8

from typing import Dict, List  # noqa: F401
import importlib
import pkgutil

from trackbot_api.apis.messages_api_base import BaseMessagesApi
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
from trackbot_api.models.create_message_request import CreateMessageRequest
from trackbot_api.models.create_message_response import CreateMessageResponse
from trackbot_api.models.error import Error
from trackbot_api.models.message import Message


router = APIRouter()

ns_pkg = trackbot_api.impl
for _, name, _ in pkgutil.iter_modules(ns_pkg.__path__, ns_pkg.__name__ + "."):
    importlib.import_module(name)


@router.post(
    "/conversations/{conversationId}/messages",
    responses={
        201: {"model": CreateMessageResponse, "description": "Message created successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Messages"],
    summary="Create a new message",
    response_model_by_alias=True,
)
async def create_message(
    conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")] = Path(..., description="Conversation ID", ge=1),
    create_message_request: CreateMessageRequest = Body(None, description=""),
) -> CreateMessageResponse:
    """Send a new message in a conversation"""
    if not BaseMessagesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseMessagesApi.subclasses[0]().create_message(conversationId, create_message_request)


@router.delete(
    "/messages/{messageId}",
    responses={
        204: {"description": "Operation completed successfully with no content to return"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Messages"],
    summary="Delete message",
    response_model_by_alias=True,
)
async def delete_message(
    messageId: Annotated[int, Field(strict=True, ge=1, description="Message ID")] = Path(..., description="Message ID", ge=1),
) -> None:
    """Delete a message from a conversation"""
    if not BaseMessagesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseMessagesApi.subclasses[0]().delete_message(messageId)


@router.get(
    "/messages/{messageId}",
    responses={
        200: {"model": Message, "description": "Message details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Messages"],
    summary="Get message by ID",
    response_model_by_alias=True,
)
async def get_message_by_id(
    messageId: Annotated[int, Field(strict=True, ge=1, description="Message ID")] = Path(..., description="Message ID", ge=1),
) -> Message:
    """Retrieve a specific message by its ID"""
    if not BaseMessagesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseMessagesApi.subclasses[0]().get_message_by_id(messageId)


@router.get(
    "/conversations/{conversationId}/messages",
    responses={
        200: {"model": List[Message], "description": "List of messages retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Messages"],
    summary="List messages in a conversation",
    response_model_by_alias=True,
)
async def list_messages(
    conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")] = Path(..., description="Conversation ID", ge=1),
) -> List[Message]:
    """Retrieve all messages in a specific conversation"""
    if not BaseMessagesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseMessagesApi.subclasses[0]().list_messages(conversationId)
