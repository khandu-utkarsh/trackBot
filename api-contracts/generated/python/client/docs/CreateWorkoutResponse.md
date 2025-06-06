# CreateWorkoutResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | ID of the created workout | 

## Example

```python
from trackbot_client.models.create_workout_response import CreateWorkoutResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CreateWorkoutResponse from a JSON string
create_workout_response_instance = CreateWorkoutResponse.from_json(json)
# print the JSON string representation of the object
print(CreateWorkoutResponse.to_json())

# convert the object into a dict
create_workout_response_dict = create_workout_response_instance.to_dict()
# create an instance of CreateWorkoutResponse from a dict
create_workout_response_from_dict = CreateWorkoutResponse.from_dict(create_workout_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


