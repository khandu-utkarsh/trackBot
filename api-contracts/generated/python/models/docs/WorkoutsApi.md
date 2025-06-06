# trackbot_models.WorkoutsApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_workout**](WorkoutsApi.md#create_workout) | **POST** /users/{userId}/workouts | Create a new workout
[**delete_workout**](WorkoutsApi.md#delete_workout) | **DELETE** /workouts/{workoutId} | Delete workout
[**get_workout_by_id**](WorkoutsApi.md#get_workout_by_id) | **GET** /workouts/{workoutId} | Get workout by ID
[**list_workouts**](WorkoutsApi.md#list_workouts) | **GET** /users/{userId}/workouts | List workouts for a user
[**update_workout**](WorkoutsApi.md#update_workout) | **PUT** /workouts/{workoutId} | Update workout


# **create_workout**
> CreateWorkoutResponse create_workout(user_id, create_workout_request)

Create a new workout

Create a new workout session for a user

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.create_workout_request import CreateWorkoutRequest
from trackbot_models.models.create_workout_response import CreateWorkoutResponse
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
    api_instance = trackbot_models.WorkoutsApi(api_client)
    user_id = 1 # int | User ID
    create_workout_request = trackbot_models.CreateWorkoutRequest() # CreateWorkoutRequest | 

    try:
        # Create a new workout
        api_response = api_instance.create_workout(user_id, create_workout_request)
        print("The response of WorkoutsApi->create_workout:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorkoutsApi->create_workout: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 
 **create_workout_request** | [**CreateWorkoutRequest**](CreateWorkoutRequest.md)|  | 

### Return type

[**CreateWorkoutResponse**](CreateWorkoutResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Workout created successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_workout**
> delete_workout(workout_id)

Delete workout

Delete a workout and all associated exercises

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
    api_instance = trackbot_models.WorkoutsApi(api_client)
    workout_id = 1 # int | Workout ID

    try:
        # Delete workout
        api_instance.delete_workout(workout_id)
    except Exception as e:
        print("Exception when calling WorkoutsApi->delete_workout: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workout_id** | **int**| Workout ID | 

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

# **get_workout_by_id**
> Workout get_workout_by_id(workout_id)

Get workout by ID

Retrieve a specific workout by its ID

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.workout import Workout
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
    api_instance = trackbot_models.WorkoutsApi(api_client)
    workout_id = 1 # int | Workout ID

    try:
        # Get workout by ID
        api_response = api_instance.get_workout_by_id(workout_id)
        print("The response of WorkoutsApi->get_workout_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorkoutsApi->get_workout_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workout_id** | **int**| Workout ID | 

### Return type

[**Workout**](Workout.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Workout details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_workouts**
> List[Workout] list_workouts(user_id, year=year, month=month, day=day)

List workouts for a user

Retrieve all workouts for a specific user with optional date filtering

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.workout import Workout
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
    api_instance = trackbot_models.WorkoutsApi(api_client)
    user_id = 1 # int | User ID
    year = '2024' # str | Filter by year (YYYY format) (optional)
    month = '01' # str | Filter by month (MM format) (optional)
    day = '15' # str | Filter by day (DD format) (optional)

    try:
        # List workouts for a user
        api_response = api_instance.list_workouts(user_id, year=year, month=month, day=day)
        print("The response of WorkoutsApi->list_workouts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorkoutsApi->list_workouts: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 
 **year** | **str**| Filter by year (YYYY format) | [optional] 
 **month** | **str**| Filter by month (MM format) | [optional] 
 **day** | **str**| Filter by day (DD format) | [optional] 

### Return type

[**List[Workout]**](Workout.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of workouts retrieved successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_workout**
> Workout update_workout(workout_id, update_workout_request)

Update workout

Update an existing workout

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.update_workout_request import UpdateWorkoutRequest
from trackbot_models.models.workout import Workout
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
    api_instance = trackbot_models.WorkoutsApi(api_client)
    workout_id = 1 # int | Workout ID
    update_workout_request = trackbot_models.UpdateWorkoutRequest() # UpdateWorkoutRequest | 

    try:
        # Update workout
        api_response = api_instance.update_workout(workout_id, update_workout_request)
        print("The response of WorkoutsApi->update_workout:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WorkoutsApi->update_workout: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workout_id** | **int**| Workout ID | 
 **update_workout_request** | [**UpdateWorkoutRequest**](UpdateWorkoutRequest.md)|  | 

### Return type

[**Workout**](Workout.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Workout updated successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

