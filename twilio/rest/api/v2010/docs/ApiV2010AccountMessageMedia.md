# ApiV2010AccountMessageMedia

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created this resource | [optional] 
**ContentType** | Pointer to **NullableString** | The default mime-type of the media | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was last updated | [optional] 
**ParentSid** | Pointer to **NullableString** | The SID of the resource that created the media | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies this resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of this resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountMessageMedia

`func NewApiV2010AccountMessageMedia() *ApiV2010AccountMessageMedia`

NewApiV2010AccountMessageMedia instantiates a new ApiV2010AccountMessageMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountMessageMediaWithDefaults

`func NewApiV2010AccountMessageMediaWithDefaults() *ApiV2010AccountMessageMedia`

NewApiV2010AccountMessageMediaWithDefaults instantiates a new ApiV2010AccountMessageMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountMessageMedia) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountMessageMedia) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountMessageMedia) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountMessageMedia) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountMessageMedia) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountMessageMedia) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContentType

`func (o *ApiV2010AccountMessageMedia) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ApiV2010AccountMessageMedia) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ApiV2010AccountMessageMedia) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ApiV2010AccountMessageMedia) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ApiV2010AccountMessageMedia) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ApiV2010AccountMessageMedia) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountMessageMedia) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountMessageMedia) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountMessageMedia) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountMessageMedia) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountMessageMedia) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountMessageMedia) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountMessageMedia) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountMessageMedia) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountMessageMedia) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountMessageMedia) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountMessageMedia) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountMessageMedia) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetParentSid

`func (o *ApiV2010AccountMessageMedia) GetParentSid() string`

GetParentSid returns the ParentSid field if non-nil, zero value otherwise.

### GetParentSidOk

`func (o *ApiV2010AccountMessageMedia) GetParentSidOk() (*string, bool)`

GetParentSidOk returns a tuple with the ParentSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSid

`func (o *ApiV2010AccountMessageMedia) SetParentSid(v string)`

SetParentSid sets ParentSid field to given value.

### HasParentSid

`func (o *ApiV2010AccountMessageMedia) HasParentSid() bool`

HasParentSid returns a boolean if a field has been set.

### SetParentSidNil

`func (o *ApiV2010AccountMessageMedia) SetParentSidNil(b bool)`

 SetParentSidNil sets the value for ParentSid to be an explicit nil

### UnsetParentSid
`func (o *ApiV2010AccountMessageMedia) UnsetParentSid()`

UnsetParentSid ensures that no value is present for ParentSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountMessageMedia) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountMessageMedia) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountMessageMedia) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountMessageMedia) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountMessageMedia) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountMessageMedia) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountMessageMedia) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountMessageMedia) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountMessageMedia) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountMessageMedia) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountMessageMedia) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountMessageMedia) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


