# WorkoutsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createWorkout**](#createworkout) | **POST** /users/{userId}/workouts | Create a new workout|
|[**deleteWorkout**](#deleteworkout) | **DELETE** /workouts/{workoutId} | Delete workout|
|[**getWorkoutById**](#getworkoutbyid) | **GET** /workouts/{workoutId} | Get workout by ID|
|[**listWorkouts**](#listworkouts) | **GET** /users/{userId}/workouts | List workouts for a user|
|[**updateWorkout**](#updateworkout) | **PUT** /workouts/{workoutId} | Update workout|

# **createWorkout**
> CreateWorkoutResponse createWorkout(createWorkoutRequest)

Create a new workout session for a user

### Example

```typescript
import {
    WorkoutsApi,
    Configuration,
    CreateWorkoutRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new WorkoutsApi(configuration);

let userId: number; //User ID (default to undefined)
let createWorkoutRequest: CreateWorkoutRequest; //

const { status, data } = await apiInstance.createWorkout(
    userId,
    createWorkoutRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createWorkoutRequest** | **CreateWorkoutRequest**|  | |
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**CreateWorkoutResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Workout created successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteWorkout**
> deleteWorkout()

Delete a workout and all associated exercises

### Example

```typescript
import {
    WorkoutsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new WorkoutsApi(configuration);

let workoutId: number; //Workout ID (default to undefined)

const { status, data } = await apiInstance.deleteWorkout(
    workoutId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **workoutId** | [**number**] | Workout ID | defaults to undefined|


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

# **getWorkoutById**
> Workout getWorkoutById()

Retrieve a specific workout by its ID

### Example

```typescript
import {
    WorkoutsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new WorkoutsApi(configuration);

let workoutId: number; //Workout ID (default to undefined)

const { status, data } = await apiInstance.getWorkoutById(
    workoutId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **workoutId** | [**number**] | Workout ID | defaults to undefined|


### Return type

**Workout**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Workout details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listWorkouts**
> Array<Workout> listWorkouts()

Retrieve all workouts for a specific user with optional date filtering

### Example

```typescript
import {
    WorkoutsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new WorkoutsApi(configuration);

let userId: number; //User ID (default to undefined)
let year: string; //Filter by year (YYYY format) (optional) (default to undefined)
let month: string; //Filter by month (MM format) (optional) (default to undefined)
let day: string; //Filter by day (DD format) (optional) (default to undefined)

const { status, data } = await apiInstance.listWorkouts(
    userId,
    year,
    month,
    day
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **userId** | [**number**] | User ID | defaults to undefined|
| **year** | [**string**] | Filter by year (YYYY format) | (optional) defaults to undefined|
| **month** | [**string**] | Filter by month (MM format) | (optional) defaults to undefined|
| **day** | [**string**] | Filter by day (DD format) | (optional) defaults to undefined|


### Return type

**Array<Workout>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of workouts retrieved successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateWorkout**
> Workout updateWorkout(updateWorkoutRequest)

Update an existing workout

### Example

```typescript
import {
    WorkoutsApi,
    Configuration,
    UpdateWorkoutRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new WorkoutsApi(configuration);

let workoutId: number; //Workout ID (default to undefined)
let updateWorkoutRequest: UpdateWorkoutRequest; //

const { status, data } = await apiInstance.updateWorkout(
    workoutId,
    updateWorkoutRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateWorkoutRequest** | **UpdateWorkoutRequest**|  | |
| **workoutId** | [**number**] | Workout ID | defaults to undefined|


### Return type

**Workout**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Workout updated successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

