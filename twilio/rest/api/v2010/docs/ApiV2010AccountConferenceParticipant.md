# ApiV2010AccountConferenceParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**CallSidToCoach** | Pointer to **NullableString** | The SID of the participant who is being &#x60;coached&#x60; | [optional] 
**Coaching** | Pointer to **NullableBool** | Indicates if the participant changed to coach | [optional] 
**ConferenceSid** | Pointer to **NullableString** | The SID of the conference the participant is in | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**EndConferenceOnExit** | Pointer to **NullableBool** | Whether the conference ends when the participant leaves | [optional] 
**Hold** | Pointer to **NullableBool** | Whether the participant is on hold | [optional] 
**Label** | Pointer to **NullableString** | The label of this participant | [optional] 
**Muted** | Pointer to **NullableBool** | Whether the participant is muted | [optional] 
**StartConferenceOnEnter** | Pointer to **NullableBool** | Whether the conference starts when the participant joins the conference | [optional] 
**Status** | Pointer to **NullableString** | The status of the participant&#39;s call in a session | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountConferenceParticipant

`func NewApiV2010AccountConferenceParticipant() *ApiV2010AccountConferenceParticipant`

NewApiV2010AccountConferenceParticipant instantiates a new ApiV2010AccountConferenceParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountConferenceParticipantWithDefaults

`func NewApiV2010AccountConferenceParticipantWithDefaults() *ApiV2010AccountConferenceParticipant`

NewApiV2010AccountConferenceParticipantWithDefaults instantiates a new ApiV2010AccountConferenceParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountConferenceParticipant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountConferenceParticipant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountConferenceParticipant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountConferenceParticipant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountConferenceParticipant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountConferenceParticipant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallSid

`func (o *ApiV2010AccountConferenceParticipant) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountConferenceParticipant) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountConferenceParticipant) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountConferenceParticipant) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountConferenceParticipant) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountConferenceParticipant) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetCallSidToCoach

`func (o *ApiV2010AccountConferenceParticipant) GetCallSidToCoach() string`

GetCallSidToCoach returns the CallSidToCoach field if non-nil, zero value otherwise.

### GetCallSidToCoachOk

`func (o *ApiV2010AccountConferenceParticipant) GetCallSidToCoachOk() (*string, bool)`

GetCallSidToCoachOk returns a tuple with the CallSidToCoach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSidToCoach

`func (o *ApiV2010AccountConferenceParticipant) SetCallSidToCoach(v string)`

SetCallSidToCoach sets CallSidToCoach field to given value.

### HasCallSidToCoach

`func (o *ApiV2010AccountConferenceParticipant) HasCallSidToCoach() bool`

HasCallSidToCoach returns a boolean if a field has been set.

### SetCallSidToCoachNil

`func (o *ApiV2010AccountConferenceParticipant) SetCallSidToCoachNil(b bool)`

 SetCallSidToCoachNil sets the value for CallSidToCoach to be an explicit nil

### UnsetCallSidToCoach
`func (o *ApiV2010AccountConferenceParticipant) UnsetCallSidToCoach()`

UnsetCallSidToCoach ensures that no value is present for CallSidToCoach, not even an explicit nil
### GetCoaching

`func (o *ApiV2010AccountConferenceParticipant) GetCoaching() bool`

GetCoaching returns the Coaching field if non-nil, zero value otherwise.

### GetCoachingOk

`func (o *ApiV2010AccountConferenceParticipant) GetCoachingOk() (*bool, bool)`

GetCoachingOk returns a tuple with the Coaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaching

`func (o *ApiV2010AccountConferenceParticipant) SetCoaching(v bool)`

SetCoaching sets Coaching field to given value.

### HasCoaching

`func (o *ApiV2010AccountConferenceParticipant) HasCoaching() bool`

HasCoaching returns a boolean if a field has been set.

### SetCoachingNil

`func (o *ApiV2010AccountConferenceParticipant) SetCoachingNil(b bool)`

 SetCoachingNil sets the value for Coaching to be an explicit nil

### UnsetCoaching
`func (o *ApiV2010AccountConferenceParticipant) UnsetCoaching()`

UnsetCoaching ensures that no value is present for Coaching, not even an explicit nil
### GetConferenceSid

`func (o *ApiV2010AccountConferenceParticipant) GetConferenceSid() string`

GetConferenceSid returns the ConferenceSid field if non-nil, zero value otherwise.

### GetConferenceSidOk

`func (o *ApiV2010AccountConferenceParticipant) GetConferenceSidOk() (*string, bool)`

GetConferenceSidOk returns a tuple with the ConferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceSid

`func (o *ApiV2010AccountConferenceParticipant) SetConferenceSid(v string)`

SetConferenceSid sets ConferenceSid field to given value.

### HasConferenceSid

`func (o *ApiV2010AccountConferenceParticipant) HasConferenceSid() bool`

HasConferenceSid returns a boolean if a field has been set.

### SetConferenceSidNil

`func (o *ApiV2010AccountConferenceParticipant) SetConferenceSidNil(b bool)`

 SetConferenceSidNil sets the value for ConferenceSid to be an explicit nil

### UnsetConferenceSid
`func (o *ApiV2010AccountConferenceParticipant) UnsetConferenceSid()`

UnsetConferenceSid ensures that no value is present for ConferenceSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountConferenceParticipant) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountConferenceParticipant) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountConferenceParticipant) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountConferenceParticipant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountConferenceParticipant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountConferenceParticipant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountConferenceParticipant) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountConferenceParticipant) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountConferenceParticipant) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountConferenceParticipant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountConferenceParticipant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountConferenceParticipant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndConferenceOnExit

`func (o *ApiV2010AccountConferenceParticipant) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *ApiV2010AccountConferenceParticipant) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *ApiV2010AccountConferenceParticipant) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.

