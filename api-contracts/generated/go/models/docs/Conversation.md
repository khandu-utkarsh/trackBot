# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the conversation. Created by the database. | 
**UserId** | **int64** | ID of the user who owns this conversation. Obtained from the user table. | 
**Title** | **string** | Title of the conversation. This is the title of the conversation. | 
**CreatedAt** | **time.Time** | Timestamp when the conversation was created. Created by the database. | 

## Methods

### NewConversation

`func NewConversation(id int64, userId int64, title string, createdAt time.Time, ) *Conversation`

NewConversation instantiates a new Conversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithDefaults

`func NewConversationWithDefaults() *Conversation`

NewConversationWithDefaults instantiates a new Conversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Conversation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conversation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conversation) SetId(v int64)`

SetId sets Id field to given value.


### GetUserId

`func (o *Conversation) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Conversation) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Conversation) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetTitle

`func (o *Conversation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Conversation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Conversation) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCreatedAt

`func (o *Conversation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Conversation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Conversation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


