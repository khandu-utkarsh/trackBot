from pydantic import BaseModel, Field
from typing import Literal, List, Dict, Any, Optional
from langchain_core.messages import HumanMessage, AIMessage, SystemMessage, ToolMessage, BaseMessage

class Message(BaseModel):
    """Base message model that can convert to/from LangChain messages."""
    role: Literal["human", "ai", "system", "tool"] = Field(description="Role of the message sender")
    content: str = Field(description="Content of the message")
    tool_calls: Optional[List[Dict[str, Any]]] = Field(default=None, description="Tool calls if this is an AI message with tools")
    tool_call_id: Optional[str] = Field(default=None, description="Tool call ID if this is a tool message")
    
    def to_langchain(self) -> BaseMessage:
        """Convert to LangChain message format."""
        if self.role == "human":
            return HumanMessage(content=self.content)
        elif self.role == "ai":
            msg = AIMessage(content=self.content)
            if self.tool_calls:
                msg.tool_calls = self.tool_calls
            return msg
        elif self.role == "system":
            return SystemMessage(content=self.content)
        elif self.role == "tool":
            return ToolMessage(content=self.content, tool_call_id=self.tool_call_id or "")
        else:
            raise ValueError(f"Unknown message role: {self.role}")
    
    @classmethod
    def from_langchain(cls, message: BaseMessage) -> "Message":
        """Create from LangChain message."""
        if isinstance(message, HumanMessage):
            return cls(role="human", content=message.content)
        elif isinstance(message, AIMessage):
            tool_calls = getattr(message, 'tool_calls', None)
            return cls(role="ai", content=message.content, tool_calls=tool_calls)
        elif isinstance(message, SystemMessage):
            return cls(role="system", content=message.content)
        elif isinstance(message, ToolMessage):
            return cls(role="tool", content=message.content, tool_call_id=message.tool_call_id)
        else:
            raise ValueError(f"Unknown message type: {type(message)}")

class MessageHistory(BaseModel):
    """Collection of messages with utility methods."""
    messages: List[Message] = Field(default_factory=list, description="List of messages")
    
    def to_langchain(self) -> List[BaseMessage]:
        """Convert all messages to LangChain format."""
        return [msg.to_langchain() for msg in self.messages]
    
    @classmethod
    def from_langchain(cls, messages: List[BaseMessage]) -> "MessageHistory":
        """Create from list of LangChain messages."""
        return cls(messages=[Message.from_langchain(msg) for msg in messages])
    
    def add_message(self, message: Message) -> None:
        """Add a message to the history."""
        self.messages.append(message)
    
    def get_last_message(self) -> Optional[Message]:
        """Get the last message in the history."""
        return self.messages[-1] if self.messages else None 