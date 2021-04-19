# ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AddOnConfigurationSid** | Pointer to **NullableString** | The SID of the Add-on configuration | [optional] 
**AddOnResultSid** | Pointer to **NullableString** | The SID of the AddOnResult to which the payload belongs | [optional] 
**AddOnSid** | Pointer to **NullableString** | The SID of the Add-on to which the result belongs | [optional] 
**ContentType** | Pointer to **NullableString** | The MIME type of the payload | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Label** | Pointer to **NullableString** | The string that describes the payload | [optional] 
**ReferenceSid** | Pointer to **NullableString** | The SID of the recording to which the AddOnResult resource that contains the payload belongs | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 

## Methods

### NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload

`func NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload() *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload`

NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload instantiates a new ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayloadWithDefaults

`func NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayloadWithDefaults() *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload`

NewApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayloadWithDefaults instantiates a new ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnConfigurationSid() string`

GetAddOnConfigurationSid returns the AddOnConfigurationSid field if non-nil, zero value otherwise.

### GetAddOnConfigurationSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnConfigurationSidOk() (*string, bool)`

GetAddOnConfigurationSidOk returns a tuple with the AddOnConfigurationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnConfigurationSid(v string)`

SetAddOnConfigurationSid sets AddOnConfigurationSid field to given value.

### HasAddOnConfigurationSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasAddOnConfigurationSid() bool`

HasAddOnConfigurationSid returns a boolean if a field has been set.

### SetAddOnConfigurationSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnConfigurationSidNil(b bool)`

 SetAddOnConfigurationSidNil sets the value for AddOnConfigurationSid to be an explicit nil

### UnsetAddOnConfigurationSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetAddOnConfigurationSid()`

UnsetAddOnConfigurationSid ensures that no value is present for AddOnConfigurationSid, not even an explicit nil
### GetAddOnResultSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnResultSid() string`

GetAddOnResultSid returns the AddOnResultSid field if non-nil, zero value otherwise.

### GetAddOnResultSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnResultSidOk() (*string, bool)`

GetAddOnResultSidOk returns a tuple with the AddOnResultSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnResultSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnResultSid(v string)`

SetAddOnResultSid sets AddOnResultSid field to given value.

### HasAddOnResultSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasAddOnResultSid() bool`

HasAddOnResultSid returns a boolean if a field has been set.

### SetAddOnResultSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnResultSidNil(b bool)`

 SetAddOnResultSidNil sets the value for AddOnResultSid to be an explicit nil

### UnsetAddOnResultSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetAddOnResultSid()`

UnsetAddOnResultSid ensures that no value is present for AddOnResultSid, not even an explicit nil
### GetAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnSid() string`

GetAddOnSid returns the AddOnSid field if non-nil, zero value otherwise.

### GetAddOnSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetAddOnSidOk() (*string, bool)`

GetAddOnSidOk returns a tuple with the AddOnSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnSid(v string)`

SetAddOnSid sets AddOnSid field to given value.

### HasAddOnSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasAddOnSid() bool`

HasAddOnSid returns a boolean if a field has been set.

### SetAddOnSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetAddOnSidNil(b bool)`

 SetAddOnSidNil sets the value for AddOnSid to be an explicit nil

### UnsetAddOnSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetAddOnSid()`

UnsetAddOnSid ensures that no value is present for AddOnSid, not even an explicit nil
### GetContentType

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLabel

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetReferenceSid() string`

GetReferenceSid returns the ReferenceSid field if non-nil, zero value otherwise.

### GetReferenceSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetReferenceSidOk() (*string, bool)`

GetReferenceSidOk returns a tuple with the ReferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetReferenceSid(v string)`

SetReferenceSid sets ReferenceSid field to given value.

### HasReferenceSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasReferenceSid() bool`

HasReferenceSid returns a boolean if a field has been set.

### SetReferenceSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetReferenceSidNil(b bool)`

 SetReferenceSidNil sets the value for ReferenceSid to be an explicit nil

### UnsetReferenceSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetReferenceSid()`

UnsetReferenceSid ensures that no value is present for ReferenceSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


