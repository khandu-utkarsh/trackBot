## @trackbot-app/api-client@1.0.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install @trackbot-app/api-client@1.0.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConversationsApi* | [**createConversation**](docs/ConversationsApi.md#createconversation) | **POST** /users/{userId}/conversations | Create a new conversation
*ConversationsApi* | [**deleteConversation**](docs/ConversationsApi.md#deleteconversation) | **DELETE** /conversations/{conversationId} | Delete conversation
*ConversationsApi* | [**getConversationById**](docs/ConversationsApi.md#getconversationbyid) | **GET** /conversations/{conversationId} | Get conversation by ID
*ConversationsApi* | [**listConversations**](docs/ConversationsApi.md#listconversations) | **GET** /users/{userId}/conversations | List conversations for a user
*ExercisesApi* | [**createExercise**](docs/ExercisesApi.md#createexercise) | **POST** /workouts/{workoutId}/exercises | Create a new exercise
*ExercisesApi* | [**deleteExercise**](docs/ExercisesApi.md#deleteexercise) | **DELETE** /exercises/{exerciseId} | Delete exercise
*ExercisesApi* | [**getExerciseById**](docs/ExercisesApi.md#getexercisebyid) | **GET** /exercises/{exerciseId} | Get exercise by ID
*ExercisesApi* | [**listExercises**](docs/ExercisesApi.md#listexercises) | **GET** /workouts/{workoutId}/exercises | List exercises for a workout
*ExercisesApi* | [**updateExercise**](docs/ExercisesApi.md#updateexercise) | **PUT** /exercises/{exerciseId} | Update exercise
*MessagesApi* | [**createMessage**](docs/MessagesApi.md#createmessage) | **POST** /conversations/{conversationId}/messages | Create a new message
*MessagesApi* | [**deleteMessage**](docs/MessagesApi.md#deletemessage) | **DELETE** /messages/{messageId} | Delete message
*MessagesApi* | [**getMessageById**](docs/MessagesApi.md#getmessagebyid) | **GET** /messages/{messageId} | Get message by ID
*MessagesApi* | [**listMessages**](docs/MessagesApi.md#listmessages) | **GET** /conversations/{conversationId}/messages | List messages in a conversation
*UsersApi* | [**createUser**](docs/UsersApi.md#createuser) | **POST** /users | Create a new user
*UsersApi* | [**deleteUser**](docs/UsersApi.md#deleteuser) | **DELETE** /users/{userId} | Delete user
*UsersApi* | [**getUserByEmail**](docs/UsersApi.md#getuserbyemail) | **GET** /users/email/{email} | Get user by email
*UsersApi* | [**getUserById**](docs/UsersApi.md#getuserbyid) | **GET** /users/{userId} | Get user by ID
*UsersApi* | [**listUsers**](docs/UsersApi.md#listusers) | **GET** /users | List all users
*WorkoutsApi* | [**createWorkout**](docs/WorkoutsApi.md#createworkout) | **POST** /users/{userId}/workouts | Create a new workout
*WorkoutsApi* | [**deleteWorkout**](docs/WorkoutsApi.md#deleteworkout) | **DELETE** /workouts/{workoutId} | Delete workout
*WorkoutsApi* | [**getWorkoutById**](docs/WorkoutsApi.md#getworkoutbyid) | **GET** /workouts/{workoutId} | Get workout by ID
*WorkoutsApi* | [**listWorkouts**](docs/WorkoutsApi.md#listworkouts) | **GET** /users/{userId}/workouts | List workouts for a user
*WorkoutsApi* | [**updateWorkout**](docs/WorkoutsApi.md#updateworkout) | **PUT** /workouts/{workoutId} | Update workout


### Documentation For Models

 - [BaseExercise](docs/BaseExercise.md)
 - [CardioExercise](docs/CardioExercise.md)
 - [Conversation](docs/Conversation.md)
 - [CreateCardioExerciseRequest](docs/CreateCardioExerciseRequest.md)
 - [CreateConversationRequest](docs/CreateConversationRequest.md)
 - [CreateConversationResponse](docs/CreateConversationResponse.md)
 - [CreateExerciseRequest](docs/CreateExerciseRequest.md)
 - [CreateExerciseResponse](docs/CreateExerciseResponse.md)
 - [CreateMessageRequest](docs/CreateMessageRequest.md)
 - [CreateMessageResponse](docs/CreateMessageResponse.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreateWeightExerciseRequest](docs/CreateWeightExerciseRequest.md)
 - [CreateWorkoutRequest](docs/CreateWorkoutRequest.md)
 - [CreateWorkoutResponse](docs/CreateWorkoutResponse.md)
 - [Exercise](docs/Exercise.md)
 - [ExerciseType](docs/ExerciseType.md)
 - [Message](docs/Message.md)
 - [MessageType](docs/MessageType.md)
 - [ModelError](docs/ModelError.md)
 - [UpdateExerciseRequest](docs/UpdateExerciseRequest.md)
 - [UpdateWorkoutRequest](docs/UpdateWorkoutRequest.md)
 - [User](docs/User.md)
 - [WeightExercise](docs/WeightExercise.md)
 - [Workout](docs/Workout.md)
 - [WorkoutListParams](docs/WorkoutListParams.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization

Endpoints do not require authorization.

