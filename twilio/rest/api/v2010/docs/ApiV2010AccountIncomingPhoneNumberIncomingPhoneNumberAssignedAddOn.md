# ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | A JSON string that represents the current configuration | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Description** | Pointer to **NullableString** | A short description of the Add-on functionality | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**ResourceSid** | Pointer to **NullableString** | The SID of the Phone Number that installed this Add-on | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnWithDefaults

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnWithDefaults() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnWithDefaults instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConfiguration

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDescription

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetResourceSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetResourceSid() string`

GetResourceSid returns the ResourceSid field if non-nil, zero value otherwise.

### GetResourceSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetResourceSidOk() (*string, bool)`

GetResourceSidOk returns a tuple with the ResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetResourceSid(v string)`

SetResourceSid sets ResourceSid field to given value.

### HasResourceSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasResourceSid() bool`

HasResourceSid returns a boolean if a field has been set.

### SetResourceSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetResourceSidNil(b bool)`

 SetResourceSidNil sets the value for ResourceSid to be an explicit nil

### UnsetResourceSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetResourceSid()`

UnsetResourceSid ensures that no value is present for ResourceSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUniqueName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


