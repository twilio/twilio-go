# ServerlessV1ServiceEnvironmentVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Variable resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Variable resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Variable resource was last updated | [optional] 
**EnvironmentSid** | Pointer to **NullableString** | The SID of the Environment in which the Variable exists | [optional] 
**Key** | Pointer to **NullableString** | A string by which the Variable resource can be referenced | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Variable resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Variable resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Variable resource | [optional] 
**Value** | Pointer to **NullableString** | A string that contains the actual value of the Variable | [optional] 

## Methods

### NewServerlessV1ServiceEnvironmentVariable

`func NewServerlessV1ServiceEnvironmentVariable() *ServerlessV1ServiceEnvironmentVariable`

NewServerlessV1ServiceEnvironmentVariable instantiates a new ServerlessV1ServiceEnvironmentVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceEnvironmentVariableWithDefaults

`func NewServerlessV1ServiceEnvironmentVariableWithDefaults() *ServerlessV1ServiceEnvironmentVariable`

NewServerlessV1ServiceEnvironmentVariableWithDefaults instantiates a new ServerlessV1ServiceEnvironmentVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceEnvironmentVariable) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceEnvironmentVariable) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceEnvironmentVariable) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceEnvironmentVariable) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceEnvironmentVariable) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceEnvironmentVariable) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ServerlessV1ServiceEnvironmentVariable) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ServerlessV1ServiceEnvironmentVariable) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ServerlessV1ServiceEnvironmentVariable) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentVariable) GetEnvironmentSid() string`

GetEnvironmentSid returns the EnvironmentSid field if non-nil, zero value otherwise.

### GetEnvironmentSidOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetEnvironmentSidOk() (*string, bool)`

GetEnvironmentSidOk returns a tuple with the EnvironmentSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentVariable) SetEnvironmentSid(v string)`

SetEnvironmentSid sets EnvironmentSid field to given value.

### HasEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentVariable) HasEnvironmentSid() bool`

HasEnvironmentSid returns a boolean if a field has been set.

### SetEnvironmentSidNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetEnvironmentSidNil(b bool)`

 SetEnvironmentSidNil sets the value for EnvironmentSid to be an explicit nil

### UnsetEnvironmentSid
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetEnvironmentSid()`

UnsetEnvironmentSid ensures that no value is present for EnvironmentSid, not even an explicit nil
### GetKey

`func (o *ServerlessV1ServiceEnvironmentVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ServerlessV1ServiceEnvironmentVariable) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ServerlessV1ServiceEnvironmentVariable) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceEnvironmentVariable) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceEnvironmentVariable) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceEnvironmentVariable) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceEnvironmentVariable) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceEnvironmentVariable) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceEnvironmentVariable) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceEnvironmentVariable) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceEnvironmentVariable) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceEnvironmentVariable) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValue

`func (o *ServerlessV1ServiceEnvironmentVariable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServerlessV1ServiceEnvironmentVariable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServerlessV1ServiceEnvironmentVariable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServerlessV1ServiceEnvironmentVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ServerlessV1ServiceEnvironmentVariable) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ServerlessV1ServiceEnvironmentVariable) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


