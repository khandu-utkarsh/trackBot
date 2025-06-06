# CreateUserResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | ID of the created user. Created by the database. | 
**email** | **str** | User&#39;s email address. This is the primary key for the user and obtained from the Google Auth. | 

## Example

```python
from trackbot_client.models.create_user_response import CreateUserResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CreateUserResponse from a JSON string
create_user_response_instance = CreateUserResponse.from_json(json)
# print the JSON string representation of the object
print(CreateUserResponse.to_json())

# convert the object into a dict
create_user_response_dict = create_user_response_instance.to_dict()
# create an instance of CreateUserResponse from a dict
create_user_response_from_dict = CreateUserResponse.from_dict(create_user_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


