# BaseExercise


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the exercise. Created by the database. | 
**workout_id** | **int** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**name** | **str** | Name of the exercise | 
**type** | [**ExerciseType**](ExerciseType.md) |  | 
**notes** | **str** | Additional notes about the exercise | [optional] 
**created_at** | **datetime** | Timestamp when the exercise was created. Created by the database. | 

## Example

```python
from trackbot_models.models.base_exercise import BaseExercise

# TODO update the JSON string below
json = "{}"
# create an instance of BaseExercise from a JSON string
base_exercise_instance = BaseExercise.from_json(json)
# print the JSON string representation of the object
print BaseExercise.to_json()

# convert the object into a dict
base_exercise_dict = base_exercise_instance.to_dict()
# create an instance of BaseExercise from a dict
base_exercise_from_dict = BaseExercise.from_dict(base_exercise_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


