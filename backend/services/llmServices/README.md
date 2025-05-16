# Workout LLM Service

This service uses LangChain and OpenAI to process natural language requests for workout-related actions and forwards them to the main workout application service.

## Setup

1. Create a virtual environment:
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

3. Create a `.env` file with the following variables:
```env
OPENAI_API_KEY=your_openai_api_key
MODEL_NAME=gpt-3.5-turbo
TEMPERATURE=0.7
MAX_TOKENS=1000
WORKOUT_APP_URL=http://localhost:8080
HOST=0.0.0.0
PORT=8081
```

## Running the Service

Start the service:
```bash
python app/main.py
```

The service will be available at `http://localhost:8081`

## API Usage

Send a POST request to `/api/v1/process` with the following JSON body:
```json
{
    "input": "Create a new workout for tomorrow with 30 minutes of cardio"
}
```

The service will:
1. Process the natural language input using LangChain and OpenAI
2. Determine the appropriate action and parameters
3. Forward the request to the main workout application service
4. Return the response from the workout application service

## Available Actions

The LLM service can handle the following types of requests:
- Creating new workouts
- Getting workout history
- Updating existing workouts
- Deleting workouts
- Getting workout statistics 