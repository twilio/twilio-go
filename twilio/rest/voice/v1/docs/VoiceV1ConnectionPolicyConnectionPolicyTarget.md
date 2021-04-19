# VoiceV1ConnectionPolicyConnectionPolicyTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ConnectionPolicySid** | Pointer to **NullableString** | The SID of the Connection Policy that owns the Target | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the target is enabled | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Priority** | Pointer to **NullableInt32** | The relative importance of the target | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Target** | Pointer to **NullableString** | The SIP address you want Twilio to route your calls to | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**Weight** | Pointer to **NullableInt32** | The value that determines the relative load the Target should receive compared to others with the same priority | [optional] 

## Methods

### NewVoiceV1ConnectionPolicyConnectionPolicyTarget

`func NewVoiceV1ConnectionPolicyConnectionPolicyTarget() *VoiceV1ConnectionPolicyConnectionPolicyTarget`

NewVoiceV1ConnectionPolicyConnectionPolicyTarget instantiates a new VoiceV1ConnectionPolicyConnectionPolicyTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceV1ConnectionPolicyConnectionPolicyTargetWithDefaults

`func NewVoiceV1ConnectionPolicyConnectionPolicyTargetWithDefaults() *VoiceV1ConnectionPolicyConnectionPolicyTarget`

NewVoiceV1ConnectionPolicyConnectionPolicyTargetWithDefaults instantiates a new VoiceV1ConnectionPolicyConnectionPolicyTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConnectionPolicySid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetConnectionPolicySid() string`

GetConnectionPolicySid returns the ConnectionPolicySid field if non-nil, zero value otherwise.

### GetConnectionPolicySidOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetConnectionPolicySidOk() (*string, bool)`

GetConnectionPolicySidOk returns a tuple with the ConnectionPolicySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPolicySid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetConnectionPolicySid(v string)`

SetConnectionPolicySid sets ConnectionPolicySid field to given value.

### HasConnectionPolicySid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasConnectionPolicySid() bool`

HasConnectionPolicySid returns a boolean if a field has been set.

### SetConnectionPolicySidNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetConnectionPolicySidNil(b bool)`

 SetConnectionPolicySidNil sets the value for ConnectionPolicySid to be an explicit nil

### UnsetConnectionPolicySid
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetConnectionPolicySid()`

UnsetConnectionPolicySid ensures that no value is present for ConnectionPolicySid, not even an explicit nil
### GetDateCreated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnabled

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFriendlyName

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetPriority

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTarget

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetUrl

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWeight

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *VoiceV1ConnectionPolicyConnectionPolicyTarget) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


