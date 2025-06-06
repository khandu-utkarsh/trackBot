# UpdateExerciseRequest


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
from trackbot_models.models.update_exercise_request import UpdateExerciseRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateExerciseRequest from a JSON string
update_exercise_request_instance = UpdateExerciseRequest.from_json(json)
# print the JSON string representation of the object
print UpdateExerciseRequest.to_json()

# convert the object into a dict
update_exercise_request_dict = update_exercise_request_instance.to_dict()
# create an instance of UpdateExerciseRequest from a dict
update_exercise_request_from_dict = UpdateExerciseRequest.from_dict(update_exercise_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


