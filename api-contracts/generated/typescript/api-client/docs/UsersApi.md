# UsersApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createUser**](#createuser) | **POST** /users | Create a new user|
|[**deleteUser**](#deleteuser) | **DELETE** /users/{userId} | Delete user|
|[**getUserByEmail**](#getuserbyemail) | **GET** /users/email/{email} | Get user by email|
|[**getUserById**](#getuserbyid) | **GET** /users/{userId} | Get user by ID|
|[**listUsers**](#listusers) | **GET** /users | List all users|

# **createUser**
> User createUser(createUserRequest)

Create a new user account

### Example

```typescript
import {
    UsersApi,
    Configuration,
    CreateUserRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new UsersApi(configuration);

let createUserRequest: CreateUserRequest; //

const { status, data } = await apiInstance.createUser(
    createUserRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createUserRequest** | **CreateUserRequest**|  | |


### Return type

**User**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | User created successfully |  -  |
|**200** | User already exists |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**409** | Conflict with current state of resource |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteUser**
> deleteUser()

Delete a user and all associated data

### Example

```typescript
import {
    UsersApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new UsersApi(configuration);

let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.deleteUser(
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **userId** | [**number**] | User ID | defaults to undefined|


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

# **getUserByEmail**
> User getUserByEmail()

Retrieve a user by their email address

### Example

```typescript
import {
    UsersApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new UsersApi(configuration);

let email: string; //User email address (default to undefined)

const { status, data } = await apiInstance.getUserByEmail(
    email
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **email** | [**string**] | User email address | defaults to undefined|


### Return type

**User**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | User details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getUserById**
> User getUserById()

Retrieve a specific user by their ID

### Example

```typescript
import {
    UsersApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new UsersApi(configuration);

let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.getUserById(
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**User**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | User details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listUsers**
> Array<User> listUsers()

Retrieve a list of all users in the system

### Example

```typescript
import {
    UsersApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new UsersApi(configuration);

const { status, data } = await apiInstance.listUsers();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<User>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of users retrieved successfully |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

