# Microservices Integration: Workout App + LLM Services

This document explains how the Workout App Service (Go) integrates with the LLM Service (Python/FastAPI) to provide AI-powered fitness coaching through chat conversations.

## Architecture Overview

```
Frontend (React) 
    ↓ HTTP Requests
Workout App Service (Go:8080)
    ↓ HTTP Requests  
LLM Service (Python/FastAPI:8081)
    ↓ API Calls
OpenAI/LangChain/LangGraph
```

## Services

### 1. Workout App Service (Go)
- **Port**: 8080
- **Database**: PostgreSQL
- **Responsibilities**:
  - User management
  - Workout and exercise tracking
  - Conversation and message storage
  - Authentication and authorization
  - HTTP client for LLM service communication

### 2. LLM Service (Python/FastAPI)
- **Port**: 8081
- **Dependencies**: LangChain, LangGraph, OpenAI
- **Responsibilities**:
  - AI chat processing
  - Fitness coaching responses
  - Workout parsing and analysis
  - LangGraph agentic workflows

## Integration Flow

### 1. User Creates a Message
```
POST /api/users/{userID}/conversations/{conversationID}/messages
{
  "content": "I did 3 sets of 10 push-ups today",
  "message_type": "user"
}
```

### 2. Workout App Service Processing
1. **Validates** user and conversation ownership
2. **Stores** the user message in PostgreSQL
3. **Triggers** AI response processing (async)
4. **Returns** the created message immediately

### 3. AI Response Processing (Async)
1. **Retrieves** conversation history from database
2. **Calls** LLM Service with conversation context
3. **Receives** AI-generated response
4. **Stores** assistant message in database

### 4. LLM Service Processing
```
POST /api/v1/chat/message
{
  "messages": [
    {"role": "user", "content": "I did 3 sets of 10 push-ups today"}
  ],
  "user_id": "123",
  "conversation_id": "456",
  "context": {}
}
```

1. **Analyzes** message type (chat vs workout logging)
2. **Processes** with LangChain/LangGraph
3. **Returns** structured response

## Key Components

### LLM Client (Go)
```go
type LLMClient struct {
    baseURL    string
    httpClient *http.Client
}

func (c *LLMClient) ProcessChatMessage(
    ctx context.Context, 
    messages []models.Message, 
    userID, conversationID int64, 
    context map[string]interface{}
) (*ChatResponse, error)
```

### Message Handler Integration
```go
// If this is a user message, trigger AI response
if message.MessageType == models.MessageTypeUser && h.llmClient != nil {
    go h.processAIResponse(ctx, userID, conversationID)
}
```

### LLM Service Chat Endpoint
```python
@router.post("/message", response_model=ChatResponse)
async def process_chat_message(request: ChatRequest):
    response = await llm_service.process_chat_message(
        user_message=user_message,
        conversation_history=request.messages,
        user_id=request.user_id,
        conversation_id=request.conversation_id,
        context=request.context or {}
    )
```

## Environment Variables

### Workout App Service
```env
LLM_SERVICE_URL=http://llm-service:8081
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=workout_app
```

### LLM Service
```env
OPENAI_API_KEY=your_api_key
MODEL_PROVIDER=openai
MODEL_NAME=gpt-4o-mini
OPENAI_MODEL_TEMPERATURE=0.7
OPENAI_MODEL_MAX_TOKENS=1000
WORKOUT_APP_URL=http://workout-app:8080
```

## Running the Services

### Using Docker Compose
```bash
# Set your OpenAI API key
export OPENAI_API_KEY=your_api_key_here

# Start all services
docker-compose up --build
```

### Manual Setup
```bash
# Terminal 1: Start PostgreSQL
docker run -d -p 5432:5432 -e POSTGRES_DB=workout_app -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password postgres:15

# Terminal 2: Start LLM Service
cd backend/services/llmServices
pip install -r requirements.txt
python app/main.py

# Terminal 3: Start Workout App Service
cd backend/services/workoutAppServices
go run cmd/api/*.go
```

## API Endpoints

### Workout App Service
- `GET /api/users/{userID}/conversations` - List conversations
- `POST /api/users/{userID}/conversations` - Create conversation
- `GET /api/users/{userID}/conversations/{conversationID}/messages` - List messages
- `POST /api/users/{userID}/conversations/{conversationID}/messages` - Create message (triggers AI)

### LLM Service
- `POST /api/v1/chat/message` - Process chat message
- `GET /api/v1/chat/health` - Health check

## Features

### AI Fitness Coach
- **Personalized responses** based on conversation history
- **Workout logging** detection and parsing
- **Exercise guidance** and form tips
- **Nutrition advice** and meal planning
- **Progress tracking** and motivation

### LangGraph Integration
- **Agentic workflows** for complex fitness planning
- **Multi-step reasoning** for workout recommendations
- **Context-aware** responses based on user history
- **Structured data extraction** from natural language

## Error Handling

### Graceful Degradation
- If LLM service is unavailable, user messages are still stored
- AI responses are processed asynchronously to avoid blocking
- Health checks ensure service availability

### Retry Logic
- HTTP client includes timeout and retry mechanisms
- Circuit breaker pattern can be implemented for resilience

## Security

### Authentication
- User ownership verification for all operations
- Conversation access control
- Message-level security checks

### Data Privacy
- User data is isolated by user ID
- Conversation history is limited for LLM processing
- No sensitive data is logged

## Future Enhancements

### Real-time Updates
- WebSocket integration for live AI responses
- Push notifications for workout reminders

### Advanced AI Features
- Multi-modal input (images, voice)
- Personalized workout plan generation
- Integration with fitness trackers

### Monitoring
- Service health monitoring
- Performance metrics
- Error tracking and alerting 