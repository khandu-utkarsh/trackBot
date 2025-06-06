# Message


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Unique identifier for the message. Created by the database. | 
**conversation_id** | **int** | ID of the conversation this message belongs to. Obtained from the conversation table. | 
**user_id** | **int** | ID of the user who sent this message. Obtained from the user table. | 
**content** | **str** | Content of the message. This is the message that the user or assistant sends. | 
**message_type** | [**MessageType**](MessageType.md) |  | 
**created_at** | **datetime** | Timestamp when the message was created. Created by the database. | 

## Example

```python
from trackbot_client.models.message import Message

# TODO update the JSON string below
json = "{}"
# create an instance of Message from a JSON string
message_instance = Message.from_json(json)
# print the JSON string representation of the object
print(Message.to_json())

# convert the object into a dict
message_dict = message_instance.to_dict()
# create an instance of Message from a dict
message_from_dict = Message.from_dict(message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


