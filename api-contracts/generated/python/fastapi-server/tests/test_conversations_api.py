# coding: utf-8

from fastapi.testclient import TestClient


from pydantic import Field  # noqa: F401
from typing import Any, List  # noqa: F401
from typing_extensions import Annotated  # noqa: F401
from trackbot_api.models.conversation import Conversation  # noqa: F401
from trackbot_api.models.create_conversation_request import CreateConversationRequest  # noqa: F401
from trackbot_api.models.create_conversation_response import CreateConversationResponse  # noqa: F401
from trackbot_api.models.error import Error  # noqa: F401


def test_create_conversation(client: TestClient):
    """Test case for create_conversation

    Create a new conversation
    """
    create_conversation_request = {"title":"Workout Planning Session"}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "POST",
    #    "/users/{userId}/conversations".format(userId=1),
    #    headers=headers,
    #    json=create_conversation_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_delete_conversation(client: TestClient):
    """Test case for delete_conversation

    Delete conversation
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "DELETE",
    #    "/conversations/{conversationId}".format(conversationId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_conversation_by_id(client: TestClient):
    """Test case for get_conversation_by_id

    Get conversation by ID
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/conversations/{conversationId}".format(conversationId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_list_conversations(client: TestClient):
    """Test case for list_conversations

    List conversations for a user
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/users/{userId}/conversations".format(userId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200

