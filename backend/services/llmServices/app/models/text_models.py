from pydantic import BaseModel, Field
from typing import Dict, Any

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