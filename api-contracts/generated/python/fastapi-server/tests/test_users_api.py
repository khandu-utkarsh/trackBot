# coding: utf-8

from fastapi.testclient import TestClient


from pydantic import Field, StrictStr  # noqa: F401
from typing import Any, List  # noqa: F401
from typing_extensions import Annotated  # noqa: F401
from trackbot_api.models.create_user_request import CreateUserRequest  # noqa: F401
from trackbot_api.models.create_user_response import CreateUserResponse  # noqa: F401
from trackbot_api.models.error import Error  # noqa: F401
from trackbot_api.models.user import User  # noqa: F401


def test_create_user(client: TestClient):
    """Test case for create_user

    Create a new user
    """
    create_user_request = {"email":"user@gmail.com"}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "POST",
    #    "/users",
    #    headers=headers,
    #    json=create_user_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_delete_user(client: TestClient):
    """Test case for delete_user

    Delete user
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "DELETE",
    #    "/users/{userId}".format(userId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_user_by_email(client: TestClient):
    """Test case for get_user_by_email

    Get user by email
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/users/email/{email}".format(email='user@example.com'),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_user_by_id(client: TestClient):
    """Test case for get_user_by_id

    Get user by ID
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/users/{userId}".format(userId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_list_users(client: TestClient):
    """Test case for list_users

    List all users
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/users",
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200

