from pydantic import BaseModel, Field
from typing import Dict, Any, Literal, Union, List, Optional
from datetime import datetime
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage, BaseMessage


# Models need to process the conversation
class TrackBotHumanMessage(BaseModel):
    role: Literal["user"]
    content: str

    def to_langchain(self) -> BaseMessage:
        return HumanMessage(content=self.content)

class TrackBotAIMessage(BaseModel):
    role: Literal["assistant"]
    content: str

    def to_langchain(self) -> BaseMessage:
        return AIMessage(content=self.content)

Message = Union[TrackBotHumanMessage, TrackBotAIMessage]

class ProcessMessagesRequest(BaseModel):
    messages: List[Message]
    user_id: int


# Only contains the generated output of one call.
class ProcessMessagesResponse(BaseModel):
    message: BaseMessage