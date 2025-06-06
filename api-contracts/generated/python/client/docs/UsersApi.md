# trackbot_client.UsersApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_user**](UsersApi.md#create_user) | **POST** /users | Create a new user
[**delete_user**](UsersApi.md#delete_user) | **DELETE** /users/{userId} | Delete user
[**get_user_by_email**](UsersApi.md#get_user_by_email) | **GET** /users/email/{email} | Get user by email
[**get_user_by_id**](UsersApi.md#get_user_by_id) | **GET** /users/{userId} | Get user by ID
[**list_users**](UsersApi.md#list_users) | **GET** /users | List all users


# **create_user**
> User create_user(create_user_request)

Create a new user

Create a new user account

### Example


```python
import trackbot_client
from trackbot_client.models.create_user_request import CreateUserRequest
from trackbot_client.models.user import User
from trackbot_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_client.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_client.UsersApi(api_client)
    create_user_request = trackbot_client.CreateUserRequest() # CreateUserRequest | 

    try:
        # Create a new user
        api_response = api_instance.create_user(create_user_request)
        print("The response of UsersApi->create_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->create_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_user_request** | [**CreateUserRequest**](CreateUserRequest.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | User created successfully |  -  |
**200** | User already exists |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**409** | Conflict with current state of resource |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_user**
> delete_user(user_id)

Delete user

Delete a user and all associated data

### Example


```python
import trackbot_client
from trackbot_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_client.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_client.UsersApi(api_client)
    user_id = 1 # int | User ID

    try:
        # Delete user
        api_instance.delete_user(user_id)
    except Exception as e:
        print("Exception when calling UsersApi->delete_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 

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

# **get_user_by_email**
> User get_user_by_email(email)

Get user by email

Retrieve a user by their email address

### Example


```python
import trackbot_client
from trackbot_client.models.user import User
from trackbot_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_client.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_client.UsersApi(api_client)
    email = 'user@example.com' # str | User email address

    try:
        # Get user by email
        api_response = api_instance.get_user_by_email(email)
        print("The response of UsersApi->get_user_by_email:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_user_by_email: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **str**| User email address | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user_by_id**
> User get_user_by_id(user_id)

Get user by ID

Retrieve a specific user by their ID

### Example


```python
import trackbot_client
from trackbot_client.models.user import User
from trackbot_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_client.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_client.UsersApi(api_client)
    user_id = 1 # int | User ID

    try:
        # Get user by ID
        api_response = api_instance.get_user_by_id(user_id)
        print("The response of UsersApi->get_user_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_user_by_id: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_users**
> List[User] list_users()

List all users

Retrieve a list of all users in the system

### Example


```python
import trackbot_client
from trackbot_client.models.user import User
from trackbot_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = trackbot_client.Configuration(
    host = "http://localhost:8080/api/v1"
)


# Enter a context with an instance of the API client
with trackbot_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = trackbot_client.UsersApi(api_client)

    try:
        # List all users
        api_response = api_instance.list_users()
        print("The response of UsersApi->list_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->list_users: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**List[User]**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of users retrieved successfully |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

