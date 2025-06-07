#!/bin/bash

# Generate Go models and server code from OpenAPI specification using Docker
# This script generates code for both workout and llm services

set -e

echo "Generating Go data models from OpenAPI specification using Docker..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
API_CONTRACTS_DIR="$(dirname "$SCRIPT_DIR")"
OPENAPI_FILE="$API_CONTRACTS_DIR/openapi.yaml"
OUTPUT_DIR="$API_CONTRACTS_DIR/generated/go"

# Clean previous generation
echo -e "${YELLOW} Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR/models"

# Generate Go models using Docker
echo -e "${BLUE} Generating Go API models and interfaces...${NC}"
docker run --rm \
  -v "$API_CONTRACTS_DIR:/local" \
  openapitools/openapi-generator-cli generate \
  -i /local/openapi.yaml \
  -g go \
  -o /local/generated/go/models \
  --skip-validate-spec \
  --package-name api_models \
  --global-property models,supportingFiles,apiDocs=false,modelDocs=false,modelTests=false,apiTests=false \
  --additional-properties=packageName=api_models,generateInterfaces=true

for file in .travis.yml .gitignore .openapi-generator-ignore README.md git_push.sh client.go configuration.go response.go; do
  rm -f "$OUTPUT_DIR/models/$file"
done

# Copy models to backend services (optional)
echo -e "${YELLOW} Copying generated models to backend services...${NC}"
WORKOUT_SERVICE_DIR="$API_CONTRACTS_DIR/../backend/services/workoutAppServices/internal/generated"
rm -rf "$WORKOUT_SERVICE_DIR"
mkdir -p "$WORKOUT_SERVICE_DIR"
find "$OUTPUT_DIR" -type f -name "*.go" ! -name "api_*.go" ! -name "client.go" ! -name "configuration.go" ! -name "response.go" -exec cp {} "$WORKOUT_SERVICE_DIR" \;

echo -e "${GREEN}Go code generation completed successfully!${NC}"
echo -e "${BLUE} Generated files are in: $OUTPUT_DIR${NC}"
echo -e "${BLUE} Models copied to backend services${NC}"
