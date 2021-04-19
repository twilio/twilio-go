# SupersimV1NetworkAccessProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that the Network Access Profile belongs to | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewSupersimV1NetworkAccessProfile

`func NewSupersimV1NetworkAccessProfile() *SupersimV1NetworkAccessProfile`

NewSupersimV1NetworkAccessProfile instantiates a new SupersimV1NetworkAccessProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupersimV1NetworkAccessProfileWithDefaults

`func NewSupersimV1NetworkAccessProfileWithDefaults() *SupersimV1NetworkAccessProfile`

NewSupersimV1NetworkAccessProfileWithDefaults instantiates a new SupersimV1NetworkAccessProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SupersimV1NetworkAccessProfile) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SupersimV1NetworkAccessProfile) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SupersimV1NetworkAccessProfile) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SupersimV1NetworkAccessProfile) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SupersimV1NetworkAccessProfile) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SupersimV1NetworkAccessProfile) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *SupersimV1NetworkAccessProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SupersimV1NetworkAccessProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SupersimV1NetworkAccessProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SupersimV1NetworkAccessProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SupersimV1NetworkAccessProfile) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SupersimV1NetworkAccessProfile) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SupersimV1NetworkAccessProfile) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SupersimV1NetworkAccessProfile) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SupersimV1NetworkAccessProfile) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SupersimV1NetworkAccessProfile) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SupersimV1NetworkAccessProfile) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SupersimV1NetworkAccessProfile) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *SupersimV1NetworkAccessProfile) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SupersimV1NetworkAccessProfile) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SupersimV1NetworkAccessProfile) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SupersimV1NetworkAccessProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SupersimV1NetworkAccessProfile) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SupersimV1NetworkAccessProfile) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *SupersimV1NetworkAccessProfile) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SupersimV1NetworkAccessProfile) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SupersimV1NetworkAccessProfile) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SupersimV1NetworkAccessProfile) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SupersimV1NetworkAccessProfile) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SupersimV1NetworkAccessProfile) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *SupersimV1NetworkAccessProfile) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *SupersimV1NetworkAccessProfile) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *SupersimV1NetworkAccessProfile) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *SupersimV1NetworkAccessProfile) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *SupersimV1NetworkAccessProfile) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *SupersimV1NetworkAccessProfile) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *SupersimV1NetworkAccessProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupersimV1NetworkAccessProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupersimV1NetworkAccessProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupersimV1NetworkAccessProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SupersimV1NetworkAccessProfile) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SupersimV1NetworkAccessProfile) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


