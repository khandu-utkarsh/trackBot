# CreateCardioExerciseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the exercise | 
**Type** | **string** |  | 
**Notes** | Pointer to **string** | Additional notes | [optional] 
**Distance** | **float32** | Distance in meters | 
**Duration** | **int32** | Duration in seconds | 

## Methods

### NewCreateCardioExerciseRequest

`func NewCreateCardioExerciseRequest(name string, type_ string, distance float32, duration int32, ) *CreateCardioExerciseRequest`

NewCreateCardioExerciseRequest instantiates a new CreateCardioExerciseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCardioExerciseRequestWithDefaults

`func NewCreateCardioExerciseRequestWithDefaults() *CreateCardioExerciseRequest`

NewCreateCardioExerciseRequestWithDefaults instantiates a new CreateCardioExerciseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCardioExerciseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCardioExerciseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCardioExerciseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateCardioExerciseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCardioExerciseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCardioExerciseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *CreateCardioExerciseRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateCardioExerciseRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateCardioExerciseRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateCardioExerciseRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDistance

`func (o *CreateCardioExerciseRequest) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CreateCardioExerciseRequest) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CreateCardioExerciseRequest) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetDuration

`func (o *CreateCardioExerciseRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateCardioExerciseRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateCardioExerciseRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


