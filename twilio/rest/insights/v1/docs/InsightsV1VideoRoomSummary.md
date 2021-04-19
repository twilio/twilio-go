# InsightsV1VideoRoomSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account SID associated with this room. | [optional] 
**Codecs** | Pointer to **[]string** | Codecs used by participants in the room. | [optional] 
**ConcurrentParticipants** | Pointer to **NullableInt32** | Actual number of concurrent participants. | [optional] 
**CreateTime** | Pointer to **NullableTime** | Creation time of the room. | [optional] 
**CreatedMethod** | Pointer to **NullableString** | How the room was created. | [optional] 
**DurationSec** | Pointer to **NullableInt32** | Total room duration from create time to end time. | [optional] 
**EdgeLocation** | Pointer to **NullableString** | Edge location of Twilio media servers for the room. | [optional] 
**EndReason** | Pointer to **NullableString** | Reason the room ended. | [optional] 
**EndTime** | Pointer to **NullableTime** | End time for the room. | [optional] 
**Links** | Pointer to **map[string]interface{}** | Room subresources. | [optional] 
**MaxConcurrentParticipants** | Pointer to **NullableInt32** | Maximum number of participants allowed in the room at the same time allowed by the application settings. | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** | Max number of total participants allowed by the application settings. | [optional] 
**MediaRegion** | Pointer to **NullableString** | Region of Twilio media servers for the room. | [optional] 
**ProcessingState** | Pointer to **NullableString** | Video Log Analyzer resource state. Will be either &#x60;in-progress&#x60; or &#x60;complete&#x60;. | [optional] 
**RecordingEnabled** | Pointer to **NullableBool** | Boolean indicating if recording is enabled for the room. | [optional] 
**RoomName** | Pointer to **NullableString** | room friendly name. | [optional] 
**RoomSid** | Pointer to **NullableString** | Unique identifier for the room. | [optional] 
**RoomStatus** | Pointer to **NullableString** | Status of the room. | [optional] 
**RoomType** | Pointer to **NullableString** | Type of room. | [optional] 
**StatusCallback** | Pointer to **NullableString** | Webhook provided for status callbacks. | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | HTTP method provided for status callback URL. | [optional] 
**TotalParticipantDurationSec** | Pointer to **NullableInt32** | Combined amount of participant time in the room. | [optional] 
**TotalRecordingDurationSec** | Pointer to **NullableInt32** | Combined amount of recorded seconds for participants in the room. | [optional] 
**UniqueParticipantIdentities** | Pointer to **NullableInt32** | Unique number of participant identities. | [optional] 
**UniqueParticipants** | Pointer to **NullableInt32** | Number of participants. May include duplicate identities for participants who left and rejoined. | [optional] 
**Url** | Pointer to **NullableString** | URL for the room resource. | [optional] 

## Methods

### NewInsightsV1VideoRoomSummary

`func NewInsightsV1VideoRoomSummary() *InsightsV1VideoRoomSummary`

NewInsightsV1VideoRoomSummary instantiates a new InsightsV1VideoRoomSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsV1VideoRoomSummaryWithDefaults

`func NewInsightsV1VideoRoomSummaryWithDefaults() *InsightsV1VideoRoomSummary`

NewInsightsV1VideoRoomSummaryWithDefaults instantiates a new InsightsV1VideoRoomSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *InsightsV1VideoRoomSummary) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *InsightsV1VideoRoomSummary) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *InsightsV1VideoRoomSummary) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *InsightsV1VideoRoomSummary) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *InsightsV1VideoRoomSummary) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *InsightsV1VideoRoomSummary) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCodecs

`func (o *InsightsV1VideoRoomSummary) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *InsightsV1VideoRoomSummary) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *InsightsV1VideoRoomSummary) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *InsightsV1VideoRoomSummary) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### SetCodecsNil

`func (o *InsightsV1VideoRoomSummary) SetCodecsNil(b bool)`

 SetCodecsNil sets the value for Codecs to be an explicit nil

### UnsetCodecs
`func (o *InsightsV1VideoRoomSummary) UnsetCodecs()`

UnsetCodecs ensures that no value is present for Codecs, not even an explicit nil
### GetConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) GetConcurrentParticipants() int32`

