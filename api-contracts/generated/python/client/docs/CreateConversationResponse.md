# CreateConversationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | ID of the created conversation. Created by the database. | 

## Example

```python
from trackbot_client.models.create_conversation_response import CreateConversationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CreateConversationResponse from a JSON string
create_conversation_response_instance = CreateConversationResponse.from_json(json)
# print the JSON string representation of the object
print(CreateConversationResponse.to_json())

# convert the object into a dict
create_conversation_response_dict = create_conversation_response_instance.to_dict()
# create an instance of CreateConversationResponse from a dict
create_conversation_response_from_dict = CreateConversationResponse.from_dict(create_conversation_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


