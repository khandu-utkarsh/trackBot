# Exercise


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | Unique identifier for the exercise. Created by the database. | [default to undefined]
**workout_id** | **number** | ID of the workout this exercise belongs to. Obtained from the workout table. | [default to undefined]
**name** | **string** | Name of the exercise | [default to undefined]
**type** | **string** |  | [default to undefined]
**notes** | **string** | Additional notes about the exercise | [optional] [default to undefined]
**created_at** | **string** | Timestamp when the exercise was created. Created by the database. | [default to undefined]
**distance** | **number** | Distance covered in meters | [default to undefined]
**duration** | **number** | Duration in seconds | [default to undefined]
**sets** | **number** | Number of sets | [default to undefined]
**reps** | **number** | Number of repetitions per set | [default to undefined]
**weight** | **number** | Weight in kilograms | [default to undefined]

## Example

```typescript
import { Exercise } from '@trackbot-app/api-client';

const instance: Exercise = {
    id,
    workout_id,
    name,
    type,
    notes,
    created_at,
    distance,
    duration,
    sets,
    reps,
    weight,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
