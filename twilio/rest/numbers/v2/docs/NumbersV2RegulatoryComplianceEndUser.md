# NumbersV2RegulatoryComplianceEndUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The set of parameters that compose the End Users resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Type** | Pointer to **NullableString** | The type of end user of the Bundle resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the End User resource | [optional] 

## Methods

### NewNumbersV2RegulatoryComplianceEndUser

`func NewNumbersV2RegulatoryComplianceEndUser() *NumbersV2RegulatoryComplianceEndUser`

NewNumbersV2RegulatoryComplianceEndUser instantiates a new NumbersV2RegulatoryComplianceEndUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersV2RegulatoryComplianceEndUserWithDefaults

`func NewNumbersV2RegulatoryComplianceEndUserWithDefaults() *NumbersV2RegulatoryComplianceEndUser`

NewNumbersV2RegulatoryComplianceEndUserWithDefaults instantiates a new NumbersV2RegulatoryComplianceEndUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NumbersV2RegulatoryComplianceEndUser) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NumbersV2RegulatoryComplianceEndUser) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NumbersV2RegulatoryComplianceEndUser) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *NumbersV2RegulatoryComplianceEndUser) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NumbersV2RegulatoryComplianceEndUser) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NumbersV2RegulatoryComplianceEndUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDateCreated

`func (o *NumbersV2RegulatoryComplianceEndUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NumbersV2RegulatoryComplianceEndUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NumbersV2RegulatoryComplianceEndUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NumbersV2RegulatoryComplianceEndUser) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NumbersV2RegulatoryComplianceEndUser) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NumbersV2RegulatoryComplianceEndUser) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *NumbersV2RegulatoryComplianceEndUser) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NumbersV2RegulatoryComplianceEndUser) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NumbersV2RegulatoryComplianceEndUser) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetSid

`func (o *NumbersV2RegulatoryComplianceEndUser) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NumbersV2RegulatoryComplianceEndUser) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NumbersV2RegulatoryComplianceEndUser) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *NumbersV2RegulatoryComplianceEndUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumbersV2RegulatoryComplianceEndUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NumbersV2RegulatoryComplianceEndUser) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *NumbersV2RegulatoryComplianceEndUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NumbersV2RegulatoryComplianceEndUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NumbersV2RegulatoryComplianceEndUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NumbersV2RegulatoryComplianceEndUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NumbersV2RegulatoryComplianceEndUser) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NumbersV2RegulatoryComplianceEndUser) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