GetConcurrentParticipants returns the ConcurrentParticipants field if non-nil, zero value otherwise.

### GetConcurrentParticipantsOk

`func (o *InsightsV1VideoRoomSummary) GetConcurrentParticipantsOk() (*int32, bool)`

GetConcurrentParticipantsOk returns a tuple with the ConcurrentParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) SetConcurrentParticipants(v int32)`

SetConcurrentParticipants sets ConcurrentParticipants field to given value.

### HasConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) HasConcurrentParticipants() bool`

HasConcurrentParticipants returns a boolean if a field has been set.

### SetConcurrentParticipantsNil

`func (o *InsightsV1VideoRoomSummary) SetConcurrentParticipantsNil(b bool)`

 SetConcurrentParticipantsNil sets the value for ConcurrentParticipants to be an explicit nil

### UnsetConcurrentParticipants
`func (o *InsightsV1VideoRoomSummary) UnsetConcurrentParticipants()`

UnsetConcurrentParticipants ensures that no value is present for ConcurrentParticipants, not even an explicit nil
### GetCreateTime

`func (o *InsightsV1VideoRoomSummary) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InsightsV1VideoRoomSummary) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InsightsV1VideoRoomSummary) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *InsightsV1VideoRoomSummary) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *InsightsV1VideoRoomSummary) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *InsightsV1VideoRoomSummary) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetCreatedMethod

`func (o *InsightsV1VideoRoomSummary) GetCreatedMethod() string`

GetCreatedMethod returns the CreatedMethod field if non-nil, zero value otherwise.

### GetCreatedMethodOk

`func (o *InsightsV1VideoRoomSummary) GetCreatedMethodOk() (*string, bool)`

GetCreatedMethodOk returns a tuple with the CreatedMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedMethod

`func (o *InsightsV1VideoRoomSummary) SetCreatedMethod(v string)`

SetCreatedMethod sets CreatedMethod field to given value.

### HasCreatedMethod

`func (o *InsightsV1VideoRoomSummary) HasCreatedMethod() bool`

HasCreatedMethod returns a boolean if a field has been set.

### SetCreatedMethodNil

`func (o *InsightsV1VideoRoomSummary) SetCreatedMethodNil(b bool)`

 SetCreatedMethodNil sets the value for CreatedMethod to be an explicit nil

### UnsetCreatedMethod
`func (o *InsightsV1VideoRoomSummary) UnsetCreatedMethod()`

UnsetCreatedMethod ensures that no value is present for CreatedMethod, not even an explicit nil
### GetDurationSec

`func (o *InsightsV1VideoRoomSummary) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *InsightsV1VideoRoomSummary) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *InsightsV1VideoRoomSummary) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.

### HasDurationSec

`func (o *InsightsV1VideoRoomSummary) HasDurationSec() bool`

HasDurationSec returns a boolean if a field has been set.

### SetDurationSecNil

`func (o *InsightsV1VideoRoomSummary) SetDurationSecNil(b bool)`

 SetDurationSecNil sets the value for DurationSec to be an explicit nil

### UnsetDurationSec
`func (o *InsightsV1VideoRoomSummary) UnsetDurationSec()`

UnsetDurationSec ensures that no value is present for DurationSec, not even an explicit nil
### GetEdgeLocation

`func (o *InsightsV1VideoRoomSummary) GetEdgeLocation() string`

GetEdgeLocation returns the EdgeLocation field if non-nil, zero value otherwise.

### GetEdgeLocationOk

`func (o *InsightsV1VideoRoomSummary) GetEdgeLocationOk() (*string, bool)`

GetEdgeLocationOk returns a tuple with the EdgeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeLocation

