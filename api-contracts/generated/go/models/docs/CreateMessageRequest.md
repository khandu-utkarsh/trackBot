# CreateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Content of the message. This is the message that the user or assistant sends. | 
**MessageType** | [**MessageType**](MessageType.md) |  | 

## Methods

### NewCreateMessageRequest

`func NewCreateMessageRequest(content string, messageType MessageType, ) *CreateMessageRequest`

NewCreateMessageRequest instantiates a new CreateMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageRequestWithDefaults

`func NewCreateMessageRequestWithDefaults() *CreateMessageRequest`

NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateMessageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateMessageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateMessageRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetMessageType

`func (o *CreateMessageRequest) GetMessageType() MessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *CreateMessageRequest) GetMessageTypeOk() (*MessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *CreateMessageRequest) SetMessageType(v MessageType)`

SetMessageType sets MessageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


