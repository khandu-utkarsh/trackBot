#!/bin/bash

# Production environment management script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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

# Check if .env file exists
check_env() {
    if [ ! -f .env ]; then
        print_error ".env file not found. Please create it with your production environment variables."
        print_status "Required variables:"
        echo "  OPENAI_API_KEY=your_production_openai_api_key"
        return 1
    fi
    
    # Check if OPENAI_API_KEY is set
    if ! grep -q "OPENAI_API_KEY=" .env || grep -q "OPENAI_API_KEY=your_" .env; then
        print_error "Please set a valid OPENAI_API_KEY in .env file"
        return 1
    fi
    
    return 0
}

# Start production environment
start_prod() {
    print_status "Starting production environment..."
    
    if ! check_env; then
        print_error "Please configure .env file first"
        exit 1
    fi
    
    print_status "Building and starting containers in production mode..."
    docker-compose -f docker-compose.prod.yml up --build -d
    
    print_success "Production environment started!"
    print_status "Services are running in the background"
    status
}

# Stop production environment
stop_prod() {
    print_status "Stopping production environment..."
    docker-compose -f docker-compose.prod.yml down
    print_success "Production environment stopped"
}

# Restart production environment
restart_prod() {
    print_status "Restarting production environment..."
    docker-compose -f docker-compose.prod.yml down
    docker-compose -f docker-compose.prod.yml up --build -d
    print_success "Production environment restarted!"
    status
}

# Update production environment
update_prod() {
    print_status "Updating production environment..."
    print_warning "This will rebuild all containers with latest code"
    
    read -p "Continue? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        docker-compose -f docker-compose.prod.yml down
        docker-compose -f docker-compose.prod.yml build --no-cache
        docker-compose -f docker-compose.prod.yml up -d
        print_success "Production environment updated!"
        status
    else
        print_status "Update cancelled"
    fi
}

# Show logs
logs() {
    service=${1:-""}
    if [ -z "$service" ]; then
        print_status "Showing logs for all services..."
        docker-compose -f docker-compose.prod.yml logs -f
    else
        print_status "Showing logs for $service..."
        docker-compose -f docker-compose.prod.yml logs -f "$service"
    fi
}

# Show status
status() {
    print_status "Container status:"
    docker-compose -f docker-compose.prod.yml ps
    echo
    print_status "Service URLs:"
    echo "  Frontend:    http://localhost:3000"
    echo "  Backend:     http://localhost:8080"
    echo "  LLM Service: http://localhost:8081"
    echo "  Database:    localhost:5432"
    echo
    print_status "Health checks:"
    echo "  Backend:     curl http://localhost:8080/health"
    echo "  LLM Service: curl http://localhost:8081/api/v1/chat/health"
}

# Database shell
db_shell() {
    print_status "Connecting to production database..."
    docker exec -it postgres psql -U postgres -d workout_app
}

# Backup database
backup_db() {
    timestamp=$(date +%Y%m%d_%H%M%S)
    backup_file="backup_${timestamp}.sql"
    
    print_status "Creating database backup: $backup_file"
    docker exec postgres pg_dump -U postgres workout_app > "$backup_file"
    print_success "Database backup created: $backup_file"
}

# Monitor resources
monitor() {
    print_status "Resource usage:"
    docker stats --no-stream
}

# Show help
show_help() {
    echo "Production Environment Management Script"
    echo
    echo "Usage: $0 [COMMAND]"
    echo
    echo "Commands:"
    echo "  start     Start production environment"
    echo "  stop      Stop production environment"
    echo "  restart   Restart production environment"
    echo "  update    Update production environment (rebuild)"
    echo "  logs      Show logs for all services or specific service"
    echo "  status    Show container status and service URLs"
    echo "  db        Open database shell"
    echo "  backup    Create database backup"
    echo "  monitor   Show resource usage"
    echo "  help      Show this help message"
    echo
    echo "Examples:"
    echo "  $0 start"
    echo "  $0 logs frontend"
    echo "  $0 backup"
}

# Main script logic
case "${1:-help}" in
    "start")
        start_prod
        ;;
    "stop")
        stop_prod
        ;;
    "restart")
        restart_prod
        ;;
    "update")
        update_prod
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
    "backup")
        backup_db
        ;;
    "monitor")
        monitor
        ;;
    "help"|*)
        show_help
        ;;
esac 