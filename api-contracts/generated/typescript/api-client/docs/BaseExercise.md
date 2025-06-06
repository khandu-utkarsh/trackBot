# BaseExercise


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the exercise. Created by the database. | [default to undefined]
**workout_id** | **number** | ID of the workout this exercise belongs to. Obtained from the workout table. | [default to undefined]
**name** | **string** | Name of the exercise | [default to undefined]
**type** | [**ExerciseType**](ExerciseType.md) |  | [default to undefined]
**notes** | **string** | Additional notes about the exercise | [optional] [default to undefined]
**created_at** | **string** | Timestamp when the exercise was created. Created by the database. | [default to undefined]

## Example

```typescript
import { BaseExercise } from '@trackbot-app/api-client';

const instance: BaseExercise = {
    id,
    workout_id,
    name,
    type,
    notes,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
