# FlexV1FlexFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ChannelType** | Pointer to **NullableString** | The channel type | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the chat service | [optional] 
**ContactIdentity** | Pointer to **NullableString** | The channel contact&#39;s Identity | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the Flex Flow is enabled | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Integration** | Pointer to **map[string]interface{}** | An object that contains specific parameters for the integration | [optional] 
**IntegrationType** | Pointer to **NullableString** | The integration type | [optional] 
**JanitorEnabled** | Pointer to **NullableBool** | Remove active Proxy sessions if the corresponding Task is deleted. | [optional] 
**LongLived** | Pointer to **NullableBool** | Re-use this chat channel for future interactions with a contact | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Flex Flow resource | [optional] 

## Methods

### NewFlexV1FlexFlow

`func NewFlexV1FlexFlow() *FlexV1FlexFlow`

NewFlexV1FlexFlow instantiates a new FlexV1FlexFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexV1FlexFlowWithDefaults

`func NewFlexV1FlexFlowWithDefaults() *FlexV1FlexFlow`

NewFlexV1FlexFlowWithDefaults instantiates a new FlexV1FlexFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FlexV1FlexFlow) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FlexV1FlexFlow) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FlexV1FlexFlow) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FlexV1FlexFlow) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FlexV1FlexFlow) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FlexV1FlexFlow) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelType

`func (o *FlexV1FlexFlow) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *FlexV1FlexFlow) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *FlexV1FlexFlow) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *FlexV1FlexFlow) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### SetChannelTypeNil

`func (o *FlexV1FlexFlow) SetChannelTypeNil(b bool)`

 SetChannelTypeNil sets the value for ChannelType to be an explicit nil

### UnsetChannelType
`func (o *FlexV1FlexFlow) UnsetChannelType()`

UnsetChannelType ensures that no value is present for ChannelType, not even an explicit nil
### GetChatServiceSid

`func (o *FlexV1FlexFlow) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *FlexV1FlexFlow) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *FlexV1FlexFlow) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *FlexV1FlexFlow) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *FlexV1FlexFlow) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *FlexV1FlexFlow) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetContactIdentity

`func (o *FlexV1FlexFlow) GetContactIdentity() string`

GetContactIdentity returns the ContactIdentity field if non-nil, zero value otherwise.

### GetContactIdentityOk

`func (o *FlexV1FlexFlow) GetContactIdentityOk() (*string, bool)`

GetContactIdentityOk returns a tuple with the ContactIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIdentity

`func (o *FlexV1FlexFlow) SetContactIdentity(v string)`

SetContactIdentity sets ContactIdentity field to given value.

### HasContactIdentity

`func (o *FlexV1FlexFlow) HasContactIdentity() bool`

HasContactIdentity returns a boolean if a field has been set.

### SetContactIdentityNil

`func (o *FlexV1FlexFlow) SetContactIdentityNil(b bool)`

 SetContactIdentityNil sets the value for ContactIdentity to be an explicit nil

### UnsetContactIdentity
`func (o *FlexV1FlexFlow) UnsetContactIdentity()`

UnsetContactIdentity ensures that no value is present for ContactIdentity, not even an explicit nil
### GetDateCreated

`func (o *FlexV1FlexFlow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FlexV1FlexFlow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FlexV1FlexFlow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FlexV1FlexFlow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FlexV1FlexFlow) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FlexV1FlexFlow) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FlexV1FlexFlow) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FlexV1FlexFlow) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FlexV1FlexFlow) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FlexV1FlexFlow) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FlexV1FlexFlow) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FlexV1FlexFlow) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnabled

`func (o *FlexV1FlexFlow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlexV1FlexFlow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlexV1FlexFlow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlexV1FlexFlow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *FlexV1FlexFlow) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *FlexV1FlexFlow) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFriendlyName

`func (o *FlexV1FlexFlow) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *FlexV1FlexFlow) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *FlexV1FlexFlow) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *FlexV1FlexFlow) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *FlexV1FlexFlow) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *FlexV1FlexFlow) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIntegration

`func (o *FlexV1FlexFlow) GetIntegration() map[string]interface{}`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *FlexV1FlexFlow) GetIntegrationOk() (*map[string]interface{}, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *FlexV1FlexFlow) SetIntegration(v map[string]interface{})`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *FlexV1FlexFlow) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *FlexV1FlexFlow) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *FlexV1FlexFlow) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetIntegrationType

`func (o *FlexV1FlexFlow) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *FlexV1FlexFlow) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *FlexV1FlexFlow) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *FlexV1FlexFlow) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### SetIntegrationTypeNil

`func (o *FlexV1FlexFlow) SetIntegrationTypeNil(b bool)`

 SetIntegrationTypeNil sets the value for IntegrationType to be an explicit nil

### UnsetIntegrationType
`func (o *FlexV1FlexFlow) UnsetIntegrationType()`

UnsetIntegrationType ensures that no value is present for IntegrationType, not even an explicit nil
### GetJanitorEnabled

`func (o *FlexV1FlexFlow) GetJanitorEnabled() bool`

GetJanitorEnabled returns the JanitorEnabled field if non-nil, zero value otherwise.

### GetJanitorEnabledOk

`func (o *FlexV1FlexFlow) GetJanitorEnabledOk() (*bool, bool)`

GetJanitorEnabledOk returns a tuple with the JanitorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJanitorEnabled

`func (o *FlexV1FlexFlow) SetJanitorEnabled(v bool)`

SetJanitorEnabled sets JanitorEnabled field to given value.

### HasJanitorEnabled

`func (o *FlexV1FlexFlow) HasJanitorEnabled() bool`

HasJanitorEnabled returns a boolean if a field has been set.

### SetJanitorEnabledNil

`func (o *FlexV1FlexFlow) SetJanitorEnabledNil(b bool)`

 SetJanitorEnabledNil sets the value for JanitorEnabled to be an explicit nil

### UnsetJanitorEnabled
`func (o *FlexV1FlexFlow) UnsetJanitorEnabled()`

UnsetJanitorEnabled ensures that no value is present for JanitorEnabled, not even an explicit nil
### GetLongLived

`func (o *FlexV1FlexFlow) GetLongLived() bool`

GetLongLived returns the LongLived field if non-nil, zero value otherwise.

### GetLongLivedOk

`func (o *FlexV1FlexFlow) GetLongLivedOk() (*bool, bool)`

GetLongLivedOk returns a tuple with the LongLived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLived

`func (o *FlexV1FlexFlow) SetLongLived(v bool)`

SetLongLived sets LongLived field to given value.

### HasLongLived

`func (o *FlexV1FlexFlow) HasLongLived() bool`

HasLongLived returns a boolean if a field has been set.

### SetLongLivedNil

`func (o *FlexV1FlexFlow) SetLongLivedNil(b bool)`

 SetLongLivedNil sets the value for LongLived to be an explicit nil

### UnsetLongLived
`func (o *FlexV1FlexFlow) UnsetLongLived()`

UnsetLongLived ensures that no value is present for LongLived, not even an explicit nil
### GetSid

`func (o *FlexV1FlexFlow) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *FlexV1FlexFlow) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *FlexV1FlexFlow) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *FlexV1FlexFlow) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *FlexV1FlexFlow) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *FlexV1FlexFlow) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *FlexV1FlexFlow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FlexV1FlexFlow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FlexV1FlexFlow) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FlexV1FlexFlow) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FlexV1FlexFlow) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FlexV1FlexFlow) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


