#!/bin/bash

# Development environment management script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

# Define the path to the .env files
LLM_ENV_FILE="$ROOT_DIR/backend/trackBot/.env.trackBot.docker"
DATABASE_ENV_FILE="$ROOT_DIR/backend/.env.database.docker"
FRONTEND_ENV_FILE="$ROOT_DIR/frontend/.env.frontend.docker"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if required .env files exist
check_env() {
    local missing_files=()
    local env_files=(
        "$LLM_ENV_FILE"
        "$FRONTEND_ENV_FILE"
        "$DATABASE_ENV_FILE"
    )
    
    # Check each required env file
    for env_file in "${env_files[@]}"; do
        if [ ! -f "$env_file" ]; then
            missing_files+=("$env_file")
        fi
    done
    
    #If any files are missing, simply report the error and return 1
    if [ ${#missing_files[@]} -ne 0 ]; then
        print_error "Missing environment files. Please create the files and try again."
        print_status "Missing files:"
        for missing_file in "${missing_files[@]}"; do
            echo "  - $missing_file"
        done
        return 1
    fi

    # Check if OpenAI API key is properly set
    local llm_env_file="$LLM_ENV_FILE"
    if grep -q "OPENAI_API_KEY=your_openai_api_key_here" "$llm_env_file" 2>/dev/null; then
        print_error "Please set a valid OPENAI_API_KEY in $llm_env_file"
        return 1
    fi
    
    return 0
}

# Start development environment
start_dev() {
    print_status "Starting development environment..."
    
    if ! check_env; then
        print_error "Please configure environment files first"
        exit 1
    fi
    
    print_status "Building and starting containers..."
    docker-compose up --build
}

# Stop development environment
stop_dev() {
    print_status "Stopping development environment..."
    docker-compose down
    print_success "Development environment stopped"
}

# Restart development environment
restart_dev() {
    print_status "Restarting development environment..."
    docker-compose down
    docker-compose up --build
}

# Clean development environment
clean_dev() {
    print_warning "This will remove all containers, volumes, and images"
    read -p "Are you sure? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_status "Cleaning development environment..."
        docker-compose down -v --rmi all
        docker system prune -f
        print_success "Development environment cleaned"
    else
        print_status "Clean cancelled"
    fi
}

# Show logs
logs() {
    service=${1:-""}
    if [ -z "$service" ]; then
        print_status "Showing logs for all services..."
        docker-compose logs -f
    else
        print_status "Showing logs for $service..."
        docker-compose logs -f "$service"
    fi
}

# Show status
status() {
    print_status "Container status:"
    docker-compose ps
    echo
    print_status "Service URLs:"
    echo "  Frontend:    http://localhost:3000"
    echo "  Backend:     http://localhost:8080"
    echo "  Database:    localhost:5432"
    echo
    print_status "Debug ports:"
    echo "  Python:      localhost:5678"
}

# Database shell
db_shell() {
    print_status "Connecting to database..."
    docker exec -it trackbot-db-dev psql -U postgres -d trackBot_app_dev
}

# Container shell
shell() {
    service=${1:-"backend"}
    print_status "Opening shell for $service..."
    case $service in
        "frontend")
            docker exec -it workout-frontend-dev sh
            ;;
        "backend")
            docker exec -it trackbot-dev sh
            ;;
        "db"|"postgres")
            docker exec -it trackbot-db-dev bash
            ;;
        *)
            print_error "Unknown service: $service"
            print_status "Available services: frontend, backend, db"
            exit 1
            ;;
    esac
}

# Show help
show_help() {
    echo "Development Environment Management Script"
    echo
    echo "Usage: $0 [COMMAND]"
    echo
    echo "Commands:"
    echo "  start     Start development environment"
    echo "  stop      Stop development environment"
    echo "  restart   Restart development environment"
    echo "  clean     Clean development environment (removes all data)"
    echo "  logs      Show logs for all services or specific service"
    echo "  status    Show container status and service URLs"
    echo "  db        Open database shell"
    echo "  shell     Open shell for service (frontend|backend|db)"
    echo "  help      Show this help message"
    echo
    echo "Examples:"
    echo "  $0 start"
    echo "  $0 logs frontend"
    echo "  $0 shell backend"
}

# Main script logic
case "${1:-help}" in
    "start")
        start_dev
        ;;
    "stop")
        stop_dev
        ;;
    "restart")
        restart_dev
        ;;
    "clean")
        clean_dev
        ;;
    "logs")
        logs "$2"
        ;;
    "status")
        status
        ;;
    "db")
        db_shell
        ;;
    "shell")
        shell "$2"
        ;;
    "help"|*)
        show_help
        ;;
esac 