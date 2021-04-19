# ListCredentialResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Credentials** | Pointer to [**[]ConversationsV1Credential**](ConversationsV1Credential.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListCredentialResponse

`func NewListCredentialResponse() *ListCredentialResponse`

NewListCredentialResponse instantiates a new ListCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialResponseWithDefaults

`func NewListCredentialResponseWithDefaults() *ListCredentialResponse`

NewListCredentialResponseWithDefaults instantiates a new ListCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ListCredentialResponse) GetCredentials() []ConversationsV1Credential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ListCredentialResponse) GetCredentialsOk() (*[]ConversationsV1Credential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ListCredentialResponse) SetCredentials(v []ConversationsV1Credential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ListCredentialResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetMeta

`func (o *ListCredentialResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCredentialResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCredentialResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCredentialResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


