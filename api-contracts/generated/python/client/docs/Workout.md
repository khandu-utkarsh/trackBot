# Workout


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the workout. Created by the database. | 
**user_id** | **int** | ID of the user who owns this workout. Obtained from the user table. | 
**created_at** | **datetime** | Timestamp when the workout was created. Created by the database. | 
**updated_at** | **datetime** | Timestamp when the workout was last updated. Created by the database. | 

## Example

```python
from trackbot_client.models.workout import Workout

# TODO update the JSON string below
json = "{}"
# create an instance of Workout from a JSON string
workout_instance = Workout.from_json(json)
# print the JSON string representation of the object
print(Workout.to_json())

# convert the object into a dict
workout_dict = workout_instance.to_dict()
# create an instance of Workout from a dict
workout_from_dict = Workout.from_dict(workout_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


