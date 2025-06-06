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
  --additional-properties=npmName=@trackbot-app/api-client,npmVersion=1.0.0,supportsES6=true,typescriptThreePlus=true


# Copy to frontend directory
#echo -e "${YELLOW} Copying generated types to frontend...${NC}"
#`FRONTEND_DIR="$API_CONTRACTS_DIR/../frontend/src/types/generated"
#mkdir -p "$FRONTEND_DIR"

# Copy the main types
#if [ -d "$OUTPUT_DIR/types" ]; then
#  cp -r "$OUTPUT_DIR/types/"* "$FRONTEND_DIR/"
#fi

# Copy API client to frontend
#FRONTEND_API_DIR="$API_CONTRACTS_DIR/../frontend/src/api/generated"
#mkdir -p "$FRONTEND_API_DIR"

#if [ -d "$OUTPUT_DIR/api-client" ]; then
#  cp -r "$OUTPUT_DIR/api-client/"* "$FRONTEND_API_DIR/"
#fi

# Generate package.json for the TypeScript client
#cat > "$OUTPUT_DIR/api-client/package.json" << EOF
#{
#  "name": "@trackbot-app/api-client",
#  "version": "1.0.0",
#  "description": "Generated TypeScript API client for TrackBot App",
#  "main": "dist/index.js",
#  "types": "dist/index.d.ts",
#  "scripts": {
#    "build": "tsc",
#    "prepublishOnly": "npm run build"
#  },
#  "dependencies": {
#    "axios": "^1.6.0"
#  },
#  "devDependencies": {
#    "typescript": "^5.0.0",
#    "@types/node": "^20.0.0"
#  },
#  "files": [
#    "dist",
#    "src"
#  ]
#}
#EOF


#echo -e "${GREEN} TypeScript code generation completed successfully!${NC}"
#echo -e "${BLUE} Generated files are in: $OUTPUT_DIR${NC}"
#echo -e "${BLUE} Types and API client copied to frontend${NC}"