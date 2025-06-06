# Message


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the message. Created by the database. | [default to undefined]
**conversation_id** | **number** | ID of the conversation this message belongs to. Obtained from the conversation table. | [default to undefined]
**user_id** | **number** | ID of the user who sent this message. Obtained from the user table. | [default to undefined]
**content** | **string** | Content of the message. This is the message that the user or assistant sends. | [default to undefined]
**message_type** | [**MessageType**](MessageType.md) |  | [default to undefined]
**created_at** | **string** | Timestamp when the message was created. Created by the database. | [default to undefined]

## Example

```typescript
import { Message } from '@trackbot-app/api-client';

const instance: Message = {
    id,
    conversation_id,
    user_id,
    content,
    message_type,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
