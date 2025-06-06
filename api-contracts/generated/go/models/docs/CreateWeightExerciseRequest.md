# CreateWeightExerciseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the exercise | 
**Type** | **string** |  | 
**Notes** | Pointer to **string** | Additional notes | [optional] 
**Sets** | **int32** | Number of sets | 
**Reps** | **int32** | Number of repetitions per set | 
**Weight** | **float32** | Weight in kilograms | 

## Methods

### NewCreateWeightExerciseRequest

`func NewCreateWeightExerciseRequest(name string, type_ string, sets int32, reps int32, weight float32, ) *CreateWeightExerciseRequest`

NewCreateWeightExerciseRequest instantiates a new CreateWeightExerciseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWeightExerciseRequestWithDefaults

`func NewCreateWeightExerciseRequestWithDefaults() *CreateWeightExerciseRequest`

NewCreateWeightExerciseRequestWithDefaults instantiates a new CreateWeightExerciseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWeightExerciseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWeightExerciseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWeightExerciseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateWeightExerciseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWeightExerciseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWeightExerciseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *CreateWeightExerciseRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateWeightExerciseRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateWeightExerciseRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateWeightExerciseRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSets

`func (o *CreateWeightExerciseRequest) GetSets() int32`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *CreateWeightExerciseRequest) GetSetsOk() (*int32, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *CreateWeightExerciseRequest) SetSets(v int32)`

SetSets sets Sets field to given value.


### GetReps

`func (o *CreateWeightExerciseRequest) GetReps() int32`

GetReps returns the Reps field if non-nil, zero value otherwise.

### GetRepsOk

`func (o *CreateWeightExerciseRequest) GetRepsOk() (*int32, bool)`

GetRepsOk returns a tuple with the Reps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReps

`func (o *CreateWeightExerciseRequest) SetReps(v int32)`

SetReps sets Reps field to given value.


### GetWeight

`func (o *CreateWeightExerciseRequest) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateWeightExerciseRequest) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateWeightExerciseRequest) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


