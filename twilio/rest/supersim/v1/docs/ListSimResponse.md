# ListSimResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**Sims** | Pointer to [**[]SupersimV1Sim**](SupersimV1Sim.md) |  | [optional] 

## Methods

### NewListSimResponse

`func NewListSimResponse() *ListSimResponse`

NewListSimResponse instantiates a new ListSimResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSimResponseWithDefaults

`func NewListSimResponseWithDefaults() *ListSimResponse`

NewListSimResponseWithDefaults instantiates a new ListSimResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSimResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSimResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSimResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSimResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSims

`func (o *ListSimResponse) GetSims() []SupersimV1Sim`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *ListSimResponse) GetSimsOk() (*[]SupersimV1Sim, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *ListSimResponse) SetSims(v []SupersimV1Sim)`

SetSims sets Sims field to given value.

### HasSims

`func (o *ListSimResponse) HasSims() bool`

HasSims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


