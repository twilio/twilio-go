# TaskrouterV1WorkspaceWorkspaceCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AvgTaskAcceptanceTime** | Pointer to **NullableInt32** | The average time in seconds between Task creation and acceptance | [optional] 
**EndTime** | Pointer to **NullableTime** | The end of the interval during which these statistics were calculated | [optional] 
**ReservationsAccepted** | Pointer to **NullableInt32** | The total number of Reservations accepted by Workers | [optional] 
**ReservationsCanceled** | Pointer to **NullableInt32** | The total number of Reservations that were canceled | [optional] 
**ReservationsCreated** | Pointer to **NullableInt32** | The total number of Reservations that were created for Workers | [optional] 
**ReservationsRejected** | Pointer to **NullableInt32** | The total number of Reservations that were rejected | [optional] 
**ReservationsRescinded** | Pointer to **NullableInt32** | The total number of Reservations that were rescinded | [optional] 
**ReservationsTimedOut** | Pointer to **NullableInt32** | The total number of Reservations that were timed out | [optional] 
**SplitByWaitTime** | Pointer to **map[string]interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds | [optional] 
**StartTime** | Pointer to **NullableTime** | The beginning of the interval during which these statistics were calculated | [optional] 
**TasksCanceled** | Pointer to **NullableInt32** | The total number of Tasks that were canceled | [optional] 
**TasksCompleted** | Pointer to **NullableInt32** | The total number of Tasks that were completed | [optional] 
**TasksCreated** | Pointer to **NullableInt32** | The total number of Tasks created | [optional] 
**TasksDeleted** | Pointer to **NullableInt32** | The total number of Tasks that were deleted | [optional] 
**TasksMoved** | Pointer to **NullableInt32** | The total number of Tasks that were moved from one queue to another | [optional] 
**TasksTimedOutInWorkflow** | Pointer to **NullableInt32** | The total number of Tasks that were timed out of their Workflows | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workspace statistics resource | [optional] 
**WaitDurationUntilAccepted** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks that were accepted | [optional] 
**WaitDurationUntilCanceled** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks that were canceled | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkspaceCumulativeStatistics

`func NewTaskrouterV1WorkspaceWorkspaceCumulativeStatistics() *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics`

NewTaskrouterV1WorkspaceWorkspaceCumulativeStatistics instantiates a new TaskrouterV1WorkspaceWorkspaceCumulativeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkspaceCumulativeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkspaceCumulativeStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics`

NewTaskrouterV1WorkspaceWorkspaceCumulativeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkspaceCumulativeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetAvgTaskAcceptanceTime() int32`

GetAvgTaskAcceptanceTime returns the AvgTaskAcceptanceTime field if non-nil, zero value otherwise.

### GetAvgTaskAcceptanceTimeOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetAvgTaskAcceptanceTimeOk() (*int32, bool)`

GetAvgTaskAcceptanceTimeOk returns a tuple with the AvgTaskAcceptanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetAvgTaskAcceptanceTime(v int32)`

SetAvgTaskAcceptanceTime sets AvgTaskAcceptanceTime field to given value.

### HasAvgTaskAcceptanceTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasAvgTaskAcceptanceTime() bool`

HasAvgTaskAcceptanceTime returns a boolean if a field has been set.

### SetAvgTaskAcceptanceTimeNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetAvgTaskAcceptanceTimeNil(b bool)`

 SetAvgTaskAcceptanceTimeNil sets the value for AvgTaskAcceptanceTime to be an explicit nil

### UnsetAvgTaskAcceptanceTime
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetAvgTaskAcceptanceTime()`

UnsetAvgTaskAcceptanceTime ensures that no value is present for AvgTaskAcceptanceTime, not even an explicit nil
### GetEndTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsAccepted() int32`

GetReservationsAccepted returns the ReservationsAccepted field if non-nil, zero value otherwise.

### GetReservationsAcceptedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsAcceptedOk() (*int32, bool)`

GetReservationsAcceptedOk returns a tuple with the ReservationsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsAccepted(v int32)`

SetReservationsAccepted sets ReservationsAccepted field to given value.

### HasReservationsAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsAccepted() bool`

HasReservationsAccepted returns a boolean if a field has been set.

### SetReservationsAcceptedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsAcceptedNil(b bool)`

 SetReservationsAcceptedNil sets the value for ReservationsAccepted to be an explicit nil

### UnsetReservationsAccepted
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsAccepted()`

UnsetReservationsAccepted ensures that no value is present for ReservationsAccepted, not even an explicit nil
### GetReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsCanceled() int32`

GetReservationsCanceled returns the ReservationsCanceled field if non-nil, zero value otherwise.

### GetReservationsCanceledOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsCanceledOk() (*int32, bool)`

GetReservationsCanceledOk returns a tuple with the ReservationsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsCanceled(v int32)`

SetReservationsCanceled sets ReservationsCanceled field to given value.

### HasReservationsCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsCanceled() bool`

HasReservationsCanceled returns a boolean if a field has been set.

### SetReservationsCanceledNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsCanceledNil(b bool)`

 SetReservationsCanceledNil sets the value for ReservationsCanceled to be an explicit nil

### UnsetReservationsCanceled
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsCanceled()`

UnsetReservationsCanceled ensures that no value is present for ReservationsCanceled, not even an explicit nil
### GetReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsCreated() int32`

GetReservationsCreated returns the ReservationsCreated field if non-nil, zero value otherwise.

### GetReservationsCreatedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsCreatedOk() (*int32, bool)`

GetReservationsCreatedOk returns a tuple with the ReservationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsCreated(v int32)`

SetReservationsCreated sets ReservationsCreated field to given value.

### HasReservationsCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsCreated() bool`

HasReservationsCreated returns a boolean if a field has been set.

### SetReservationsCreatedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsCreatedNil(b bool)`

 SetReservationsCreatedNil sets the value for ReservationsCreated to be an explicit nil

### UnsetReservationsCreated
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsCreated()`

UnsetReservationsCreated ensures that no value is present for ReservationsCreated, not even an explicit nil
### GetReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsRejected() int32`

GetReservationsRejected returns the ReservationsRejected field if non-nil, zero value otherwise.

### GetReservationsRejectedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsRejectedOk() (*int32, bool)`

GetReservationsRejectedOk returns a tuple with the ReservationsRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsRejected(v int32)`

SetReservationsRejected sets ReservationsRejected field to given value.

### HasReservationsRejected

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsRejected() bool`

HasReservationsRejected returns a boolean if a field has been set.

### SetReservationsRejectedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsRejectedNil(b bool)`

 SetReservationsRejectedNil sets the value for ReservationsRejected to be an explicit nil

### UnsetReservationsRejected
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsRejected()`

UnsetReservationsRejected ensures that no value is present for ReservationsRejected, not even an explicit nil
### GetReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsRescinded() int32`

GetReservationsRescinded returns the ReservationsRescinded field if non-nil, zero value otherwise.

### GetReservationsRescindedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsRescindedOk() (*int32, bool)`

GetReservationsRescindedOk returns a tuple with the ReservationsRescinded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsRescinded(v int32)`

SetReservationsRescinded sets ReservationsRescinded field to given value.

### HasReservationsRescinded

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsRescinded() bool`

HasReservationsRescinded returns a boolean if a field has been set.

### SetReservationsRescindedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsRescindedNil(b bool)`

 SetReservationsRescindedNil sets the value for ReservationsRescinded to be an explicit nil

### UnsetReservationsRescinded
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsRescinded()`

UnsetReservationsRescinded ensures that no value is present for ReservationsRescinded, not even an explicit nil
### GetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsTimedOut() int32`

GetReservationsTimedOut returns the ReservationsTimedOut field if non-nil, zero value otherwise.

### GetReservationsTimedOutOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetReservationsTimedOutOk() (*int32, bool)`

GetReservationsTimedOutOk returns a tuple with the ReservationsTimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsTimedOut(v int32)`

SetReservationsTimedOut sets ReservationsTimedOut field to given value.

### HasReservationsTimedOut

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasReservationsTimedOut() bool`

HasReservationsTimedOut returns a boolean if a field has been set.

### SetReservationsTimedOutNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetReservationsTimedOutNil(b bool)`

 SetReservationsTimedOutNil sets the value for ReservationsTimedOut to be an explicit nil

