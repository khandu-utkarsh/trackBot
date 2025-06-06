# Exercise


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the exercise. Created by the database. | 
**workout_id** | **int** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**name** | **str** | Name of the exercise | 
**type** | **str** |  | 
**notes** | **str** | Additional notes about the exercise | [optional] 
**created_at** | **datetime** | Timestamp when the exercise was created. Created by the database. | 
**distance** | **float** | Distance covered in meters | 
**duration** | **int** | Duration in seconds | 
**sets** | **int** | Number of sets | 
**reps** | **int** | Number of repetitions per set | 
**weight** | **float** | Weight in kilograms | 

## Example

```python
from trackbot_models.models.exercise import Exercise

# TODO update the JSON string below
json = "{}"
# create an instance of Exercise from a JSON string
exercise_instance = Exercise.from_json(json)
# print the JSON string representation of the object
print Exercise.to_json()

# convert the object into a dict
exercise_dict = exercise_instance.to_dict()
# create an instance of Exercise from a dict
exercise_from_dict = Exercise.from_dict(exercise_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


