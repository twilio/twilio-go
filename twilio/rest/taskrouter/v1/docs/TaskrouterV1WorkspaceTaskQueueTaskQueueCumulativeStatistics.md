# TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AvgTaskAcceptanceTime** | Pointer to **NullableInt32** | The average time in seconds between Task creation and acceptance | [optional] 
**EndTime** | Pointer to **NullableTime** | The end of the interval during which these statistics were calculated | [optional] 
**ReservationsAccepted** | Pointer to **NullableInt32** | The total number of Reservations accepted for Tasks in the TaskQueue | [optional] 
**ReservationsCanceled** | Pointer to **NullableInt32** | The total number of Reservations canceled for Tasks in the TaskQueue | [optional] 
**ReservationsCreated** | Pointer to **NullableInt32** | The total number of Reservations created for Tasks in the TaskQueue | [optional] 
**ReservationsRejected** | Pointer to **NullableInt32** | The total number of Reservations rejected for Tasks in the TaskQueue | [optional] 
**ReservationsRescinded** | Pointer to **NullableInt32** | The total number of Reservations rescinded | [optional] 
**ReservationsTimedOut** | Pointer to **NullableInt32** | The total number of Reservations that timed out for Tasks in the TaskQueue | [optional] 
**SplitByWaitTime** | Pointer to **map[string]interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds | [optional] 
**StartTime** | Pointer to **NullableTime** | The beginning of the interval during which these statistics were calculated | [optional] 
**TaskQueueSid** | Pointer to **NullableString** | The SID of the TaskQueue from which these statistics were calculated | [optional] 
**TasksCanceled** | Pointer to **NullableInt32** | The total number of Tasks canceled in the TaskQueue | [optional] 
**TasksCompleted** | Pointer to **NullableInt32** | The total number of Tasks completed in the TaskQueue | [optional] 
**TasksDeleted** | Pointer to **NullableInt32** | The total number of Tasks deleted in the TaskQueue | [optional] 
**TasksEntered** | Pointer to **NullableInt32** | The total number of Tasks entered into the TaskQueue | [optional] 
**TasksMoved** | Pointer to **NullableInt32** | The total number of Tasks that were moved from one queue to another | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskQueue statistics resource | [optional] 
**WaitDurationInQueueUntilAccepted** | Pointer to **map[string]interface{}** | The relative wait duration statistics for Tasks accepted while in the TaskQueue | [optional] 
**WaitDurationUntilAccepted** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks accepted while in the TaskQueue | [optional] 
**WaitDurationUntilCanceled** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks canceled while in the TaskQueue | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the TaskQueue | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics() *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatisticsWithDefaults() *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetAvgTaskAcceptanceTime() int32`

GetAvgTaskAcceptanceTime returns the AvgTaskAcceptanceTime field if non-nil, zero value otherwise.

### GetAvgTaskAcceptanceTimeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetAvgTaskAcceptanceTimeOk() (*int32, bool)`

GetAvgTaskAcceptanceTimeOk returns a tuple with the AvgTaskAcceptanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetAvgTaskAcceptanceTime(v int32)`

SetAvgTaskAcceptanceTime sets AvgTaskAcceptanceTime field to given value.

### HasAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasAvgTaskAcceptanceTime() bool`

HasAvgTaskAcceptanceTime returns a boolean if a field has been set.

### SetAvgTaskAcceptanceTimeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetAvgTaskAcceptanceTimeNil(b bool)`

 SetAvgTaskAcceptanceTimeNil sets the value for AvgTaskAcceptanceTime to be an explicit nil

### UnsetAvgTaskAcceptanceTime
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetAvgTaskAcceptanceTime()`

UnsetAvgTaskAcceptanceTime ensures that no value is present for AvgTaskAcceptanceTime, not even an explicit nil
### GetEndTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetReservationsAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsAccepted() int32`

GetReservationsAccepted returns the ReservationsAccepted field if non-nil, zero value otherwise.

### GetReservationsAcceptedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsAcceptedOk() (*int32, bool)`

GetReservationsAcceptedOk returns a tuple with the ReservationsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsAccepted(v int32)`

SetReservationsAccepted sets ReservationsAccepted field to given value.

### HasReservationsAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsAccepted() bool`

HasReservationsAccepted returns a boolean if a field has been set.

### SetReservationsAcceptedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsAcceptedNil(b bool)`

 SetReservationsAcceptedNil sets the value for ReservationsAccepted to be an explicit nil

