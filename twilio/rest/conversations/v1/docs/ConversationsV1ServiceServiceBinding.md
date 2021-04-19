# ConversationsV1ServiceServiceBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this binding. | [optional] 
**BindingType** | Pointer to **NullableString** | The push technology to use for the binding. | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the resource is associated with. | [optional] 
**CredentialSid** | Pointer to **NullableString** | The SID of the Credential for the binding. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**Endpoint** | Pointer to **NullableString** | The unique endpoint identifier for the Binding. | [optional] 
**Identity** | Pointer to **NullableString** | The identity of Conversation User associated with this binding. | [optional] 
**MessageTypes** | Pointer to **[]string** | The Conversation message types the binding is subscribed to. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this binding. | [optional] 

## Methods

### NewConversationsV1ServiceServiceBinding

`func NewConversationsV1ServiceServiceBinding() *ConversationsV1ServiceServiceBinding`

NewConversationsV1ServiceServiceBinding instantiates a new ConversationsV1ServiceServiceBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceBindingWithDefaults

`func NewConversationsV1ServiceServiceBindingWithDefaults() *ConversationsV1ServiceServiceBinding`

NewConversationsV1ServiceServiceBindingWithDefaults instantiates a new ConversationsV1ServiceServiceBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceBinding) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceBinding) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceBinding) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceBinding) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceBinding) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceBinding) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBindingType

`func (o *ConversationsV1ServiceServiceBinding) GetBindingType() string`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *ConversationsV1ServiceServiceBinding) GetBindingTypeOk() (*string, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *ConversationsV1ServiceServiceBinding) SetBindingType(v string)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *ConversationsV1ServiceServiceBinding) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### SetBindingTypeNil

`func (o *ConversationsV1ServiceServiceBinding) SetBindingTypeNil(b bool)`

 SetBindingTypeNil sets the value for BindingType to be an explicit nil

### UnsetBindingType
`func (o *ConversationsV1ServiceServiceBinding) UnsetBindingType()`

UnsetBindingType ensures that no value is present for BindingType, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceBinding) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceBinding) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceBinding) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceBinding) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceBinding) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceBinding) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetCredentialSid

`func (o *ConversationsV1ServiceServiceBinding) GetCredentialSid() string`

GetCredentialSid returns the CredentialSid field if non-nil, zero value otherwise.

### GetCredentialSidOk

`func (o *ConversationsV1ServiceServiceBinding) GetCredentialSidOk() (*string, bool)`

GetCredentialSidOk returns a tuple with the CredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSid

`func (o *ConversationsV1ServiceServiceBinding) SetCredentialSid(v string)`

SetCredentialSid sets CredentialSid field to given value.

### HasCredentialSid

`func (o *ConversationsV1ServiceServiceBinding) HasCredentialSid() bool`

HasCredentialSid returns a boolean if a field has been set.

### SetCredentialSidNil

`func (o *ConversationsV1ServiceServiceBinding) SetCredentialSidNil(b bool)`

 SetCredentialSidNil sets the value for CredentialSid to be an explicit nil

### UnsetCredentialSid
`func (o *ConversationsV1ServiceServiceBinding) UnsetCredentialSid()`

UnsetCredentialSid ensures that no value is present for CredentialSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceBinding) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceBinding) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceBinding) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceBinding) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceBinding) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceBinding) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceBinding) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceBinding) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceBinding) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceBinding) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceBinding) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceBinding) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndpoint

`func (o *ConversationsV1ServiceServiceBinding) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ConversationsV1ServiceServiceBinding) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ConversationsV1ServiceServiceBinding) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ConversationsV1ServiceServiceBinding) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ConversationsV1ServiceServiceBinding) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ConversationsV1ServiceServiceBinding) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIdentity

`func (o *ConversationsV1ServiceServiceBinding) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ConversationsV1ServiceServiceBinding) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ConversationsV1ServiceServiceBinding) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ConversationsV1ServiceServiceBinding) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ConversationsV1ServiceServiceBinding) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ConversationsV1ServiceServiceBinding) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetMessageTypes

`func (o *ConversationsV1ServiceServiceBinding) GetMessageTypes() []string`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *ConversationsV1ServiceServiceBinding) GetMessageTypesOk() (*[]string, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *ConversationsV1ServiceServiceBinding) SetMessageTypes(v []string)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *ConversationsV1ServiceServiceBinding) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### SetMessageTypesNil

`func (o *ConversationsV1ServiceServiceBinding) SetMessageTypesNil(b bool)`

 SetMessageTypesNil sets the value for MessageTypes to be an explicit nil

### UnsetMessageTypes
`func (o *ConversationsV1ServiceServiceBinding) UnsetMessageTypes()`

UnsetMessageTypes ensures that no value is present for MessageTypes, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceBinding) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceBinding) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceBinding) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceBinding) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceBinding) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceBinding) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceBinding) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceBinding) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceBinding) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceBinding) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceBinding) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceBinding) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


