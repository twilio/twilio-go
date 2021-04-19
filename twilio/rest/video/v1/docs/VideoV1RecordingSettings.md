# VideoV1RecordingSettings

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AwsCredentialsSid** | Pointer to **NullableString** | The SID of the stored Credential resource | [optional] 
**AwsS3Url** | Pointer to **NullableString** | The URL of the AWS S3 bucket where the recordings are stored | [optional] 
**AwsStorageEnabled** | Pointer to **NullableBool** | Whether all recordings are written to the aws_s3_url | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether all recordings are stored in an encrypted form | [optional] 
**EncryptionKeySid** | Pointer to **NullableString** | The SID of the Public Key resource used for encryption | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1RecordingSettings

`func NewVideoV1RecordingSettings() *VideoV1RecordingSettings`

NewVideoV1RecordingSettings instantiates a new VideoV1RecordingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RecordingSettingsWithDefaults

`func NewVideoV1RecordingSettingsWithDefaults() *VideoV1RecordingSettings`

NewVideoV1RecordingSettingsWithDefaults instantiates a new VideoV1RecordingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1RecordingSettings) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1RecordingSettings) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1RecordingSettings) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1RecordingSettings) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1RecordingSettings) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1RecordingSettings) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAwsCredentialsSid

`func (o *VideoV1RecordingSettings) GetAwsCredentialsSid() string`

GetAwsCredentialsSid returns the AwsCredentialsSid field if non-nil, zero value otherwise.

### GetAwsCredentialsSidOk

`func (o *VideoV1RecordingSettings) GetAwsCredentialsSidOk() (*string, bool)`

GetAwsCredentialsSidOk returns a tuple with the AwsCredentialsSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialsSid

`func (o *VideoV1RecordingSettings) SetAwsCredentialsSid(v string)`

SetAwsCredentialsSid sets AwsCredentialsSid field to given value.

### HasAwsCredentialsSid

`func (o *VideoV1RecordingSettings) HasAwsCredentialsSid() bool`

HasAwsCredentialsSid returns a boolean if a field has been set.

### SetAwsCredentialsSidNil

`func (o *VideoV1RecordingSettings) SetAwsCredentialsSidNil(b bool)`

 SetAwsCredentialsSidNil sets the value for AwsCredentialsSid to be an explicit nil

### UnsetAwsCredentialsSid
`func (o *VideoV1RecordingSettings) UnsetAwsCredentialsSid()`

UnsetAwsCredentialsSid ensures that no value is present for AwsCredentialsSid, not even an explicit nil
### GetAwsS3Url

`func (o *VideoV1RecordingSettings) GetAwsS3Url() string`

GetAwsS3Url returns the AwsS3Url field if non-nil, zero value otherwise.

### GetAwsS3UrlOk

`func (o *VideoV1RecordingSettings) GetAwsS3UrlOk() (*string, bool)`

GetAwsS3UrlOk returns a tuple with the AwsS3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3Url

`func (o *VideoV1RecordingSettings) SetAwsS3Url(v string)`

SetAwsS3Url sets AwsS3Url field to given value.

### HasAwsS3Url

`func (o *VideoV1RecordingSettings) HasAwsS3Url() bool`

HasAwsS3Url returns a boolean if a field has been set.

### SetAwsS3UrlNil

`func (o *VideoV1RecordingSettings) SetAwsS3UrlNil(b bool)`

 SetAwsS3UrlNil sets the value for AwsS3Url to be an explicit nil

### UnsetAwsS3Url
`func (o *VideoV1RecordingSettings) UnsetAwsS3Url()`

UnsetAwsS3Url ensures that no value is present for AwsS3Url, not even an explicit nil
### GetAwsStorageEnabled

`func (o *VideoV1RecordingSettings) GetAwsStorageEnabled() bool`

GetAwsStorageEnabled returns the AwsStorageEnabled field if non-nil, zero value otherwise.

### GetAwsStorageEnabledOk

`func (o *VideoV1RecordingSettings) GetAwsStorageEnabledOk() (*bool, bool)`

GetAwsStorageEnabledOk returns a tuple with the AwsStorageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageEnabled

`func (o *VideoV1RecordingSettings) SetAwsStorageEnabled(v bool)`

SetAwsStorageEnabled sets AwsStorageEnabled field to given value.

### HasAwsStorageEnabled

`func (o *VideoV1RecordingSettings) HasAwsStorageEnabled() bool`

HasAwsStorageEnabled returns a boolean if a field has been set.

### SetAwsStorageEnabledNil

`func (o *VideoV1RecordingSettings) SetAwsStorageEnabledNil(b bool)`

 SetAwsStorageEnabledNil sets the value for AwsStorageEnabled to be an explicit nil

### UnsetAwsStorageEnabled
`func (o *VideoV1RecordingSettings) UnsetAwsStorageEnabled()`

UnsetAwsStorageEnabled ensures that no value is present for AwsStorageEnabled, not even an explicit nil
### GetEncryptionEnabled

`func (o *VideoV1RecordingSettings) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *VideoV1RecordingSettings) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *VideoV1RecordingSettings) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *VideoV1RecordingSettings) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *VideoV1RecordingSettings) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *VideoV1RecordingSettings) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetEncryptionKeySid

`func (o *VideoV1RecordingSettings) GetEncryptionKeySid() string`

GetEncryptionKeySid returns the EncryptionKeySid field if non-nil, zero value otherwise.

### GetEncryptionKeySidOk

`func (o *VideoV1RecordingSettings) GetEncryptionKeySidOk() (*string, bool)`

GetEncryptionKeySidOk returns a tuple with the EncryptionKeySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeySid

`func (o *VideoV1RecordingSettings) SetEncryptionKeySid(v string)`

SetEncryptionKeySid sets EncryptionKeySid field to given value.

### HasEncryptionKeySid

`func (o *VideoV1RecordingSettings) HasEncryptionKeySid() bool`

HasEncryptionKeySid returns a boolean if a field has been set.

### SetEncryptionKeySidNil

`func (o *VideoV1RecordingSettings) SetEncryptionKeySidNil(b bool)`

 SetEncryptionKeySidNil sets the value for EncryptionKeySid to be an explicit nil

### UnsetEncryptionKeySid
`func (o *VideoV1RecordingSettings) UnsetEncryptionKeySid()`

UnsetEncryptionKeySid ensures that no value is present for EncryptionKeySid, not even an explicit nil
### GetFriendlyName

`func (o *VideoV1RecordingSettings) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VideoV1RecordingSettings) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VideoV1RecordingSettings) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VideoV1RecordingSettings) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VideoV1RecordingSettings) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VideoV1RecordingSettings) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetUrl

`func (o *VideoV1RecordingSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1RecordingSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1RecordingSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1RecordingSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1RecordingSettings) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1RecordingSettings) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


