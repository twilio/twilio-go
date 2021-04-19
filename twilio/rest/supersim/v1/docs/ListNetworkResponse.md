# ListNetworkResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**Networks** | Pointer to [**[]SupersimV1Network**](SupersimV1Network.md) |  | [optional] 

## Methods

### NewListNetworkResponse

`func NewListNetworkResponse() *ListNetworkResponse`

NewListNetworkResponse instantiates a new ListNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkResponseWithDefaults

`func NewListNetworkResponseWithDefaults() *ListNetworkResponse`

NewListNetworkResponseWithDefaults instantiates a new ListNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListNetworkResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNetworks

`func (o *ListNetworkResponse) GetNetworks() []SupersimV1Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ListNetworkResponse) GetNetworksOk() (*[]SupersimV1Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ListNetworkResponse) SetNetworks(v []SupersimV1Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ListNetworkResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


