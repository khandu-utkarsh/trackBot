# CreateWeightExerciseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name of the exercise | [default to undefined]
**type** | **string** |  | [default to undefined]
**notes** | **string** | Additional notes | [optional] [default to undefined]
**sets** | **number** | Number of sets | [default to undefined]
**reps** | **number** | Number of repetitions per set | [default to undefined]
**weight** | **number** | Weight in kilograms | [default to undefined]

## Example

```typescript
import { CreateWeightExerciseRequest } from '@trackbot-app/api-client';

const instance: CreateWeightExerciseRequest = {
    name,
    type,
    notes,
    sets,
    reps,
    weight,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
