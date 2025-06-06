# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from pydantic import Field
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.conversation import Conversation
from trackbot_api.models.create_conversation_request import CreateConversationRequest
from trackbot_api.models.create_conversation_response import CreateConversationResponse
from trackbot_api.models.error import Error


class BaseConversationsApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseConversationsApi.subclasses = BaseConversationsApi.subclasses + (cls,)
    async def create_conversation(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
        create_conversation_request: CreateConversationRequest,
    ) -> CreateConversationResponse:
        """Start a new conversation with the AI assistant"""
        ...


    async def delete_conversation(
        self,
        conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")],
    ) -> None:
        """Delete a conversation and all associated messages"""
        ...


    async def get_conversation_by_id(
        self,
        conversationId: Annotated[int, Field(strict=True, ge=1, description="Conversation ID")],
    ) -> Conversation:
        """Retrieve a specific conversation by its ID"""
        ...


    async def list_conversations(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
    ) -> List[Conversation]:
        """Retrieve all conversations for a specific user"""
        ...
