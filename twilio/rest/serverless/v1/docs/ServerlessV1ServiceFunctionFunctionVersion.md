# ServerlessV1ServiceFunctionFunctionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Function Version resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Function Version resource was created | [optional] 
**FunctionSid** | Pointer to **NullableString** | The SID of the Function resource that is the parent of the Function Version resource | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Path** | Pointer to **NullableString** | The URL-friendly string by which the Function Version resource can be referenced | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Function Version resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Function Version resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Function Version resource | [optional] 
**Visibility** | Pointer to **NullableString** | The access control that determines how the Function Version resource can be accessed | [optional] 

## Methods

### NewServerlessV1ServiceFunctionFunctionVersion

`func NewServerlessV1ServiceFunctionFunctionVersion() *ServerlessV1ServiceFunctionFunctionVersion`

NewServerlessV1ServiceFunctionFunctionVersion instantiates a new ServerlessV1ServiceFunctionFunctionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceFunctionFunctionVersionWithDefaults

`func NewServerlessV1ServiceFunctionFunctionVersionWithDefaults() *ServerlessV1ServiceFunctionFunctionVersion`

NewServerlessV1ServiceFunctionFunctionVersionWithDefaults instantiates a new ServerlessV1ServiceFunctionFunctionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetFunctionSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetFunctionSid() string`

GetFunctionSid returns the FunctionSid field if non-nil, zero value otherwise.

### GetFunctionSidOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetFunctionSidOk() (*string, bool)`

GetFunctionSidOk returns a tuple with the FunctionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetFunctionSid(v string)`

SetFunctionSid sets FunctionSid field to given value.

### HasFunctionSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasFunctionSid() bool`

HasFunctionSid returns a boolean if a field has been set.

### SetFunctionSidNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetFunctionSidNil(b bool)`

 SetFunctionSidNil sets the value for FunctionSid to be an explicit nil

### UnsetFunctionSid
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetFunctionSid()`

UnsetFunctionSid ensures that no value is present for FunctionSid, not even an explicit nil
### GetLinks

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPath

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVisibility

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServerlessV1ServiceFunctionFunctionVersion) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServerlessV1ServiceFunctionFunctionVersion) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *ServerlessV1ServiceFunctionFunctionVersion) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *ServerlessV1ServiceFunctionFunctionVersion) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


