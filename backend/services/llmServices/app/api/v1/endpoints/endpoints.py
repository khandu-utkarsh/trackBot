from fastapi import APIRouter, HTTPException
from internal.generated.trackbot_client.api import MessagesApi, ConversationsApi, ExercisesApi, WorkoutsApi, UsersApi
from internal.generated.trackbot_client.models import CreateMessageRequest, CreateConversationRequest, CreateExerciseRequest, CreateWorkoutRequest, CreateUserRequest
from internal.generated.trackbot_client.models import ListMessagesRequest, ListConversationsRequest, ListExercisesRequest, ListWorkoutsRequest, ListUsersRequest
from internal.generated.trackbot_client.models import ListMessagesResponse, ListConversationsResponse, ListExercisesResponse, ListWorkoutsResponse, ListUsersResponse
from internal.generated.trackbot_client.models import Message, Conversation, Exercise, Workout, User
from internal.generated.trackbot_client.models import UpdateMessageRequest, UpdateConversationRequest, UpdateExerciseRequest, UpdateWorkoutRequest, UpdateUserRequest
from internal.generated.trackbot_client.models import DeleteMessageRequest, DeleteConversationRequest, DeleteExerciseRequest, DeleteWorkoutRequest, DeleteUserRequest
from internal.generated.trackbot_client.models import DeleteMessageResponse, DeleteConversationResponse, DeleteExerciseResponse, DeleteWorkoutResponse, DeleteUserResponse
from models import ChatRequest, ChatResponse, AIServiceMessageRequest, AIServiceMessageResponse

from app.services.agent_service import AgentService
import logging

router = APIRouter()
agent_service = AgentService()
logger = logging.getLogger(__name__)



@router.post("/process_conversation", response_model=AIServiceMessageResponse)
async def process_messages_handler(request: ChatRequest):
    """Process messages through the LangGraph agent workflow."""
    logger.info(f"Received request for user {request.user_id}")
    
    try:
        # Convert request messages to LangChain format
        langchain_messages = [m.to_langchain() for m in request.messages]
        
        # Process through agent service
        result = await agent_service.process_messages(
            messages=langchain_messages,
            user_id=request.user_id,
            session_id=request.session_id
        )
        
        # Convert response message back to our format
        response_message = Message.from_langchain(result["messages"][-1])
        
        return ProcessMessagesResponse(
            message=response_message,
            session_id=result["session_id"],
            tools_called=result["tools_called"],
            needs_user_input=result["needs_user_input"],
            pending_input_prompt=result["pending_input_prompt"],
            status=result["status"]
        )
        
    except Exception as e:
        logger.error(f"Error processing request: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))

@router.post("/continue_conversation", response_model=ProcessMessagesResponse)
async def continue_session_handler(request: ContinueSessionRequest):
    """Continue a session after user input."""
    logger.info(f"Continuing session {request.session_id}")
    
    try:
        result = await agent_service.continue_session(
            session_id=request.session_id,
            user_response=request.user_response
        )
        
        # Convert response message back to our format
        response_message = Message.from_langchain(result["messages"][-1])
        
        return ProcessMessagesResponse(
            message=response_message,
            session_id=result["session_id"],
            tools_called=result["tools_called"],
            needs_user_input=result["needs_user_input"],
            pending_input_prompt=result["pending_input_prompt"],
            status=result["status"]
        )
        
    except ValueError as e:
        raise HTTPException(status_code=404, detail=str(e))
    except Exception as e:
        logger.error(f"Error continuing session: {str(e)}")
        raise HTTPException(status_code=500, detail=str(e))