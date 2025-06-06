# Workout App API Contracts

This directory contains the comprehensive OpenAPI specification and code generation tools for the Workout App. It serves as the **single source of truth** for all API contracts across the entire application.

## 🏗️ Structure

```
api-contracts/
├── openapi.yaml                    # Main OpenAPI 3.0 specification
├── components/                     # Reusable OpenAPI components
│   ├── schemas/                    # Data model definitions
│   │   ├── user.yaml              # User-related schemas
│   │   ├── workout.yaml           # Workout-related schemas
│   │   ├── exercise.yaml          # Exercise-related schemas
│   │   ├── conversation.yaml      # Conversation-related schemas
│   │   └── message.yaml           # Message-related schemas
│   ├── parameters/                # Reusable parameters
│   │   └── common.yaml            # Common path/query parameters
│   └── responses/                 # Reusable responses
│       └── errors.yaml            # Error response definitions
├── services/                      # Service-specific specifications
├── generated/                     # Generated code output
│   ├── go/                        # Generated Go code
│   └── typescript/                # Generated TypeScript code
├── tools/                         # Generation and tooling scripts
│   ├── package.json              # Node.js dependencies for tools
│   ├── generate-go.sh            # Go code generation script
│   └── generate-typescript.sh    # TypeScript generation script
└── docs/                          # Generated documentation
```

## 🎯 Data Models

The API includes the following core models based on your existing Go structs:

### User Management
- **User**: Core user entity with email and timestamps
- **CreateUserRequest/Response**: User creation endpoints
- **UpdateUserRequest**: User modification

### Workout Tracking
- **Workout**: Workout session entity
- **CreateWorkoutRequest/Response**: Workout creation
- **WorkoutListParams**: Advanced filtering (by date, user)

### Exercise Management
- **Exercise**: Polymorphic exercise model supporting:
  - **CardioExercise**: Distance and duration tracking
  - **WeightExercise**: Sets, reps, and weight tracking
- **ExerciseType**: Enum for cardio/weights
- **Create/UpdateExerciseRequest**: Type-specific creation

### AI Conversations
- **Conversation**: Chat conversation entity with active status
- **Message**: Individual message with type (user/assistant/system)
- **MessageType**: Enum for message sender types

## 🚀 Quick Start

### 1. Install Tools
```bash
cd api-contracts/tools
npm install
```

### 2. Generate Code
```bash
# Generate all code
npm run generate:all

# Or generate individually
npm run generate:go        # Go models and server
npm run generate:typescript # TypeScript generation script
```

### 3. Validate Specification
```bash
npm run validate           # Validate OpenAPI spec
npm run lint              # Lint with Spectral
```

### 4. Documentation
```bash
npm run docs:serve        # Serve interactive docs
npm run docs:build        # Build static docs
```

## 🔧 Code Generation

### Go Code Generation

The Go generation creates:
- **Models**: Struct definitions matching your existing models
- **Server**: Gin server stubs with handlers
- **Client**: HTTP client for service-to-service communication

Generated Go code is automatically copied to:
- `backend/services/workoutAppServices/internal/generated/`
- `backend/services/llmServices/internal/generated/`

### TypeScript Generation

The TypeScript generation creates:
- **Types**: Interface definitions for all models
- **API Client**: Axios-based HTTP client
- **React Query Hooks**: Ready-to-use hooks for React

Generated TypeScript code is automatically copied to:
- `frontend/src/types/generated/`
- `frontend/src/api/generated/`

## 🔄 Workflow

1. **Modify OpenAPI specs** in `components/schemas/` or main `openapi.yaml`
2. **Run generation** with `npm run generate:all`
3. **Generated code is automatically copied** to backend services and frontend
4. **Use the generated types/clients** in your application code

## 📖 API Documentation

The OpenAPI specification includes:

### Endpoints
- **Users**: CRUD operations, lookup by email
- **Workouts**: Per-user workout management with date filtering
- **Exercises**: Polymorphic exercise tracking (cardio/weights)
- **Conversations**: AI chat conversation management
- **Messages**: Message handling within conversations

### Features
- **Comprehensive validation**: Input validation with patterns and constraints
- **Error handling**: Standardized error responses with codes
- **Type safety**: Discriminated unions for exercise types
- **Filtering**: Advanced query parameters for workouts
- **Relationships**: Proper foreign key relationships

## 🛠️ Tools and Scripts

### Available Scripts
```bash
npm run validate          # Validate OpenAPI spec
npm run generate:go      # Generate Go code
npm run generate:typescript # Generate TypeScript code
npm run generate:all     # Generate all code
npm run docs:serve       # Serve documentation
npm run docs:build       # Build static documentation
npm run lint            # Lint OpenAPI spec
```

### Manual Generation
```bash
# Make scripts executable (one time)
chmod +x tools/generate-go.sh tools/generate-typescript.sh

# Run generation
./tools/generate-go.sh
./tools/generate-typescript.sh
```

## 📋 Examples

### Using Generated Go Models
```go
package main

import (
    "workout_app_backend/generated/models"
    "time"
)

func main() {
    user := models.User{
        Email: "user@example.com",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    cardio := models.CardioExercise{
        BaseExercise: models.BaseExercise{
            Name: "Running",
            Type: models.ExerciseTypeCardio,
        },
        Distance: 5000.0,
        Duration: 1800,
    }
}
```

### Using Generated TypeScript Client
```typescript
import { Configuration, UsersApi, ExercisesApi } from './api/generated';

const config = new Configuration({
  basePath: 'http://localhost:8080/api/v1'
});

const usersApi = new UsersApi(config);
const exercisesApi = new ExercisesApi(config);

// Create user
const user = await usersApi.createUser({
  createUserRequest: { email: 'user@example.com' }
});

// Create cardio exercise
const exercise = await exercisesApi.createExercise({
  workoutId: 1,
  createExerciseRequest: {
    name: 'Running',
    type: 'cardio',
    distance: 5000,
    duration: 1800
  }
});
```

## 🔍 Benefits

1. **Single Source of Truth**: All API contracts in one place
2. **Type Safety**: Generated types prevent runtime errors
3. **Consistency**: Same models across Go backend and TypeScript frontend
4. **Developer Experience**: Auto-completion and IntelliSense
5. **Documentation**: Always up-to-date API docs
6. **Validation**: Built-in request/response validation
7. **Maintainability**: Changes propagate automatically

## 🚨 Important Notes

- **Never edit generated code manually** - it will be overwritten
- **Always regenerate after OpenAPI changes**
- **Generated code is gitignored** in the services directories
- **The OpenAPI spec should be committed** to version control

## 🆘 Troubleshooting

### Common Issues

1. **Generation fails**: Ensure `openapi-generator-cli` is installed
2. **Invalid spec**: Run `npm run validate` to check for errors
3. **Type mismatches**: Check discriminator properties in polymorphic types
4. **Missing references**: Ensure all `$ref` paths are correct

### Getting Help

1. Validate your OpenAPI spec: `npm run validate`
2. Check the generated documentation: `npm run docs:serve`
3. Review the generated code structure in `generated/` 