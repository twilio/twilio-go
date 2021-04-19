# ListCompositionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Compositions** | Pointer to [**[]VideoV1Composition**](VideoV1Composition.md) |  | [optional] 
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 

## Methods

### NewListCompositionResponse

`func NewListCompositionResponse() *ListCompositionResponse`

NewListCompositionResponse instantiates a new ListCompositionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompositionResponseWithDefaults

`func NewListCompositionResponseWithDefaults() *ListCompositionResponse`

NewListCompositionResponseWithDefaults instantiates a new ListCompositionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositions

`func (o *ListCompositionResponse) GetCompositions() []VideoV1Composition`

GetCompositions returns the Compositions field if non-nil, zero value otherwise.

### GetCompositionsOk

`func (o *ListCompositionResponse) GetCompositionsOk() (*[]VideoV1Composition, bool)`

GetCompositionsOk returns a tuple with the Compositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositions

`func (o *ListCompositionResponse) SetCompositions(v []VideoV1Composition)`

SetCompositions sets Compositions field to given value.

### HasCompositions

`func (o *ListCompositionResponse) HasCompositions() bool`

HasCompositions returns a boolean if a field has been set.

### GetMeta

`func (o *ListCompositionResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCompositionResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCompositionResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCompositionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


