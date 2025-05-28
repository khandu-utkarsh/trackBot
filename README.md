# WorkoutWebApp

A full-stack fitness application with AI-powered coaching, built with React (Next.js), Go, Python (FastAPI), and PostgreSQL.

## Architecture

```
Frontend (Next.js) ←→ Backend (Go) ←→ LLM Service (Python/FastAPI) ←→ OpenAI
                           ↓
                    PostgreSQL Database
```

## Features

- 🏋️ **Workout Tracking**: Log and manage your workouts
- 🤖 **AI Fitness Coach**: Get personalized advice and workout plans
- 💬 **Chat Interface**: Natural language interaction with AI
- 📊 **Progress Tracking**: Monitor your fitness journey
- 🔐 **User Authentication**: Secure user management

## Quick Start

### Development Setup (Recommended for debugging)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd WorkoutWebApp
   ```

2. **Set up environment variables**
   ```bash
   # Create .env file in the root directory
   echo "OPENAI_API_KEY=your_openai_api_key_here" > .env
   ```

3. **Start development environment**
   ```bash
   # This uses docker-compose.yml (development by default)
   docker-compose up --build
   ```

4. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - LLM Service: http://localhost:8081
   - Database: localhost:5432

### Production Setup

```bash
# Use the production configuration
docker-compose -f docker-compose.prod.yml up --build
```

## Development Features

### Hot Reloading
- **Frontend**: Next.js development server with hot reload
- **Backend**: Air for Go live reloading
- **LLM Service**: Uvicorn with auto-reload

### Debugging
- **Go Backend**: Delve debugger on port 2345
- **Python LLM Service**: debugpy on port 5678 (set ENABLE_DEBUGGER=true)
- **Frontend**: Standard Next.js debugging

### Volume Mounts
All source code is mounted as volumes for instant code changes without rebuilds.

## API Endpoints

### Backend (Go) - Port 8080
- `GET /health` - Health check
- `GET /api/users` - List users
- `POST /api/users` - Create user
- `GET /api/users/{id}/conversations` - Get user conversations
- `POST /api/users/{id}/conversations` - Create conversation
- `GET /api/users/{id}/conversations/{id}/messages` - Get messages
- `POST /api/users/{id}/conversations/{id}/messages` - Send message

### LLM Service (Python) - Port 8081
- `POST /api/v1/chat/message` - Process chat message
- `GET /api/v1/chat/health` - Health check

## Database Schema

### Users
- `id` (SERIAL PRIMARY KEY)
- `email` (VARCHAR UNIQUE)
- `created_at`, `updated_at` (TIMESTAMP)

### Conversations
- `id` (SERIAL PRIMARY KEY)
- `user_id` (FOREIGN KEY)
- `title` (VARCHAR)
- `is_active` (BOOLEAN)
- `created_at`, `updated_at` (TIMESTAMP)

### Messages
- `id` (SERIAL PRIMARY KEY)
- `conversation_id` (FOREIGN KEY)
- `user_id` (FOREIGN KEY)
- `content` (TEXT)
- `message_type` (ENUM: user, assistant, system)
- `created_at`, `updated_at` (TIMESTAMP)

## Development Workflow

### Making Changes

1. **Frontend Changes**: Edit files in `frontend/src/` - changes reflect immediately
2. **Backend Changes**: Edit files in `backend/services/workoutAppServices/` - Air will rebuild automatically
3. **LLM Service Changes**: Edit files in `backend/services/llmServices/` - Uvicorn will reload automatically

### Debugging

#### Go Backend (Delve)
```bash
# Connect to the debugger
dlv connect localhost:2345
```

#### Python LLM Service
```bash
# Enable debugger in docker-compose.yml
ENABLE_DEBUGGER=true

# Connect with your IDE to localhost:5678
```

### Database Access
```bash
# Connect to PostgreSQL
docker exec -it workout-db-dev psql -U postgres -d workout_app_dev
```

### Logs
```bash
# View logs for specific service
docker-compose logs -f frontend
docker-compose logs -f workout-app
docker-compose logs -f llm-service
docker-compose logs -f postgres
```

## Environment Variables

### Development (.env)
```bash
OPENAI_API_KEY=your_openai_api_key_here
```

### Frontend
- `REACT_APP_API_URL` - Backend API URL
- `REACT_APP_WS_URL` - WebSocket URL
- `NODE_ENV` - Environment mode

### Backend (Go)
- `PORT` - Server port
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` - Database config
- `LLM_SERVICE_URL` - LLM service URL
- `GO_ENV` - Environment mode

### LLM Service (Python)
- `MODEL_PROVIDER` - AI provider (openai)
- `MODEL_NAME` - Model name (gpt-4o-mini)
- `OPENAI_API_KEY` - OpenAI API key
- `PYTHON_ENV` - Environment mode

## Troubleshooting

### Common Issues

1. **Port conflicts**: Make sure ports 3000, 8080, 8081, 5432 are available
2. **OpenAI API Key**: Ensure your API key is valid and has sufficient credits
3. **Database connection**: Wait for PostgreSQL to be ready (health check)

### Reset Development Environment
```bash
# Stop and remove containers, volumes, and images
docker-compose down -v --rmi all

# Rebuild everything
docker-compose up --build
```

### View Container Status
```bash
docker-compose ps
```

### Access Container Shell
```bash
docker exec -it workout-frontend-dev sh
docker exec -it workout-app-dev sh
docker exec -it llm-service-dev bash
```

## Production Deployment

For production deployment, use `docker-compose.prod.yml` which:
- Uses optimized production builds
- Removes development tools and debuggers
- Uses production-ready configurations
- Implements proper health checks

```bash
# Production deployment
export OPENAI_API_KEY=your_production_api_key
docker-compose -f docker-compose.prod.yml up -d
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes using the development setup
4. Test thoroughly
5. Submit a pull request

## License

[Your License Here]