# CardioExercise


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

## Example

```python
from trackbot_client.models.cardio_exercise import CardioExercise

# TODO update the JSON string below
json = "{}"
# create an instance of CardioExercise from a JSON string
cardio_exercise_instance = CardioExercise.from_json(json)
# print the JSON string representation of the object
print(CardioExercise.to_json())

# convert the object into a dict
cardio_exercise_dict = cardio_exercise_instance.to_dict()
# create an instance of CardioExercise from a dict
cardio_exercise_from_dict = CardioExercise.from_dict(cardio_exercise_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


