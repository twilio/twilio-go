# IpMessagingV2ServiceBinding

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**BindingType** | Pointer to **NullableString** |  | [optional] 
**CredentialSid** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**DateUpdated** | Pointer to **NullableTime** |  | [optional] 
**Endpoint** | Pointer to **NullableString** |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**MessageTypes** | Pointer to **[]string** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpMessagingV2ServiceBinding

`func NewIpMessagingV2ServiceBinding() *IpMessagingV2ServiceBinding`

NewIpMessagingV2ServiceBinding instantiates a new IpMessagingV2ServiceBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV2ServiceBindingWithDefaults

`func NewIpMessagingV2ServiceBindingWithDefaults() *IpMessagingV2ServiceBinding`

NewIpMessagingV2ServiceBindingWithDefaults instantiates a new IpMessagingV2ServiceBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV2ServiceBinding) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV2ServiceBinding) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV2ServiceBinding) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV2ServiceBinding) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV2ServiceBinding) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV2ServiceBinding) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBindingType

`func (o *IpMessagingV2ServiceBinding) GetBindingType() string`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *IpMessagingV2ServiceBinding) GetBindingTypeOk() (*string, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *IpMessagingV2ServiceBinding) SetBindingType(v string)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *IpMessagingV2ServiceBinding) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### SetBindingTypeNil

`func (o *IpMessagingV2ServiceBinding) SetBindingTypeNil(b bool)`

 SetBindingTypeNil sets the value for BindingType to be an explicit nil

### UnsetBindingType
`func (o *IpMessagingV2ServiceBinding) UnsetBindingType()`

UnsetBindingType ensures that no value is present for BindingType, not even an explicit nil
### GetCredentialSid

`func (o *IpMessagingV2ServiceBinding) GetCredentialSid() string`

GetCredentialSid returns the CredentialSid field if non-nil, zero value otherwise.

### GetCredentialSidOk

`func (o *IpMessagingV2ServiceBinding) GetCredentialSidOk() (*string, bool)`

GetCredentialSidOk returns a tuple with the CredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSid

`func (o *IpMessagingV2ServiceBinding) SetCredentialSid(v string)`

SetCredentialSid sets CredentialSid field to given value.

### HasCredentialSid

`func (o *IpMessagingV2ServiceBinding) HasCredentialSid() bool`

HasCredentialSid returns a boolean if a field has been set.

### SetCredentialSidNil

`func (o *IpMessagingV2ServiceBinding) SetCredentialSidNil(b bool)`

 SetCredentialSidNil sets the value for CredentialSid to be an explicit nil

### UnsetCredentialSid
`func (o *IpMessagingV2ServiceBinding) UnsetCredentialSid()`

UnsetCredentialSid ensures that no value is present for CredentialSid, not even an explicit nil
### GetDateCreated

`func (o *IpMessagingV2ServiceBinding) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IpMessagingV2ServiceBinding) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IpMessagingV2ServiceBinding) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IpMessagingV2ServiceBinding) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *IpMessagingV2ServiceBinding) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *IpMessagingV2ServiceBinding) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *IpMessagingV2ServiceBinding) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *IpMessagingV2ServiceBinding) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *IpMessagingV2ServiceBinding) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *IpMessagingV2ServiceBinding) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *IpMessagingV2ServiceBinding) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *IpMessagingV2ServiceBinding) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndpoint

`func (o *IpMessagingV2ServiceBinding) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IpMessagingV2ServiceBinding) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IpMessagingV2ServiceBinding) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IpMessagingV2ServiceBinding) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *IpMessagingV2ServiceBinding) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *IpMessagingV2ServiceBinding) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIdentity

`func (o *IpMessagingV2ServiceBinding) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IpMessagingV2ServiceBinding) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IpMessagingV2ServiceBinding) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IpMessagingV2ServiceBinding) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *IpMessagingV2ServiceBinding) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *IpMessagingV2ServiceBinding) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLinks

`func (o *IpMessagingV2ServiceBinding) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV2ServiceBinding) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV2ServiceBinding) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV2ServiceBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV2ServiceBinding) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV2ServiceBinding) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMessageTypes

`func (o *IpMessagingV2ServiceBinding) GetMessageTypes() []string`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *IpMessagingV2ServiceBinding) GetMessageTypesOk() (*[]string, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *IpMessagingV2ServiceBinding) SetMessageTypes(v []string)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *IpMessagingV2ServiceBinding) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### SetMessageTypesNil

`func (o *IpMessagingV2ServiceBinding) SetMessageTypesNil(b bool)`

 SetMessageTypesNil sets the value for MessageTypes to be an explicit nil

### UnsetMessageTypes
`func (o *IpMessagingV2ServiceBinding) UnsetMessageTypes()`

UnsetMessageTypes ensures that no value is present for MessageTypes, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV2ServiceBinding) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV2ServiceBinding) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV2ServiceBinding) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV2ServiceBinding) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV2ServiceBinding) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV2ServiceBinding) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *IpMessagingV2ServiceBinding) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpMessagingV2ServiceBinding) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpMessagingV2ServiceBinding) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpMessagingV2ServiceBinding) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IpMessagingV2ServiceBinding) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IpMessagingV2ServiceBinding) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *IpMessagingV2ServiceBinding) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV2ServiceBinding) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV2ServiceBinding) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV2ServiceBinding) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV2ServiceBinding) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV2ServiceBinding) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


