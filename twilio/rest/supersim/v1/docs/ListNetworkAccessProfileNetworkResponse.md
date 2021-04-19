# ListNetworkAccessProfileNetworkResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**Networks** | Pointer to [**[]SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork.md) |  | [optional] 

## Methods

### NewListNetworkAccessProfileNetworkResponse

`func NewListNetworkAccessProfileNetworkResponse() *ListNetworkAccessProfileNetworkResponse`

NewListNetworkAccessProfileNetworkResponse instantiates a new ListNetworkAccessProfileNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkAccessProfileNetworkResponseWithDefaults

`func NewListNetworkAccessProfileNetworkResponseWithDefaults() *ListNetworkAccessProfileNetworkResponse`

NewListNetworkAccessProfileNetworkResponseWithDefaults instantiates a new ListNetworkAccessProfileNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListNetworkAccessProfileNetworkResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkAccessProfileNetworkResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkAccessProfileNetworkResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkAccessProfileNetworkResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNetworks

`func (o *ListNetworkAccessProfileNetworkResponse) GetNetworks() []SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ListNetworkAccessProfileNetworkResponse) GetNetworksOk() (*[]SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ListNetworkAccessProfileNetworkResponse) SetNetworks(v []SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ListNetworkAccessProfileNetworkResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


