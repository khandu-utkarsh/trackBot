# ExercisesApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createExercise**](#createexercise) | **POST** /workouts/{workoutId}/exercises | Create a new exercise|
|[**deleteExercise**](#deleteexercise) | **DELETE** /exercises/{exerciseId} | Delete exercise|
|[**getExerciseById**](#getexercisebyid) | **GET** /exercises/{exerciseId} | Get exercise by ID|
|[**listExercises**](#listexercises) | **GET** /workouts/{workoutId}/exercises | List exercises for a workout|
|[**updateExercise**](#updateexercise) | **PUT** /exercises/{exerciseId} | Update exercise|

# **createExercise**
> CreateExerciseResponse createExercise(createExerciseRequest)

Add a new exercise to a workout (cardio or weight training)

### Example

```typescript
import {
    ExercisesApi,
    Configuration,
    CreateExerciseRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ExercisesApi(configuration);

let workoutId: number; //Workout ID (default to undefined)
let createExerciseRequest: CreateExerciseRequest; //

const { status, data } = await apiInstance.createExercise(
    workoutId,
    createExerciseRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createExerciseRequest** | **CreateExerciseRequest**|  | |
| **workoutId** | [**number**] | Workout ID | defaults to undefined|


### Return type

**CreateExerciseResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Exercise created successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteExercise**
> deleteExercise()

Delete an exercise from a workout

### Example

```typescript
import {
    ExercisesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ExercisesApi(configuration);

let exerciseId: number; //Exercise ID (default to undefined)

const { status, data } = await apiInstance.deleteExercise(
    exerciseId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **exerciseId** | [**number**] | Exercise ID | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** | Operation completed successfully with no content to return |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getExerciseById**
> Exercise getExerciseById()

Retrieve a specific exercise by its ID

### Example

```typescript
import {
    ExercisesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ExercisesApi(configuration);

let exerciseId: number; //Exercise ID (default to undefined)

const { status, data } = await apiInstance.getExerciseById(
    exerciseId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **exerciseId** | [**number**] | Exercise ID | defaults to undefined|


### Return type

**Exercise**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Exercise details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listExercises**
> Array<Exercise> listExercises()

Retrieve all exercises for a specific workout

### Example

```typescript
import {
    ExercisesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ExercisesApi(configuration);

let workoutId: number; //Workout ID (default to undefined)

const { status, data } = await apiInstance.listExercises(
    workoutId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **workoutId** | [**number**] | Workout ID | defaults to undefined|


### Return type

**Array<Exercise>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of exercises retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateExercise**
> Exercise updateExercise(updateExerciseRequest)

Update an existing exercise

### Example

```typescript
import {
    ExercisesApi,
    Configuration,
    UpdateExerciseRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ExercisesApi(configuration);

let exerciseId: number; //Exercise ID (default to undefined)
let updateExerciseRequest: UpdateExerciseRequest; //

const { status, data } = await apiInstance.updateExercise(
    exerciseId,
    updateExerciseRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateExerciseRequest** | **UpdateExerciseRequest**|  | |
| **exerciseId** | [**number**] | Exercise ID | defaults to undefined|


### Return type

**Exercise**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Exercise updated successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

