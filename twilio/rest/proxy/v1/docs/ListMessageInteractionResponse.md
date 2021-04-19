# ListMessageInteractionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Interactions** | Pointer to [**[]ProxyV1ServiceSessionParticipantMessageInteraction**](ProxyV1ServiceSessionParticipantMessageInteraction.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListMessageInteractionResponse

`func NewListMessageInteractionResponse() *ListMessageInteractionResponse`

NewListMessageInteractionResponse instantiates a new ListMessageInteractionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessageInteractionResponseWithDefaults

`func NewListMessageInteractionResponseWithDefaults() *ListMessageInteractionResponse`

NewListMessageInteractionResponseWithDefaults instantiates a new ListMessageInteractionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractions

`func (o *ListMessageInteractionResponse) GetInteractions() []ProxyV1ServiceSessionParticipantMessageInteraction`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *ListMessageInteractionResponse) GetInteractionsOk() (*[]ProxyV1ServiceSessionParticipantMessageInteraction, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *ListMessageInteractionResponse) SetInteractions(v []ProxyV1ServiceSessionParticipantMessageInteraction)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *ListMessageInteractionResponse) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetMeta

`func (o *ListMessageInteractionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMessageInteractionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMessageInteractionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMessageInteractionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


