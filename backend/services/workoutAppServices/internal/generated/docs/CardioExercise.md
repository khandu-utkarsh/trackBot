# CardioExercise

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

## Methods

### NewCardioExercise

`func NewCardioExercise(id int64, workoutId int64, name string, type_ string, createdAt time.Time, distance float32, duration int32, ) *CardioExercise`

NewCardioExercise instantiates a new CardioExercise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardioExerciseWithDefaults

`func NewCardioExerciseWithDefaults() *CardioExercise`

NewCardioExerciseWithDefaults instantiates a new CardioExercise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardioExercise) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardioExercise) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardioExercise) SetId(v int64)`

SetId sets Id field to given value.


### GetWorkoutId

`func (o *CardioExercise) GetWorkoutId() int64`

GetWorkoutId returns the WorkoutId field if non-nil, zero value otherwise.

### GetWorkoutIdOk

`func (o *CardioExercise) GetWorkoutIdOk() (*int64, bool)`

GetWorkoutIdOk returns a tuple with the WorkoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutId

`func (o *CardioExercise) SetWorkoutId(v int64)`

SetWorkoutId sets WorkoutId field to given value.


### GetName

`func (o *CardioExercise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardioExercise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardioExercise) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CardioExercise) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardioExercise) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardioExercise) SetType(v string)`

SetType sets Type field to given value.


### GetNotes

`func (o *CardioExercise) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CardioExercise) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CardioExercise) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CardioExercise) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CardioExercise) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CardioExercise) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CardioExercise) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDistance

`func (o *CardioExercise) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CardioExercise) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CardioExercise) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetDuration

`func (o *CardioExercise) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CardioExercise) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CardioExercise) SetDuration(v int32)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