### UnsetReservationsAccepted
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsAccepted()`

UnsetReservationsAccepted ensures that no value is present for ReservationsAccepted, not even an explicit nil
### GetReservationsCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsCanceled() int32`

GetReservationsCanceled returns the ReservationsCanceled field if non-nil, zero value otherwise.

### GetReservationsCanceledOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsCanceledOk() (*int32, bool)`

GetReservationsCanceledOk returns a tuple with the ReservationsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsCanceled(v int32)`

SetReservationsCanceled sets ReservationsCanceled field to given value.

### HasReservationsCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsCanceled() bool`

HasReservationsCanceled returns a boolean if a field has been set.

### SetReservationsCanceledNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsCanceledNil(b bool)`

 SetReservationsCanceledNil sets the value for ReservationsCanceled to be an explicit nil

### UnsetReservationsCanceled
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsCanceled()`

UnsetReservationsCanceled ensures that no value is present for ReservationsCanceled, not even an explicit nil
### GetReservationsCreated

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsCreated() int32`

GetReservationsCreated returns the ReservationsCreated field if non-nil, zero value otherwise.

### GetReservationsCreatedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsCreatedOk() (*int32, bool)`

GetReservationsCreatedOk returns a tuple with the ReservationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCreated

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsCreated(v int32)`

SetReservationsCreated sets ReservationsCreated field to given value.

### HasReservationsCreated

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsCreated() bool`

HasReservationsCreated returns a boolean if a field has been set.

### SetReservationsCreatedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsCreatedNil(b bool)`

 SetReservationsCreatedNil sets the value for ReservationsCreated to be an explicit nil

### UnsetReservationsCreated
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsCreated()`

UnsetReservationsCreated ensures that no value is present for ReservationsCreated, not even an explicit nil
### GetReservationsRejected

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsRejected() int32`

GetReservationsRejected returns the ReservationsRejected field if non-nil, zero value otherwise.

### GetReservationsRejectedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsRejectedOk() (*int32, bool)`

GetReservationsRejectedOk returns a tuple with the ReservationsRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRejected

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsRejected(v int32)`

SetReservationsRejected sets ReservationsRejected field to given value.

### HasReservationsRejected

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsRejected() bool`

HasReservationsRejected returns a boolean if a field has been set.

### SetReservationsRejectedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsRejectedNil(b bool)`

 SetReservationsRejectedNil sets the value for ReservationsRejected to be an explicit nil

### UnsetReservationsRejected
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsRejected()`

UnsetReservationsRejected ensures that no value is present for ReservationsRejected, not even an explicit nil
### GetReservationsRescinded

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsRescinded() int32`

GetReservationsRescinded returns the ReservationsRescinded field if non-nil, zero value otherwise.

### GetReservationsRescindedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsRescindedOk() (*int32, bool)`

GetReservationsRescindedOk returns a tuple with the ReservationsRescinded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRescinded

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsRescinded(v int32)`

SetReservationsRescinded sets ReservationsRescinded field to given value.

### HasReservationsRescinded

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsRescinded() bool`

HasReservationsRescinded returns a boolean if a field has been set.

### SetReservationsRescindedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsRescindedNil(b bool)`

 SetReservationsRescindedNil sets the value for ReservationsRescinded to be an explicit nil

### UnsetReservationsRescinded
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsRescinded()`

UnsetReservationsRescinded ensures that no value is present for ReservationsRescinded, not even an explicit nil
### GetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsTimedOut() int32`

GetReservationsTimedOut returns the ReservationsTimedOut field if non-nil, zero value otherwise.

### GetReservationsTimedOutOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetReservationsTimedOutOk() (*int32, bool)`

GetReservationsTimedOutOk returns a tuple with the ReservationsTimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsTimedOut(v int32)`

SetReservationsTimedOut sets ReservationsTimedOut field to given value.

### HasReservationsTimedOut

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasReservationsTimedOut() bool`

HasReservationsTimedOut returns a boolean if a field has been set.

### SetReservationsTimedOutNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetReservationsTimedOutNil(b bool)`

 SetReservationsTimedOutNil sets the value for ReservationsTimedOut to be an explicit nil

### UnsetReservationsTimedOut
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetReservationsTimedOut()`

UnsetReservationsTimedOut ensures that no value is present for ReservationsTimedOut, not even an explicit nil
### GetSplitByWaitTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetSplitByWaitTime() map[string]interface{}`

GetSplitByWaitTime returns the SplitByWaitTime field if non-nil, zero value otherwise.

### GetSplitByWaitTimeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetSplitByWaitTimeOk() (*map[string]interface{}, bool)`

