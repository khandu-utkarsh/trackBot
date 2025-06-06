# BaseExercise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the exercise. Created by the database. | 
**WorkoutId** | **int64** | ID of the workout this exercise belongs to. Obtained from the workout table. | 
**Name** | **string** | Name of the exercise | 
**Type** | [**ExerciseType**](ExerciseType.md) |  | 
**Notes** | Pointer to **string** | Additional notes about the exercise | [optional] 
**CreatedAt** | **time.Time** | Timestamp when the exercise was created. Created by the database. | 

## Methods

### NewBaseExercise

`func NewBaseExercise(id int64, workoutId int64, name string, type_ ExerciseType, createdAt time.Time, ) *BaseExercise`

NewBaseExercise instantiates a new BaseExercise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseExerciseWithDefaults

`func NewBaseExerciseWithDefaults() *BaseExercise`

NewBaseExerciseWithDefaults instantiates a new BaseExercise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseExercise) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseExercise) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseExercise) SetId(v int64)`

SetId sets Id field to given value.


### GetWorkoutId

`func (o *BaseExercise) GetWorkoutId() int64`

GetWorkoutId returns the WorkoutId field if non-nil, zero value otherwise.

### GetWorkoutIdOk

`func (o *BaseExercise) GetWorkoutIdOk() (*int64, bool)`

GetWorkoutIdOk returns a tuple with the WorkoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutId

`func (o *BaseExercise) SetWorkoutId(v int64)`

SetWorkoutId sets WorkoutId field to given value.


### GetName

`func (o *BaseExercise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseExercise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseExercise) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BaseExercise) GetType() ExerciseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseExercise) GetTypeOk() (*ExerciseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseExercise) SetType(v ExerciseType)`

SetType sets Type field to given value.


### GetNotes

`func (o *BaseExercise) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BaseExercise) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BaseExercise) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BaseExercise) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BaseExercise) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseExercise) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseExercise) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


