# TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActivityDurations** | Pointer to **[]map[string]interface{}** | The minimum, average, maximum, and total time that Workers spent in each Activity | [optional] 
**EndTime** | Pointer to **NullableTime** | The end of the interval during which these statistics were calculated | [optional] 
**ReservationsAccepted** | Pointer to **NullableInt32** | The total number of Reservations that were accepted | [optional] 
**ReservationsCanceled** | Pointer to **NullableInt32** | The total number of Reservations that were canceled | [optional] 
**ReservationsCreated** | Pointer to **NullableInt32** | The total number of Reservations that were created | [optional] 
**ReservationsRejected** | Pointer to **NullableInt32** | The total number of Reservations that were rejected | [optional] 
**ReservationsRescinded** | Pointer to **NullableInt32** | The total number of Reservations that were rescinded | [optional] 
**ReservationsTimedOut** | Pointer to **NullableInt32** | The total number of Reservations that were timed out | [optional] 
**StartTime** | Pointer to **NullableTime** | The beginning of the interval during which these statistics were calculated | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workers statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Workers | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics

`func NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics() *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics`

NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics instantiates a new TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics`

NewTaskrouterV1WorkspaceWorkerWorkersCumulativeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActivityDurations

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetActivityDurations() []map[string]interface{}`

GetActivityDurations returns the ActivityDurations field if non-nil, zero value otherwise.

### GetActivityDurationsOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetActivityDurationsOk() (*[]map[string]interface{}, bool)`

GetActivityDurationsOk returns a tuple with the ActivityDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDurations

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetActivityDurations(v []map[string]interface{})`

SetActivityDurations sets ActivityDurations field to given value.

### HasActivityDurations

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasActivityDurations() bool`

HasActivityDurations returns a boolean if a field has been set.

### SetActivityDurationsNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetActivityDurationsNil(b bool)`

 SetActivityDurationsNil sets the value for ActivityDurations to be an explicit nil

### UnsetActivityDurations
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetActivityDurations()`

UnsetActivityDurations ensures that no value is present for ActivityDurations, not even an explicit nil
### GetEndTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsAccepted() int32`

GetReservationsAccepted returns the ReservationsAccepted field if non-nil, zero value otherwise.

### GetReservationsAcceptedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsAcceptedOk() (*int32, bool)`

GetReservationsAcceptedOk returns a tuple with the ReservationsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsAccepted(v int32)`

SetReservationsAccepted sets ReservationsAccepted field to given value.

### HasReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsAccepted() bool`

HasReservationsAccepted returns a boolean if a field has been set.

### SetReservationsAcceptedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsAcceptedNil(b bool)`

 SetReservationsAcceptedNil sets the value for ReservationsAccepted to be an explicit nil

### UnsetReservationsAccepted
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsAccepted()`

UnsetReservationsAccepted ensures that no value is present for ReservationsAccepted, not even an explicit nil
### GetReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsCanceled() int32`

GetReservationsCanceled returns the ReservationsCanceled field if non-nil, zero value otherwise.

### GetReservationsCanceledOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsCanceledOk() (*int32, bool)`

GetReservationsCanceledOk returns a tuple with the ReservationsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsCanceled(v int32)`

SetReservationsCanceled sets ReservationsCanceled field to given value.

### HasReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsCanceled() bool`

HasReservationsCanceled returns a boolean if a field has been set.

### SetReservationsCanceledNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsCanceledNil(b bool)`

 SetReservationsCanceledNil sets the value for ReservationsCanceled to be an explicit nil

### UnsetReservationsCanceled
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsCanceled()`

UnsetReservationsCanceled ensures that no value is present for ReservationsCanceled, not even an explicit nil
### GetReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsCreated() int32`

GetReservationsCreated returns the ReservationsCreated field if non-nil, zero value otherwise.

### GetReservationsCreatedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsCreatedOk() (*int32, bool)`

GetReservationsCreatedOk returns a tuple with the ReservationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsCreated(v int32)`

SetReservationsCreated sets ReservationsCreated field to given value.

### HasReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsCreated() bool`

HasReservationsCreated returns a boolean if a field has been set.

### SetReservationsCreatedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsCreatedNil(b bool)`

 SetReservationsCreatedNil sets the value for ReservationsCreated to be an explicit nil

### UnsetReservationsCreated
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsCreated()`

UnsetReservationsCreated ensures that no value is present for ReservationsCreated, not even an explicit nil
### GetReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsRejected() int32`

GetReservationsRejected returns the ReservationsRejected field if non-nil, zero value otherwise.

### GetReservationsRejectedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsRejectedOk() (*int32, bool)`

GetReservationsRejectedOk returns a tuple with the ReservationsRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsRejected(v int32)`

SetReservationsRejected sets ReservationsRejected field to given value.

### HasReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsRejected() bool`

HasReservationsRejected returns a boolean if a field has been set.

### SetReservationsRejectedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsRejectedNil(b bool)`

 SetReservationsRejectedNil sets the value for ReservationsRejected to be an explicit nil

### UnsetReservationsRejected
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsRejected()`

UnsetReservationsRejected ensures that no value is present for ReservationsRejected, not even an explicit nil
### GetReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsRescinded() int32`

GetReservationsRescinded returns the ReservationsRescinded field if non-nil, zero value otherwise.

### GetReservationsRescindedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsRescindedOk() (*int32, bool)`

GetReservationsRescindedOk returns a tuple with the ReservationsRescinded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsRescinded(v int32)`

SetReservationsRescinded sets ReservationsRescinded field to given value.

### HasReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsRescinded() bool`

HasReservationsRescinded returns a boolean if a field has been set.

### SetReservationsRescindedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsRescindedNil(b bool)`

 SetReservationsRescindedNil sets the value for ReservationsRescinded to be an explicit nil

### UnsetReservationsRescinded
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsRescinded()`

UnsetReservationsRescinded ensures that no value is present for ReservationsRescinded, not even an explicit nil
### GetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsTimedOut() int32`

GetReservationsTimedOut returns the ReservationsTimedOut field if non-nil, zero value otherwise.

### GetReservationsTimedOutOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetReservationsTimedOutOk() (*int32, bool)`

GetReservationsTimedOutOk returns a tuple with the ReservationsTimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsTimedOut(v int32)`

SetReservationsTimedOut sets ReservationsTimedOut field to given value.

### HasReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasReservationsTimedOut() bool`

HasReservationsTimedOut returns a boolean if a field has been set.

### SetReservationsTimedOutNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetReservationsTimedOutNil(b bool)`

 SetReservationsTimedOutNil sets the value for ReservationsTimedOut to be an explicit nil

### UnsetReservationsTimedOut
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetReservationsTimedOut()`

UnsetReservationsTimedOut ensures that no value is present for ReservationsTimedOut, not even an explicit nil
### GetStartTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


