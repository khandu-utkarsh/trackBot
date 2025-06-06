# coding: utf-8

from typing import Dict, List  # noqa: F401
import importlib
import pkgutil

from trackbot_api.apis.exercises_api_base import BaseExercisesApi
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
from trackbot_api.models.create_exercise_request import CreateExerciseRequest
from trackbot_api.models.create_exercise_response import CreateExerciseResponse
from trackbot_api.models.error import Error
from trackbot_api.models.exercise import Exercise
from trackbot_api.models.update_exercise_request import UpdateExerciseRequest


router = APIRouter()

ns_pkg = trackbot_api.impl
for _, name, _ in pkgutil.iter_modules(ns_pkg.__path__, ns_pkg.__name__ + "."):
    importlib.import_module(name)


@router.post(
    "/workouts/{workoutId}/exercises",
    responses={
        201: {"model": CreateExerciseResponse, "description": "Exercise created successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Exercises"],
    summary="Create a new exercise",
    response_model_by_alias=True,
)
async def create_exercise(
    workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")] = Path(..., description="Workout ID", ge=1),
    create_exercise_request: CreateExerciseRequest = Body(None, description=""),
) -> CreateExerciseResponse:
    """Add a new exercise to a workout (cardio or weight training)"""
    if not BaseExercisesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseExercisesApi.subclasses[0]().create_exercise(workoutId, create_exercise_request)


@router.delete(
    "/exercises/{exerciseId}",
    responses={
        204: {"description": "Operation completed successfully with no content to return"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Exercises"],
    summary="Delete exercise",
    response_model_by_alias=True,
)
async def delete_exercise(
    exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")] = Path(..., description="Exercise ID", ge=1),
) -> None:
    """Delete an exercise from a workout"""
    if not BaseExercisesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseExercisesApi.subclasses[0]().delete_exercise(exerciseId)


@router.get(
    "/exercises/{exerciseId}",
    responses={
        200: {"model": Exercise, "description": "Exercise details retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Exercises"],
    summary="Get exercise by ID",
    response_model_by_alias=True,
)
async def get_exercise_by_id(
    exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")] = Path(..., description="Exercise ID", ge=1),
) -> Exercise:
    """Retrieve a specific exercise by its ID"""
    if not BaseExercisesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseExercisesApi.subclasses[0]().get_exercise_by_id(exerciseId)


@router.get(
    "/workouts/{workoutId}/exercises",
    responses={
        200: {"model": List[Exercise], "description": "List of exercises retrieved successfully"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Exercises"],
    summary="List exercises for a workout",
    response_model_by_alias=True,
)
async def list_exercises(
    workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")] = Path(..., description="Workout ID", ge=1),
) -> List[Exercise]:
    """Retrieve all exercises for a specific workout"""
    if not BaseExercisesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseExercisesApi.subclasses[0]().list_exercises(workoutId)


@router.put(
    "/exercises/{exerciseId}",
    responses={
        200: {"model": Exercise, "description": "Exercise updated successfully"},
        400: {"model": Error, "description": "Bad request - invalid input or parameters"},
        404: {"model": Error, "description": "Resource not found"},
        500: {"model": Error, "description": "Internal server error"},
    },
    tags=["Exercises"],
    summary="Update exercise",
    response_model_by_alias=True,
)
async def update_exercise(
    exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")] = Path(..., description="Exercise ID", ge=1),
    update_exercise_request: UpdateExerciseRequest = Body(None, description=""),
) -> Exercise:
    """Update an existing exercise"""
    if not BaseExercisesApi.subclasses:
        raise HTTPException(status_code=500, detail="Not implemented")
    return await BaseExercisesApi.subclasses[0]().update_exercise(exerciseId, update_exercise_request)
