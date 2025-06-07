#!/bin/bash

# Generate TypeScript types and API client from OpenAPI specification
# This script generates types and client for the frontend

set -e

echo "Generating TypeScript code from OpenAPI specification..."

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Directories
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
API_CONTRACTS_DIR="$(dirname "$SCRIPT_DIR")"
OPENAPI_FILE="$API_CONTRACTS_DIR/openapi.yaml"
OUTPUT_DIR="$API_CONTRACTS_DIR/generated/typescript"

# Clean previous generation
echo -e "${YELLOW} Cleaning previous generated files...${NC}"
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Generate TypeScript types and API client
echo -e "${BLUE} Generating TypeScript types and API client...${NC}"
docker run --rm \
  -v "$API_CONTRACTS_DIR:/local" \
  openapitools/openapi-generator-cli generate \
  -i /local/openapi.yaml \
  -g typescript-axios \
  -o /local/generated/typescript/api-client \
  --additional-properties=npmName=@trackbot-app/api-client,npmVersion=1.0.0,supportsES6=true,typescriptThreePlus=true \
  --global-property=apiTests=false,modelTests=false,apiDocs=false,modelDocs=false \
  --ignore-file-override=/local/.openapi-generator-ignore



for file in .travis.yml .gitignore .openapi-generator-ignore README.md git_push.sh .npmignore package-lock.json tsconfig.json tsconfig.esm.json package.json; do
  rm -f "$OUTPUT_DIR/api-client/$file"
done


# Copy to frontend directory
echo -e "${YELLOW} Copying generated types to frontend...${NC}"
FRONTEND_DIR="$API_CONTRACTS_DIR/../frontend/src/lib/types/generated"
rm -rf "$WORKOUT_SERVICE_DIR"
mkdir -p "$FRONTEND_DIR"

# Copy the main types
# Copy only .ts files from the generated client
if [ -d "$OUTPUT_DIR/api-client" ]; then
  find "$OUTPUT_DIR/api-client" -name '*.ts' -exec cp {} "$FRONTEND_DIR/" \;
fi

echo -e "${GREEN} TypeScript code generation completed successfully!${NC}"
echo -e "${BLUE} Generated files are in: $OUTPUT_DIR${NC}"
echo -e "${BLUE} Types and API client copied to frontend${NC}"