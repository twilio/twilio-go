# ListSampleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 
**Samples** | Pointer to [**[]AutopilotV1AssistantTaskSample**](AutopilotV1AssistantTaskSample.md) |  | [optional] 

## Methods

### NewListSampleResponse

`func NewListSampleResponse() *ListSampleResponse`

NewListSampleResponse instantiates a new ListSampleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSampleResponseWithDefaults

`func NewListSampleResponseWithDefaults() *ListSampleResponse`

NewListSampleResponseWithDefaults instantiates a new ListSampleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSampleResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSampleResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSampleResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSampleResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSamples

`func (o *ListSampleResponse) GetSamples() []AutopilotV1AssistantTaskSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ListSampleResponse) GetSamplesOk() (*[]AutopilotV1AssistantTaskSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ListSampleResponse) SetSamples(v []AutopilotV1AssistantTaskSample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ListSampleResponse) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


