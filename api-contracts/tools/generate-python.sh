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

# Check if OpenAPI generator is available
if ! command -v openapi-generator-cli &> /dev/null; then
    echo -e "${RED}openapi-generator-cli not found. Please install it first:${NC}"
    echo -e "${BLUE}pip install openapi-generator-cli${NC}"
    exit 1
fi

# Clean previous generation
echo -e "${YELLOW} Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Generate Python Pydantic models
echo -e "${BLUE} Generating Python Pydantic models...${NC}"
openapi-generator-cli generate \
  -i "$OPENAPI_FILE" \
  -g python-pydantic-v1 \
  -o "$OUTPUT_DIR/models" \
  --package-name trackbot_models \
  --additional-properties=packageName=trackbot_models,packageVersion=1.0.0,packageUrl=https://github.com/yourorg/trackbot

# Generate Python FastAPI server (optional)
echo -e "${BLUE} Generating FastAPI server code...${NC}"
openapi-generator-cli generate \
  -i "$OPENAPI_FILE" \
  -g python-fastapi \
  -o "$OUTPUT_DIR/fastapi-server" \
  --package-name trackbot_api \
  --additional-properties=packageName=trackbot_api,packageVersion=1.0.0,generateSourceCodeOnly=false

# Generate Python HTTP client
echo -e "${BLUE} Generating Python HTTP client...${NC}"
openapi-generator-cli generate \
  -i "$OPENAPI_FILE" \
  -g python \
  -o "$OUTPUT_DIR/client" \
  --package-name trackbot_client \
  --additional-properties=packageName=trackbot_client,packageVersion=1.0.0,library=urllib3

# Generate advanced Pydantic v2 models (if datamodel-codegen is available)
echo -e "${BLUE} Generating Pydantic v2 models with datamodel-codegen...${NC}"
if command -v datamodel-codegen &> /dev/null; then
    datamodel-codegen \
        --input "$OPENAPI_FILE" \
        --input-file-type openapi \
        --output "$OUTPUT_DIR/pydantic-v2/models.py" \
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
    
    echo -e "${GREEN} Pydantic v2 models generated with datamodel-codegen${NC}"
else
    echo -e "${YELLOW}⚠️ datamodel-codegen not found. Install with: pip install datamodel-code-generator[http]${NC}"
fi


# Copy models to potential Python services
#echo -e "${YELLOW} Setting up Python service directories...${NC}"

# Create Python service directories (if they don't exist)
#PYTHON_SERVICES_DIR="$API_CONTRACTS_DIR/../backend/services/pythonServices"
#mkdir -p "$PYTHON_SERVICES_DIR/shared/models"

# Copy Pydantic models
#if [ -d "$OUTPUT_DIR/models" ]; then
#    cp -r "$OUTPUT_DIR/models/"* "$PYTHON_SERVICES_DIR/shared/models/"
#fi

#if ! command -v datamodel-codegen &> /dev/null; then
#    echo -e "\n${YELLOW}💡 Tip: Install datamodel-codegen for better Pydantic v2 support:${NC}"
#    echo -e "${BLUE}pip install 'datamodel-code-generator[http]'${NC}"
#fi 