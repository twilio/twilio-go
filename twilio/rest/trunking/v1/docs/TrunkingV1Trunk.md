# TrunkingV1Trunk

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AuthType** | Pointer to **NullableString** | The types of authentication mapped to the domain | [optional] 
**AuthTypeSet** | Pointer to **[]string** | Reserved | [optional] 
**CnamLookupEnabled** | Pointer to **NullableBool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DisasterRecoveryMethod** | Pointer to **NullableString** | The HTTP method we use to call the disaster_recovery_url | [optional] 
**DisasterRecoveryUrl** | Pointer to **NullableString** | The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL | [optional] 
**DomainName** | Pointer to **NullableString** | The unique address you reserve on Twilio to which you route your SIP traffic | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Recording** | Pointer to **map[string]interface{}** | The recording settings for the trunk | [optional] 
**Secure** | Pointer to **NullableBool** | Whether Secure Trunking is enabled for the trunk | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TransferMode** | Pointer to **NullableString** | The call transfer settings for the trunk | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewTrunkingV1Trunk

`func NewTrunkingV1Trunk() *TrunkingV1Trunk`

NewTrunkingV1Trunk instantiates a new TrunkingV1Trunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkingV1TrunkWithDefaults

`func NewTrunkingV1TrunkWithDefaults() *TrunkingV1Trunk`

NewTrunkingV1TrunkWithDefaults instantiates a new TrunkingV1Trunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrunkingV1Trunk) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrunkingV1Trunk) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrunkingV1Trunk) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrunkingV1Trunk) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrunkingV1Trunk) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrunkingV1Trunk) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAuthType

`func (o *TrunkingV1Trunk) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TrunkingV1Trunk) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TrunkingV1Trunk) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *TrunkingV1Trunk) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *TrunkingV1Trunk) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *TrunkingV1Trunk) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetAuthTypeSet

`func (o *TrunkingV1Trunk) GetAuthTypeSet() []string`

GetAuthTypeSet returns the AuthTypeSet field if non-nil, zero value otherwise.

### GetAuthTypeSetOk

`func (o *TrunkingV1Trunk) GetAuthTypeSetOk() (*[]string, bool)`

GetAuthTypeSetOk returns a tuple with the AuthTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTypeSet

`func (o *TrunkingV1Trunk) SetAuthTypeSet(v []string)`

SetAuthTypeSet sets AuthTypeSet field to given value.

### HasAuthTypeSet

`func (o *TrunkingV1Trunk) HasAuthTypeSet() bool`

HasAuthTypeSet returns a boolean if a field has been set.

### SetAuthTypeSetNil

`func (o *TrunkingV1Trunk) SetAuthTypeSetNil(b bool)`

 SetAuthTypeSetNil sets the value for AuthTypeSet to be an explicit nil

### UnsetAuthTypeSet
`func (o *TrunkingV1Trunk) UnsetAuthTypeSet()`

UnsetAuthTypeSet ensures that no value is present for AuthTypeSet, not even an explicit nil
### GetCnamLookupEnabled

`func (o *TrunkingV1Trunk) GetCnamLookupEnabled() bool`

GetCnamLookupEnabled returns the CnamLookupEnabled field if non-nil, zero value otherwise.

### GetCnamLookupEnabledOk

`func (o *TrunkingV1Trunk) GetCnamLookupEnabledOk() (*bool, bool)`

GetCnamLookupEnabledOk returns a tuple with the CnamLookupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamLookupEnabled

`func (o *TrunkingV1Trunk) SetCnamLookupEnabled(v bool)`

SetCnamLookupEnabled sets CnamLookupEnabled field to given value.

### HasCnamLookupEnabled

`func (o *TrunkingV1Trunk) HasCnamLookupEnabled() bool`

HasCnamLookupEnabled returns a boolean if a field has been set.

### SetCnamLookupEnabledNil

`func (o *TrunkingV1Trunk) SetCnamLookupEnabledNil(b bool)`

 SetCnamLookupEnabledNil sets the value for CnamLookupEnabled to be an explicit nil

### UnsetCnamLookupEnabled
`func (o *TrunkingV1Trunk) UnsetCnamLookupEnabled()`

UnsetCnamLookupEnabled ensures that no value is present for CnamLookupEnabled, not even an explicit nil
### GetDateCreated

`func (o *TrunkingV1Trunk) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrunkingV1Trunk) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrunkingV1Trunk) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrunkingV1Trunk) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrunkingV1Trunk) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrunkingV1Trunk) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrunkingV1Trunk) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrunkingV1Trunk) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrunkingV1Trunk) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrunkingV1Trunk) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrunkingV1Trunk) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrunkingV1Trunk) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDisasterRecoveryMethod

`func (o *TrunkingV1Trunk) GetDisasterRecoveryMethod() string`

GetDisasterRecoveryMethod returns the DisasterRecoveryMethod field if non-nil, zero value otherwise.

### GetDisasterRecoveryMethodOk

`func (o *TrunkingV1Trunk) GetDisasterRecoveryMethodOk() (*string, bool)`

GetDisasterRecoveryMethodOk returns a tuple with the DisasterRecoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRecoveryMethod

