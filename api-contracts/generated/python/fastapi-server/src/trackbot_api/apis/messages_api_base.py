# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from pydantic import Field
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.create_message_request import CreateMessageRequest
from trackbot_api.models.create_message_response import CreateMessageResponse
from trackbot_api.models.error import Error
from trackbot_api.models.message import Message


class BaseMessagesApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseMessagesApi.subclasses = BaseMessagesApi.subclasses + (cls,)
    async def create_message(
        self,
        conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")],
        create_message_request: CreateMessageRequest,
    ) -> CreateMessageResponse:
        """Send a new message in a conversation"""
        ...


    async def delete_message(
        self,
        messageId: Annotated[int, Field(strict=True, ge=1, description="Message ID")],
    ) -> None:
        """Delete a message from a conversation"""
        ...


    async def get_message_by_id(
        self,
        messageId: Annotated[int, Field(strict=True, ge=1, description="Message ID")],
    ) -> Message:
        """Retrieve a specific message by its ID"""
        ...


    async def list_messages(
        self,
        conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")],
    ) -> List[Message]:
        """Retrieve all messages in a specific conversation"""
        ...
