from datetime import datetime, UTC
from sqlalchemy import Column, Integer, String, DateTime, ForeignKey, Text, Boolean, JSON, Enum
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base
from typing import Dict, Any, List, Optional
Base = declarative_base()

class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    email = Column(String(255), unique=True, nullable=False)
    name = Column(String(255))
    picture = Column(String(255))
    created_at = Column(DateTime, default=datetime.now(UTC))
    updated_at = Column(DateTime, default=datetime.now(UTC), onupdate=datetime.now(UTC))

    exercises = relationship("Exercise", back_populates="user")
    conversations = relationship("Conversation", back_populates="user")

class Exercise(Base):
    __tablename__ = 'exercises'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey('users.id'), nullable=False)
    name = Column(String(255), nullable=False)
    description = Column(Text)
    sets = Column(Integer)
    reps = Column(Integer)
    weight = Column(Integer)
    distance = Column(Integer) # in meters
    duration = Column(Integer)  # in seconds
    created_at = Column(DateTime, default=datetime.now(UTC))
    updated_at = Column(DateTime, default=datetime.now(UTC), onupdate=datetime.now(UTC))

    user = relationship("User", back_populates="exercises")

class Conversation(Base):
    __tablename__ = 'conversations'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey('users.id'), nullable=False)
    title = Column(String(255))
    created_at = Column(DateTime, default=datetime.now(UTC))
    updated_at = Column(DateTime, default=datetime.now(UTC), onupdate=datetime.now(UTC))

    user = relationship("User", back_populates="conversations")
    messages = relationship("Message", back_populates="conversation")

class Message(Base):
    __tablename__ = 'messages'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey('users.id'), nullable=False)    
    conversation_id = Column(Integer, ForeignKey('conversations.id'), nullable=False)
    type = Column(Enum('user', 'assistant', "other", name='message_type_enum'), nullable=False)
    langchain_message = Column(Text, nullable=False)
    created_at = Column(DateTime, default=datetime.now(UTC))

    user = relationship("User")
    conversation = relationship("Conversation", back_populates="messages") 


class PersistedAgentState(Base):
    __tablename__ = 'persisted_agent_states'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey('users.id'), nullable=False)
    conversation_id = Column(Integer, ForeignKey('conversations.id'), nullable=False)
    tools_called = Column(JSON, nullable=False)
    state = Column(JSON, nullable=False)
    pending_input_prompt = Column(Text)
    status = Column(Text, nullable=False)
    next_action = Column(Text)
    created_at = Column(DateTime, default=datetime.now(UTC))
    updated_at = Column(DateTime, default=datetime.now(UTC), onupdate=datetime.now(UTC))

    conversation = relationship("Conversation", back_populates="states")