from typing import Dict, Any, Optional, Union
from langchain_core.tools import tool
from models.trackBot_models import Exercise
from sqlalchemy.orm import Session
import logging

logger = logging.getLogger(__name__)


from pydantic import BaseModel, Field, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing_extensions import Annotated
from typing import Optional, Set
from typing_extensions import Self

class CreateCardioExerciseRequest(BaseModel):
    """
    CreateCardioExerciseRequest
    """
    user_id: StrictInt = Field(description="ID of the user creating the exercise")
    name: StrictStr = Field(description="Name of the exercise")
    description: Optional[StrictStr] = Field(default=None, description="Additional notes")
    distance: Union[Annotated[float, Field(strict=True, ge=0)], Annotated[int, Field(strict=True, ge=0)]] = Field(description="Distance in meters")
    duration: Annotated[int, Field(strict=True, ge=1)] = Field(description="Duration in seconds")

class CreateStrengthExerciseRequest(BaseModel):
    """
    CreateStrengthExerciseRequest
    """ # noqa: E501
    user_id: StrictInt = Field(description="ID of the user creating the exercise")
    name: StrictStr = Field(description="Name of the exercise")
    description: Optional[StrictStr] = Field(default=None, description="Additional notes")
    reps: Annotated[int, Field(strict=True, ge=1)] = Field(description="Repetitions per set")
    weight: Union[Annotated[float, Field(strict=True, ge=0)], Annotated[int, Field(strict=True, ge=0)]] = Field(description="Weight in kilograms")


@tool("create_cardio_exercise", description="Create a new cardio exercise", args_schema=CreateCardioExerciseRequest)
async def create_cardio_exercise(user_id: int, name: str, distance: Union[int, float], duration: int, description: Optional[str] = None) -> str:
    """
    Create a new cardio exercise.
    
    Args:
        user_id: ID of the user creating the exercise
        name: Name of the exercise
        distance: Distance in meters
        duration: Duration in seconds
        description: Optional description of the exercise
        
    Returns:
        str: Success or error message with exercise details
    """
    try:
        # Create new exercise record
        exercise = Exercise(
            user_id=user_id,
            name=name,
            description=description,
            duration=duration,
            distance=distance
        )
        
        # Get database session
        db = Session()
        try:
            db.add(exercise)
            db.commit()
            db.refresh(exercise)
            logger.info(f"Created cardio exercise: {exercise.name}")
            return f"Successfully created cardio exercise: {exercise.name} (ID: {exercise.id})"
        finally:
            db.close()
            
    except Exception as e:
        logger.error(f"Error creating cardio exercise: {str(e)}")
        return f"Failed to create cardio exercise: {str(e)}"

@tool("create_strength_exercise", description="Create a new strength exercise", args_schema=CreateStrengthExerciseRequest)
async def create_strength_exercise(user_id: int, name: str, reps: int, weight: Union[int, float], description: Optional[str] = None) -> str:
    """
    Create a new strength exercise.
    
    Args:
        user_id: ID of the user creating the exercise
        name: Name of the exercise
        reps: Number of repetitions
        weight: Weight in kilograms
        notes: Optional notes about the exercise
        
    Returns:
        str: Success or error message with exercise details
    """
    try:
        # Create new exercise record
        exercise = Exercise(
            user_id=user_id,
            name=name,
            description=description,
            reps=reps,
            weight=weight,
        )
        
        # Get database session
        db = Session()
        try:
            db.add(exercise)
            db.commit()
            db.refresh(exercise)
            logger.info(f"Created strength exercise: {exercise.name}")
            return f"Successfully created strength exercise: {exercise.name} (ID: {exercise.id})"
        finally:
            db.close()
            
    except Exception as e:
        logger.error(f"Error creating strength exercise: {str(e)}")
        return f"Failed to create strength exercise: {str(e)}"

def get_available_tools() -> Dict[str, Any]:
    """
    Get all available tools for the agent.
    
    Returns:
        Dictionary mapping tool names to tool functions
    """
    return {
        # Exercise management tools
        "create_cardio_exercise": create_cardio_exercise,
        "create_strength_exercise": create_strength_exercise,
    }