# ApiV2010AccountSipSipCredentialList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique sid that identifies this account | [optional] 
**DateCreated** | Pointer to **NullableString** | The date this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The date this resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | Human readable descriptive text | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this credential | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | The list of credentials associated with this credential list. | [optional] 
**Uri** | Pointer to **NullableString** | The URI for this resource | [optional] 

## Methods

### NewApiV2010AccountSipSipCredentialList

`func NewApiV2010AccountSipSipCredentialList() *ApiV2010AccountSipSipCredentialList`

NewApiV2010AccountSipSipCredentialList instantiates a new ApiV2010AccountSipSipCredentialList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountSipSipCredentialListWithDefaults

`func NewApiV2010AccountSipSipCredentialListWithDefaults() *ApiV2010AccountSipSipCredentialList`

NewApiV2010AccountSipSipCredentialListWithDefaults instantiates a new ApiV2010AccountSipSipCredentialList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountSipSipCredentialList) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountSipSipCredentialList) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountSipSipCredentialList) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountSipSipCredentialList) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountSipSipCredentialList) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountSipSipCredentialList) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountSipSipCredentialList) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountSipSipCredentialList) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountSipSipCredentialList) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountSipSipCredentialList) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountSipSipCredentialList) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountSipSipCredentialList) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountSipSipCredentialList) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountSipSipCredentialList) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountSipSipCredentialList) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountSipSipCredentialList) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountSipSipCredentialList) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountSipSipCredentialList) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountSipSipCredentialList) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountSipSipCredentialList) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountSipSipCredentialList) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountSipSipCredentialList) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountSipSipCredentialList) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountSipSipCredentialList) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountSipSipCredentialList) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountSipSipCredentialList) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountSipSipCredentialList) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountSipSipCredentialList) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountSipSipCredentialList) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountSipSipCredentialList) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountSipSipCredentialList) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountSipSipCredentialList) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountSipSipCredentialList) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountSipSipCredentialList) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountSipSipCredentialList) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountSipSipCredentialList) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountSipSipCredentialList) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountSipSipCredentialList) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountSipSipCredentialList) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountSipSipCredentialList) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountSipSipCredentialList) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountSipSipCredentialList) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


