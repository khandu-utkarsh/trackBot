# CreateCardioExerciseRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the exercise | 
**type** | **str** |  | 
**notes** | **str** | Additional notes | [optional] 
**distance** | **float** | Distance in meters | 
**duration** | **int** | Duration in seconds | 

## Example

```python
from trackbot_models.models.create_cardio_exercise_request import CreateCardioExerciseRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateCardioExerciseRequest from a JSON string
create_cardio_exercise_request_instance = CreateCardioExerciseRequest.from_json(json)
# print the JSON string representation of the object
print CreateCardioExerciseRequest.to_json()

# convert the object into a dict
create_cardio_exercise_request_dict = create_cardio_exercise_request_instance.to_dict()
# create an instance of CreateCardioExerciseRequest from a dict
create_cardio_exercise_request_from_dict = CreateCardioExerciseRequest.from_dict(create_cardio_exercise_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


