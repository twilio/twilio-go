# ApiV2010AccountRecordingRecordingAddOnResult

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AddOnConfigurationSid** | Pointer to **NullableString** | The SID of the Add-on configuration | [optional] 
**AddOnSid** | Pointer to **NullableString** | The SID of the Add-on to which the result belongs | [optional] 
**DateCompleted** | Pointer to **NullableString** | The date and time in GMT that the result was completed | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**ReferenceSid** | Pointer to **NullableString** | The SID of the recording to which the AddOnResult resource belongs | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the result | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 

## Methods

### NewApiV2010AccountRecordingRecordingAddOnResult

`func NewApiV2010AccountRecordingRecordingAddOnResult() *ApiV2010AccountRecordingRecordingAddOnResult`

NewApiV2010AccountRecordingRecordingAddOnResult instantiates a new ApiV2010AccountRecordingRecordingAddOnResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountRecordingRecordingAddOnResultWithDefaults

`func NewApiV2010AccountRecordingRecordingAddOnResultWithDefaults() *ApiV2010AccountRecordingRecordingAddOnResult`

NewApiV2010AccountRecordingRecordingAddOnResultWithDefaults instantiates a new ApiV2010AccountRecordingRecordingAddOnResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAddOnConfigurationSid() string`

GetAddOnConfigurationSid returns the AddOnConfigurationSid field if non-nil, zero value otherwise.

### GetAddOnConfigurationSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAddOnConfigurationSidOk() (*string, bool)`

GetAddOnConfigurationSidOk returns a tuple with the AddOnConfigurationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAddOnConfigurationSid(v string)`

SetAddOnConfigurationSid sets AddOnConfigurationSid field to given value.

### HasAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasAddOnConfigurationSid() bool`

HasAddOnConfigurationSid returns a boolean if a field has been set.

### SetAddOnConfigurationSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAddOnConfigurationSidNil(b bool)`

 SetAddOnConfigurationSidNil sets the value for AddOnConfigurationSid to be an explicit nil

### UnsetAddOnConfigurationSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetAddOnConfigurationSid()`

UnsetAddOnConfigurationSid ensures that no value is present for AddOnConfigurationSid, not even an explicit nil
### GetAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAddOnSid() string`

GetAddOnSid returns the AddOnSid field if non-nil, zero value otherwise.

### GetAddOnSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetAddOnSidOk() (*string, bool)`

GetAddOnSidOk returns a tuple with the AddOnSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAddOnSid(v string)`

SetAddOnSid sets AddOnSid field to given value.

### HasAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasAddOnSid() bool`

HasAddOnSid returns a boolean if a field has been set.

### SetAddOnSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetAddOnSidNil(b bool)`

 SetAddOnSidNil sets the value for AddOnSid to be an explicit nil

### UnsetAddOnSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetAddOnSid()`

UnsetAddOnSid ensures that no value is present for AddOnSid, not even an explicit nil
### GetDateCompleted

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateCompleted() string`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateCompletedOk() (*string, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateCompleted(v string)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### SetDateCompletedNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateCompletedNil(b bool)`

 SetDateCompletedNil sets the value for DateCompleted to be an explicit nil

### UnsetDateCompleted
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetDateCompleted()`

UnsetDateCompleted ensures that no value is present for DateCompleted, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetReferenceSid() string`

GetReferenceSid returns the ReferenceSid field if non-nil, zero value otherwise.

### GetReferenceSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetReferenceSidOk() (*string, bool)`

GetReferenceSidOk returns a tuple with the ReferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetReferenceSid(v string)`

SetReferenceSid sets ReferenceSid field to given value.

### HasReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasReferenceSid() bool`

HasReferenceSid returns a boolean if a field has been set.

### SetReferenceSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetReferenceSidNil(b bool)`

 SetReferenceSidNil sets the value for ReferenceSid to be an explicit nil

### UnsetReferenceSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetReferenceSid()`

UnsetReferenceSid ensures that no value is present for ReferenceSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResult) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountRecordingRecordingAddOnResult) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


