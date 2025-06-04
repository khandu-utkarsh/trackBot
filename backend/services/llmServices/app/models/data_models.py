from pydantic import BaseModel, Field
from typing import Dict, Any
from datetime import datetime
from typing import List, Optional


class ProcessTextRequest(BaseModel):
    input: str

class ProcessTextResponse(BaseModel):
    output: str

class WorkoutRequest(BaseModel):
    user_input: str = Field(description="User's workout description or request")
    user_id: str = Field(description="ID of the user making the request")
    context: Dict[str, Any] = Field(default_factory=dict, description="Additional context for the request")

class WorkoutAction(BaseModel):
    action: str = Field(description="The action to perform (create_workout, get_history, update_workout, delete_workout, get_statistics)")
    parameters: Dict[str, Any] = Field(description="Parameters required for the action") 


class Exercise(BaseModel):
    workout_id: str = Field(description="ID of the workout", default=None)
    name: str = Field(description="Name of the exercise")
    sets: int = Field(description="Number of sets")
    reps: int = Field(description="Number of repetitions per set")
    weight: Optional[float] = Field(description="Weight used (if applicable)", default=None)
    notes: Optional[str] = Field(description="Additional notes about the exercise", default=None)
    distance: Optional[float] = Field(description="Distance covered (if applicable)", default=None)
    time: Optional[int] = Field(description="Time taken (if applicable)", default=None)


class Workout(BaseModel):
    workout_id: str = Field(description="ID of the workout", default=None)
    date: datetime = Field(description="Date and time of the workout")
    name: str = Field(description="Name of the workout")
    exercises: List[Exercise] = Field(description="List of exercises in the workout")
    user_id: str = Field(description="ID of the user who created the workout")
