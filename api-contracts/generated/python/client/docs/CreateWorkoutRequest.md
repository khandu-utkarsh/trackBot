# CreateWorkoutRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **int** | ID of the user creating the workout | [optional] 

## Example

```python
from trackbot_client.models.create_workout_request import CreateWorkoutRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateWorkoutRequest from a JSON string
create_workout_request_instance = CreateWorkoutRequest.from_json(json)
# print the JSON string representation of the object
print(CreateWorkoutRequest.to_json())

# convert the object into a dict
create_workout_request_dict = create_workout_request_instance.to_dict()
# create an instance of CreateWorkoutRequest from a dict
create_workout_request_from_dict = CreateWorkoutRequest.from_dict(create_workout_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


