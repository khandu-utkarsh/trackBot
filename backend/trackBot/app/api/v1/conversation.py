from fastapi import APIRouter, Depends, HTTPException, Request
from sqlalchemy.orm import Session
from typing import List
from app.config.database import get_db
import app.models.models as openapiModels
from pydantic import BaseModel
from datetime import datetime
from langchain_core.messages import HumanMessage, BaseMessage, AIMessage, SystemMessage, ToolMessage
from app.models.trackBot_models import User, Conversation, Message
from app.models.trackBot_models import PersistedAgentState
from app.agent.trackBot_agent import TrackBotAgent
from app.agent.state.state import AgentState
import logging
from middleware.auth import verify_token
from fastapi import Depends
from langchain_core.messages import messages_from_dict
import json


logger = logging.getLogger(__name__)

router = APIRouter(dependencies=[Depends(verify_token)])

# Conversation endpoints
@router.post("/users/{user_id}/conversations/", response_model=openapiModels.CreateConversationResponse)
def create_conversation(user_id: int, conversation: openapiModels.CreateConversationRequest, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")
    
    db_conversation = Conversation(user_id=user_id, title=conversation.title)
    db.add(db_conversation)
    db.commit()
    db.refresh(db_conversation)

    #Once conversation is created, we need to create the message, used to create the conversation.



    response = openapiModels.CreateConversationResponse(
        id=db_conversation.id,
        title=db_conversation.title,
        user_id=db_conversation.user_id,
        updated_at=db_conversation.updated_at
    )

    return response

@router.get("/users/{user_id}/conversations/", response_model=openapiModels.ListConversationsResponse)
def get_user_conversations(user_id: int, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")
    
    conversations = []
    for conversation in user.conversations:
        oc = openapiModels.Conversation(
            id=conversation.id,
            title=conversation.title,
            user_id=conversation.user_id,
            updated_at=conversation.updated_at
        )
        conversations.append(oc)

    response = openapiModels.ListConversationsResponse(conversations=conversations)
    return response

@router.get("/users/{user_id}/conversations/{conversation_id}", response_model=openapiModels.Conversation)
def get_conversation(user_id: int, conversation_id: int, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")
    
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()
    
    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")

    #Create the agent state in the database as well.
    agent_state = PersistedAgentState(
        conversation_id=conversation_id,
        user_id=user_id,
        tools_called=[],
        pending_input_prompt=None,
        status="started",
        next_action="process_messages"
    )

    db.add(agent_state)
    db.commit()
    db.refresh(agent_state)



    response = openapiModels.Conversation(
        id=conversation.id,
        title=conversation.title,
        user_id=conversation.user_id,
        updated_at=conversation.updated_at
    )
    return response

@router.put("/users/{user_id}/conversations/{conversation_id}", response_model=openapiModels.Conversation)
def update_conversation(user_id: int, conversation_id: int, request: openapiModels.UpdateConversationRequest, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")
    
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()
    
    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")
    
    for key, value in request.dict(exclude_unset=True).items():
        setattr(conversation, key, value)
    
    db.commit()
    db.refresh(conversation)

    response = openapiModels.Conversation(
        id=conversation.id,
        title=conversation.title,
        user_id=conversation.user_id,
        updated_at=conversation.updated_at
    )
    return response

@router.delete("/users/{user_id}/conversations/{conversation_id}", response_model=openapiModels.DeleteConversationResponse)
def delete_conversation(user_id: int, conversation_id: int, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")
    
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()
    
    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")
    
    messages_count = len(conversation.messages)
    db.delete(conversation)
    db.commit()

    response = openapiModels.DeleteConversationResponse(
        id=conversation.id,
        title=conversation.title,
        deleted_at=conversation.updated_at,
        messages_deleted_count=messages_count
    )
    return response

# Message endpoints
async def _validate_conversation_access(user_id: int, conversation_id: int, db: Session) -> Conversation:
    """Validate user and conversation access."""
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()

    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        raise HTTPException(status_code=404, detail="User not found")

    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")

    if user.id != conversation.user_id: 
        raise HTTPException(status_code=403, detail="User does not have access to this conversation")
    
    return conversation

def _convert_db_messages_to_langchain(conversation: Conversation) -> List[BaseMessage]:
    """Convert database messages to LangChain message format."""
    langChainMessages = []
    for message in conversation.messages:
        msg_dict = json.loads(message.langchain_message)
        msg_type = msg_dict["type"]
        if msg_type == "human":
            msg = HumanMessage(**msg_dict)
            langChainMessages.append(msg)
        elif msg_type == "ai":
            msg = AIMessage(**msg_dict)
            langChainMessages.append(msg)
        elif msg_type == "system":
            msg = SystemMessage(**msg_dict)
            langChainMessages.append(msg)
        elif msg_type == "tool":
            msg = ToolMessage(**msg_dict)
            langChainMessages.append(msg)
        else:
            raise ValueError(f"Unknown message type: {msg_type}")
    return langChainMessages

def _log_message_to_db(user_id: int, conversation_id: int, message: BaseMessage, message_type: str, db: Session) -> Message:
    """Log a message to the database."""
    db_message = Message(
        user_id=user_id,
        conversation_id=conversation_id,
        type=message_type,
        langchain_message=message.model_dump_json()
    )
    db.add(db_message)
    db.commit()
    db.refresh(db_message)
    return db_message

def _convert_to_openapi_message(db_message: Message) -> openapiModels.Message:
    """Convert database message to OpenAPI model."""
    return openapiModels.Message(
        id=db_message.id,
        conversation_id=db_message.conversation_id,
        user_id=db_message.user_id,
        langchain_message=db_message.langchain_message,
        message_type=db_message.type,
        created_at=db_message.created_at
    )

@router.post("/users/{user_id}/conversations/{conversation_id}/messages/", response_model=openapiModels.ListMessagesResponse)
async def create_message(user_id: int, conversation_id: int, message: openapiModels.CreateMessageRequest, db: Session = Depends(get_db)):
    # Step 1: Validate access
    conversation = await _validate_conversation_access(user_id, conversation_id, db)
    
    # Step 2: Convert existing messages to LangChain format
    langChainMessages = _convert_db_messages_to_langchain(conversation)
    
    # Step 3: Create and log the new user message
    human_message = HumanMessage(content=message.langchain_message)
    db_message = _log_message_to_db(user_id, conversation_id, human_message, "user", db)
    
    # Step 4: Create agent state
    state: AgentState = AgentState(
        user_id=user_id,
        conversation_id=conversation_id,
        messages=langChainMessages + [human_message],
        tools_called=[],
        pending_input_prompt=None,
        status="started",
        next_action="process_messages"
    )
    
    # Step 5: Initialize and run the agent
    agent = TrackBotAgent.create(user_id=user_id, conversation_id=conversation_id)
    original_messages_count = len(state["messages"])
    
    try:
        # Run the agent
        result = await agent.run(state)
        
        # Step 6: Log agent responses to DB and prepare response

        for i in range(original_messages_count, len(result["messages"])):
            message = result["messages"][i]
            
            # Determine message type
            message_type = "other"
            if message.type == "ai":
                message_type = "assistant"
            elif message.type == "human":
                message_type = "user"
            
            # Log to DB and convert to OpenAPI model
            db_message = _log_message_to_db(user_id, conversation_id, message, message_type, db)

        # Step 7: Preparing the responses
        responseMessages = []
        for message in conversation.messages:
            responseMessages.append(_convert_to_openapi_message(message))
        
        responseMessages.sort(key=lambda x: x.created_at, reverse=True)

        return openapiModels.ListMessagesResponse(messages=responseMessages)
        
    except Exception as e:
        logger.error(f"Error in agent execution: {e}")
        raise HTTPException(status_code=500, detail=str(e))
 
@router.get("/users/{user_id}/conversations/{conversation_id}/messages/", response_model=openapiModels.ListMessagesResponse)
def get_conversation_messages(user_id: int, conversation_id: int, limit: int = 50, offset: int = 0, db: Session = Depends(get_db)):
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()
    
    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")
    
    db_messages = db.query(Message).filter(
        Message.conversation_id == conversation_id
    ).offset(offset).limit(limit).all()
    
    messages = []
    for message in db_messages:
        om = openapiModels.Message(
            id=message.id,
            conversation_id=message.conversation_id,
            user_id=message.user_id,
            langchain_message=message.langchain_message,
            message_type=message.type,
            created_at=message.created_at
        )
        messages.append(om)

    response = openapiModels.ListMessagesResponse(messages=messages)
    return response

@router.get("/users/{user_id}/conversations/{conversation_id}/messages/{message_id}", response_model=openapiModels.Message)
def get_message(user_id: int, conversation_id: int, message_id: int, db: Session = Depends(get_db)):
    conversation = db.query(Conversation).filter(
        Conversation.id == conversation_id,
        Conversation.user_id == user_id
    ).first()
    
    if conversation is None:
        raise HTTPException(status_code=404, detail="Conversation not found")
    
    message = db.query(Message).filter(
        Message.id == message_id,
        Message.conversation_id == conversation_id
    ).first()
    
    if message is None:
        raise HTTPException(status_code=404, detail="Message not found")
    
    response = openapiModels.Message(
        id=message.id,
        conversation_id=message.conversation_id,
        user_id=message.user_id,
        langchain_message=message.langchain_message,
        message_type=message.type,
        created_at=message.created_at
    )
    return response