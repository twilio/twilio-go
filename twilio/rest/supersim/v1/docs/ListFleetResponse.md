# ListFleetResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Fleets** | Pointer to [**[]SupersimV1Fleet**](SupersimV1Fleet.md) |  | [optional] 
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 

## Methods

### NewListFleetResponse

`func NewListFleetResponse() *ListFleetResponse`

NewListFleetResponse instantiates a new ListFleetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFleetResponseWithDefaults

`func NewListFleetResponseWithDefaults() *ListFleetResponse`

NewListFleetResponseWithDefaults instantiates a new ListFleetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleets

`func (o *ListFleetResponse) GetFleets() []SupersimV1Fleet`

GetFleets returns the Fleets field if non-nil, zero value otherwise.

### GetFleetsOk

`func (o *ListFleetResponse) GetFleetsOk() (*[]SupersimV1Fleet, bool)`

GetFleetsOk returns a tuple with the Fleets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleets

`func (o *ListFleetResponse) SetFleets(v []SupersimV1Fleet)`

SetFleets sets Fleets field to given value.

### HasFleets

`func (o *ListFleetResponse) HasFleets() bool`

HasFleets returns a boolean if a field has been set.

### GetMeta

`func (o *ListFleetResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFleetResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFleetResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFleetResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


