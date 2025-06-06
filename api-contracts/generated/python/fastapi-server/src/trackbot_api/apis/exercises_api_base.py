# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from pydantic import Field
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.create_exercise_request import CreateExerciseRequest
from trackbot_api.models.create_exercise_response import CreateExerciseResponse
from trackbot_api.models.error import Error
from trackbot_api.models.exercise import Exercise
from trackbot_api.models.update_exercise_request import UpdateExerciseRequest


class BaseExercisesApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseExercisesApi.subclasses = BaseExercisesApi.subclasses + (cls,)
    async def create_exercise(
        self,
        workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")],
        create_exercise_request: CreateExerciseRequest,
    ) -> CreateExerciseResponse:
        """Add a new exercise to a workout (cardio or weight training)"""
        ...


    async def delete_exercise(
        self,
        exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")],
    ) -> None:
        """Delete an exercise from a workout"""
        ...


    async def get_exercise_by_id(
        self,
        exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")],
    ) -> Exercise:
        """Retrieve a specific exercise by its ID"""
        ...


    async def list_exercises(
        self,
        workoutId: Annotated[int, Field(strict=True, ge=1, description="Workout ID")],
    ) -> List[Exercise]:
        """Retrieve all exercises for a specific workout"""
        ...


    async def update_exercise(
        self,
        exerciseId: Annotated[int, Field(strict=True, ge=1, description="Exercise ID")],
        update_exercise_request: UpdateExerciseRequest,
    ) -> Exercise:
        """Update an existing exercise"""
        ...
