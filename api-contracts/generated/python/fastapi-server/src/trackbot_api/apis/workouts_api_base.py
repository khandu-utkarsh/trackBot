# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from pydantic import Field, field_validator
from typing import Any, List, Optional
from typing_extensions import Annotated
from trackbot_api.models.create_workout_request import CreateWorkoutRequest
from trackbot_api.models.create_workout_response import CreateWorkoutResponse
from trackbot_api.models.error import Error
from trackbot_api.models.update_workout_request import UpdateWorkoutRequest
from trackbot_api.models.workout import Workout


class BaseWorkoutsApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseWorkoutsApi.subclasses = BaseWorkoutsApi.subclasses + (cls,)
    async def create_workout(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
        create_workout_request: CreateWorkoutRequest,
    ) -> CreateWorkoutResponse:
        """Create a new workout session for a user"""
        ...


    async def delete_workout(
        self,
        workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")],
    ) -> None:
        """Delete a workout and all associated exercises"""
        ...


    async def get_workout_by_id(
        self,
        workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")],
    ) -> Workout:
        """Retrieve a specific workout by its ID"""
        ...


    async def list_workouts(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
        year: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by year (YYYY format)")],
        month: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by month (MM format)")],
        day: Annotated[Optional[Annotated[str, Field(strict=True)]], Field(description="Filter by day (DD format)")],
    ) -> List[Workout]:
        """Retrieve all workouts for a specific user with optional date filtering"""
        ...


    async def update_workout(
        self,
        workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")],
        update_workout_request: UpdateWorkoutRequest,
    ) -> Workout:
        """Update an existing workout"""
        ...
