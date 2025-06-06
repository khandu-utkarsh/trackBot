# trackbot_models.MessagesApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_message**](MessagesApi.md#create_message) | **POST** /conversations/{conversationId}/messages | Create a new message
[**delete_message**](MessagesApi.md#delete_message) | **DELETE** /messages/{messageId} | Delete message
[**get_message_by_id**](MessagesApi.md#get_message_by_id) | **GET** /messages/{messageId} | Get message by ID
[**list_messages**](MessagesApi.md#list_messages) | **GET** /conversations/{conversationId}/messages | List messages in a conversation


# **create_message**
> CreateMessageResponse create_message(conversation_id, create_message_request)

Create a new message

Send a new message in a conversation

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.create_message_request import CreateMessageRequest
from trackbot_models.models.create_message_response import CreateMessageResponse
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
    api_instance = trackbot_models.MessagesApi(api_client)
    conversation_id = 1 # int | Conversation ID
    create_message_request = trackbot_models.CreateMessageRequest() # CreateMessageRequest | 

    try:
        # Create a new message
        api_response = api_instance.create_message(conversation_id, create_message_request)
        print("The response of MessagesApi->create_message:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MessagesApi->create_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_id** | **int**| Conversation ID | 
 **create_message_request** | [**CreateMessageRequest**](CreateMessageRequest.md)|  | 

### Return type

[**CreateMessageResponse**](CreateMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Message created successfully |  -  |
**400** | Bad request - invalid input or parameters |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_message**
> delete_message(message_id)

Delete message

Delete a message from a conversation

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
    api_instance = trackbot_models.MessagesApi(api_client)
    message_id = 1 # int | Message ID

    try:
        # Delete message
        api_instance.delete_message(message_id)
    except Exception as e:
        print("Exception when calling MessagesApi->delete_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message_id** | **int**| Message ID | 

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

# **get_message_by_id**
> Message get_message_by_id(message_id)

Get message by ID

Retrieve a specific message by its ID

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.message import Message
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
    api_instance = trackbot_models.MessagesApi(api_client)
    message_id = 1 # int | Message ID

    try:
        # Get message by ID
        api_response = api_instance.get_message_by_id(message_id)
        print("The response of MessagesApi->get_message_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MessagesApi->get_message_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message_id** | **int**| Message ID | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Message details retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_messages**
> List[Message] list_messages(conversation_id)

List messages in a conversation

Retrieve all messages in a specific conversation

### Example

```python
import time
import os
import trackbot_models
from trackbot_models.models.message import Message
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
    api_instance = trackbot_models.MessagesApi(api_client)
    conversation_id = 1 # int | Conversation ID

    try:
        # List messages in a conversation
        api_response = api_instance.list_messages(conversation_id)
        print("The response of MessagesApi->list_messages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MessagesApi->list_messages: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_id** | **int**| Conversation ID | 

### Return type

[**List[Message]**](Message.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of messages retrieved successfully |  -  |
**404** | Resource not found |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

