# Conversation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the conversation. Created by the database. | [default to undefined]
**user_id** | **number** | ID of the user who owns this conversation. Obtained from the user table. | [default to undefined]
**title** | **string** | Title of the conversation. This is the title of the conversation. | [default to undefined]
**created_at** | **string** | Timestamp when the conversation was created. Created by the database. | [default to undefined]

## Example

```typescript
import { Conversation } from '@trackbot-app/api-client';

const instance: Conversation = {
    id,
    user_id,
    title,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
