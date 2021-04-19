# VideoV1CompositionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AwsCredentialsSid** | Pointer to **NullableString** | The SID of the stored Credential resource | [optional] 
**AwsS3Url** | Pointer to **NullableString** | The URL of the AWS S3 bucket where the compositions are stored | [optional] 
**AwsStorageEnabled** | Pointer to **NullableBool** | Whether all compositions are written to the aws_s3_url | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether all compositions are stored in an encrypted form | [optional] 
**EncryptionKeySid** | Pointer to **NullableString** | The SID of the Public Key resource used for encryption | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1CompositionSettings

`func NewVideoV1CompositionSettings() *VideoV1CompositionSettings`

NewVideoV1CompositionSettings instantiates a new VideoV1CompositionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1CompositionSettingsWithDefaults

`func NewVideoV1CompositionSettingsWithDefaults() *VideoV1CompositionSettings`

NewVideoV1CompositionSettingsWithDefaults instantiates a new VideoV1CompositionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1CompositionSettings) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1CompositionSettings) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1CompositionSettings) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1CompositionSettings) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1CompositionSettings) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1CompositionSettings) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAwsCredentialsSid

`func (o *VideoV1CompositionSettings) GetAwsCredentialsSid() string`

GetAwsCredentialsSid returns the AwsCredentialsSid field if non-nil, zero value otherwise.

### GetAwsCredentialsSidOk

`func (o *VideoV1CompositionSettings) GetAwsCredentialsSidOk() (*string, bool)`

GetAwsCredentialsSidOk returns a tuple with the AwsCredentialsSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialsSid

`func (o *VideoV1CompositionSettings) SetAwsCredentialsSid(v string)`

SetAwsCredentialsSid sets AwsCredentialsSid field to given value.

### HasAwsCredentialsSid

`func (o *VideoV1CompositionSettings) HasAwsCredentialsSid() bool`

HasAwsCredentialsSid returns a boolean if a field has been set.

### SetAwsCredentialsSidNil

`func (o *VideoV1CompositionSettings) SetAwsCredentialsSidNil(b bool)`

 SetAwsCredentialsSidNil sets the value for AwsCredentialsSid to be an explicit nil

### UnsetAwsCredentialsSid
`func (o *VideoV1CompositionSettings) UnsetAwsCredentialsSid()`

UnsetAwsCredentialsSid ensures that no value is present for AwsCredentialsSid, not even an explicit nil
### GetAwsS3Url

`func (o *VideoV1CompositionSettings) GetAwsS3Url() string`

GetAwsS3Url returns the AwsS3Url field if non-nil, zero value otherwise.

### GetAwsS3UrlOk

`func (o *VideoV1CompositionSettings) GetAwsS3UrlOk() (*string, bool)`

GetAwsS3UrlOk returns a tuple with the AwsS3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3Url

`func (o *VideoV1CompositionSettings) SetAwsS3Url(v string)`

SetAwsS3Url sets AwsS3Url field to given value.

### HasAwsS3Url

`func (o *VideoV1CompositionSettings) HasAwsS3Url() bool`

HasAwsS3Url returns a boolean if a field has been set.

### SetAwsS3UrlNil

`func (o *VideoV1CompositionSettings) SetAwsS3UrlNil(b bool)`

 SetAwsS3UrlNil sets the value for AwsS3Url to be an explicit nil

### UnsetAwsS3Url
`func (o *VideoV1CompositionSettings) UnsetAwsS3Url()`

UnsetAwsS3Url ensures that no value is present for AwsS3Url, not even an explicit nil
### GetAwsStorageEnabled

`func (o *VideoV1CompositionSettings) GetAwsStorageEnabled() bool`

GetAwsStorageEnabled returns the AwsStorageEnabled field if non-nil, zero value otherwise.

### GetAwsStorageEnabledOk

`func (o *VideoV1CompositionSettings) GetAwsStorageEnabledOk() (*bool, bool)`

GetAwsStorageEnabledOk returns a tuple with the AwsStorageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageEnabled

`func (o *VideoV1CompositionSettings) SetAwsStorageEnabled(v bool)`

SetAwsStorageEnabled sets AwsStorageEnabled field to given value.

### HasAwsStorageEnabled

`func (o *VideoV1CompositionSettings) HasAwsStorageEnabled() bool`

HasAwsStorageEnabled returns a boolean if a field has been set.

### SetAwsStorageEnabledNil

`func (o *VideoV1CompositionSettings) SetAwsStorageEnabledNil(b bool)`

 SetAwsStorageEnabledNil sets the value for AwsStorageEnabled to be an explicit nil

### UnsetAwsStorageEnabled
`func (o *VideoV1CompositionSettings) UnsetAwsStorageEnabled()`

UnsetAwsStorageEnabled ensures that no value is present for AwsStorageEnabled, not even an explicit nil
### GetEncryptionEnabled

`func (o *VideoV1CompositionSettings) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *VideoV1CompositionSettings) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *VideoV1CompositionSettings) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *VideoV1CompositionSettings) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *VideoV1CompositionSettings) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *VideoV1CompositionSettings) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetEncryptionKeySid

`func (o *VideoV1CompositionSettings) GetEncryptionKeySid() string`

GetEncryptionKeySid returns the EncryptionKeySid field if non-nil, zero value otherwise.

### GetEncryptionKeySidOk

`func (o *VideoV1CompositionSettings) GetEncryptionKeySidOk() (*string, bool)`

GetEncryptionKeySidOk returns a tuple with the EncryptionKeySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeySid

`func (o *VideoV1CompositionSettings) SetEncryptionKeySid(v string)`

SetEncryptionKeySid sets EncryptionKeySid field to given value.

### HasEncryptionKeySid

`func (o *VideoV1CompositionSettings) HasEncryptionKeySid() bool`

HasEncryptionKeySid returns a boolean if a field has been set.

### SetEncryptionKeySidNil

`func (o *VideoV1CompositionSettings) SetEncryptionKeySidNil(b bool)`

 SetEncryptionKeySidNil sets the value for EncryptionKeySid to be an explicit nil

### UnsetEncryptionKeySid
`func (o *VideoV1CompositionSettings) UnsetEncryptionKeySid()`

UnsetEncryptionKeySid ensures that no value is present for EncryptionKeySid, not even an explicit nil
### GetFriendlyName

`func (o *VideoV1CompositionSettings) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VideoV1CompositionSettings) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VideoV1CompositionSettings) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VideoV1CompositionSettings) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VideoV1CompositionSettings) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VideoV1CompositionSettings) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetUrl

`func (o *VideoV1CompositionSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1CompositionSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1CompositionSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1CompositionSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1CompositionSettings) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1CompositionSettings) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


