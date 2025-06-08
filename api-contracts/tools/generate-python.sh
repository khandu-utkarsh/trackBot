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
CLIENT_PACKAGE_NAME="trackbot_client"

TARGET_DIR="$API_CONTRACTS_DIR/../backend/services/llmServices/app/$CLIENT_PACKAGE_NAME"


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
echo -e "${YELLOW} Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"


## Commenting out the client generation for now
## TODO: Uncomment this when we have a use case for it

# Generate Python HTTP client
echo -e "${BLUE} Generating Python HTTP client...${NC}"
docker run --rm \
  -v "$API_CONTRACTS_DIR:/local" \
  openapitools/openapi-generator-cli generate \
  -i /local/openapi.yaml \
  -g python \
  -o /local/generated/python/client \
  --package-name trackbot_client \
  --additional-properties=packageName=trackbot_client,packageVersion=1.0.0,library=urllib3 \
  --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false \
  --ignore-file-override=/local/.openapi-generator-ignore


UNNEEDED_FILES=(
  .travis.yml .gitignore .openapi-generator-ignore README.md git_push.sh
  .gitlab-ci.yml requirements.txt pyproject.toml tox.ini setup.py setup.cfg
  .github test-requirements.txt
)

for file in "${UNNEEDED_FILES[@]}"; do
  rm -rf "$OUTPUT_DIR/client/$file"
done


# Copy models to potential Python services
echo -e "${YELLOW} Setting up Python service directories...${NC}"


rm -rf "$TARGET_DIR"
mkdir -p "$(dirname "$TARGET_DIR")"
cp -r "$OUTPUT_DIR/client/$CLIENT_PACKAGE_NAME" "$TARGET_DIR"

echo -e "${GREEN}Python code generation completed successfully!${NC}"
echo -e "${BLUE} Generated files are in: $OUTPUT_DIR${NC}"
echo -e "${BLUE} Models copied to backend services${NC}"

