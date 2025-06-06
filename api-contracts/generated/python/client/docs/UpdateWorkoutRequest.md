# UpdateWorkoutRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **int** | ID of the user who owns the workout | [optional] 

## Example

```python
from trackbot_client.models.update_workout_request import UpdateWorkoutRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateWorkoutRequest from a JSON string
update_workout_request_instance = UpdateWorkoutRequest.from_json(json)
# print the JSON string representation of the object
print(UpdateWorkoutRequest.to_json())

# convert the object into a dict
update_workout_request_dict = update_workout_request_instance.to_dict()
# create an instance of UpdateWorkoutRequest from a dict
update_workout_request_from_dict = UpdateWorkoutRequest.from_dict(update_workout_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


