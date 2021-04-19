# ListFlowRevisionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 
**Revisions** | Pointer to [**[]StudioV2FlowFlowRevision**](StudioV2FlowFlowRevision.md) |  | [optional] 

## Methods

### NewListFlowRevisionResponse

`func NewListFlowRevisionResponse() *ListFlowRevisionResponse`

NewListFlowRevisionResponse instantiates a new ListFlowRevisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlowRevisionResponseWithDefaults

`func NewListFlowRevisionResponseWithDefaults() *ListFlowRevisionResponse`

NewListFlowRevisionResponseWithDefaults instantiates a new ListFlowRevisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListFlowRevisionResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFlowRevisionResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFlowRevisionResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFlowRevisionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRevisions

`func (o *ListFlowRevisionResponse) GetRevisions() []StudioV2FlowFlowRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ListFlowRevisionResponse) GetRevisionsOk() (*[]StudioV2FlowFlowRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ListFlowRevisionResponse) SetRevisions(v []StudioV2FlowFlowRevision)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *ListFlowRevisionResponse) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


