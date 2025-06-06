# WorkoutListParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | ID of the user to filter workouts | [optional] 
**Year** | Pointer to **string** | Filter by year (YYYY format) | [optional] 
**Month** | Pointer to **string** | Filter by month (MM format) | [optional] 
**Day** | Pointer to **string** | Filter by day (DD format) | [optional] 

## Methods

### NewWorkoutListParams

`func NewWorkoutListParams() *WorkoutListParams`

NewWorkoutListParams instantiates a new WorkoutListParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkoutListParamsWithDefaults

`func NewWorkoutListParamsWithDefaults() *WorkoutListParams`

NewWorkoutListParamsWithDefaults instantiates a new WorkoutListParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *WorkoutListParams) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkoutListParams) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkoutListParams) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkoutListParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetYear

`func (o *WorkoutListParams) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *WorkoutListParams) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *WorkoutListParams) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *WorkoutListParams) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *WorkoutListParams) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *WorkoutListParams) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *WorkoutListParams) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *WorkoutListParams) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDay

`func (o *WorkoutListParams) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *WorkoutListParams) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *WorkoutListParams) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *WorkoutListParams) HasDay() bool`

HasDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


