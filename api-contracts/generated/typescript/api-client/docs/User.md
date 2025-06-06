# User


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the user. Created by the database. | [default to undefined]
**email** | **string** | User\&#39;s email address. This is the primary key for the user and obtained from the Google Auth. | [default to undefined]
**created_at** | **string** | Timestamp when the user was created. Created by the database. | [default to undefined]

## Example

```typescript
import { User } from '@trackbot-app/api-client';

const instance: User = {
    id,
    email,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
