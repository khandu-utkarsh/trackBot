from typing import Dict, Any, Optional
from langchain_core.tools import tool
import logging
from internal.generated.trackbot_client.api_client import ApiClient
from internal.generated.trackbot_client.api.workouts_api import WorkoutsApi
from internal.generated.trackbot_client.api.exercises_api import ExercisesApi
from internal.generated.trackbot_client.models.create_workout_request import CreateWorkoutRequest
from internal.generated.trackbot_client.models.create_exercise_request import CreateExerciseRequest
logger = logging.getLogger(__name__)

# Initialize API client and APIs
api_client = ApiClient()
workouts_api = WorkoutsApi(api_client)
exercises_api = ExercisesApi(api_client)

@tool
async def create_workout(CreateWorkoutRequest: CreateWorkoutRequest) -> str:
    """
    Create a new workout.
    
    Args:
        CreateWorkoutRequest: CreateWorkoutRequest object
        
    Returns:
        Workout creation status and information
    """
    logger.info(f"Creating workout: {CreateWorkoutRequest.name}")
    try:
        workout = workouts_api.create_workout(CreateWorkoutRequest)
        return f"Successfully created workout: {workout.name} (ID: {workout.id})"
    except Exception as e:
        return f"Error creating workout: {str(e)}"

@tool
async def get_workout(workout_id: str) -> str:
    """
    Get a workout by ID.
    
    Args:
        workout_id: ID of the workout to retrieve
        
    Returns:
        Workout information
    """
    logger.info(f"Getting workout: {workout_id}")
    try:
        workout = workouts_api.get_workout_by_id(workout_id)
        return f"Workout: {workout.name} (ID: {workout.id}) - {workout.description}"
    except Exception as e:
        return f"Error getting workout: {str(e)}"

@tool
async def get_user_workouts(user_id: str, limit: Optional[int] = None, offset: Optional[int] = None) -> str:
    """
    Get workouts for a specific user.
    
    Args:
        user_id: ID of the user
        limit: Maximum number of workouts to return
        offset: Offset for pagination
        
    Returns:
        List of user workouts
    """
    logger.info(f"Getting workouts for user: {user_id}")
    try:
        workouts = workouts_api.users_user_id_workouts_get(user_id, limit=limit, offset=offset)
        workout_list = [f"{w.name} (ID: {w.id})" for w in workouts.workouts]
        return f"Workouts for user {user_id}: {', '.join(workout_list)}"
    except Exception as e:
        return f"Error getting user workouts: {str(e)}"

@tool
async def delete_workout(workout_id: str) -> str:
    """
    Delete a workout by ID.
    
    Args:
        workout_id: ID of the workout to delete
        
    Returns:
        Deletion status
    """
    logger.info(f"Deleting workout: {workout_id}")
    try:
        workouts_api.delete_workout(workout_id)
        return f"Successfully deleted workout: {workout_id}"
    except Exception as e:
        return f"Error deleting workout: {str(e)}"

@tool
async def create_exercise(createExerciseRequest: CreateExerciseRequest) -> str:
    """
    Create a new exercise.
    
    Args:
        createExerciseRequest: CreateExerciseRequest object
        
    Returns:
        Exercise creation status and information
    """
    logger.info(f"Creating exercise: {createExerciseRequest.name}")
    try:
        exercise = exercises_api.create_exercise(createExerciseRequest)
        return f"Successfully created exercise: {exercise.name} (ID: {exercise.id})"
    except Exception as e:
        return f"Error creating exercise: {str(e)}"

@tool
async def get_exercise(exercise_id: str) -> str:
    """
    Get an exercise by ID.
    
    Args:
        exercise_id: ID of the exercise to retrieve
        
    Returns:
        Exercise information
    """
    logger.info(f"Getting exercise: {exercise_id}")
    try:
        exercise = exercises_api.get_exercise_by_id(exercise_id)
        return f"Exercise: {exercise.name} (ID: {exercise.id}, Type: {exercise.type})"
    except Exception as e:
        return f"Error getting exercise: {str(e)}"

@tool
async def list_exercises(limit: Optional[int] = None, offset: Optional[int] = None) -> str:
    """
    List exercises with optional pagination.
    
    Args:
        limit: Maximum number of exercises to return
        offset: Offset for pagination
        
    Returns:
        List of exercises
    """
    logger.info("Listing exercises")
    try:
        exercises = exercises_api.list_exercises(limit=limit, offset=offset)
        exercise_list = [f"{e.name} (ID: {e.id}, Type: {e.type})" for e in exercises.exercises]
        return f"Exercises: {', '.join(exercise_list)}"
    except Exception as e:
        return f"Error listing exercises: {str(e)}"

@tool
async def delete_exercise(exercise_id: str) -> str:
    """
    Delete an exercise by ID.
    
    Args:
        exercise_id: ID of the exercise to delete
        
    Returns:
        Deletion status
    """
    logger.info(f"Deleting exercise: {exercise_id}")
    try:
        exercises_api.delete_exercise(exercise_id)
        return f"Successfully deleted exercise: {exercise_id}"
    except Exception as e:
        return f"Error deleting exercise: {str(e)}"

def get_available_tools() -> Dict[str, Any]:
    """
    Get all available tools for the agent.
    
    Returns:
        Dictionary mapping tool names to tool functions
    """
    return {
        # Workout management tools
        "create_workout": create_workout,
        "get_workout": get_workout,
        "get_user_workouts": get_user_workouts,
        "delete_workout": delete_workout,
        # Exercise management tools
        "create_exercise": create_exercise,
        "get_exercise": get_exercise,
        "list_exercises": list_exercises,
        "delete_exercise": delete_exercise,
    } 