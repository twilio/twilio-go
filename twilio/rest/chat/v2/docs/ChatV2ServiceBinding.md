# ChatV2ServiceBinding

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**BindingType** | Pointer to **NullableString** | The push technology to use for the binding | [optional] 
**CredentialSid** | Pointer to **NullableString** | The SID of the Credential for the binding | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Endpoint** | Pointer to **NullableString** | The unique endpoint identifier for the Binding | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Binding&#39;s User | [optional] 
**MessageTypes** | Pointer to **[]string** | The Programmable Chat message types the binding is subscribed to | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Binding resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Binding resource | [optional] 

## Methods

### NewChatV2ServiceBinding

`func NewChatV2ServiceBinding() *ChatV2ServiceBinding`

NewChatV2ServiceBinding instantiates a new ChatV2ServiceBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceBindingWithDefaults

`func NewChatV2ServiceBindingWithDefaults() *ChatV2ServiceBinding`

NewChatV2ServiceBindingWithDefaults instantiates a new ChatV2ServiceBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2ServiceBinding) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2ServiceBinding) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2ServiceBinding) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2ServiceBinding) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2ServiceBinding) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2ServiceBinding) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBindingType

`func (o *ChatV2ServiceBinding) GetBindingType() string`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *ChatV2ServiceBinding) GetBindingTypeOk() (*string, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *ChatV2ServiceBinding) SetBindingType(v string)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *ChatV2ServiceBinding) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### SetBindingTypeNil

`func (o *ChatV2ServiceBinding) SetBindingTypeNil(b bool)`

 SetBindingTypeNil sets the value for BindingType to be an explicit nil

### UnsetBindingType
`func (o *ChatV2ServiceBinding) UnsetBindingType()`

UnsetBindingType ensures that no value is present for BindingType, not even an explicit nil
### GetCredentialSid

`func (o *ChatV2ServiceBinding) GetCredentialSid() string`

GetCredentialSid returns the CredentialSid field if non-nil, zero value otherwise.

### GetCredentialSidOk

`func (o *ChatV2ServiceBinding) GetCredentialSidOk() (*string, bool)`

GetCredentialSidOk returns a tuple with the CredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSid

`func (o *ChatV2ServiceBinding) SetCredentialSid(v string)`

SetCredentialSid sets CredentialSid field to given value.

### HasCredentialSid

`func (o *ChatV2ServiceBinding) HasCredentialSid() bool`

HasCredentialSid returns a boolean if a field has been set.

### SetCredentialSidNil

`func (o *ChatV2ServiceBinding) SetCredentialSidNil(b bool)`

 SetCredentialSidNil sets the value for CredentialSid to be an explicit nil

### UnsetCredentialSid
`func (o *ChatV2ServiceBinding) UnsetCredentialSid()`

UnsetCredentialSid ensures that no value is present for CredentialSid, not even an explicit nil
### GetDateCreated

`func (o *ChatV2ServiceBinding) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV2ServiceBinding) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV2ServiceBinding) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV2ServiceBinding) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV2ServiceBinding) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV2ServiceBinding) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV2ServiceBinding) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV2ServiceBinding) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV2ServiceBinding) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV2ServiceBinding) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV2ServiceBinding) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV2ServiceBinding) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndpoint

`func (o *ChatV2ServiceBinding) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ChatV2ServiceBinding) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ChatV2ServiceBinding) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ChatV2ServiceBinding) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ChatV2ServiceBinding) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ChatV2ServiceBinding) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIdentity

`func (o *ChatV2ServiceBinding) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ChatV2ServiceBinding) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ChatV2ServiceBinding) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ChatV2ServiceBinding) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ChatV2ServiceBinding) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ChatV2ServiceBinding) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLinks

`func (o *ChatV2ServiceBinding) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChatV2ServiceBinding) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChatV2ServiceBinding) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChatV2ServiceBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ChatV2ServiceBinding) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ChatV2ServiceBinding) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMessageTypes

`func (o *ChatV2ServiceBinding) GetMessageTypes() []string`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *ChatV2ServiceBinding) GetMessageTypesOk() (*[]string, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *ChatV2ServiceBinding) SetMessageTypes(v []string)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *ChatV2ServiceBinding) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### SetMessageTypesNil

`func (o *ChatV2ServiceBinding) SetMessageTypesNil(b bool)`

 SetMessageTypesNil sets the value for MessageTypes to be an explicit nil

### UnsetMessageTypes
`func (o *ChatV2ServiceBinding) UnsetMessageTypes()`

UnsetMessageTypes ensures that no value is present for MessageTypes, not even an explicit nil
### GetServiceSid

`func (o *ChatV2ServiceBinding) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV2ServiceBinding) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV2ServiceBinding) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV2ServiceBinding) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV2ServiceBinding) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV2ServiceBinding) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV2ServiceBinding) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV2ServiceBinding) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV2ServiceBinding) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV2ServiceBinding) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV2ServiceBinding) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV2ServiceBinding) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ChatV2ServiceBinding) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2ServiceBinding) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2ServiceBinding) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2ServiceBinding) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2ServiceBinding) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2ServiceBinding) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


