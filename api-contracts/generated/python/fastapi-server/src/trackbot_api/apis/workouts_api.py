# coding: utf-8

from typing import Dict, List  # noqa: F401
import importlib
import pkgutil

from trackbot_api.apis.workouts_api_base import BaseWorkoutsApi
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
from pydantic import Field, field_validator
from typing import Any, List, Optional
from typing_extensions import Annotated
from trackbot_api.models.create_workout_request import CreateWorkoutRequest
from trackbot_api.models.create_workout_response import CreateWorkoutResponse
from trackbot_api.models.error import Error
from trackbot_api.models.update_workout_request import UpdateWorkoutRequest
from trackbot_api.models.workout import Workout


router = APIRouter()

ns_pkg = trackbot_api.impl
for _, name, _ in pkgutil.iter_modules(ns_pkg.__path__, ns_pkg.__name__ + "."):
    importlib.import_module(name)


@router.post(
    "/users/{userId}/workouts",
    responses={
        201: {"model": CreateWorkoutResponse, "description": "Workout created successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Workouts"],
    summary="Create a new workout",
    response_model_by_alias=True,
)
async def create_workout(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
    create_workout_request: CreateWorkoutRequest = Body(None, description=""),
) -> CreateWorkoutResponse:
    """Create a new workout session for a user"""
    if not BaseWorkoutsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseWorkoutsApi.subclasses[0]().create_workout(userId, create_workout_request)


@router.delete(
    "/workouts/{workoutId}",
    responses={
        204: {"description": "Operation completed successfully with no content to return"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Workouts"],
    summary="Delete workout",
    response_model_by_alias=True,
)
async def delete_workout(
    workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")] = Path(..., description="Workout ID", ge=1),
) -> None:
    """Delete a workout and all associated exercises"""
    if not BaseWorkoutsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseWorkoutsApi.subclasses[0]().delete_workout(workoutId)


@router.get(
    "/workouts/{workoutId}",
    responses={
        200: {"model": Workout, "description": "Workout details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Workouts"],
    summary="Get workout by ID",
    response_model_by_alias=True,
)
async def get_workout_by_id(
    workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")] = Path(..., description="Workout ID", ge=1),
) -> Workout:
    """Retrieve a specific workout by its ID"""
    if not BaseWorkoutsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseWorkoutsApi.subclasses[0]().get_workout_by_id(workoutId)


@router.get(
    "/users/{userId}/workouts",
    responses={
        200: {"model": List[Workout], "description": "List of workouts retrieved successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Workouts"],
    summary="List workouts for a user",
    response_model_by_alias=True,
)
async def list_workouts(
    userId: Annotated[int, Field(strict=True, ge=1, description="User ID")] = Path(..., description="User ID", ge=1),
    year: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by year (YYYY format)")] = Query(None, description="Filter by year (YYYY format)", alias="year", regex=r"/^\d{4}$/"),
    month: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by month (MM format)")] = Query(None, description="Filter by month (MM format)", alias="month", regex=r"/^(0[1-9]|1[0-2])$/"),
    day: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by day (DD format)")] = Query(None, description="Filter by day (DD format)", alias="day", regex=r"/^(0[1-9]|[12][0-9]|3[01])$/"),
) -> List[Workout]:
    """Retrieve all workouts for a specific user with optional date filtering"""
    if not BaseWorkoutsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseWorkoutsApi.subclasses[0]().list_workouts(userId, year, month, day)


@router.put(
    "/workouts/{workoutId}",
    responses={
        200: {"model": Workout, "description": "Workout updated successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Workouts"],
    summary="Update workout",
    response_model_by_alias=True,
)
async def update_workout(
    workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")] = Path(..., description="Workout ID", ge=1),
    update_workout_request: UpdateWorkoutRequest = Body(None, description=""),
) -> Workout:
    """Update an existing workout"""
    if not BaseWorkoutsApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseWorkoutsApi.subclasses[0]().update_workout(workoutId, update_workout_request)