### UnsetReservationsTimedOut
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetReservationsTimedOut()`

UnsetReservationsTimedOut ensures that no value is present for ReservationsTimedOut, not even an explicit nil
### GetSplitByWaitTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetSplitByWaitTime() map[string]interface{}`

GetSplitByWaitTime returns the SplitByWaitTime field if non-nil, zero value otherwise.

### GetSplitByWaitTimeOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetSplitByWaitTimeOk() (*map[string]interface{}, bool)`

GetSplitByWaitTimeOk returns a tuple with the SplitByWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByWaitTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetSplitByWaitTime(v map[string]interface{})`

SetSplitByWaitTime sets SplitByWaitTime field to given value.

### HasSplitByWaitTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasSplitByWaitTime() bool`

HasSplitByWaitTime returns a boolean if a field has been set.

### SetSplitByWaitTimeNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetSplitByWaitTimeNil(b bool)`

 SetSplitByWaitTimeNil sets the value for SplitByWaitTime to be an explicit nil

### UnsetSplitByWaitTime
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetSplitByWaitTime()`

UnsetSplitByWaitTime ensures that no value is present for SplitByWaitTime, not even an explicit nil
### GetStartTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTasksCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCanceled() int32`

GetTasksCanceled returns the TasksCanceled field if non-nil, zero value otherwise.

### GetTasksCanceledOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCanceledOk() (*int32, bool)`

GetTasksCanceledOk returns a tuple with the TasksCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCanceled(v int32)`

SetTasksCanceled sets TasksCanceled field to given value.

### HasTasksCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksCanceled() bool`

HasTasksCanceled returns a boolean if a field has been set.

### SetTasksCanceledNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCanceledNil(b bool)`

 SetTasksCanceledNil sets the value for TasksCanceled to be an explicit nil

### UnsetTasksCanceled
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksCanceled()`

UnsetTasksCanceled ensures that no value is present for TasksCanceled, not even an explicit nil
### GetTasksCompleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCompleted() int32`

GetTasksCompleted returns the TasksCompleted field if non-nil, zero value otherwise.

### GetTasksCompletedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCompletedOk() (*int32, bool)`

GetTasksCompletedOk returns a tuple with the TasksCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCompleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCompleted(v int32)`

SetTasksCompleted sets TasksCompleted field to given value.

### HasTasksCompleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksCompleted() bool`

HasTasksCompleted returns a boolean if a field has been set.

### SetTasksCompletedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCompletedNil(b bool)`

 SetTasksCompletedNil sets the value for TasksCompleted to be an explicit nil

### UnsetTasksCompleted
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksCompleted()`

UnsetTasksCompleted ensures that no value is present for TasksCompleted, not even an explicit nil
### GetTasksCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCreated() int32`

GetTasksCreated returns the TasksCreated field if non-nil, zero value otherwise.

### GetTasksCreatedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksCreatedOk() (*int32, bool)`

GetTasksCreatedOk returns a tuple with the TasksCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCreated(v int32)`

SetTasksCreated sets TasksCreated field to given value.

### HasTasksCreated

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksCreated() bool`

HasTasksCreated returns a boolean if a field has been set.

### SetTasksCreatedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksCreatedNil(b bool)`

 SetTasksCreatedNil sets the value for TasksCreated to be an explicit nil

### UnsetTasksCreated
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksCreated()`

UnsetTasksCreated ensures that no value is present for TasksCreated, not even an explicit nil
### GetTasksDeleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksDeleted() int32`

GetTasksDeleted returns the TasksDeleted field if non-nil, zero value otherwise.

### GetTasksDeletedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksDeletedOk() (*int32, bool)`

GetTasksDeletedOk returns a tuple with the TasksDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksDeleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksDeleted(v int32)`

SetTasksDeleted sets TasksDeleted field to given value.

### HasTasksDeleted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksDeleted() bool`

HasTasksDeleted returns a boolean if a field has been set.

### SetTasksDeletedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksDeletedNil(b bool)`

 SetTasksDeletedNil sets the value for TasksDeleted to be an explicit nil

### UnsetTasksDeleted
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksDeleted()`

UnsetTasksDeleted ensures that no value is present for TasksDeleted, not even an explicit nil
### GetTasksMoved

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksMoved() int32`

