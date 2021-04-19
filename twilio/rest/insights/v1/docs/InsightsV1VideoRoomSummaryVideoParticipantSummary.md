# InsightsV1VideoRoomSummaryVideoParticipantSummary

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account SID associated with the room. | [optional] 
**Codecs** | Pointer to **[]string** | Codecs detected from the participant. | [optional] 
**DurationSec** | Pointer to **NullableInt32** | Amount of time in seconds the participant was in the room. | [optional] 
**EdgeLocation** | Pointer to **NullableString** | Name of the edge location the participant connected to. | [optional] 
**EndReason** | Pointer to **NullableString** | Reason the participant left the room. | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | Errors encountered by the participant. | [optional] 
**ErrorCodeUrl** | Pointer to **NullableString** | Twilio error code dictionary link. | [optional] 
**JoinTime** | Pointer to **NullableTime** | When the participant joined the room. | [optional] 
**LeaveTime** | Pointer to **NullableTime** | When the participant left the room | [optional] 
**MediaRegion** | Pointer to **NullableString** | Twilio media region the participant connected to. | [optional] 
**ParticipantIdentity** | Pointer to **NullableString** | The application-defined string that uniquely identifies the participant within a Room. | [optional] 
**ParticipantSid** | Pointer to **NullableString** | Unique identifier for the participant. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Object containing information about the participant&#39;s data from the room. | [optional] 
**PublisherInfo** | Pointer to **map[string]interface{}** | Object containing information about the SDK name and version. | [optional] 
**RoomSid** | Pointer to **NullableString** | Unique identifier for the room. | [optional] 
**Status** | Pointer to **NullableString** | Status of the room. | [optional] 
**Url** | Pointer to **NullableString** | URL of the participant resource. | [optional] 

## Methods

### NewInsightsV1VideoRoomSummaryVideoParticipantSummary

`func NewInsightsV1VideoRoomSummaryVideoParticipantSummary() *InsightsV1VideoRoomSummaryVideoParticipantSummary`

NewInsightsV1VideoRoomSummaryVideoParticipantSummary instantiates a new InsightsV1VideoRoomSummaryVideoParticipantSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsV1VideoRoomSummaryVideoParticipantSummaryWithDefaults

`func NewInsightsV1VideoRoomSummaryVideoParticipantSummaryWithDefaults() *InsightsV1VideoRoomSummaryVideoParticipantSummary`

NewInsightsV1VideoRoomSummaryVideoParticipantSummaryWithDefaults instantiates a new InsightsV1VideoRoomSummaryVideoParticipantSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCodecs

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### SetCodecsNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetCodecsNil(b bool)`

 SetCodecsNil sets the value for Codecs to be an explicit nil

### UnsetCodecs
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetCodecs()`

UnsetCodecs ensures that no value is present for Codecs, not even an explicit nil
### GetDurationSec

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.

### HasDurationSec

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasDurationSec() bool`

HasDurationSec returns a boolean if a field has been set.

### SetDurationSecNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetDurationSecNil(b bool)`

 SetDurationSecNil sets the value for DurationSec to be an explicit nil

### UnsetDurationSec
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetDurationSec()`

UnsetDurationSec ensures that no value is present for DurationSec, not even an explicit nil
### GetEdgeLocation

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetEdgeLocation() string`

GetEdgeLocation returns the EdgeLocation field if non-nil, zero value otherwise.

### GetEdgeLocationOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetEdgeLocationOk() (*string, bool)`

GetEdgeLocationOk returns a tuple with the EdgeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeLocation

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetEdgeLocation(v string)`

SetEdgeLocation sets EdgeLocation field to given value.

### HasEdgeLocation

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasEdgeLocation() bool`

HasEdgeLocation returns a boolean if a field has been set.

### SetEdgeLocationNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetEdgeLocationNil(b bool)`

 SetEdgeLocationNil sets the value for EdgeLocation to be an explicit nil

### UnsetEdgeLocation
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetEdgeLocation()`

UnsetEdgeLocation ensures that no value is present for EdgeLocation, not even an explicit nil
### GetEndReason

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### SetEndReasonNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetEndReasonNil(b bool)`

 SetEndReasonNil sets the value for EndReason to be an explicit nil

### UnsetEndReason
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetEndReason()`

UnsetEndReason ensures that no value is present for EndReason, not even an explicit nil
### GetErrorCode

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorCodeUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetErrorCodeUrl() string`

GetErrorCodeUrl returns the ErrorCodeUrl field if non-nil, zero value otherwise.

### GetErrorCodeUrlOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetErrorCodeUrlOk() (*string, bool)`

GetErrorCodeUrlOk returns a tuple with the ErrorCodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCodeUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetErrorCodeUrl(v string)`

SetErrorCodeUrl sets ErrorCodeUrl field to given value.

### HasErrorCodeUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasErrorCodeUrl() bool`

