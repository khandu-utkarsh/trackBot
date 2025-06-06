# WeightExercise


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the exercise. Created by the database. | 
**workout_id** | **int** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**name** | **str** | Name of the exercise | 
**type** | **str** |  | 
**notes** | **str** | Additional notes about the exercise | [optional] 
**created_at** | **datetime** | Timestamp when the exercise was created. Created by the database. | 
**sets** | **int** | Number of sets | 
**reps** | **int** | Number of repetitions per set | 
**weight** | **float** | Weight in kilograms | 

## Example

```python
from trackbot_client.models.weight_exercise import WeightExercise

# TODO update the JSON string below
json = "{}"
# create an instance of WeightExercise from a JSON string
weight_exercise_instance = WeightExercise.from_json(json)
# print the JSON string representation of the object
print(WeightExercise.to_json())

# convert the object into a dict
weight_exercise_dict = weight_exercise_instance.to_dict()
# create an instance of WeightExercise from a dict
weight_exercise_from_dict = WeightExercise.from_dict(weight_exercise_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


