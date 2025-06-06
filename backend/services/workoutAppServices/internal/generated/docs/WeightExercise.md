# WeightExercise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the exercise. Created by the database. | 
**WorkoutId** | **int64** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**Name** | **string** | Name of the exercise | 
**Type** | **string** |  | 
**Notes** | Pointer to **string** | Additional notes about the exercise | [optional] 
**CreatedAt** | **time.Time** | Timestamp when the exercise was created. Created by the database. | 
**Sets** | **int32** | Number of sets | 
**Reps** | **int32** | Number of repetitions per set | 
**Weight** | **float32** | Weight in kilograms | 

## Methods

### NewWeightExercise

`func NewWeightExercise(id int64, workoutId int64, name string, type_ string, createdAt time.Time, sets int32, reps int32, weight float32, ) *WeightExercise`

NewWeightExercise instantiates a new WeightExercise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeightExerciseWithDefaults

`func NewWeightExerciseWithDefaults() *WeightExercise`

NewWeightExerciseWithDefaults instantiates a new WeightExercise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WeightExercise) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WeightExercise) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WeightExercise) SetId(v int64)`

SetId sets Id field to given value.


### GetWorkoutId

`func (o *WeightExercise) GetWorkoutId() int64`

GetWorkoutId returns the WorkoutId field if non-nil, zero value otherwise.

### GetWorkoutIdOk

`func (o *WeightExercise) GetWorkoutIdOk() (*int64, bool)`

GetWorkoutIdOk returns a tuple with the WorkoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutId

`func (o *WeightExercise) SetWorkoutId(v int64)`

SetWorkoutId sets WorkoutId field to given value.


### GetName

`func (o *WeightExercise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WeightExercise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WeightExercise) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WeightExercise) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeightExercise) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeightExercise) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *WeightExercise) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *WeightExercise) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *WeightExercise) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *WeightExercise) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WeightExercise) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WeightExercise) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WeightExercise) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSets

`func (o *WeightExercise) GetSets() int32`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *WeightExercise) GetSetsOk() (*int32, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *WeightExercise) SetSets(v int32)`

SetSets sets Sets field to given value.


### GetReps

`func (o *WeightExercise) GetReps() int32`

GetReps returns the Reps field if non-nil, zero value otherwise.

### GetRepsOk

`func (o *WeightExercise) GetRepsOk() (*int32, bool)`

GetRepsOk returns a tuple with the Reps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReps

`func (o *WeightExercise) SetReps(v int32)`

SetReps sets Reps field to given value.


### GetWeight

`func (o *WeightExercise) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WeightExercise) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WeightExercise) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