`func (o *TrunkingV1Trunk) SetDisasterRecoveryMethod(v string)`

SetDisasterRecoveryMethod sets DisasterRecoveryMethod field to given value.

### HasDisasterRecoveryMethod

`func (o *TrunkingV1Trunk) HasDisasterRecoveryMethod() bool`

HasDisasterRecoveryMethod returns a boolean if a field has been set.

### SetDisasterRecoveryMethodNil

`func (o *TrunkingV1Trunk) SetDisasterRecoveryMethodNil(b bool)`

 SetDisasterRecoveryMethodNil sets the value for DisasterRecoveryMethod to be an explicit nil

### UnsetDisasterRecoveryMethod
`func (o *TrunkingV1Trunk) UnsetDisasterRecoveryMethod()`

UnsetDisasterRecoveryMethod ensures that no value is present for DisasterRecoveryMethod, not even an explicit nil
### GetDisasterRecoveryUrl

`func (o *TrunkingV1Trunk) GetDisasterRecoveryUrl() string`

GetDisasterRecoveryUrl returns the DisasterRecoveryUrl field if non-nil, zero value otherwise.

### GetDisasterRecoveryUrlOk

`func (o *TrunkingV1Trunk) GetDisasterRecoveryUrlOk() (*string, bool)`

GetDisasterRecoveryUrlOk returns a tuple with the DisasterRecoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRecoveryUrl

`func (o *TrunkingV1Trunk) SetDisasterRecoveryUrl(v string)`

SetDisasterRecoveryUrl sets DisasterRecoveryUrl field to given value.

### HasDisasterRecoveryUrl

`func (o *TrunkingV1Trunk) HasDisasterRecoveryUrl() bool`

HasDisasterRecoveryUrl returns a boolean if a field has been set.

### SetDisasterRecoveryUrlNil

`func (o *TrunkingV1Trunk) SetDisasterRecoveryUrlNil(b bool)`

 SetDisasterRecoveryUrlNil sets the value for DisasterRecoveryUrl to be an explicit nil

### UnsetDisasterRecoveryUrl
`func (o *TrunkingV1Trunk) UnsetDisasterRecoveryUrl()`

UnsetDisasterRecoveryUrl ensures that no value is present for DisasterRecoveryUrl, not even an explicit nil
### GetDomainName

`func (o *TrunkingV1Trunk) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *TrunkingV1Trunk) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *TrunkingV1Trunk) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *TrunkingV1Trunk) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *TrunkingV1Trunk) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *TrunkingV1Trunk) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetFriendlyName

`func (o *TrunkingV1Trunk) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrunkingV1Trunk) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrunkingV1Trunk) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrunkingV1Trunk) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrunkingV1Trunk) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrunkingV1Trunk) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TrunkingV1Trunk) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrunkingV1Trunk) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrunkingV1Trunk) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrunkingV1Trunk) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TrunkingV1Trunk) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TrunkingV1Trunk) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRecording

`func (o *TrunkingV1Trunk) GetRecording() map[string]interface{}`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *TrunkingV1Trunk) GetRecordingOk() (*map[string]interface{}, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *TrunkingV1Trunk) SetRecording(v map[string]interface{})`

SetRecording sets Recording field to given value.

### HasRecording

`func (o *TrunkingV1Trunk) HasRecording() bool`

HasRecording returns a boolean if a field has been set.

### SetRecordingNil

`func (o *TrunkingV1Trunk) SetRecordingNil(b bool)`

 SetRecordingNil sets the value for Recording to be an explicit nil

### UnsetRecording
`func (o *TrunkingV1Trunk) UnsetRecording()`

UnsetRecording ensures that no value is present for Recording, not even an explicit nil
### GetSecure

`func (o *TrunkingV1Trunk) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *TrunkingV1Trunk) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *TrunkingV1Trunk) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *TrunkingV1Trunk) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### SetSecureNil

`func (o *TrunkingV1Trunk) SetSecureNil(b bool)`

 SetSecureNil sets the value for Secure to be an explicit nil

### UnsetSecure
`func (o *TrunkingV1Trunk) UnsetSecure()`

UnsetSecure ensures that no value is present for Secure, not even an explicit nil
### GetSid

`func (o *TrunkingV1Trunk) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrunkingV1Trunk) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrunkingV1Trunk) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrunkingV1Trunk) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrunkingV1Trunk) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrunkingV1Trunk) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTransferMode

`func (o *TrunkingV1Trunk) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *TrunkingV1Trunk) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *TrunkingV1Trunk) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *TrunkingV1Trunk) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### SetTransferModeNil

`func (o *TrunkingV1Trunk) SetTransferModeNil(b bool)`

 SetTransferModeNil sets the value for TransferMode to be an explicit nil

### UnsetTransferMode
`func (o *TrunkingV1Trunk) UnsetTransferMode()`

UnsetTransferMode ensures that no value is present for TransferMode, not even an explicit nil
### GetUrl

`func (o *TrunkingV1Trunk) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrunkingV1Trunk) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrunkingV1Trunk) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrunkingV1Trunk) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrunkingV1Trunk) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrunkingV1Trunk) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