`func (o *InsightsV1VideoRoomSummary) SetEdgeLocation(v string)`

SetEdgeLocation sets EdgeLocation field to given value.

### HasEdgeLocation

`func (o *InsightsV1VideoRoomSummary) HasEdgeLocation() bool`

HasEdgeLocation returns a boolean if a field has been set.

### SetEdgeLocationNil

`func (o *InsightsV1VideoRoomSummary) SetEdgeLocationNil(b bool)`

 SetEdgeLocationNil sets the value for EdgeLocation to be an explicit nil

### UnsetEdgeLocation
`func (o *InsightsV1VideoRoomSummary) UnsetEdgeLocation()`

UnsetEdgeLocation ensures that no value is present for EdgeLocation, not even an explicit nil
### GetEndReason

`func (o *InsightsV1VideoRoomSummary) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *InsightsV1VideoRoomSummary) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *InsightsV1VideoRoomSummary) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *InsightsV1VideoRoomSummary) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### SetEndReasonNil

`func (o *InsightsV1VideoRoomSummary) SetEndReasonNil(b bool)`

 SetEndReasonNil sets the value for EndReason to be an explicit nil

### UnsetEndReason
`func (o *InsightsV1VideoRoomSummary) UnsetEndReason()`

UnsetEndReason ensures that no value is present for EndReason, not even an explicit nil
### GetEndTime

`func (o *InsightsV1VideoRoomSummary) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InsightsV1VideoRoomSummary) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InsightsV1VideoRoomSummary) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InsightsV1VideoRoomSummary) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *InsightsV1VideoRoomSummary) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *InsightsV1VideoRoomSummary) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetLinks

`func (o *InsightsV1VideoRoomSummary) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InsightsV1VideoRoomSummary) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InsightsV1VideoRoomSummary) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InsightsV1VideoRoomSummary) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *InsightsV1VideoRoomSummary) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *InsightsV1VideoRoomSummary) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMaxConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) GetMaxConcurrentParticipants() int32`

GetMaxConcurrentParticipants returns the MaxConcurrentParticipants field if non-nil, zero value otherwise.

### GetMaxConcurrentParticipantsOk

`func (o *InsightsV1VideoRoomSummary) GetMaxConcurrentParticipantsOk() (*int32, bool)`

GetMaxConcurrentParticipantsOk returns a tuple with the MaxConcurrentParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) SetMaxConcurrentParticipants(v int32)`

SetMaxConcurrentParticipants sets MaxConcurrentParticipants field to given value.

### HasMaxConcurrentParticipants

`func (o *InsightsV1VideoRoomSummary) HasMaxConcurrentParticipants() bool`

HasMaxConcurrentParticipants returns a boolean if a field has been set.

### SetMaxConcurrentParticipantsNil

`func (o *InsightsV1VideoRoomSummary) SetMaxConcurrentParticipantsNil(b bool)`

 SetMaxConcurrentParticipantsNil sets the value for MaxConcurrentParticipants to be an explicit nil

### UnsetMaxConcurrentParticipants
`func (o *InsightsV1VideoRoomSummary) UnsetMaxConcurrentParticipants()`

UnsetMaxConcurrentParticipants ensures that no value is present for MaxConcurrentParticipants, not even an explicit nil
### GetMaxParticipants

`func (o *InsightsV1VideoRoomSummary) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *InsightsV1VideoRoomSummary) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *InsightsV1VideoRoomSummary) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *InsightsV1VideoRoomSummary) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *InsightsV1VideoRoomSummary) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *InsightsV1VideoRoomSummary) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetMediaRegion

`func (o *InsightsV1VideoRoomSummary) GetMediaRegion() string`

GetMediaRegion returns the MediaRegion field if non-nil, zero value otherwise.

### GetMediaRegionOk

`func (o *InsightsV1VideoRoomSummary) GetMediaRegionOk() (*string, bool)`

GetMediaRegionOk returns a tuple with the MediaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaRegion

