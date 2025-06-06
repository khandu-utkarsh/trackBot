# CreateWorkoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | ID of the user creating the workout | [optional] 

## Methods

### NewCreateWorkoutRequest

`func NewCreateWorkoutRequest() *CreateWorkoutRequest`

NewCreateWorkoutRequest instantiates a new CreateWorkoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkoutRequestWithDefaults

`func NewCreateWorkoutRequestWithDefaults() *CreateWorkoutRequest`

NewCreateWorkoutRequestWithDefaults instantiates a new CreateWorkoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateWorkoutRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateWorkoutRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateWorkoutRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateWorkoutRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


