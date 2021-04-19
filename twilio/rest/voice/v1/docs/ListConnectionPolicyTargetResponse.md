# ListConnectionPolicyTargetResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListByocTrunkResponseMeta**](ListByocTrunkResponse_meta.md) |  | [optional] 
**Targets** | Pointer to [**[]VoiceV1ConnectionPolicyConnectionPolicyTarget**](VoiceV1ConnectionPolicyConnectionPolicyTarget.md) |  | [optional] 

## Methods

### NewListConnectionPolicyTargetResponse

`func NewListConnectionPolicyTargetResponse() *ListConnectionPolicyTargetResponse`

NewListConnectionPolicyTargetResponse instantiates a new ListConnectionPolicyTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectionPolicyTargetResponseWithDefaults

`func NewListConnectionPolicyTargetResponseWithDefaults() *ListConnectionPolicyTargetResponse`

NewListConnectionPolicyTargetResponseWithDefaults instantiates a new ListConnectionPolicyTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListConnectionPolicyTargetResponse) GetMeta() ListByocTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConnectionPolicyTargetResponse) GetMetaOk() (*ListByocTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConnectionPolicyTargetResponse) SetMeta(v ListByocTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConnectionPolicyTargetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTargets

`func (o *ListConnectionPolicyTargetResponse) GetTargets() []VoiceV1ConnectionPolicyConnectionPolicyTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListConnectionPolicyTargetResponse) GetTargetsOk() (*[]VoiceV1ConnectionPolicyConnectionPolicyTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListConnectionPolicyTargetResponse) SetTargets(v []VoiceV1ConnectionPolicyConnectionPolicyTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListConnectionPolicyTargetResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