### HasEndConferenceOnExit

`func (o *ApiV2010AccountConferenceParticipant) HasEndConferenceOnExit() bool`

HasEndConferenceOnExit returns a boolean if a field has been set.

### SetEndConferenceOnExitNil

`func (o *ApiV2010AccountConferenceParticipant) SetEndConferenceOnExitNil(b bool)`

 SetEndConferenceOnExitNil sets the value for EndConferenceOnExit to be an explicit nil

### UnsetEndConferenceOnExit
`func (o *ApiV2010AccountConferenceParticipant) UnsetEndConferenceOnExit()`

UnsetEndConferenceOnExit ensures that no value is present for EndConferenceOnExit, not even an explicit nil
### GetHold

`func (o *ApiV2010AccountConferenceParticipant) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *ApiV2010AccountConferenceParticipant) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *ApiV2010AccountConferenceParticipant) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *ApiV2010AccountConferenceParticipant) HasHold() bool`

HasHold returns a boolean if a field has been set.

### SetHoldNil

`func (o *ApiV2010AccountConferenceParticipant) SetHoldNil(b bool)`

 SetHoldNil sets the value for Hold to be an explicit nil

### UnsetHold
`func (o *ApiV2010AccountConferenceParticipant) UnsetHold()`

UnsetHold ensures that no value is present for Hold, not even an explicit nil
### GetLabel

`func (o *ApiV2010AccountConferenceParticipant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiV2010AccountConferenceParticipant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiV2010AccountConferenceParticipant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiV2010AccountConferenceParticipant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ApiV2010AccountConferenceParticipant) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ApiV2010AccountConferenceParticipant) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetMuted

`func (o *ApiV2010AccountConferenceParticipant) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ApiV2010AccountConferenceParticipant) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ApiV2010AccountConferenceParticipant) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ApiV2010AccountConferenceParticipant) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### SetMutedNil

`func (o *ApiV2010AccountConferenceParticipant) SetMutedNil(b bool)`

 SetMutedNil sets the value for Muted to be an explicit nil

### UnsetMuted
`func (o *ApiV2010AccountConferenceParticipant) UnsetMuted()`

UnsetMuted ensures that no value is present for Muted, not even an explicit nil
### GetStartConferenceOnEnter

`func (o *ApiV2010AccountConferenceParticipant) GetStartConferenceOnEnter() bool`

GetStartConferenceOnEnter returns the StartConferenceOnEnter field if non-nil, zero value otherwise.

### GetStartConferenceOnEnterOk

`func (o *ApiV2010AccountConferenceParticipant) GetStartConferenceOnEnterOk() (*bool, bool)`

GetStartConferenceOnEnterOk returns a tuple with the StartConferenceOnEnter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConferenceOnEnter

`func (o *ApiV2010AccountConferenceParticipant) SetStartConferenceOnEnter(v bool)`

SetStartConferenceOnEnter sets StartConferenceOnEnter field to given value.

### HasStartConferenceOnEnter

`func (o *ApiV2010AccountConferenceParticipant) HasStartConferenceOnEnter() bool`

HasStartConferenceOnEnter returns a boolean if a field has been set.

### SetStartConferenceOnEnterNil

`func (o *ApiV2010AccountConferenceParticipant) SetStartConferenceOnEnterNil(b bool)`

 SetStartConferenceOnEnterNil sets the value for StartConferenceOnEnter to be an explicit nil

### UnsetStartConferenceOnEnter
`func (o *ApiV2010AccountConferenceParticipant) UnsetStartConferenceOnEnter()`

UnsetStartConferenceOnEnter ensures that no value is present for StartConferenceOnEnter, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountConferenceParticipant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountConferenceParticipant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountConferenceParticipant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountConferenceParticipant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountConferenceParticipant) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountConferenceParticipant) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountConferenceParticipant) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountConferenceParticipant) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountConferenceParticipant) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountConferenceParticipant) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountConferenceParticipant) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountConferenceParticipant) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