GetSplitByWaitTimeOk returns a tuple with the SplitByWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByWaitTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetSplitByWaitTime(v map[string]interface{})`

SetSplitByWaitTime sets SplitByWaitTime field to given value.

### HasSplitByWaitTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasSplitByWaitTime() bool`

HasSplitByWaitTime returns a boolean if a field has been set.

### SetSplitByWaitTimeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetSplitByWaitTimeNil(b bool)`

 SetSplitByWaitTimeNil sets the value for SplitByWaitTime to be an explicit nil

### UnsetSplitByWaitTime
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetSplitByWaitTime()`

UnsetSplitByWaitTime ensures that no value is present for SplitByWaitTime, not even an explicit nil
### GetStartTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTaskQueueSid() string`

GetTaskQueueSid returns the TaskQueueSid field if non-nil, zero value otherwise.

### GetTaskQueueSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTaskQueueSidOk() (*string, bool)`

GetTaskQueueSidOk returns a tuple with the TaskQueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTaskQueueSid(v string)`

SetTaskQueueSid sets TaskQueueSid field to given value.

### HasTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTaskQueueSid() bool`

HasTaskQueueSid returns a boolean if a field has been set.

### SetTaskQueueSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTaskQueueSidNil(b bool)`

 SetTaskQueueSidNil sets the value for TaskQueueSid to be an explicit nil

### UnsetTaskQueueSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTaskQueueSid()`

UnsetTaskQueueSid ensures that no value is present for TaskQueueSid, not even an explicit nil
### GetTasksCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksCanceled() int32`

GetTasksCanceled returns the TasksCanceled field if non-nil, zero value otherwise.

### GetTasksCanceledOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksCanceledOk() (*int32, bool)`

GetTasksCanceledOk returns a tuple with the TasksCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksCanceled(v int32)`

SetTasksCanceled sets TasksCanceled field to given value.

### HasTasksCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTasksCanceled() bool`

HasTasksCanceled returns a boolean if a field has been set.

### SetTasksCanceledNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksCanceledNil(b bool)`

 SetTasksCanceledNil sets the value for TasksCanceled to be an explicit nil

### UnsetTasksCanceled
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTasksCanceled()`

UnsetTasksCanceled ensures that no value is present for TasksCanceled, not even an explicit nil
### GetTasksCompleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksCompleted() int32`

GetTasksCompleted returns the TasksCompleted field if non-nil, zero value otherwise.

### GetTasksCompletedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksCompletedOk() (*int32, bool)`

GetTasksCompletedOk returns a tuple with the TasksCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCompleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksCompleted(v int32)`

SetTasksCompleted sets TasksCompleted field to given value.

### HasTasksCompleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTasksCompleted() bool`

HasTasksCompleted returns a boolean if a field has been set.

### SetTasksCompletedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksCompletedNil(b bool)`

 SetTasksCompletedNil sets the value for TasksCompleted to be an explicit nil

### UnsetTasksCompleted
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTasksCompleted()`

UnsetTasksCompleted ensures that no value is present for TasksCompleted, not even an explicit nil
### GetTasksDeleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksDeleted() int32`

GetTasksDeleted returns the TasksDeleted field if non-nil, zero value otherwise.

### GetTasksDeletedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksDeletedOk() (*int32, bool)`

GetTasksDeletedOk returns a tuple with the TasksDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksDeleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksDeleted(v int32)`

SetTasksDeleted sets TasksDeleted field to given value.

### HasTasksDeleted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTasksDeleted() bool`

HasTasksDeleted returns a boolean if a field has been set.

### SetTasksDeletedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksDeletedNil(b bool)`

 SetTasksDeletedNil sets the value for TasksDeleted to be an explicit nil

### UnsetTasksDeleted
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTasksDeleted()`

UnsetTasksDeleted ensures that no value is present for TasksDeleted, not even an explicit nil
### GetTasksEntered

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksEntered() int32`

GetTasksEntered returns the TasksEntered field if non-nil, zero value otherwise.

### GetTasksEnteredOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksEnteredOk() (*int32, bool)`

GetTasksEnteredOk returns a tuple with the TasksEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksEntered

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksEntered(v int32)`

SetTasksEntered sets TasksEntered field to given value.

### HasTasksEntered

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTasksEntered() bool`

HasTasksEntered returns a boolean if a field has been set.

### SetTasksEnteredNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksEnteredNil(b bool)`

 SetTasksEnteredNil sets the value for TasksEntered to be an explicit nil

