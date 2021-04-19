# ListTrunkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListTrunkResponseMeta**](ListTrunkResponse_meta.md) |  | [optional] 
**Trunks** | Pointer to [**[]TrunkingV1Trunk**](TrunkingV1Trunk.md) |  | [optional] 

## Methods

### NewListTrunkResponse

`func NewListTrunkResponse() *ListTrunkResponse`

NewListTrunkResponse instantiates a new ListTrunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTrunkResponseWithDefaults

`func NewListTrunkResponseWithDefaults() *ListTrunkResponse`

NewListTrunkResponseWithDefaults instantiates a new ListTrunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListTrunkResponse) GetMeta() ListTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTrunkResponse) GetMetaOk() (*ListTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTrunkResponse) SetMeta(v ListTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTrunkResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTrunks

`func (o *ListTrunkResponse) GetTrunks() []TrunkingV1Trunk`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *ListTrunkResponse) GetTrunksOk() (*[]TrunkingV1Trunk, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *ListTrunkResponse) SetTrunks(v []TrunkingV1Trunk)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *ListTrunkResponse) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


