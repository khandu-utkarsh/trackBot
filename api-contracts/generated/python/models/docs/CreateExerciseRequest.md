# CreateExerciseRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the exercise | 
**type** | **str** |  | 
**notes** | **str** | Additional notes | [optional] 
**distance** | **float** | Distance in meters | 
**duration** | **int** | Duration in seconds | 
**sets** | **int** | Number of sets | 
**reps** | **int** | Number of repetitions per set | 
**weight** | **float** | Weight in kilograms | 

## Example

```python
from trackbot_models.models.create_exercise_request import CreateExerciseRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateExerciseRequest from a JSON string
create_exercise_request_instance = CreateExerciseRequest.from_json(json)
# print the JSON string representation of the object
print CreateExerciseRequest.to_json()

# convert the object into a dict
create_exercise_request_dict = create_exercise_request_instance.to_dict()
# create an instance of CreateExerciseRequest from a dict
create_exercise_request_from_dict = CreateExerciseRequest.from_dict(create_exercise_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


