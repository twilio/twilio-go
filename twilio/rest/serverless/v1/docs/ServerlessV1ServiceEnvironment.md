# ServerlessV1ServiceEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Environment resource | [optional] 
**BuildSid** | Pointer to **NullableString** | The SID of the build deployed in the environment | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Environment resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Environment resource was last updated | [optional] 
**DomainName** | Pointer to **NullableString** | The base domain name for all Functions and Assets deployed in the Environment | [optional] 
**DomainSuffix** | Pointer to **NullableString** | A URL-friendly name that represents the environment | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the Environment resource&#39;s nested resources | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Environment resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Environment resource | [optional] 
**UniqueName** | Pointer to **NullableString** | A user-defined string that uniquely identifies the Environment resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Environment resource | [optional] 

## Methods

### NewServerlessV1ServiceEnvironment

`func NewServerlessV1ServiceEnvironment() *ServerlessV1ServiceEnvironment`

NewServerlessV1ServiceEnvironment instantiates a new ServerlessV1ServiceEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceEnvironmentWithDefaults

`func NewServerlessV1ServiceEnvironmentWithDefaults() *ServerlessV1ServiceEnvironment`

NewServerlessV1ServiceEnvironmentWithDefaults instantiates a new ServerlessV1ServiceEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceEnvironment) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceEnvironment) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceEnvironment) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceEnvironment) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceEnvironment) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceEnvironment) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBuildSid

`func (o *ServerlessV1ServiceEnvironment) GetBuildSid() string`

GetBuildSid returns the BuildSid field if non-nil, zero value otherwise.

### GetBuildSidOk

`func (o *ServerlessV1ServiceEnvironment) GetBuildSidOk() (*string, bool)`

GetBuildSidOk returns a tuple with the BuildSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSid

`func (o *ServerlessV1ServiceEnvironment) SetBuildSid(v string)`

SetBuildSid sets BuildSid field to given value.

### HasBuildSid

`func (o *ServerlessV1ServiceEnvironment) HasBuildSid() bool`

HasBuildSid returns a boolean if a field has been set.

### SetBuildSidNil

`func (o *ServerlessV1ServiceEnvironment) SetBuildSidNil(b bool)`

 SetBuildSidNil sets the value for BuildSid to be an explicit nil

### UnsetBuildSid
`func (o *ServerlessV1ServiceEnvironment) UnsetBuildSid()`

UnsetBuildSid ensures that no value is present for BuildSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceEnvironment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceEnvironment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceEnvironment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceEnvironment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceEnvironment) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceEnvironment) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ServerlessV1ServiceEnvironment) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ServerlessV1ServiceEnvironment) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ServerlessV1ServiceEnvironment) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ServerlessV1ServiceEnvironment) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ServerlessV1ServiceEnvironment) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ServerlessV1ServiceEnvironment) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDomainName

`func (o *ServerlessV1ServiceEnvironment) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ServerlessV1ServiceEnvironment) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ServerlessV1ServiceEnvironment) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ServerlessV1ServiceEnvironment) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ServerlessV1ServiceEnvironment) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ServerlessV1ServiceEnvironment) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetDomainSuffix

`func (o *ServerlessV1ServiceEnvironment) GetDomainSuffix() string`

GetDomainSuffix returns the DomainSuffix field if non-nil, zero value otherwise.

### GetDomainSuffixOk

`func (o *ServerlessV1ServiceEnvironment) GetDomainSuffixOk() (*string, bool)`

GetDomainSuffixOk returns a tuple with the DomainSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSuffix

`func (o *ServerlessV1ServiceEnvironment) SetDomainSuffix(v string)`

SetDomainSuffix sets DomainSuffix field to given value.

### HasDomainSuffix

`func (o *ServerlessV1ServiceEnvironment) HasDomainSuffix() bool`

HasDomainSuffix returns a boolean if a field has been set.

### SetDomainSuffixNil

`func (o *ServerlessV1ServiceEnvironment) SetDomainSuffixNil(b bool)`

 SetDomainSuffixNil sets the value for DomainSuffix to be an explicit nil

### UnsetDomainSuffix
`func (o *ServerlessV1ServiceEnvironment) UnsetDomainSuffix()`

UnsetDomainSuffix ensures that no value is present for DomainSuffix, not even an explicit nil
### GetLinks

`func (o *ServerlessV1ServiceEnvironment) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessV1ServiceEnvironment) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessV1ServiceEnvironment) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessV1ServiceEnvironment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ServerlessV1ServiceEnvironment) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ServerlessV1ServiceEnvironment) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceEnvironment) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceEnvironment) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceEnvironment) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceEnvironment) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceEnvironment) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceEnvironment) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceEnvironment) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceEnvironment) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceEnvironment) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceEnvironment) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceEnvironment) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceEnvironment) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *ServerlessV1ServiceEnvironment) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ServerlessV1ServiceEnvironment) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ServerlessV1ServiceEnvironment) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ServerlessV1ServiceEnvironment) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *ServerlessV1ServiceEnvironment) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *ServerlessV1ServiceEnvironment) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceEnvironment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceEnvironment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceEnvironment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceEnvironment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceEnvironment) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceEnvironment) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


