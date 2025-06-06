#!/bin/bash

# Generate Python Pydantic models and FastAPI client from OpenAPI specification
# This script generates Pydantic models and optionally FastAPI server code

set -e

echo "Generating Python code from OpenAPI specification..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
API_CONTRACTS_DIR="$(dirname "$SCRIPT_DIR")"
OPENAPI_FILE="$API_CONTRACTS_DIR/openapi.yaml"
OUTPUT_DIR="$API_CONTRACTS_DIR/generated/python"

echo "SCRIPT_DIR: $SCRIPT_DIR"
echo "API_CONTRACTS_DIR: $API_CONTRACTS_DIR"
echo "OPENAPI_FILE: $OPENAPI_FILE"
echo "OUTPUT_DIR: $OUTPUT_DIR"

# Check if Docker is available
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Docker not found. Please install Docker first.${NC}"
    exit 1
fi

# Clean previous generation
echo -e "${YELLOW} Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"


## Commenting out the client generation for now
## TODO: Uncomment this when we have a use case for it

# Generate Python HTTP client
#echo -e "${BLUE} Generating Python HTTP client...${NC}"
#docker run --rm \
#  -v "$API_CONTRACTS_DIR:/local" \
#  openapitools/openapi-generator-cli generate \
#  -i /local/openapi.yaml \
#  -g python \
#  -o /local/generated/python/client \
#  --package-name trackbot_client \
#  --additional-properties=packageName=trackbot_client,packageVersion=1.0.0,library=urllib3 \
#  --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false \
#  --ignore-file-override=/local/.openapi-generator-ignore

# Generate advanced Pydantic v2 models using Docker
echo -e "${BLUE} Generating Pydantic v2 models with datamodel-codegen...${NC}"
mkdir -p "$OUTPUT_DIR/pydantic-v2"

docker run --rm \
  -v "$API_CONTRACTS_DIR:/workspace" \
  -w /workspace \
  python:3.11-slim \
  sh -c "
    pip install 'datamodel-code-generator[http]' && \
    datamodel-codegen \
        --input /workspace/openapi.yaml \
        --input-file-type openapi \
        --output /workspace/generated/python/pydantic-v2/models.py \
        --output-model-type pydantic_v2.BaseModel \
        --field-constraints \
        --use-annotated \
        --use-generic-container-types \
        --use-schema-description \
        --use-field-description \
        --use-default-kwarg \
        --snake-case-field \
        --strict-nullable \
        --target-python-version 3.11
  "

echo -e "${GREEN} Pydantic v2 models generated with datamodel-codegen${NC}"


# Copy models to potential Python services
echo -e "${YELLOW} Setting up Python service directories...${NC}"

# Create Python service directories (if they don't exist)
PYTHON_SERVICES_DIR="$API_CONTRACTS_DIR/../backend/services/llmServices/internal/generated/models"
echo "PYTHON_SERVICES_DIR: $PYTHON_SERVICES_DIR"
mkdir -p "$PYTHON_SERVICES_DIR"

# Copy Pydantic models
if [ -d "$OUTPUT_DIR/pydantic-v2" ]; then
    cp -r "$OUTPUT_DIR/pydantic-v2/"* "$PYTHON_SERVICES_DIR"
fi

echo -e "${GREEN}Python code generation completed successfully!${NC}"
echo -e "${BLUE} Generated files are in: $OUTPUT_DIR${NC}"
echo -e "${BLUE} Models copied to backend services${NC}"

