# ApiV2010AccountCall

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created this resource | [optional] 
**Annotation** | Pointer to **NullableString** | The annotation provided for the call | [optional] 
**AnsweredBy** | Pointer to **NullableString** | Either &#x60;human&#x60; or &#x60;machine&#x60; if this call was initiated with answering machine detection. Empty otherwise. | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API Version used to create the call | [optional] 
**CallerName** | Pointer to **NullableString** | The caller&#39;s name if this call was an incoming call to a phone number with caller ID Lookup enabled. Otherwise, empty. | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was last updated | [optional] 
**Direction** | Pointer to **NullableString** | A string describing the direction of the call. &#x60;inbound&#x60; for inbound calls, &#x60;outbound-api&#x60; for calls initiated via the REST API or &#x60;outbound-dial&#x60; for calls initiated by a &#x60;Dial&#x60; verb. | [optional] 
**Duration** | Pointer to **NullableString** | The length of the call in seconds. | [optional] 
**EndTime** | Pointer to **NullableString** | The end time of the call. Null if the call did not complete successfully. | [optional] 
**ForwardedFrom** | Pointer to **NullableString** | The forwarding phone number if this call was an incoming call forwarded from another number (depends on carrier supporting forwarding). Otherwise, empty. | [optional] 
**From** | Pointer to **NullableString** | The phone number, SIP address or Client identifier that made this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. | [optional] 
**FromFormatted** | Pointer to **NullableString** | The calling phone number, SIP address, or Client identifier formatted for display. | [optional] 
**GroupSid** | Pointer to **NullableString** | The Group SID associated with this call. If no Group is associated with the call, the field is empty. | [optional] 
**ParentCallSid** | Pointer to **NullableString** | The SID that identifies the call that created this leg. | [optional] 
**PhoneNumberSid** | Pointer to **NullableString** | If the call was inbound, this is the SID of the IncomingPhoneNumber resource that received the call. If the call was outbound, it is the SID of the OutgoingCallerId resource from which the call was placed. | [optional] 
**Price** | Pointer to **NullableString** | The charge for this call, in the currency associated with the account. Populated after the call is completed. May not be immediately available. | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which &#x60;Price&#x60; is measured. | [optional] 
**QueueTime** | Pointer to **NullableString** | The wait time in milliseconds before the call is placed. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies this resource | [optional] 
**StartTime** | Pointer to **NullableString** | The start time of the call. Null if the call has not yet been dialed. | [optional] 
**Status** | Pointer to **NullableString** | The status of this call. | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related subresources identified by their relative URIs | [optional] 
**To** | Pointer to **NullableString** | The phone number, SIP address or Client identifier that received this call. Phone numbers are in E.164 format (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. | [optional] 
**ToFormatted** | Pointer to **NullableString** | The phone number, SIP address or Client identifier that received this call. Formatted for display. | [optional] 
**TrunkSid** | Pointer to **NullableString** | The (optional) unique identifier of the trunk resource that was used for this call. | [optional] 
**Uri** | Pointer to **NullableString** | The URI of this resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountCall

`func NewApiV2010AccountCall() *ApiV2010AccountCall`

NewApiV2010AccountCall instantiates a new ApiV2010AccountCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountCallWithDefaults

`func NewApiV2010AccountCallWithDefaults() *ApiV2010AccountCall`

NewApiV2010AccountCallWithDefaults instantiates a new ApiV2010AccountCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountCall) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountCall) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountCall) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountCall) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountCall) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountCall) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAnnotation

`func (o *ApiV2010AccountCall) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ApiV2010AccountCall) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ApiV2010AccountCall) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ApiV2010AccountCall) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ApiV2010AccountCall) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ApiV2010AccountCall) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
### GetAnsweredBy

`func (o *ApiV2010AccountCall) GetAnsweredBy() string`

GetAnsweredBy returns the AnsweredBy field if non-nil, zero value otherwise.

### GetAnsweredByOk

`func (o *ApiV2010AccountCall) GetAnsweredByOk() (*string, bool)`

GetAnsweredByOk returns a tuple with the AnsweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredBy

`func (o *ApiV2010AccountCall) SetAnsweredBy(v string)`

SetAnsweredBy sets AnsweredBy field to given value.

### HasAnsweredBy

`func (o *ApiV2010AccountCall) HasAnsweredBy() bool`

HasAnsweredBy returns a boolean if a field has been set.

### SetAnsweredByNil

`func (o *ApiV2010AccountCall) SetAnsweredByNil(b bool)`

 SetAnsweredByNil sets the value for AnsweredBy to be an explicit nil

### UnsetAnsweredBy
`func (o *ApiV2010AccountCall) UnsetAnsweredBy()`

UnsetAnsweredBy ensures that no value is present for AnsweredBy, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountCall) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountCall) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountCall) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountCall) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountCall) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountCall) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetCallerName

`func (o *ApiV2010AccountCall) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *ApiV2010AccountCall) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *ApiV2010AccountCall) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *ApiV2010AccountCall) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### SetCallerNameNil

