# coding: utf-8

from fastapi.testclient import TestClient


from pydantic import Field, field_validator  # noqa: F401
from typing import Any, List, Optional  # noqa: F401
from typing_extensions import Annotated  # noqa: F401
from trackbot_api.models.create_workout_request import CreateWorkoutRequest  # noqa: F401
from trackbot_api.models.create_workout_response import CreateWorkoutResponse  # noqa: F401
from trackbot_api.models.error import Error  # noqa: F401
from trackbot_api.models.update_workout_request import UpdateWorkoutRequest  # noqa: F401
from trackbot_api.models.workout import Workout  # noqa: F401


def test_create_workout(client: TestClient):
    """Test case for create_workout

    Create a new workout
    """
    create_workout_request = {"user_id":1}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "POST",
    #    "/users/{userId}/workouts".format(userId=1),
    #    headers=headers,
    #    json=create_workout_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_delete_workout(client: TestClient):
    """Test case for delete_workout

    Delete workout
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "DELETE",
    #    "/workouts/{workoutId}".format(workoutId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_workout_by_id(client: TestClient):
    """Test case for get_workout_by_id

    Get workout by ID
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/workouts/{workoutId}".format(workoutId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_list_workouts(client: TestClient):
    """Test case for list_workouts

    List workouts for a user
    """
    params = [("year", '2024'),     ("month", '01'),     ("day", '15')]
    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/users/{userId}/workouts".format(userId=1),
    #    headers=headers,
    #    params=params,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_update_workout(client: TestClient):
    """Test case for update_workout

    Update workout
    """
    update_workout_request = {"user_id":1}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "PUT",
    #    "/workouts/{workoutId}".format(workoutId=1),
    #    headers=headers,
    #    json=update_workout_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200

