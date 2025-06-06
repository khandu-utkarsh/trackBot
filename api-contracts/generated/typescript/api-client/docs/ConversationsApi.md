# ConversationsApi

All URIs are relative to *http://localhost:8080/api/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createConversation**](#createconversation) | **POST** /users/{userId}/conversations | Create a new conversation|
|[**deleteConversation**](#deleteconversation) | **DELETE** /conversations/{conversationId} | Delete conversation|
|[**getConversationById**](#getconversationbyid) | **GET** /conversations/{conversationId} | Get conversation by ID|
|[**listConversations**](#listconversations) | **GET** /users/{userId}/conversations | List conversations for a user|

# **createConversation**
> CreateConversationResponse createConversation(createConversationRequest)

Start a new conversation with the AI assistant

### Example

```typescript
import {
    ConversationsApi,
    Configuration,
    CreateConversationRequest
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ConversationsApi(configuration);

let userId: number; //User ID (default to undefined)
let createConversationRequest: CreateConversationRequest; //

const { status, data } = await apiInstance.createConversation(
    userId,
    createConversationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createConversationRequest** | **CreateConversationRequest**|  | |
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**CreateConversationResponse**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Conversation created successfully |  -  |
|**400** | Bad request - invalid input or parameters |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteConversation**
> deleteConversation()

Delete a conversation and all associated messages

### Example

```typescript
import {
    ConversationsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ConversationsApi(configuration);

let conversationId: number; //Conversation ID (default to undefined)

const { status, data } = await apiInstance.deleteConversation(
    conversationId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **conversationId** | [**number**] | Conversation ID | defaults to undefined|


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

# **getConversationById**
> Conversation getConversationById()

Retrieve a specific conversation by its ID

### Example

```typescript
import {
    ConversationsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ConversationsApi(configuration);

let conversationId: number; //Conversation ID (default to undefined)

const { status, data } = await apiInstance.getConversationById(
    conversationId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **conversationId** | [**number**] | Conversation ID | defaults to undefined|


### Return type

**Conversation**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Conversation details retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listConversations**
> Array<Conversation> listConversations()

Retrieve all conversations for a specific user

### Example

```typescript
import {
    ConversationsApi,
    Configuration
} from '@trackbot-app/api-client';

const configuration = new Configuration();
const apiInstance = new ConversationsApi(configuration);

let userId: number; //User ID (default to undefined)

const { status, data } = await apiInstance.listConversations(
    userId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **userId** | [**number**] | User ID | defaults to undefined|


### Return type

**Array<Conversation>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | List of conversations retrieved successfully |  -  |
|**404** | Resource not found |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

