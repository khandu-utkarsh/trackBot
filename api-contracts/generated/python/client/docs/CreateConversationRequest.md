# CreateConversationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**title** | **str** | Title of the conversation | 

## Example

```python
from trackbot_client.models.create_conversation_request import CreateConversationRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateConversationRequest from a JSON string
create_conversation_request_instance = CreateConversationRequest.from_json(json)
# print the JSON string representation of the object
print(CreateConversationRequest.to_json())

# convert the object into a dict
create_conversation_request_dict = create_conversation_request_instance.to_dict()
# create an instance of CreateConversationRequest from a dict
create_conversation_request_from_dict = CreateConversationRequest.from_dict(create_conversation_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


