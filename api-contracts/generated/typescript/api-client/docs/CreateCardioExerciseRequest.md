# CreateCardioExerciseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name of the exercise | [default to undefined]
**type** | **string** |  | [default to undefined]
**notes** | **string** | Additional notes | [optional] [default to undefined]
**distance** | **number** | Distance in meters | [default to undefined]
**duration** | **number** | Duration in seconds | [default to undefined]

## Example

```typescript
import { CreateCardioExerciseRequest } from '@trackbot-app/api-client';

const instance: CreateCardioExerciseRequest = {
    name,
    type,
    notes,
    distance,
    duration,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
