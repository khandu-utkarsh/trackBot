# CreateExerciseResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | ID of the created exercise | 

## Example

```python
from trackbot_models.models.create_exercise_response import CreateExerciseResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CreateExerciseResponse from a JSON string
create_exercise_response_instance = CreateExerciseResponse.from_json(json)
# print the JSON string representation of the object
print CreateExerciseResponse.to_json()

# convert the object into a dict
create_exercise_response_dict = create_exercise_response_instance.to_dict()
# create an instance of CreateExerciseResponse from a dict
create_exercise_response_from_dict = CreateExerciseResponse.from_dict(create_exercise_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


