# coding: utf-8

from fastapi.testclient import TestClient


from pydantic import Field  # noqa: F401
from typing import Any, List  # noqa: F401
from typing_extensions import Annotated  # noqa: F401
from trackbot_api.models.create_message_request import CreateMessageRequest  # noqa: F401
from trackbot_api.models.create_message_response import CreateMessageResponse  # noqa: F401
from trackbot_api.models.error import Error  # noqa: F401
from trackbot_api.models.message import Message  # noqa: F401


def test_create_message(client: TestClient):
    """Test case for create_message

    Create a new message
    """
    create_message_request = {"message_type":"user","content":"I want to start a new workout plan"}

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "POST",
    #    "/conversations/{conversationId}/messages".format(conversationId=1),
    #    headers=headers,
    #    json=create_message_request,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_delete_message(client: TestClient):
    """Test case for delete_message

    Delete message
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "DELETE",
    #    "/messages/{messageId}".format(messageId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_get_message_by_id(client: TestClient):
    """Test case for get_message_by_id

    Get message by ID
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/messages/{messageId}".format(messageId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200


def test_list_messages(client: TestClient):
    """Test case for list_messages

    List messages in a conversation
    """

    headers = {
    }
    # uncomment below to make a request
    #response = client.request(
    #    "GET",
    #    "/conversations/{conversationId}/messages".format(conversationId=1),
    #    headers=headers,
    #)

    # uncomment below to assert the status code of the HTTP response
    #assert response.status_code == 200

