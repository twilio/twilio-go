# ListConnectionPolicyResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**ConnectionPolicies** | Pointer to [**[]VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md) |  | [optional] 
**Meta** | Pointer to [**ListByocTrunkResponseMeta**](ListByocTrunkResponse_meta.md) |  | [optional] 

## Methods

### NewListConnectionPolicyResponse

`func NewListConnectionPolicyResponse() *ListConnectionPolicyResponse`

NewListConnectionPolicyResponse instantiates a new ListConnectionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectionPolicyResponseWithDefaults

`func NewListConnectionPolicyResponseWithDefaults() *ListConnectionPolicyResponse`

NewListConnectionPolicyResponseWithDefaults instantiates a new ListConnectionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionPolicies

`func (o *ListConnectionPolicyResponse) GetConnectionPolicies() []VoiceV1ConnectionPolicy`

GetConnectionPolicies returns the ConnectionPolicies field if non-nil, zero value otherwise.

### GetConnectionPoliciesOk

`func (o *ListConnectionPolicyResponse) GetConnectionPoliciesOk() (*[]VoiceV1ConnectionPolicy, bool)`

GetConnectionPoliciesOk returns a tuple with the ConnectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPolicies

`func (o *ListConnectionPolicyResponse) SetConnectionPolicies(v []VoiceV1ConnectionPolicy)`

SetConnectionPolicies sets ConnectionPolicies field to given value.

### HasConnectionPolicies

`func (o *ListConnectionPolicyResponse) HasConnectionPolicies() bool`

HasConnectionPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListConnectionPolicyResponse) GetMeta() ListByocTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConnectionPolicyResponse) GetMetaOk() (*ListByocTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConnectionPolicyResponse) SetMeta(v ListByocTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConnectionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


