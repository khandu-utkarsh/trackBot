# Conversation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the conversation. Created by the database. | 
**user_id** | **int** | ID of the user who owns this conversation. Obtained from the user table. | 
**title** | **str** | Title of the conversation. This is the title of the conversation. | 
**created_at** | **datetime** | Timestamp when the conversation was created. Created by the database. | 

## Example

```python
from trackbot_client.models.conversation import Conversation

# TODO update the JSON string below
json = "{}"
# create an instance of Conversation from a JSON string
conversation_instance = Conversation.from_json(json)
# print the JSON string representation of the object
print(Conversation.to_json())

# convert the object into a dict
conversation_dict = conversation_instance.to_dict()
# create an instance of Conversation from a dict
conversation_from_dict = Conversation.from_dict(conversation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


