# VerifyV2ServiceEntityNewFactor

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account Sid. | [optional] 
**Binding** | Pointer to **map[string]interface{}** | Unique external identifier of the Entity | [optional] 
**Config** | Pointer to **map[string]interface{}** | Binding for a &#x60;factor_type&#x60;. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date this Factor was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Factor was updated | [optional] 
**EntitySid** | Pointer to **NullableString** | Entity Sid. | [optional] 
**FactorType** | Pointer to **NullableString** | The Type of this Factor | [optional] 
**FriendlyName** | Pointer to **NullableString** | A human readable description of this resource. | [optional] 
**Identity** | Pointer to **NullableString** | Unique external identifier of the Entity | [optional] 
**ServiceSid** | Pointer to **NullableString** | Service Sid. | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Factor. | [optional] 
**Status** | Pointer to **NullableString** | The Status of this Factor | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewVerifyV2ServiceEntityNewFactor

`func NewVerifyV2ServiceEntityNewFactor() *VerifyV2ServiceEntityNewFactor`

NewVerifyV2ServiceEntityNewFactor instantiates a new VerifyV2ServiceEntityNewFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceEntityNewFactorWithDefaults

`func NewVerifyV2ServiceEntityNewFactorWithDefaults() *VerifyV2ServiceEntityNewFactor`

NewVerifyV2ServiceEntityNewFactorWithDefaults instantiates a new VerifyV2ServiceEntityNewFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceEntityNewFactor) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceEntityNewFactor) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceEntityNewFactor) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceEntityNewFactor) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceEntityNewFactor) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceEntityNewFactor) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBinding

`func (o *VerifyV2ServiceEntityNewFactor) GetBinding() map[string]interface{}`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *VerifyV2ServiceEntityNewFactor) GetBindingOk() (*map[string]interface{}, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *VerifyV2ServiceEntityNewFactor) SetBinding(v map[string]interface{})`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *VerifyV2ServiceEntityNewFactor) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### SetBindingNil

`func (o *VerifyV2ServiceEntityNewFactor) SetBindingNil(b bool)`

 SetBindingNil sets the value for Binding to be an explicit nil

### UnsetBinding
`func (o *VerifyV2ServiceEntityNewFactor) UnsetBinding()`

UnsetBinding ensures that no value is present for Binding, not even an explicit nil
### GetConfig

`func (o *VerifyV2ServiceEntityNewFactor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VerifyV2ServiceEntityNewFactor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VerifyV2ServiceEntityNewFactor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VerifyV2ServiceEntityNewFactor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *VerifyV2ServiceEntityNewFactor) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *VerifyV2ServiceEntityNewFactor) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceEntityNewFactor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceEntityNewFactor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceEntityNewFactor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceEntityNewFactor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceEntityNewFactor) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceEntityNewFactor) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceEntityNewFactor) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceEntityNewFactor) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceEntityNewFactor) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceEntityNewFactor) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceEntityNewFactor) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceEntityNewFactor) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEntitySid

`func (o *VerifyV2ServiceEntityNewFactor) GetEntitySid() string`

GetEntitySid returns the EntitySid field if non-nil, zero value otherwise.

### GetEntitySidOk

`func (o *VerifyV2ServiceEntityNewFactor) GetEntitySidOk() (*string, bool)`

GetEntitySidOk returns a tuple with the EntitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySid

`func (o *VerifyV2ServiceEntityNewFactor) SetEntitySid(v string)`

SetEntitySid sets EntitySid field to given value.

### HasEntitySid

`func (o *VerifyV2ServiceEntityNewFactor) HasEntitySid() bool`

HasEntitySid returns a boolean if a field has been set.

### SetEntitySidNil

`func (o *VerifyV2ServiceEntityNewFactor) SetEntitySidNil(b bool)`

 SetEntitySidNil sets the value for EntitySid to be an explicit nil

### UnsetEntitySid
`func (o *VerifyV2ServiceEntityNewFactor) UnsetEntitySid()`

UnsetEntitySid ensures that no value is present for EntitySid, not even an explicit nil
### GetFactorType

`func (o *VerifyV2ServiceEntityNewFactor) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *VerifyV2ServiceEntityNewFactor) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *VerifyV2ServiceEntityNewFactor) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *VerifyV2ServiceEntityNewFactor) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *VerifyV2ServiceEntityNewFactor) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *VerifyV2ServiceEntityNewFactor) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetFriendlyName

`func (o *VerifyV2ServiceEntityNewFactor) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VerifyV2ServiceEntityNewFactor) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VerifyV2ServiceEntityNewFactor) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VerifyV2ServiceEntityNewFactor) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VerifyV2ServiceEntityNewFactor) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VerifyV2ServiceEntityNewFactor) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentity

`func (o *VerifyV2ServiceEntityNewFactor) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VerifyV2ServiceEntityNewFactor) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VerifyV2ServiceEntityNewFactor) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VerifyV2ServiceEntityNewFactor) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *VerifyV2ServiceEntityNewFactor) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *VerifyV2ServiceEntityNewFactor) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceEntityNewFactor) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceEntityNewFactor) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceEntityNewFactor) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceEntityNewFactor) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceEntityNewFactor) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceEntityNewFactor) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceEntityNewFactor) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceEntityNewFactor) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceEntityNewFactor) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceEntityNewFactor) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceEntityNewFactor) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceEntityNewFactor) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceEntityNewFactor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceEntityNewFactor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceEntityNewFactor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceEntityNewFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceEntityNewFactor) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceEntityNewFactor) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceEntityNewFactor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceEntityNewFactor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceEntityNewFactor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceEntityNewFactor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceEntityNewFactor) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceEntityNewFactor) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


