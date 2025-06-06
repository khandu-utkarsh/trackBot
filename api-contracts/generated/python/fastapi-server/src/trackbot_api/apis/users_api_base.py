# coding: utf-8

from typing import ClassVar, Dict, List, Tuple  # noqa: F401

from pydantic import Field, StrictStr
from typing import Any, List
from typing_extensions import Annotated
from trackbot_api.models.create_user_request import CreateUserRequest
from trackbot_api.models.create_user_response import CreateUserResponse
from trackbot_api.models.error import Error
from trackbot_api.models.user import User


class BaseUsersApi:
    subclasses: ClassVar[Tuple] = ()

    def __init_subclass__(cls, **kwargs):
        super().__init_subclass__(**kwargs)
        BaseUsersApi.subclasses = BaseUsersApi.subclasses + (cls,)
    async def create_user(
        self,
        create_user_request: CreateUserRequest,
    ) -> User:
        """Create a new user account"""
        ...


    async def delete_user(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
    ) -> None:
        """Delete a user and all associated data"""
        ...


    async def get_user_by_email(
        self,
        email: Annotated[StrictStr, Field(description="User email address")],
    ) -> User:
        """Retrieve a user by their email address"""
        ...


    async def get_user_by_id(
        self,
        userId: Annotated[int, Field(strict=True, ge=1, description="User ID")],
    ) -> User:
        """Retrieve a specific user by their ID"""
        ...


    async def list_users(
        self,
    ) -> List[User]:
        """Retrieve a list of all users in the system"""
        ...
