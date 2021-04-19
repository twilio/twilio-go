# ListSyncStreamResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Streams** | Pointer to [**[]SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md) |  | [optional] 

## Methods

### NewListSyncStreamResponse

`func NewListSyncStreamResponse() *ListSyncStreamResponse`

NewListSyncStreamResponse instantiates a new ListSyncStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncStreamResponseWithDefaults

`func NewListSyncStreamResponseWithDefaults() *ListSyncStreamResponse`

NewListSyncStreamResponseWithDefaults instantiates a new ListSyncStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSyncStreamResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSyncStreamResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSyncStreamResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSyncStreamResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStreams

`func (o *ListSyncStreamResponse) GetStreams() []SyncV1ServiceSyncStream`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *ListSyncStreamResponse) GetStreamsOk() (*[]SyncV1ServiceSyncStream, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *ListSyncStreamResponse) SetStreams(v []SyncV1ServiceSyncStream)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *ListSyncStreamResponse) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


