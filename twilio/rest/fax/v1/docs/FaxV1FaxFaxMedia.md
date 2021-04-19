# FaxV1FaxFaxMedia

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ContentType** | Pointer to **NullableString** | The content type of the stored fax media | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FaxSid** | Pointer to **NullableString** | The SID of the fax the FaxMedia resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the FaxMedia resource | [optional] 

## Methods

### NewFaxV1FaxFaxMedia

`func NewFaxV1FaxFaxMedia() *FaxV1FaxFaxMedia`

NewFaxV1FaxFaxMedia instantiates a new FaxV1FaxFaxMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxV1FaxFaxMediaWithDefaults

`func NewFaxV1FaxFaxMediaWithDefaults() *FaxV1FaxFaxMedia`

NewFaxV1FaxFaxMediaWithDefaults instantiates a new FaxV1FaxFaxMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FaxV1FaxFaxMedia) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FaxV1FaxFaxMedia) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FaxV1FaxFaxMedia) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FaxV1FaxFaxMedia) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FaxV1FaxFaxMedia) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FaxV1FaxFaxMedia) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContentType

`func (o *FaxV1FaxFaxMedia) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FaxV1FaxFaxMedia) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FaxV1FaxFaxMedia) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FaxV1FaxFaxMedia) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *FaxV1FaxFaxMedia) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *FaxV1FaxFaxMedia) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDateCreated

`func (o *FaxV1FaxFaxMedia) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FaxV1FaxFaxMedia) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FaxV1FaxFaxMedia) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FaxV1FaxFaxMedia) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FaxV1FaxFaxMedia) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FaxV1FaxFaxMedia) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FaxV1FaxFaxMedia) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FaxV1FaxFaxMedia) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FaxV1FaxFaxMedia) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FaxV1FaxFaxMedia) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FaxV1FaxFaxMedia) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FaxV1FaxFaxMedia) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFaxSid

`func (o *FaxV1FaxFaxMedia) GetFaxSid() string`

GetFaxSid returns the FaxSid field if non-nil, zero value otherwise.

### GetFaxSidOk

`func (o *FaxV1FaxFaxMedia) GetFaxSidOk() (*string, bool)`

GetFaxSidOk returns a tuple with the FaxSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxSid

`func (o *FaxV1FaxFaxMedia) SetFaxSid(v string)`

SetFaxSid sets FaxSid field to given value.

### HasFaxSid

`func (o *FaxV1FaxFaxMedia) HasFaxSid() bool`

HasFaxSid returns a boolean if a field has been set.

### SetFaxSidNil

`func (o *FaxV1FaxFaxMedia) SetFaxSidNil(b bool)`

 SetFaxSidNil sets the value for FaxSid to be an explicit nil

### UnsetFaxSid
`func (o *FaxV1FaxFaxMedia) UnsetFaxSid()`

UnsetFaxSid ensures that no value is present for FaxSid, not even an explicit nil
### GetSid

`func (o *FaxV1FaxFaxMedia) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *FaxV1FaxFaxMedia) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *FaxV1FaxFaxMedia) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *FaxV1FaxFaxMedia) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *FaxV1FaxFaxMedia) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *FaxV1FaxFaxMedia) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *FaxV1FaxFaxMedia) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FaxV1FaxFaxMedia) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FaxV1FaxFaxMedia) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FaxV1FaxFaxMedia) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FaxV1FaxFaxMedia) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FaxV1FaxFaxMedia) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


