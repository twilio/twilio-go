# ListInteractionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Interactions** | Pointer to [**[]ProxyV1ServiceSessionInteraction**](ProxyV1ServiceSessionInteraction.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListInteractionResponse

`func NewListInteractionResponse() *ListInteractionResponse`

NewListInteractionResponse instantiates a new ListInteractionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInteractionResponseWithDefaults

`func NewListInteractionResponseWithDefaults() *ListInteractionResponse`

NewListInteractionResponseWithDefaults instantiates a new ListInteractionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractions

`func (o *ListInteractionResponse) GetInteractions() []ProxyV1ServiceSessionInteraction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *ListInteractionResponse) GetInteractionsOk() (*[]ProxyV1ServiceSessionInteraction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *ListInteractionResponse) SetInteractions(v []ProxyV1ServiceSessionInteraction)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *ListInteractionResponse) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetMeta

`func (o *ListInteractionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInteractionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInteractionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInteractionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


