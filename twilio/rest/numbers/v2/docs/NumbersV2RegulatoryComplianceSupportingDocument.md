# NumbersV2RegulatoryComplianceSupportingDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | The set of parameters that compose the Supporting Documents resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**MimeType** | Pointer to **NullableString** | The image type of the file | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The verification status of the Supporting Document resource | [optional] 
**Type** | Pointer to **NullableString** | The type of the Supporting Document | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Supporting Document resource | [optional] 

## Methods

### NewNumbersV2RegulatoryComplianceSupportingDocument

`func NewNumbersV2RegulatoryComplianceSupportingDocument() *NumbersV2RegulatoryComplianceSupportingDocument`

NewNumbersV2RegulatoryComplianceSupportingDocument instantiates a new NumbersV2RegulatoryComplianceSupportingDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersV2RegulatoryComplianceSupportingDocumentWithDefaults

`func NewNumbersV2RegulatoryComplianceSupportingDocumentWithDefaults() *NumbersV2RegulatoryComplianceSupportingDocument`

NewNumbersV2RegulatoryComplianceSupportingDocumentWithDefaults instantiates a new NumbersV2RegulatoryComplianceSupportingDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDateCreated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetMimeType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NumbersV2RegulatoryComplianceSupportingDocument) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NumbersV2RegulatoryComplianceSupportingDocument) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