`func (o *InsightsV1VideoRoomSummary) SetMediaRegion(v string)`

SetMediaRegion sets MediaRegion field to given value.

### HasMediaRegion

`func (o *InsightsV1VideoRoomSummary) HasMediaRegion() bool`

HasMediaRegion returns a boolean if a field has been set.

### SetMediaRegionNil

`func (o *InsightsV1VideoRoomSummary) SetMediaRegionNil(b bool)`

 SetMediaRegionNil sets the value for MediaRegion to be an explicit nil

### UnsetMediaRegion
`func (o *InsightsV1VideoRoomSummary) UnsetMediaRegion()`

UnsetMediaRegion ensures that no value is present for MediaRegion, not even an explicit nil
### GetProcessingState

`func (o *InsightsV1VideoRoomSummary) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *InsightsV1VideoRoomSummary) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *InsightsV1VideoRoomSummary) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *InsightsV1VideoRoomSummary) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *InsightsV1VideoRoomSummary) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *InsightsV1VideoRoomSummary) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetRecordingEnabled

`func (o *InsightsV1VideoRoomSummary) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *InsightsV1VideoRoomSummary) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *InsightsV1VideoRoomSummary) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.

### HasRecordingEnabled

`func (o *InsightsV1VideoRoomSummary) HasRecordingEnabled() bool`

HasRecordingEnabled returns a boolean if a field has been set.

### SetRecordingEnabledNil

`func (o *InsightsV1VideoRoomSummary) SetRecordingEnabledNil(b bool)`

 SetRecordingEnabledNil sets the value for RecordingEnabled to be an explicit nil

### UnsetRecordingEnabled
`func (o *InsightsV1VideoRoomSummary) UnsetRecordingEnabled()`

UnsetRecordingEnabled ensures that no value is present for RecordingEnabled, not even an explicit nil
### GetRoomName

`func (o *InsightsV1VideoRoomSummary) GetRoomName() string`

GetRoomName returns the RoomName field if non-nil, zero value otherwise.

### GetRoomNameOk

`func (o *InsightsV1VideoRoomSummary) GetRoomNameOk() (*string, bool)`

GetRoomNameOk returns a tuple with the RoomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomName

`func (o *InsightsV1VideoRoomSummary) SetRoomName(v string)`

SetRoomName sets RoomName field to given value.

### HasRoomName

`func (o *InsightsV1VideoRoomSummary) HasRoomName() bool`

HasRoomName returns a boolean if a field has been set.

### SetRoomNameNil

`func (o *InsightsV1VideoRoomSummary) SetRoomNameNil(b bool)`

 SetRoomNameNil sets the value for RoomName to be an explicit nil

### UnsetRoomName
`func (o *InsightsV1VideoRoomSummary) UnsetRoomName()`

UnsetRoomName ensures that no value is present for RoomName, not even an explicit nil
### GetRoomSid

