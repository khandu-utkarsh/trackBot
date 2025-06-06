# User


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the user. Created by the database. | 
**email** | **str** | User&#39;s email address. This is the primary key for the user and obtained from the Google Auth. | 
**created_at** | **datetime** | Timestamp when the user was created. Created by the database. | 

## Example

```python
from trackbot_models.models.user import User

# TODO update the JSON string below
json = "{}"
# create an instance of User from a JSON string
user_instance = User.from_json(json)
# print the JSON string representation of the object
print User.to_json()

# convert the object into a dict
user_dict = user_instance.to_dict()
# create an instance of User from a dict
user_from_dict = User.from_dict(user_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


