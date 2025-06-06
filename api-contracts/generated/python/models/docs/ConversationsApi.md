# trackbot_models.ConversationsApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_conversation**](ConversationsApi.md#create_conversation) | **POST** /users/{userId}/conversations | Create a new conversation
[**delete_conversation**](ConversationsApi.md#delete_conversation) | **DELETE** /conversations/{conversationId} | Delete conversation
[**get_conversation_by_id**](ConversationsApi.md#get_conversation_by_id) | **GET** /conversations/{conversationId} | Get conversation by ID
[**list_conversations**](ConversationsApi.md#list_conversations) | **GET** /users/{userId}/conversations | List conversations for a user


# **create_conversation**
> CreateConversationResponse create_conversation(user_id, create_conversation_request)

Create a new conversation

Start a new conversation with the AI assistant

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.create_conversation_request import CreateConversationRequest
from trackbot_models.models.create_conversation_response import CreateConversationResponse
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
    api_instance = trackbot_models.ConversationsApi(api_client)
    user_id = 1 # int | User ID
    create_conversation_request = trackbot_models.CreateConversationRequest() # CreateConversationRequest | 

    try:
        # Create a new conversation
        api_response = api_instance.create_conversation(user_id, create_conversation_request)
        print("The response of ConversationsApi->create_conversation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ConversationsApi->create_conversation: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 
 **create_conversation_request** | [**CreateConversationRequest**](CreateConversationRequest.md)|  | 

### Return type

[**CreateConversationResponse**](CreateConversationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Conversation created successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_conversation**
> delete_conversation(conversation_id)

Delete conversation

Delete a conversation and all associated messages

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
    api_instance = trackbot_models.ConversationsApi(api_client)
    conversation_id = 1 # int | Conversation ID

    try:
        # Delete conversation
        api_instance.delete_conversation(conversation_id)
    except Exception as e:
        print("Exception when calling ConversationsApi->delete_conversation: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_id** | **int**| Conversation ID | 

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

# **get_conversation_by_id**
> Conversation get_conversation_by_id(conversation_id)

Get conversation by ID

Retrieve a specific conversation by its ID

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.conversation import Conversation
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
    api_instance = trackbot_models.ConversationsApi(api_client)
    conversation_id = 1 # int | Conversation ID

    try:
        # Get conversation by ID
        api_response = api_instance.get_conversation_by_id(conversation_id)
        print("The response of ConversationsApi->get_conversation_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ConversationsApi->get_conversation_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_id** | **int**| Conversation ID | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Conversation details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_conversations**
> List[Conversation] list_conversations(user_id)

List conversations for a user

Retrieve all conversations for a specific user

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.conversation import Conversation
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
    api_instance = trackbot_models.ConversationsApi(api_client)
    user_id = 1 # int | User ID

    try:
        # List conversations for a user
        api_response = api_instance.list_conversations(user_id)
        print("The response of ConversationsApi->list_conversations:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ConversationsApi->list_conversations: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **int**| User ID | 

### Return type

[**List[Conversation]**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of conversations retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

