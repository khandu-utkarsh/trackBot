from typing import Dict, Any, Optional, Union
from langchain_core.tools import tool
import logging
from trackbot_client.api_client import ApiClient
from trackbot_client.api.workouts_api import WorkoutsApi
from trackbot_client.api.exercises_api import ExercisesApi
from trackbot_client.models.create_workout_request import CreateWorkoutRequest
from trackbot_client.models.create_cardio_exercise_request import CreateCardioExerciseRequest
from trackbot_client.models.create_strength_exercise_request import CreateStrengthExerciseRequest
from trackbot_client.configuration import Configuration

logger = logging.getLogger(__name__)


# Initialize API client and APIs with correct host
api_client = ApiClient(configuration=Configuration(host="http://workout-app:8080/api"))
workouts_api = WorkoutsApi(api_client)
exercises_api = ExercisesApi(api_client)

@tool("create_workout", description="Create a new workout", args_schema=CreateWorkoutRequest)
async def create_workout(input: Dict[str, Any]) -> str:
    """
    Create a new workout.
    
    Args:
        input: Dictionary containing workout details (name, description, etc.)
        
    Returns:
        str: Success or error message with workout details
    """
    try:
        workout_request = CreateWorkoutRequest.model_validate(input)
        logger.info(f"Creating workout: {workout_request.name}")
        
        workout = workouts_api.create_workout(workout_request)
        return f"Successfully created workout: {workout.name} (ID: {workout.id})"
    except Exception as e:
        logger.error(f"Error creating workout: {str(e)}")
        return f"Failed to create workout: {str(e)}"

@tool("create_cardio_exercise", description="Create a new cardio exercise", args_schema=CreateCardioExerciseRequest)
async def create_cardio_exercise(user_id: int, workout_id: int, name: str, distance: Union[int, float], duration: int, notes: Optional[str] = None) -> str:
    """
    Create a new cardio exercise.
    
    Args:
        input: Dictionary containing cardio exercise details (name, distance, time, etc.)
        
    Returns:
        str: Success or error message with exercise details
    """
    try:
        cardio_request = CreateCardioExerciseRequest(user_id=user_id, workout_id=workout_id, name=name, distance=distance, duration=duration, notes=notes)
        logger.info(f"Creating cardio exercise: {cardio_request.name}")
        
        exercise = exercises_api.create_cardio_exercise(user_id=user_id, workout_id=workout_id, create_cardio_exercise_request=cardio_request)
        return f"Successfully created cardio exercise: {exercise.name} (ID: {exercise.id})"
    except Exception as e:
        logger.error(f"Error creating cardio exercise: {str(e)}")
        return f"Failed to create cardio exercise: {str(e)}"

@tool("create_strength_exercise", description="Create a new strength exercise", args_schema=CreateStrengthExerciseRequest)
async def create_strength_exercise(user_id: int, workout_id: int, name: str, reps: int, weight: Union[int, float], notes: Optional[str] = None) -> str:
    """
    Create a new strength exercise.
    
    Args:
        input: Dictionary containing strength exercise details (name, reps, weights, etc.)
        
    Returns:
        str: Success or error message with exercise details
    """
    try:
        strength_request = CreateStrengthExerciseRequest(user_id=user_id, workout_id=workout_id, name=name, reps=reps, weight=weight, notes=notes)
        logger.info(f"Creating strength exercise: {strength_request.name}")
        
        exercise = exercises_api.create_strength_exercise(user_id=user_id, workout_id=workout_id, create_strength_exercise_request=strength_request)
        return f"Successfully created strength exercise: {exercise.name} (ID: {exercise.id})"
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
        # Workout management tools
        "create_workout": create_workout,
        # Exercise management tools
        "create_cardio_exercise": create_cardio_exercise,
        "create_strength_exercise": create_strength_exercise,
    }