`func (o *InsightsV1VideoRoomSummary) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *InsightsV1VideoRoomSummary) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *InsightsV1VideoRoomSummary) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *InsightsV1VideoRoomSummary) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *InsightsV1VideoRoomSummary) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *InsightsV1VideoRoomSummary) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetRoomStatus

`func (o *InsightsV1VideoRoomSummary) GetRoomStatus() string`

GetRoomStatus returns the RoomStatus field if non-nil, zero value otherwise.

### GetRoomStatusOk

`func (o *InsightsV1VideoRoomSummary) GetRoomStatusOk() (*string, bool)`

GetRoomStatusOk returns a tuple with the RoomStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomStatus

`func (o *InsightsV1VideoRoomSummary) SetRoomStatus(v string)`

SetRoomStatus sets RoomStatus field to given value.

### HasRoomStatus

`func (o *InsightsV1VideoRoomSummary) HasRoomStatus() bool`

HasRoomStatus returns a boolean if a field has been set.

### SetRoomStatusNil

`func (o *InsightsV1VideoRoomSummary) SetRoomStatusNil(b bool)`

 SetRoomStatusNil sets the value for RoomStatus to be an explicit nil

### UnsetRoomStatus
`func (o *InsightsV1VideoRoomSummary) UnsetRoomStatus()`

UnsetRoomStatus ensures that no value is present for RoomStatus, not even an explicit nil
### GetRoomType

`func (o *InsightsV1VideoRoomSummary) GetRoomType() string`

GetRoomType returns the RoomType field if non-nil, zero value otherwise.

### GetRoomTypeOk

`func (o *InsightsV1VideoRoomSummary) GetRoomTypeOk() (*string, bool)`

GetRoomTypeOk returns a tuple with the RoomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomType

`func (o *InsightsV1VideoRoomSummary) SetRoomType(v string)`

SetRoomType sets RoomType field to given value.

### HasRoomType

`func (o *InsightsV1VideoRoomSummary) HasRoomType() bool`

HasRoomType returns a boolean if a field has been set.

### SetRoomTypeNil

`func (o *InsightsV1VideoRoomSummary) SetRoomTypeNil(b bool)`

 SetRoomTypeNil sets the value for RoomType to be an explicit nil

### UnsetRoomType
`func (o *InsightsV1VideoRoomSummary) UnsetRoomType()`

UnsetRoomType ensures that no value is present for RoomType, not even an explicit nil
### GetStatusCallback

`func (o *InsightsV1VideoRoomSummary) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *InsightsV1VideoRoomSummary) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *InsightsV1VideoRoomSummary) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *InsightsV1VideoRoomSummary) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *InsightsV1VideoRoomSummary) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *InsightsV1VideoRoomSummary) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *InsightsV1VideoRoomSummary) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *InsightsV1VideoRoomSummary) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *InsightsV1VideoRoomSummary) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *InsightsV1VideoRoomSummary) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *InsightsV1VideoRoomSummary) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *InsightsV1VideoRoomSummary) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetTotalParticipantDurationSec

`func (o *InsightsV1VideoRoomSummary) GetTotalParticipantDurationSec() int32`

GetTotalParticipantDurationSec returns the TotalParticipantDurationSec field if non-nil, zero value otherwise.

### GetTotalParticipantDurationSecOk

`func (o *InsightsV1VideoRoomSummary) GetTotalParticipantDurationSecOk() (*int32, bool)`

GetTotalParticipantDurationSecOk returns a tuple with the TotalParticipantDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParticipantDurationSec

`func (o *InsightsV1VideoRoomSummary) SetTotalParticipantDurationSec(v int32)`

SetTotalParticipantDurationSec sets TotalParticipantDurationSec field to given value.

### HasTotalParticipantDurationSec

`func (o *InsightsV1VideoRoomSummary) HasTotalParticipantDurationSec() bool`

HasTotalParticipantDurationSec returns a boolean if a field has been set.

### SetTotalParticipantDurationSecNil

`func (o *InsightsV1VideoRoomSummary) SetTotalParticipantDurationSecNil(b bool)`

 SetTotalParticipantDurationSecNil sets the value for TotalParticipantDurationSec to be an explicit nil

### UnsetTotalParticipantDurationSec
`func (o *InsightsV1VideoRoomSummary) UnsetTotalParticipantDurationSec()`

UnsetTotalParticipantDurationSec ensures that no value is present for TotalParticipantDurationSec, not even an explicit nil
### GetTotalRecordingDurationSec

`func (o *InsightsV1VideoRoomSummary) GetTotalRecordingDurationSec() int32`

GetTotalRecordingDurationSec returns the TotalRecordingDurationSec field if non-nil, zero value otherwise.

### GetTotalRecordingDurationSecOk

`func (o *InsightsV1VideoRoomSummary) GetTotalRecordingDurationSecOk() (*int32, bool)`

GetTotalRecordingDurationSecOk returns a tuple with the TotalRecordingDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordingDurationSec

