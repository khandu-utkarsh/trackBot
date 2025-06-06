# Workout


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the workout. Created by the database. | [default to undefined]
**user_id** | **number** | ID of the user who owns this workout. Obtained from the user table. | [default to undefined]
**created_at** | **string** | Timestamp when the workout was created. Created by the database. | [default to undefined]
**updated_at** | **string** | Timestamp when the workout was last updated. Created by the database. | [default to undefined]

## Example

```typescript
import { Workout } from '@trackbot-app/api-client';

const instance: Workout = {
    id,
    user_id,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
