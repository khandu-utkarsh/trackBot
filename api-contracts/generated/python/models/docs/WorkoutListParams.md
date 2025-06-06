# WorkoutListParams


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **int** | ID of the user to filter workouts | [optional] 
**year** | **str** | Filter by year (YYYY format) | [optional] 
**month** | **str** | Filter by month (MM format) | [optional] 
**day** | **str** | Filter by day (DD format) | [optional] 

## Example

```python
from trackbot_models.models.workout_list_params import WorkoutListParams

# TODO update the JSON string below
json = "{}"
# create an instance of WorkoutListParams from a JSON string
workout_list_params_instance = WorkoutListParams.from_json(json)
# print the JSON string representation of the object
print WorkoutListParams.to_json()

# convert the object into a dict
workout_list_params_dict = workout_list_params_instance.to_dict()
# create an instance of WorkoutListParams from a dict
workout_list_params_from_dict = WorkoutListParams.from_dict(workout_list_params_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


