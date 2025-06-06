# WorkoutListParams


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **number** | ID of the user to filter workouts | [optional] [default to undefined]
**year** | **string** | Filter by year (YYYY format) | [optional] [default to undefined]
**month** | **string** | Filter by month (MM format) | [optional] [default to undefined]
**day** | **string** | Filter by day (DD format) | [optional] [default to undefined]

## Example

```typescript
import { WorkoutListParams } from '@trackbot-app/api-client';

const instance: WorkoutListParams = {
    user_id,
    year,
    month,
    day,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
