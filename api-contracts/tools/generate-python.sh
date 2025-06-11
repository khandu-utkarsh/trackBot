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

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
API_CONTRACTS_DIR="$(dirname "$SCRIPT_DIR")"
OUTPUT_DIR="$API_CONTRACTS_DIR/generated/python"
CLIENT_PACKAGE_NAME="openapi-trackbot-models"
TARGET_DIR="$API_CONTRACTS_DIR/../backend/trackBot/app/models/"
# Debug information
echo "SCRIPT_DIR: $SCRIPT_DIR"
echo "API_CONTRACTS_DIR: $API_CONTRACTS_DIR"
echo "OPENAPI_FILE: $OPENAPI_FILE"
echo "OUTPUT_DIR: $OUTPUT_DIR"
echo "TARGET_DIR: $TARGET_DIR"

# Check if Docker is available
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Docker not found. Please install Docker first.${NC}"
    exit 1
fi

# Clean previous generation
echo -e "${YELLOW}Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Generate Python HTTP client
echo -e "${BLUE}Generating Python HTTP client...${NC}"
docker run --rm \
  -v "$API_CONTRACTS_DIR:/workspace" \
  -w /workspace \
  koxudaxi/datamodel-code-generator:latest \
  --input openapi.yaml \
  --input-file-type openapi \
  --output generated/python/models.py \
  --use-standard-collections \
  --target-python-version 3.10 \
  --output-model-type pydantic_v2.BaseModel


if [ $? -ne 0 ]; then
    echo -e "${RED}Failed to generate Python models. Please check the error above.${NC}"
    exit 1
fi

echo -e "${BLUE}Python models generated successfully!${NC}"

UNNEEDED_FILES=(
  .travis.yml .gitignore .openapi-generator-ignore README.md git_push.sh
  .gitlab-ci.yml requirements.txt pyproject.toml tox.ini setup.py setup.cfg
  .github test-requirements.txt
)

for file in "${UNNEEDED_FILES[@]}"; do
  rm -rf "$API_CONTRACTS_DIR/generated/python/$file"
done

# Copy models to potential Python services
echo -e "${YELLOW}Setting up Python service directories...${NC}"

# Create target directory structure
#echo -e "${BLUE}Creating target directory structure...${NC}"
#mkdir -p "$TARGET_DIR"

# Copy the generated models
echo -e "${BLUE}Copying generated models...${NC}"
cp "$OUTPUT_DIR/models.py" "$TARGET_DIR/"

if [ $? -ne 0 ]; then
    echo -e "${RED}Failed to copy models to target directory.${NC}"
    exit 1
fi

echo -e "${GREEN}Python code generation completed successfully!${NC}"
echo -e "${BLUE}Generated files are in: $OUTPUT_DIR${NC}"
echo -e "${BLUE}Models copied to: $TARGET_DIR${NC}"

