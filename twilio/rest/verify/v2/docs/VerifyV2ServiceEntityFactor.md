# VerifyV2ServiceEntityFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account Sid. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configurations for a &#x60;factor_type&#x60;. | [optional] 
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

### NewVerifyV2ServiceEntityFactor

`func NewVerifyV2ServiceEntityFactor() *VerifyV2ServiceEntityFactor`

NewVerifyV2ServiceEntityFactor instantiates a new VerifyV2ServiceEntityFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceEntityFactorWithDefaults

`func NewVerifyV2ServiceEntityFactorWithDefaults() *VerifyV2ServiceEntityFactor`

NewVerifyV2ServiceEntityFactorWithDefaults instantiates a new VerifyV2ServiceEntityFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceEntityFactor) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceEntityFactor) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceEntityFactor) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceEntityFactor) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceEntityFactor) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceEntityFactor) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConfig

`func (o *VerifyV2ServiceEntityFactor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VerifyV2ServiceEntityFactor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VerifyV2ServiceEntityFactor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VerifyV2ServiceEntityFactor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *VerifyV2ServiceEntityFactor) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *VerifyV2ServiceEntityFactor) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceEntityFactor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceEntityFactor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceEntityFactor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceEntityFactor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceEntityFactor) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceEntityFactor) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceEntityFactor) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceEntityFactor) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceEntityFactor) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceEntityFactor) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceEntityFactor) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceEntityFactor) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEntitySid

`func (o *VerifyV2ServiceEntityFactor) GetEntitySid() string`

GetEntitySid returns the EntitySid field if non-nil, zero value otherwise.

### GetEntitySidOk

`func (o *VerifyV2ServiceEntityFactor) GetEntitySidOk() (*string, bool)`

GetEntitySidOk returns a tuple with the EntitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySid

`func (o *VerifyV2ServiceEntityFactor) SetEntitySid(v string)`

SetEntitySid sets EntitySid field to given value.

### HasEntitySid

`func (o *VerifyV2ServiceEntityFactor) HasEntitySid() bool`

HasEntitySid returns a boolean if a field has been set.

### SetEntitySidNil

`func (o *VerifyV2ServiceEntityFactor) SetEntitySidNil(b bool)`

 SetEntitySidNil sets the value for EntitySid to be an explicit nil

### UnsetEntitySid
`func (o *VerifyV2ServiceEntityFactor) UnsetEntitySid()`

UnsetEntitySid ensures that no value is present for EntitySid, not even an explicit nil
### GetFactorType

`func (o *VerifyV2ServiceEntityFactor) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *VerifyV2ServiceEntityFactor) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *VerifyV2ServiceEntityFactor) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *VerifyV2ServiceEntityFactor) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *VerifyV2ServiceEntityFactor) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *VerifyV2ServiceEntityFactor) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetFriendlyName

`func (o *VerifyV2ServiceEntityFactor) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VerifyV2ServiceEntityFactor) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VerifyV2ServiceEntityFactor) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VerifyV2ServiceEntityFactor) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VerifyV2ServiceEntityFactor) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VerifyV2ServiceEntityFactor) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentity

`func (o *VerifyV2ServiceEntityFactor) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VerifyV2ServiceEntityFactor) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VerifyV2ServiceEntityFactor) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VerifyV2ServiceEntityFactor) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *VerifyV2ServiceEntityFactor) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *VerifyV2ServiceEntityFactor) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceEntityFactor) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceEntityFactor) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceEntityFactor) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceEntityFactor) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceEntityFactor) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceEntityFactor) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceEntityFactor) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceEntityFactor) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceEntityFactor) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceEntityFactor) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceEntityFactor) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceEntityFactor) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VerifyV2ServiceEntityFactor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyV2ServiceEntityFactor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyV2ServiceEntityFactor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyV2ServiceEntityFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VerifyV2ServiceEntityFactor) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VerifyV2ServiceEntityFactor) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceEntityFactor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceEntityFactor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceEntityFactor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceEntityFactor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceEntityFactor) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceEntityFactor) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


