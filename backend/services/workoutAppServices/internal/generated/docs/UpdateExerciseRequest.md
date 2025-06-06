# UpdateExerciseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the exercise | 
**Type** | **string** |  | 
**Notes** | Pointer to **string** | Additional notes | [optional] 
**Distance** | **float32** | Distance in meters | 
**Duration** | **int32** | Duration in seconds | 
**Sets** | **int32** | Number of sets | 
**Reps** | **int32** | Number of repetitions per set | 
**Weight** | **float32** | Weight in kilograms | 

## Methods

### NewUpdateExerciseRequest

`func NewUpdateExerciseRequest(name string, type_ string, distance float32, duration int32, sets int32, reps int32, weight float32, ) *UpdateExerciseRequest`

NewUpdateExerciseRequest instantiates a new UpdateExerciseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExerciseRequestWithDefaults

`func NewUpdateExerciseRequestWithDefaults() *UpdateExerciseRequest`

NewUpdateExerciseRequestWithDefaults instantiates a new UpdateExerciseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateExerciseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExerciseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExerciseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateExerciseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateExerciseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateExerciseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *UpdateExerciseRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateExerciseRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateExerciseRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateExerciseRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDistance

`func (o *UpdateExerciseRequest) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *UpdateExerciseRequest) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *UpdateExerciseRequest) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetDuration

`func (o *UpdateExerciseRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateExerciseRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateExerciseRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSets

`func (o *UpdateExerciseRequest) GetSets() int32`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *UpdateExerciseRequest) GetSetsOk() (*int32, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *UpdateExerciseRequest) SetSets(v int32)`

SetSets sets Sets field to given value.


### GetReps

`func (o *UpdateExerciseRequest) GetReps() int32`

GetReps returns the Reps field if non-nil, zero value otherwise.

### GetRepsOk

`func (o *UpdateExerciseRequest) GetRepsOk() (*int32, bool)`

GetRepsOk returns a tuple with the Reps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReps

`func (o *UpdateExerciseRequest) SetReps(v int32)`

SetReps sets Reps field to given value.


### GetWeight

`func (o *UpdateExerciseRequest) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateExerciseRequest) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateExerciseRequest) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


