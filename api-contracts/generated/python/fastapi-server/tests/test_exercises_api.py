# coding: utf-8

from fastapi.testclient import TestClient


from pydantic import Field  # noqa: F401
from typing import Any, List  # noqa: F401
from typing_extensions import Annotated  # noqa: F401
from trackbot_api.models.create_exercise_request import CreateExerciseRequest  # noqa: F401
from trackbot_api.models.create_exercise_response import CreateExerciseResponse  # noqa: F401
from trackbot_api.models.error import Error  # noqa: F401
from trackbot_api.models.exercise import Exercise  # noqa: F401
from trackbot_api.models.update_exercise_request import UpdateExerciseRequest  # noqa: F401


def test_create_exercise(client: TestClient):
    """Test case for create_exercise

    Create a new exercise
    """
    create_exercise_request = {"duration":1800,"notes":"Morning run in the park","distance":5000.0,"name":"Running","type":"cardio"}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "POST",
    #    "/workouts/{workoutId}/exercises".format(workoutId=1),
    #    headers=headers,
    #    json=create_exercise_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_delete_exercise(client: TestClient):
    """Test case for delete_exercise

    Delete exercise
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "DELETE",
    #    "/exercises/{exerciseId}".format(exerciseId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_exercise_by_id(client: TestClient):
    """Test case for get_exercise_by_id

    Get exercise by ID
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/exercises/{exerciseId}".format(exerciseId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_list_exercises(client: TestClient):
    """Test case for list_exercises

    List exercises for a workout
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/workouts/{workoutId}/exercises".format(workoutId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_update_exercise(client: TestClient):
    """Test case for update_exercise

    Update exercise
    """
    update_exercise_request = {"duration":1800,"notes":"Morning run in the park","distance":5000.0,"name":"Running","type":"cardio"}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "PUT",
    #    "/exercises/{exerciseId}".format(exerciseId=1),
    #    headers=headers,
    #    json=update_exercise_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200

