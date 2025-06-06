# CreateWeightExerciseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name of the exercise | 
**type** | **str** |  | 
**notes** | **str** | Additional notes | [optional] 
**sets** | **int** | Number of sets | 
**reps** | **int** | Number of repetitions per set | 
**weight** | **float** | Weight in kilograms | 

## Example

```python
from trackbot_client.models.create_weight_exercise_request import CreateWeightExerciseRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateWeightExerciseRequest from a JSON string
create_weight_exercise_request_instance = CreateWeightExerciseRequest.from_json(json)
# print the JSON string representation of the object
print(CreateWeightExerciseRequest.to_json())

# convert the object into a dict
create_weight_exercise_request_dict = create_weight_exercise_request_instance.to_dict()
# create an instance of CreateWeightExerciseRequest from a dict
create_weight_exercise_request_from_dict = CreateWeightExerciseRequest.from_dict(create_weight_exercise_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


