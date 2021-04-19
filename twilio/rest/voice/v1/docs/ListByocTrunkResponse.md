# ListByocTrunkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByocTrunks** | Pointer to [**[]VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md) |  | [optional] 
**Meta** | Pointer to [**ListByocTrunkResponseMeta**](ListByocTrunkResponse_meta.md) |  | [optional] 

## Methods

### NewListByocTrunkResponse

`func NewListByocTrunkResponse() *ListByocTrunkResponse`

NewListByocTrunkResponse instantiates a new ListByocTrunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListByocTrunkResponseWithDefaults

`func NewListByocTrunkResponseWithDefaults() *ListByocTrunkResponse`

NewListByocTrunkResponseWithDefaults instantiates a new ListByocTrunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByocTrunks

`func (o *ListByocTrunkResponse) GetByocTrunks() []VoiceV1ByocTrunk`

GetByocTrunks returns the ByocTrunks field if non-nil, zero value otherwise.

### GetByocTrunksOk

`func (o *ListByocTrunkResponse) GetByocTrunksOk() (*[]VoiceV1ByocTrunk, bool)`

GetByocTrunksOk returns a tuple with the ByocTrunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByocTrunks

`func (o *ListByocTrunkResponse) SetByocTrunks(v []VoiceV1ByocTrunk)`

SetByocTrunks sets ByocTrunks field to given value.

### HasByocTrunks

`func (o *ListByocTrunkResponse) HasByocTrunks() bool`

HasByocTrunks returns a boolean if a field has been set.

### GetMeta

`func (o *ListByocTrunkResponse) GetMeta() ListByocTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListByocTrunkResponse) GetMetaOk() (*ListByocTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListByocTrunkResponse) SetMeta(v ListByocTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListByocTrunkResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


