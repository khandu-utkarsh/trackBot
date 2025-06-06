# MessagesApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createMessage**](#createmessage) | **POST** /conversations/{conversationId}/messages | Create a new message|
|[**deleteMessage**](#deletemessage) | **DELETE** /messages/{messageId} | Delete message|
|[**getMessageById**](#getmessagebyid) | **GET** /messages/{messageId} | Get message by ID|
|[**listMessages**](#listmessages) | **GET** /conversations/{conversationId}/messages | List messages in a conversation|

# **createMessage**
> CreateMessageResponse createMessage(createMessageRequest)

Send a new message in a conversation

### Example

```typescript
import {
    MessagesApi,
    Configuration,
    CreateMessageRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let conversationId: number; //Conversation ID (default to undefined)
let createMessageRequest: CreateMessageRequest; //

const { status, data } = await apiInstance.createMessage(
    conversationId,
    createMessageRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createMessageRequest** | **CreateMessageRequest**|  | |
| **conversationId** | [**number**] | Conversation ID | defaults to undefined|


### Return type

**CreateMessageResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Message created successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteMessage**
> deleteMessage()

Delete a message from a conversation

### Example

```typescript
import {
    MessagesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let messageId: number; //Message ID (default to undefined)

const { status, data } = await apiInstance.deleteMessage(
    messageId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **messageId** | [**number**] | Message ID | defaults to undefined|


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

# **getMessageById**
> Message getMessageById()

Retrieve a specific message by its ID

### Example

```typescript
import {
    MessagesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let messageId: number; //Message ID (default to undefined)

const { status, data } = await apiInstance.getMessageById(
    messageId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **messageId** | [**number**] | Message ID | defaults to undefined|


### Return type

**Message**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Message details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listMessages**
> Array<Message> listMessages()

Retrieve all messages in a specific conversation

### Example

```typescript
import {
    MessagesApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let conversationId: number; //Conversation ID (default to undefined)

const { status, data } = await apiInstance.listMessages(
    conversationId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **conversationId** | [**number**] | Conversation ID | defaults to undefined|


### Return type

**Array<Message>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of messages retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