`func (o *ApiV2010AccountCall) SetCallerNameNil(b bool)`

 SetCallerNameNil sets the value for CallerName to be an explicit nil

### UnsetCallerName
`func (o *ApiV2010AccountCall) UnsetCallerName()`

UnsetCallerName ensures that no value is present for CallerName, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountCall) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountCall) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountCall) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountCall) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountCall) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountCall) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountCall) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountCall) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountCall) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountCall) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountCall) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountCall) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDirection

`func (o *ApiV2010AccountCall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ApiV2010AccountCall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ApiV2010AccountCall) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ApiV2010AccountCall) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *ApiV2010AccountCall) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *ApiV2010AccountCall) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetDuration

`func (o *ApiV2010AccountCall) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2010AccountCall) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2010AccountCall) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2010AccountCall) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2010AccountCall) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2010AccountCall) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndTime

`func (o *ApiV2010AccountCall) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApiV2010AccountCall) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApiV2010AccountCall) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApiV2010AccountCall) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ApiV2010AccountCall) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ApiV2010AccountCall) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetForwardedFrom

`func (o *ApiV2010AccountCall) GetForwardedFrom() string`

GetForwardedFrom returns the ForwardedFrom field if non-nil, zero value otherwise.

### GetForwardedFromOk

`func (o *ApiV2010AccountCall) GetForwardedFromOk() (*string, bool)`

GetForwardedFromOk returns a tuple with the ForwardedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedFrom

`func (o *ApiV2010AccountCall) SetForwardedFrom(v string)`

SetForwardedFrom sets ForwardedFrom field to given value.

### HasForwardedFrom

`func (o *ApiV2010AccountCall) HasForwardedFrom() bool`

HasForwardedFrom returns a boolean if a field has been set.

### SetForwardedFromNil

`func (o *ApiV2010AccountCall) SetForwardedFromNil(b bool)`

 SetForwardedFromNil sets the value for ForwardedFrom to be an explicit nil

### UnsetForwardedFrom
`func (o *ApiV2010AccountCall) UnsetForwardedFrom()`

UnsetForwardedFrom ensures that no value is present for ForwardedFrom, not even an explicit nil
### GetFrom

`func (o *ApiV2010AccountCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiV2010AccountCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiV2010AccountCall) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiV2010AccountCall) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ApiV2010AccountCall) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ApiV2010AccountCall) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetFromFormatted

`func (o *ApiV2010AccountCall) GetFromFormatted() string`

GetFromFormatted returns the FromFormatted field if non-nil, zero value otherwise.

### GetFromFormattedOk

`func (o *ApiV2010AccountCall) GetFromFormattedOk() (*string, bool)`

GetFromFormattedOk returns a tuple with the FromFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFormatted

`func (o *ApiV2010AccountCall) SetFromFormatted(v string)`

SetFromFormatted sets FromFormatted field to given value.

### HasFromFormatted

`func (o *ApiV2010AccountCall) HasFromFormatted() bool`

HasFromFormatted returns a boolean if a field has been set.

### SetFromFormattedNil

`func (o *ApiV2010AccountCall) SetFromFormattedNil(b bool)`

 SetFromFormattedNil sets the value for FromFormatted to be an explicit nil

### UnsetFromFormatted
`func (o *ApiV2010AccountCall) UnsetFromFormatted()`

UnsetFromFormatted ensures that no value is present for FromFormatted, not even an explicit nil
### GetGroupSid

`func (o *ApiV2010AccountCall) GetGroupSid() string`

GetGroupSid returns the GroupSid field if non-nil, zero value otherwise.

### GetGroupSidOk

`func (o *ApiV2010AccountCall) GetGroupSidOk() (*string, bool)`

GetGroupSidOk returns a tuple with the GroupSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSid

`func (o *ApiV2010AccountCall) SetGroupSid(v string)`

SetGroupSid sets GroupSid field to given value.

### HasGroupSid

`func (o *ApiV2010AccountCall) HasGroupSid() bool`

HasGroupSid returns a boolean if a field has been set.

### SetGroupSidNil

`func (o *ApiV2010AccountCall) SetGroupSidNil(b bool)`

 SetGroupSidNil sets the value for GroupSid to be an explicit nil

### UnsetGroupSid
`func (o *ApiV2010AccountCall) UnsetGroupSid()`

UnsetGroupSid ensures that no value is present for GroupSid, not even an explicit nil
### GetParentCallSid

`func (o *ApiV2010AccountCall) GetParentCallSid() string`

GetParentCallSid returns the ParentCallSid field if non-nil, zero value otherwise.

### GetParentCallSidOk

`func (o *ApiV2010AccountCall) GetParentCallSidOk() (*string, bool)`

GetParentCallSidOk returns a tuple with the ParentCallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCallSid

`func (o *ApiV2010AccountCall) SetParentCallSid(v string)`

