# coding: utf-8

from typing import Dict, List  # noqa: F401
import importlib
import pkgutil

from trackbot_api.apis.users_api_base import BaseUsersApi
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
from pydantic import Field, StrictStr
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.create_user_request import CreateUserRequest
from trackbot_api.models.create_user_response import CreateUserResponse
from trackbot_api.models.error import Error
from trackbot_api.models.user import User


router = APIRouter()

ns_pkg = trackbot_api.impl
for _, name, _ in pkgutil.iter_modules(ns_pkg.__path__, ns_pkg.__name__ + "."):
    importlib.import_module(name)


@router.post(
    "/users",
    responses={
        201: {"model": CreateUserResponse, "description": "User created successfully"},
        200: {"model": User, "description": "User already exists"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        409: {"model": Error, "description": "Conflict with current state of resource"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Users"],
    summary="Create a new user",
    response_model_by_alias=True,
)
async def create_user(
    create_user_request: CreateUserRequest = Body(None, description=""),
) -> User:
    """Create a new user account"""
    if not BaseUsersApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseUsersApi.subclasses[0]().create_user(create_user_request)


@router.delete(
    "/users/{userId}",
    responses={
        204: {"description": "Operation completed successfully with no content to return"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Users"],
    summary="Delete user",
    response_model_by_alias=True,
)
async def delete_user(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
) -> None:
    """Delete a user and all associated data"""
    if not BaseUsersApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseUsersApi.subclasses[0]().delete_user(userId)


@router.get(
    "/users/email/{email}",
    responses={
        200: {"model": User, "description": "User details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Users"],
    summary="Get user by email",
    response_model_by_alias=True,
)
async def get_user_by_email(
    email: Annotated[StrictStr, Field(description="User email address")] = Path(..., description="User email address"),
) -> User:
    """Retrieve a user by their email address"""
    if not BaseUsersApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseUsersApi.subclasses[0]().get_user_by_email(email)


@router.get(
    "/users/{userId}",
    responses={
        200: {"model": User, "description": "User details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Users"],
    summary="Get user by ID",
    response_model_by_alias=True,
)
async def get_user_by_id(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
) -> User:
    """Retrieve a specific user by their ID"""
    if not BaseUsersApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseUsersApi.subclasses[0]().get_user_by_id(userId)


@router.get(
    "/users",
    responses={
        200: {"model": List[User], "description": "List of users retrieved successfully"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Users"],
    summary="List all users",
    response_model_by_alias=True,
)
async def list_users(
) -> List[User]:
    """Retrieve a list of all users in the system"""
    if not BaseUsersApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseUsersApi.subclasses[0]().list_users()
