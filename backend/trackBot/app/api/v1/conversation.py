from fastapi import APIRouter, Depends, HTTPException, Request
from sqlalchemy.orm import Session
from typing import List
from app.config.database import get_db
import app.models.models as openapiModels
from pydantic import BaseModel
from datetime import datetime
from langchain_core.messages import HumanMessage, BaseMessage
from app.models.trackBot_models import User, Conversation, Message
from app.models.trackBot_models import PersistedAgentState
from app.agent.trackBot_agent import TrackBotAgent
from app.agent.state.state import AgentState
import logging
from middleware.auth import verify_token
from fastapi import Depends

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
@router.post("/users/{user_id}/conversations/{conversation_id}/messages/", response_model=openapiModels.ListMessagesResponse)
async def create_message(user_id: int, conversation_id: int, message: openapiModels.CreateMessageRequest, db: Session = Depends(get_db)):
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

    state: AgentState = AgentState(
        user_id=user_id,
        conversation_id=conversation_id,
        messages=[BaseMessage.model_validate_json(message.langchain_message) for message in conversation.messages],
        tools_called=[],
        pending_input_prompt=None,
        status="started",
        next_action="process_messages"
    )

    human_message = HumanMessage(content=message.langchain_message)
    db_message = Message(user_id=user_id, conversation_id=conversation_id, type="user", langchain_message=human_message.model_dump_json())
    db.add(db_message)
    db.commit()
    db.refresh(db_message)

    state["messages"].append(human_message)

    # Initialize the agent with PostgreSQL checkpointing
    agent = TrackBotAgent.create(user_id=user_id, conversation_id=conversation_id)
    
    # Create initial state as a dictionary matching AgentState type
 

    original_messages_count = len(state["messages"])
    try:
        # Run the agent with a unique thread ID
        result = await agent.run(state)
        
        # Create assistant's response message
        responseMessages = []
        for i in range(original_messages_count, len(result["messages"])):
            message = result["messages"][i]
            db_message_reply = Message(
                user_id=user_id,
                conversation_id=conversation_id,
                type= "assistant" if message.type == "ai" else "other",
                langchain_message=message.model_dump_json()
            )
            db.add(db_message_reply)
            db.commit()
            db.refresh(db_message_reply)

            om = openapiModels.Message(
                id=db_message_reply.id,
                conversation_id=db_message_reply.conversation_id,
                user_id=db_message_reply.user_id,
                langchain_message=db_message_reply.langchain_message,
                message_type=db_message_reply.type,
                created_at=db_message_reply.created_at
            )
            responseMessages.append(om)

        response = openapiModels.ListMessagesResponse(messages=responseMessages)
        return response
        
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