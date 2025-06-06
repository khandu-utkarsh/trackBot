# CreateUserResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | ID of the created user. Created by the database. | [default to undefined]
**email** | **string** | User\&#39;s email address. This is the primary key for the user and obtained from the Google Auth. | [default to undefined]

## Example

```typescript
import { CreateUserResponse } from '@trackbot-app/api-client';

const instance: CreateUserResponse = {
    id,
    email,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