GetTasksMoved returns the TasksMoved field if non-nil, zero value otherwise.

### GetTasksMovedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksMovedOk() (*int32, bool)`

GetTasksMovedOk returns a tuple with the TasksMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksMoved

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksMoved(v int32)`

SetTasksMoved sets TasksMoved field to given value.

### HasTasksMoved

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksMoved() bool`

HasTasksMoved returns a boolean if a field has been set.

### SetTasksMovedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksMovedNil(b bool)`

 SetTasksMovedNil sets the value for TasksMoved to be an explicit nil

### UnsetTasksMoved
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksMoved()`

UnsetTasksMoved ensures that no value is present for TasksMoved, not even an explicit nil
### GetTasksTimedOutInWorkflow

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksTimedOutInWorkflow() int32`

GetTasksTimedOutInWorkflow returns the TasksTimedOutInWorkflow field if non-nil, zero value otherwise.

### GetTasksTimedOutInWorkflowOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetTasksTimedOutInWorkflowOk() (*int32, bool)`

GetTasksTimedOutInWorkflowOk returns a tuple with the TasksTimedOutInWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksTimedOutInWorkflow

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksTimedOutInWorkflow(v int32)`

SetTasksTimedOutInWorkflow sets TasksTimedOutInWorkflow field to given value.

### HasTasksTimedOutInWorkflow

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasTasksTimedOutInWorkflow() bool`

HasTasksTimedOutInWorkflow returns a boolean if a field has been set.

### SetTasksTimedOutInWorkflowNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetTasksTimedOutInWorkflowNil(b bool)`

 SetTasksTimedOutInWorkflowNil sets the value for TasksTimedOutInWorkflow to be an explicit nil

### UnsetTasksTimedOutInWorkflow
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetTasksTimedOutInWorkflow()`

UnsetTasksTimedOutInWorkflow ensures that no value is present for TasksTimedOutInWorkflow, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWaitDurationUntilAccepted() map[string]interface{}`

GetWaitDurationUntilAccepted returns the WaitDurationUntilAccepted field if non-nil, zero value otherwise.

### GetWaitDurationUntilAcceptedOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWaitDurationUntilAcceptedOk() (*map[string]interface{}, bool)`

GetWaitDurationUntilAcceptedOk returns a tuple with the WaitDurationUntilAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWaitDurationUntilAccepted(v map[string]interface{})`

SetWaitDurationUntilAccepted sets WaitDurationUntilAccepted field to given value.

### HasWaitDurationUntilAccepted

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasWaitDurationUntilAccepted() bool`

HasWaitDurationUntilAccepted returns a boolean if a field has been set.

### SetWaitDurationUntilAcceptedNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWaitDurationUntilAcceptedNil(b bool)`

 SetWaitDurationUntilAcceptedNil sets the value for WaitDurationUntilAccepted to be an explicit nil

### UnsetWaitDurationUntilAccepted
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetWaitDurationUntilAccepted()`

UnsetWaitDurationUntilAccepted ensures that no value is present for WaitDurationUntilAccepted, not even an explicit nil
### GetWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWaitDurationUntilCanceled() map[string]interface{}`

GetWaitDurationUntilCanceled returns the WaitDurationUntilCanceled field if non-nil, zero value otherwise.

### GetWaitDurationUntilCanceledOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWaitDurationUntilCanceledOk() (*map[string]interface{}, bool)`

GetWaitDurationUntilCanceledOk returns a tuple with the WaitDurationUntilCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWaitDurationUntilCanceled(v map[string]interface{})`

SetWaitDurationUntilCanceled sets WaitDurationUntilCanceled field to given value.

### HasWaitDurationUntilCanceled

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasWaitDurationUntilCanceled() bool`

HasWaitDurationUntilCanceled returns a boolean if a field has been set.

### SetWaitDurationUntilCanceledNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWaitDurationUntilCanceledNil(b bool)`

 SetWaitDurationUntilCanceledNil sets the value for WaitDurationUntilCanceled to be an explicit nil

### UnsetWaitDurationUntilCanceled
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetWaitDurationUntilCanceled()`

UnsetWaitDurationUntilCanceled ensures that no value is present for WaitDurationUntilCanceled, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkspaceCumulativeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


