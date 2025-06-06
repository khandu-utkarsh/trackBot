# trackbot_models.ExercisesApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_exercise**](ExercisesApi.md#create_exercise) | **POST** /workouts/{workoutId}/exercises | Create a new exercise
[**delete_exercise**](ExercisesApi.md#delete_exercise) | **DELETE** /exercises/{exerciseId} | Delete exercise
[**get_exercise_by_id**](ExercisesApi.md#get_exercise_by_id) | **GET** /exercises/{exerciseId} | Get exercise by ID
[**list_exercises**](ExercisesApi.md#list_exercises) | **GET** /workouts/{workoutId}/exercises | List exercises for a workout
[**update_exercise**](ExercisesApi.md#update_exercise) | **PUT** /exercises/{exerciseId} | Update exercise


# **create_exercise**
> CreateExerciseResponse create_exercise(workout_id, create_exercise_request)

Create a new exercise

Add a new exercise to a workout (cardio or weight training)

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.create_exercise_request import CreateExerciseRequest
from trackbot_models.models.create_exercise_response import CreateExerciseResponse
from trackbot_models.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_models.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_models.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_models.ExercisesApi(api_client)
    workout_id = 1 # int | Workout ID
    create_exercise_request = trackbot_models.CreateExerciseRequest() # CreateExerciseRequest | 

    try:
        # Create a new exercise
        api_response = api_instance.create_exercise(workout_id, create_exercise_request)
        print("The response of ExercisesApi->create_exercise:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExercisesApi->create_exercise: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workout_id** | **int**| Workout ID | 
 **create_exercise_request** | [**CreateExerciseRequest**](CreateExerciseRequest.md)|  | 

### Return type

[**CreateExerciseResponse**](CreateExerciseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Exercise created successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_exercise**
> delete_exercise(exercise_id)

Delete exercise

Delete an exercise from a workout

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_models.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_models.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_models.ExercisesApi(api_client)
    exercise_id = 1 # int | Exercise ID

    try:
        # Delete exercise
        api_instance.delete_exercise(exercise_id)
    except Exception as e:
        print("Exception when calling ExercisesApi->delete_exercise: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exercise_id** | **int**| Exercise ID | 

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
**204** | Operation completed successfully with no content to return |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_exercise_by_id**
> Exercise get_exercise_by_id(exercise_id)

Get exercise by ID

Retrieve a specific exercise by its ID

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.exercise import Exercise
from trackbot_models.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_models.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_models.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_models.ExercisesApi(api_client)
    exercise_id = 1 # int | Exercise ID

    try:
        # Get exercise by ID
        api_response = api_instance.get_exercise_by_id(exercise_id)
        print("The response of ExercisesApi->get_exercise_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExercisesApi->get_exercise_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exercise_id** | **int**| Exercise ID | 

### Return type

[**Exercise**](Exercise.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Exercise details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_exercises**
> List[Exercise] list_exercises(workout_id)

List exercises for a workout

Retrieve all exercises for a specific workout

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.exercise import Exercise
from trackbot_models.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_models.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_models.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_models.ExercisesApi(api_client)
    workout_id = 1 # int | Workout ID

    try:
        # List exercises for a workout
        api_response = api_instance.list_exercises(workout_id)
        print("The response of ExercisesApi->list_exercises:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExercisesApi->list_exercises: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workout_id** | **int**| Workout ID | 

### Return type

[**List[Exercise]**](Exercise.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of exercises retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_exercise**
> Exercise update_exercise(exercise_id, update_exercise_request)

Update exercise

Update an existing exercise

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.exercise import Exercise
from trackbot_models.models.update_exercise_request import UpdateExerciseRequest
from trackbot_models.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_models.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_models.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_models.ExercisesApi(api_client)
    exercise_id = 1 # int | Exercise ID
    update_exercise_request = trackbot_models.UpdateExerciseRequest() # UpdateExerciseRequest | 

    try:
        # Update exercise
        api_response = api_instance.update_exercise(exercise_id, update_exercise_request)
        print("The response of ExercisesApi->update_exercise:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExercisesApi->update_exercise: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exercise_id** | **int**| Exercise ID | 
 **update_exercise_request** | [**UpdateExerciseRequest**](UpdateExerciseRequest.md)|  | 

### Return type

[**Exercise**](Exercise.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Exercise updated successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