HasErrorCodeUrl returns a boolean if a field has been set.

### SetErrorCodeUrlNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetErrorCodeUrlNil(b bool)`

 SetErrorCodeUrlNil sets the value for ErrorCodeUrl to be an explicit nil

### UnsetErrorCodeUrl
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetErrorCodeUrl()`

UnsetErrorCodeUrl ensures that no value is present for ErrorCodeUrl, not even an explicit nil
### GetJoinTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetJoinTime() time.Time`

GetJoinTime returns the JoinTime field if non-nil, zero value otherwise.

### GetJoinTimeOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetJoinTimeOk() (*time.Time, bool)`

GetJoinTimeOk returns a tuple with the JoinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetJoinTime(v time.Time)`

SetJoinTime sets JoinTime field to given value.

### HasJoinTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasJoinTime() bool`

HasJoinTime returns a boolean if a field has been set.

### SetJoinTimeNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetJoinTimeNil(b bool)`

 SetJoinTimeNil sets the value for JoinTime to be an explicit nil

### UnsetJoinTime
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetJoinTime()`

UnsetJoinTime ensures that no value is present for JoinTime, not even an explicit nil
### GetLeaveTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetLeaveTime() time.Time`

GetLeaveTime returns the LeaveTime field if non-nil, zero value otherwise.

### GetLeaveTimeOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetLeaveTimeOk() (*time.Time, bool)`

GetLeaveTimeOk returns a tuple with the LeaveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetLeaveTime(v time.Time)`

SetLeaveTime sets LeaveTime field to given value.

### HasLeaveTime

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasLeaveTime() bool`

HasLeaveTime returns a boolean if a field has been set.

### SetLeaveTimeNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetLeaveTimeNil(b bool)`

 SetLeaveTimeNil sets the value for LeaveTime to be an explicit nil

### UnsetLeaveTime
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetLeaveTime()`

UnsetLeaveTime ensures that no value is present for LeaveTime, not even an explicit nil
### GetMediaRegion

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetMediaRegion() string`

GetMediaRegion returns the MediaRegion field if non-nil, zero value otherwise.

### GetMediaRegionOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetMediaRegionOk() (*string, bool)`

GetMediaRegionOk returns a tuple with the MediaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaRegion

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetMediaRegion(v string)`

SetMediaRegion sets MediaRegion field to given value.

### HasMediaRegion

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasMediaRegion() bool`

HasMediaRegion returns a boolean if a field has been set.

### SetMediaRegionNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetMediaRegionNil(b bool)`

 SetMediaRegionNil sets the value for MediaRegion to be an explicit nil

### UnsetMediaRegion
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetMediaRegion()`

UnsetMediaRegion ensures that no value is present for MediaRegion, not even an explicit nil
### GetParticipantIdentity

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetParticipantIdentity() string`

GetParticipantIdentity returns the ParticipantIdentity field if non-nil, zero value otherwise.

### GetParticipantIdentityOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetParticipantIdentityOk() (*string, bool)`

GetParticipantIdentityOk returns a tuple with the ParticipantIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantIdentity

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetParticipantIdentity(v string)`

SetParticipantIdentity sets ParticipantIdentity field to given value.

### HasParticipantIdentity

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasParticipantIdentity() bool`

HasParticipantIdentity returns a boolean if a field has been set.

### SetParticipantIdentityNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetParticipantIdentityNil(b bool)`

 SetParticipantIdentityNil sets the value for ParticipantIdentity to be an explicit nil

### UnsetParticipantIdentity
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetParticipantIdentity()`

UnsetParticipantIdentity ensures that no value is present for ParticipantIdentity, not even an explicit nil
### GetParticipantSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetProperties

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPublisherInfo

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetPublisherInfo() map[string]interface{}`

GetPublisherInfo returns the PublisherInfo field if non-nil, zero value otherwise.

### GetPublisherInfoOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetPublisherInfoOk() (*map[string]interface{}, bool)`

GetPublisherInfoOk returns a tuple with the PublisherInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherInfo

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetPublisherInfo(v map[string]interface{})`

SetPublisherInfo sets PublisherInfo field to given value.

### HasPublisherInfo

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasPublisherInfo() bool`

HasPublisherInfo returns a boolean if a field has been set.

### SetPublisherInfoNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetPublisherInfoNil(b bool)`

 SetPublisherInfoNil sets the value for PublisherInfo to be an explicit nil

### UnsetPublisherInfo
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetPublisherInfo()`

UnsetPublisherInfo ensures that no value is present for PublisherInfo, not even an explicit nil
### GetRoomSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetStatus

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InsightsV1VideoRoomSummaryVideoParticipantSummary) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