### UnsetTasksEntered
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTasksEntered()`

UnsetTasksEntered ensures that no value is present for TasksEntered, not even an explicit nil
### GetTasksMoved

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksMoved() int32`

GetTasksMoved returns the TasksMoved field if non-nil, zero value otherwise.

### GetTasksMovedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetTasksMovedOk() (*int32, bool)`

GetTasksMovedOk returns a tuple with the TasksMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksMoved

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksMoved(v int32)`

SetTasksMoved sets TasksMoved field to given value.

### HasTasksMoved

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasTasksMoved() bool`

HasTasksMoved returns a boolean if a field has been set.

### SetTasksMovedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetTasksMovedNil(b bool)`

 SetTasksMovedNil sets the value for TasksMoved to be an explicit nil

### UnsetTasksMoved
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetTasksMoved()`

UnsetTasksMoved ensures that no value is present for TasksMoved, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWaitDurationInQueueUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationInQueueUntilAccepted() map[string]interface{}`

GetWaitDurationInQueueUntilAccepted returns the WaitDurationInQueueUntilAccepted field if non-nil, zero value otherwise.

### GetWaitDurationInQueueUntilAcceptedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationInQueueUntilAcceptedOk() (*map[string]interface{}, bool)`

GetWaitDurationInQueueUntilAcceptedOk returns a tuple with the WaitDurationInQueueUntilAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationInQueueUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationInQueueUntilAccepted(v map[string]interface{})`

SetWaitDurationInQueueUntilAccepted sets WaitDurationInQueueUntilAccepted field to given value.

### HasWaitDurationInQueueUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasWaitDurationInQueueUntilAccepted() bool`

HasWaitDurationInQueueUntilAccepted returns a boolean if a field has been set.

### SetWaitDurationInQueueUntilAcceptedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationInQueueUntilAcceptedNil(b bool)`

 SetWaitDurationInQueueUntilAcceptedNil sets the value for WaitDurationInQueueUntilAccepted to be an explicit nil

### UnsetWaitDurationInQueueUntilAccepted
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetWaitDurationInQueueUntilAccepted()`

UnsetWaitDurationInQueueUntilAccepted ensures that no value is present for WaitDurationInQueueUntilAccepted, not even an explicit nil
### GetWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationUntilAccepted() map[string]interface{}`

GetWaitDurationUntilAccepted returns the WaitDurationUntilAccepted field if non-nil, zero value otherwise.

### GetWaitDurationUntilAcceptedOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationUntilAcceptedOk() (*map[string]interface{}, bool)`

GetWaitDurationUntilAcceptedOk returns a tuple with the WaitDurationUntilAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationUntilAccepted(v map[string]interface{})`

SetWaitDurationUntilAccepted sets WaitDurationUntilAccepted field to given value.

### HasWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasWaitDurationUntilAccepted() bool`

HasWaitDurationUntilAccepted returns a boolean if a field has been set.

### SetWaitDurationUntilAcceptedNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationUntilAcceptedNil(b bool)`

 SetWaitDurationUntilAcceptedNil sets the value for WaitDurationUntilAccepted to be an explicit nil

### UnsetWaitDurationUntilAccepted
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetWaitDurationUntilAccepted()`

UnsetWaitDurationUntilAccepted ensures that no value is present for WaitDurationUntilAccepted, not even an explicit nil
### GetWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationUntilCanceled() map[string]interface{}`

GetWaitDurationUntilCanceled returns the WaitDurationUntilCanceled field if non-nil, zero value otherwise.

### GetWaitDurationUntilCanceledOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWaitDurationUntilCanceledOk() (*map[string]interface{}, bool)`

GetWaitDurationUntilCanceledOk returns a tuple with the WaitDurationUntilCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationUntilCanceled(v map[string]interface{})`

SetWaitDurationUntilCanceled sets WaitDurationUntilCanceled field to given value.

### HasWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasWaitDurationUntilCanceled() bool`

HasWaitDurationUntilCanceled returns a boolean if a field has been set.

### SetWaitDurationUntilCanceledNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWaitDurationUntilCanceledNil(b bool)`

 SetWaitDurationUntilCanceledNil sets the value for WaitDurationUntilCanceled to be an explicit nil

### UnsetWaitDurationUntilCanceled
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetWaitDurationUntilCanceled()`

UnsetWaitDurationUntilCanceled ensures that no value is present for WaitDurationUntilCanceled, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