`func (o *InsightsV1VideoRoomSummary) SetTotalRecordingDurationSec(v int32)`

SetTotalRecordingDurationSec sets TotalRecordingDurationSec field to given value.

### HasTotalRecordingDurationSec

`func (o *InsightsV1VideoRoomSummary) HasTotalRecordingDurationSec() bool`

HasTotalRecordingDurationSec returns a boolean if a field has been set.

### SetTotalRecordingDurationSecNil

`func (o *InsightsV1VideoRoomSummary) SetTotalRecordingDurationSecNil(b bool)`

 SetTotalRecordingDurationSecNil sets the value for TotalRecordingDurationSec to be an explicit nil

### UnsetTotalRecordingDurationSec
`func (o *InsightsV1VideoRoomSummary) UnsetTotalRecordingDurationSec()`

UnsetTotalRecordingDurationSec ensures that no value is present for TotalRecordingDurationSec, not even an explicit nil
### GetUniqueParticipantIdentities

`func (o *InsightsV1VideoRoomSummary) GetUniqueParticipantIdentities() int32`

GetUniqueParticipantIdentities returns the UniqueParticipantIdentities field if non-nil, zero value otherwise.

### GetUniqueParticipantIdentitiesOk

`func (o *InsightsV1VideoRoomSummary) GetUniqueParticipantIdentitiesOk() (*int32, bool)`

GetUniqueParticipantIdentitiesOk returns a tuple with the UniqueParticipantIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueParticipantIdentities

`func (o *InsightsV1VideoRoomSummary) SetUniqueParticipantIdentities(v int32)`

SetUniqueParticipantIdentities sets UniqueParticipantIdentities field to given value.

### HasUniqueParticipantIdentities

`func (o *InsightsV1VideoRoomSummary) HasUniqueParticipantIdentities() bool`

HasUniqueParticipantIdentities returns a boolean if a field has been set.

### SetUniqueParticipantIdentitiesNil

`func (o *InsightsV1VideoRoomSummary) SetUniqueParticipantIdentitiesNil(b bool)`

 SetUniqueParticipantIdentitiesNil sets the value for UniqueParticipantIdentities to be an explicit nil

### UnsetUniqueParticipantIdentities
`func (o *InsightsV1VideoRoomSummary) UnsetUniqueParticipantIdentities()`

UnsetUniqueParticipantIdentities ensures that no value is present for UniqueParticipantIdentities, not even an explicit nil
### GetUniqueParticipants

`func (o *InsightsV1VideoRoomSummary) GetUniqueParticipants() int32`

GetUniqueParticipants returns the UniqueParticipants field if non-nil, zero value otherwise.

### GetUniqueParticipantsOk

`func (o *InsightsV1VideoRoomSummary) GetUniqueParticipantsOk() (*int32, bool)`

GetUniqueParticipantsOk returns a tuple with the UniqueParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueParticipants

`func (o *InsightsV1VideoRoomSummary) SetUniqueParticipants(v int32)`

SetUniqueParticipants sets UniqueParticipants field to given value.

### HasUniqueParticipants

`func (o *InsightsV1VideoRoomSummary) HasUniqueParticipants() bool`

HasUniqueParticipants returns a boolean if a field has been set.

### SetUniqueParticipantsNil

`func (o *InsightsV1VideoRoomSummary) SetUniqueParticipantsNil(b bool)`

 SetUniqueParticipantsNil sets the value for UniqueParticipants to be an explicit nil

### UnsetUniqueParticipants
`func (o *InsightsV1VideoRoomSummary) UnsetUniqueParticipants()`

UnsetUniqueParticipants ensures that no value is present for UniqueParticipants, not even an explicit nil
### GetUrl

`func (o *InsightsV1VideoRoomSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InsightsV1VideoRoomSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InsightsV1VideoRoomSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InsightsV1VideoRoomSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InsightsV1VideoRoomSummary) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InsightsV1VideoRoomSummary) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


