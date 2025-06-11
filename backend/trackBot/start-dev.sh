#!/bin/bash
echo "Starting LLM service in development mode..."
if [ "$ENABLE_DEBUGGER" = "true" ]; then
    echo "Starting with debugger on port 5678..."
    python -m debugpy --listen 0.0.0.0:5678 --wait-for-client -m uvicorn app.main:app --host 0.0.0.0 --port 8081 --reload
else
    echo "Starting with hot reload..."
    uvicorn app.main:app --host 0.0.0.0 --port 8081 --reload
fi 