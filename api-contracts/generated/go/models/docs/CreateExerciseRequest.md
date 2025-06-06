# CreateExerciseRequest

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

### NewCreateExerciseRequest

`func NewCreateExerciseRequest(name string, type_ string, distance float32, duration int32, sets int32, reps int32, weight float32, ) *CreateExerciseRequest`

NewCreateExerciseRequest instantiates a new CreateExerciseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExerciseRequestWithDefaults

`func NewCreateExerciseRequestWithDefaults() *CreateExerciseRequest`

NewCreateExerciseRequestWithDefaults instantiates a new CreateExerciseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExerciseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExerciseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExerciseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateExerciseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateExerciseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateExerciseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *CreateExerciseRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateExerciseRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateExerciseRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateExerciseRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDistance

`func (o *CreateExerciseRequest) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CreateExerciseRequest) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CreateExerciseRequest) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetDuration

`func (o *CreateExerciseRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateExerciseRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateExerciseRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSets

`func (o *CreateExerciseRequest) GetSets() int32`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *CreateExerciseRequest) GetSetsOk() (*int32, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *CreateExerciseRequest) SetSets(v int32)`

SetSets sets Sets field to given value.


### GetReps

`func (o *CreateExerciseRequest) GetReps() int32`

GetReps returns the Reps field if non-nil, zero value otherwise.

### GetRepsOk

`func (o *CreateExerciseRequest) GetRepsOk() (*int32, bool)`

GetRepsOk returns a tuple with the Reps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReps

`func (o *CreateExerciseRequest) SetReps(v int32)`

SetReps sets Reps field to given value.


### GetWeight

`func (o *CreateExerciseRequest) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateExerciseRequest) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateExerciseRequest) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


