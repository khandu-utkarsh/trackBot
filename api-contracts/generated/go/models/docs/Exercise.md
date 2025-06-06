# Exercise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the exercise. Created by the database. | 
**WorkoutId** | **int64** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**Name** | **string** | Name of the exercise | 
**Type** | **string** |  | 
**Notes** | Pointer to **string** | Additional notes about the exercise | [optional] 
**CreatedAt** | **time.Time** | Timestamp when the exercise was created. Created by the database. | 
**Distance** | **float32** | Distance covered in meters | 
**Duration** | **int32** | Duration in seconds | 
**Sets** | **int32** | Number of sets | 
**Reps** | **int32** | Number of repetitions per set | 
**Weight** | **float32** | Weight in kilograms | 

## Methods

### NewExercise

`func NewExercise(id int64, workoutId int64, name string, type_ string, createdAt time.Time, distance float32, duration int32, sets int32, reps int32, weight float32, ) *Exercise`

NewExercise instantiates a new Exercise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExerciseWithDefaults

`func NewExerciseWithDefaults() *Exercise`

NewExerciseWithDefaults instantiates a new Exercise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Exercise) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Exercise) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Exercise) SetId(v int64)`

SetId sets Id field to given value.


### GetWorkoutId

`func (o *Exercise) GetWorkoutId() int64`

GetWorkoutId returns the WorkoutId field if non-nil, zero value otherwise.

### GetWorkoutIdOk

`func (o *Exercise) GetWorkoutIdOk() (*int64, bool)`

GetWorkoutIdOk returns a tuple with the WorkoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutId

`func (o *Exercise) SetWorkoutId(v int64)`

SetWorkoutId sets WorkoutId field to given value.


### GetName

`func (o *Exercise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Exercise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Exercise) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Exercise) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Exercise) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Exercise) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *Exercise) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Exercise) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Exercise) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Exercise) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Exercise) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Exercise) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Exercise) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDistance

`func (o *Exercise) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Exercise) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Exercise) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetDuration

`func (o *Exercise) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Exercise) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Exercise) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSets

`func (o *Exercise) GetSets() int32`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *Exercise) GetSetsOk() (*int32, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *Exercise) SetSets(v int32)`

SetSets sets Sets field to given value.


### GetReps

`func (o *Exercise) GetReps() int32`

GetReps returns the Reps field if non-nil, zero value otherwise.

### GetRepsOk

`func (o *Exercise) GetRepsOk() (*int32, bool)`

GetRepsOk returns a tuple with the Reps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReps

`func (o *Exercise) SetReps(v int32)`

SetReps sets Reps field to given value.


### GetWeight

`func (o *Exercise) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Exercise) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Exercise) SetWeight(v float32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


