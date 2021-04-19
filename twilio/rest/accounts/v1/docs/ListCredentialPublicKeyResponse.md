# ListCredentialPublicKeyResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Credentials** | Pointer to [**[]AccountsV1CredentialCredentialPublicKey**](AccountsV1CredentialCredentialPublicKey.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialAwsResponseMeta**](ListCredentialAwsResponse_meta.md) |  | [optional] 

## Methods

### NewListCredentialPublicKeyResponse

`func NewListCredentialPublicKeyResponse() *ListCredentialPublicKeyResponse`

NewListCredentialPublicKeyResponse instantiates a new ListCredentialPublicKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialPublicKeyResponseWithDefaults

`func NewListCredentialPublicKeyResponseWithDefaults() *ListCredentialPublicKeyResponse`

NewListCredentialPublicKeyResponseWithDefaults instantiates a new ListCredentialPublicKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ListCredentialPublicKeyResponse) GetCredentials() []AccountsV1CredentialCredentialPublicKey`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ListCredentialPublicKeyResponse) GetCredentialsOk() (*[]AccountsV1CredentialCredentialPublicKey, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ListCredentialPublicKeyResponse) SetCredentials(v []AccountsV1CredentialCredentialPublicKey)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ListCredentialPublicKeyResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetMeta

`func (o *ListCredentialPublicKeyResponse) GetMeta() ListCredentialAwsResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCredentialPublicKeyResponse) GetMetaOk() (*ListCredentialAwsResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCredentialPublicKeyResponse) SetMeta(v ListCredentialAwsResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCredentialPublicKeyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