SetParentCallSid sets ParentCallSid field to given value.

### HasParentCallSid

`func (o *ApiV2010AccountCall) HasParentCallSid() bool`

HasParentCallSid returns a boolean if a field has been set.

### SetParentCallSidNil

`func (o *ApiV2010AccountCall) SetParentCallSidNil(b bool)`

 SetParentCallSidNil sets the value for ParentCallSid to be an explicit nil

### UnsetParentCallSid
`func (o *ApiV2010AccountCall) UnsetParentCallSid()`

UnsetParentCallSid ensures that no value is present for ParentCallSid, not even an explicit nil
### GetPhoneNumberSid

`func (o *ApiV2010AccountCall) GetPhoneNumberSid() string`

GetPhoneNumberSid returns the PhoneNumberSid field if non-nil, zero value otherwise.

### GetPhoneNumberSidOk

`func (o *ApiV2010AccountCall) GetPhoneNumberSidOk() (*string, bool)`

GetPhoneNumberSidOk returns a tuple with the PhoneNumberSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberSid

`func (o *ApiV2010AccountCall) SetPhoneNumberSid(v string)`

SetPhoneNumberSid sets PhoneNumberSid field to given value.

### HasPhoneNumberSid

`func (o *ApiV2010AccountCall) HasPhoneNumberSid() bool`

HasPhoneNumberSid returns a boolean if a field has been set.

### SetPhoneNumberSidNil

`func (o *ApiV2010AccountCall) SetPhoneNumberSidNil(b bool)`

 SetPhoneNumberSidNil sets the value for PhoneNumberSid to be an explicit nil

### UnsetPhoneNumberSid
`func (o *ApiV2010AccountCall) UnsetPhoneNumberSid()`

UnsetPhoneNumberSid ensures that no value is present for PhoneNumberSid, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountCall) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountCall) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountCall) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountCall) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountCall) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountCall) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountCall) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountCall) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountCall) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountCall) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountCall) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountCall) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetQueueTime

`func (o *ApiV2010AccountCall) GetQueueTime() string`

GetQueueTime returns the QueueTime field if non-nil, zero value otherwise.

### GetQueueTimeOk

`func (o *ApiV2010AccountCall) GetQueueTimeOk() (*string, bool)`

GetQueueTimeOk returns a tuple with the QueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTime

`func (o *ApiV2010AccountCall) SetQueueTime(v string)`

SetQueueTime sets QueueTime field to given value.

### HasQueueTime

`func (o *ApiV2010AccountCall) HasQueueTime() bool`

HasQueueTime returns a boolean if a field has been set.

### SetQueueTimeNil

`func (o *ApiV2010AccountCall) SetQueueTimeNil(b bool)`

 SetQueueTimeNil sets the value for QueueTime to be an explicit nil

### UnsetQueueTime
`func (o *ApiV2010AccountCall) UnsetQueueTime()`

UnsetQueueTime ensures that no value is present for QueueTime, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountCall) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountCall) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountCall) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountCall) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountCall) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountCall) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStartTime

`func (o *ApiV2010AccountCall) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApiV2010AccountCall) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApiV2010AccountCall) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApiV2010AccountCall) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ApiV2010AccountCall) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ApiV2010AccountCall) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountCall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountCall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountCall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountCall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountCall) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountCall) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountCall) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountCall) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountCall) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountCall) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountCall) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountCall) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetTo

`func (o *ApiV2010AccountCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ApiV2010AccountCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ApiV2010AccountCall) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ApiV2010AccountCall) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ApiV2010AccountCall) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ApiV2010AccountCall) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetToFormatted

`func (o *ApiV2010AccountCall) GetToFormatted() string`

GetToFormatted returns the ToFormatted field if non-nil, zero value otherwise.

### GetToFormattedOk

`func (o *ApiV2010AccountCall) GetToFormattedOk() (*string, bool)`

GetToFormattedOk returns a tuple with the ToFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFormatted

`func (o *ApiV2010AccountCall) SetToFormatted(v string)`

SetToFormatted sets ToFormatted field to given value.

### HasToFormatted

`func (o *ApiV2010AccountCall) HasToFormatted() bool`

HasToFormatted returns a boolean if a field has been set.

### SetToFormattedNil

`func (o *ApiV2010AccountCall) SetToFormattedNil(b bool)`

 SetToFormattedNil sets the value for ToFormatted to be an explicit nil

### UnsetToFormatted
`func (o *ApiV2010AccountCall) UnsetToFormatted()`

UnsetToFormatted ensures that no value is present for ToFormatted, not even an explicit nil
### GetTrunkSid

`func (o *ApiV2010AccountCall) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *ApiV2010AccountCall) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *ApiV2010AccountCall) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *ApiV2010AccountCall) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *ApiV2010AccountCall) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *ApiV2010AccountCall) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountCall) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountCall) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountCall) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountCall) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountCall) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountCall) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


