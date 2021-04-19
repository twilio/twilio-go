# ApiV2010AccountQueueMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**DateEnqueued** | Pointer to **NullableString** | The date the member was enqueued | [optional] 
**Position** | Pointer to **NullableInt32** | This member&#39;s current position in the queue. | [optional] 
**QueueSid** | Pointer to **NullableString** | The SID of the Queue the member is in | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**WaitTime** | Pointer to **NullableInt32** | The number of seconds the member has been in the queue. | [optional] 

## Methods

### NewApiV2010AccountQueueMember

`func NewApiV2010AccountQueueMember() *ApiV2010AccountQueueMember`

NewApiV2010AccountQueueMember instantiates a new ApiV2010AccountQueueMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountQueueMemberWithDefaults

`func NewApiV2010AccountQueueMemberWithDefaults() *ApiV2010AccountQueueMember`

NewApiV2010AccountQueueMemberWithDefaults instantiates a new ApiV2010AccountQueueMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallSid

`func (o *ApiV2010AccountQueueMember) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountQueueMember) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountQueueMember) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountQueueMember) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountQueueMember) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountQueueMember) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetDateEnqueued

`func (o *ApiV2010AccountQueueMember) GetDateEnqueued() string`

GetDateEnqueued returns the DateEnqueued field if non-nil, zero value otherwise.

### GetDateEnqueuedOk

`func (o *ApiV2010AccountQueueMember) GetDateEnqueuedOk() (*string, bool)`

GetDateEnqueuedOk returns a tuple with the DateEnqueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnqueued

`func (o *ApiV2010AccountQueueMember) SetDateEnqueued(v string)`

SetDateEnqueued sets DateEnqueued field to given value.

### HasDateEnqueued

`func (o *ApiV2010AccountQueueMember) HasDateEnqueued() bool`

HasDateEnqueued returns a boolean if a field has been set.

### SetDateEnqueuedNil

`func (o *ApiV2010AccountQueueMember) SetDateEnqueuedNil(b bool)`

 SetDateEnqueuedNil sets the value for DateEnqueued to be an explicit nil

### UnsetDateEnqueued
`func (o *ApiV2010AccountQueueMember) UnsetDateEnqueued()`

UnsetDateEnqueued ensures that no value is present for DateEnqueued, not even an explicit nil
### GetPosition

`func (o *ApiV2010AccountQueueMember) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ApiV2010AccountQueueMember) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ApiV2010AccountQueueMember) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ApiV2010AccountQueueMember) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *ApiV2010AccountQueueMember) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *ApiV2010AccountQueueMember) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetQueueSid

`func (o *ApiV2010AccountQueueMember) GetQueueSid() string`

GetQueueSid returns the QueueSid field if non-nil, zero value otherwise.

### GetQueueSidOk

`func (o *ApiV2010AccountQueueMember) GetQueueSidOk() (*string, bool)`

GetQueueSidOk returns a tuple with the QueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSid

`func (o *ApiV2010AccountQueueMember) SetQueueSid(v string)`

SetQueueSid sets QueueSid field to given value.

### HasQueueSid

`func (o *ApiV2010AccountQueueMember) HasQueueSid() bool`

HasQueueSid returns a boolean if a field has been set.

### SetQueueSidNil

`func (o *ApiV2010AccountQueueMember) SetQueueSidNil(b bool)`

 SetQueueSidNil sets the value for QueueSid to be an explicit nil

### UnsetQueueSid
`func (o *ApiV2010AccountQueueMember) UnsetQueueSid()`

UnsetQueueSid ensures that no value is present for QueueSid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountQueueMember) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountQueueMember) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountQueueMember) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountQueueMember) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountQueueMember) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountQueueMember) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetWaitTime

`func (o *ApiV2010AccountQueueMember) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *ApiV2010AccountQueueMember) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *ApiV2010AccountQueueMember) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *ApiV2010AccountQueueMember) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

### SetWaitTimeNil

`func (o *ApiV2010AccountQueueMember) SetWaitTimeNil(b bool)`

 SetWaitTimeNil sets the value for WaitTime to be an explicit nil

### UnsetWaitTime
`func (o *ApiV2010AccountQueueMember) UnsetWaitTime()`

UnsetWaitTime ensures that no value is present for WaitTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


