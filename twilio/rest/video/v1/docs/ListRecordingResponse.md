# ListRecordingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**Recordings** | Pointer to [**[]VideoV1Recording**](VideoV1Recording.md) |  | [optional] 

## Methods

### NewListRecordingResponse

`func NewListRecordingResponse() *ListRecordingResponse`

NewListRecordingResponse instantiates a new ListRecordingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecordingResponseWithDefaults

`func NewListRecordingResponseWithDefaults() *ListRecordingResponse`

NewListRecordingResponseWithDefaults instantiates a new ListRecordingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRecordingResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRecordingResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRecordingResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRecordingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRecordings

`func (o *ListRecordingResponse) GetRecordings() []VideoV1Recording`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *ListRecordingResponse) GetRecordingsOk() (*[]VideoV1Recording, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *ListRecordingResponse) SetRecordings(v []VideoV1Recording)`

SetRecordings sets Recordings field to given value.

### HasRecordings

`func (o *ListRecordingResponse) HasRecordings() bool`

HasRecordings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


