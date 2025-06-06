# CreateMessageRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**content** | **string** | Content of the message. This is the message that the user or assistant sends. | [default to undefined]
**message_type** | [**MessageType**](MessageType.md) |  | [default to undefined]

## Example

```typescript
import { CreateMessageRequest } from '@trackbot-app/api-client';

const instance: CreateMessageRequest = {
    content,
    message_type,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
