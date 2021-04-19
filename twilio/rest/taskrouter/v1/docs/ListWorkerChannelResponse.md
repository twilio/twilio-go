# ListWorkerChannelResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Channels** | Pointer to [**[]TaskrouterV1WorkspaceWorkerWorkerChannel**](TaskrouterV1WorkspaceWorkerWorkerChannel.md) |  | [optional] 
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 

## Methods

### NewListWorkerChannelResponse

`func NewListWorkerChannelResponse() *ListWorkerChannelResponse`

NewListWorkerChannelResponse instantiates a new ListWorkerChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkerChannelResponseWithDefaults

`func NewListWorkerChannelResponseWithDefaults() *ListWorkerChannelResponse`

NewListWorkerChannelResponseWithDefaults instantiates a new ListWorkerChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListWorkerChannelResponse) GetChannels() []TaskrouterV1WorkspaceWorkerWorkerChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListWorkerChannelResponse) GetChannelsOk() (*[]TaskrouterV1WorkspaceWorkerWorkerChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListWorkerChannelResponse) SetChannels(v []TaskrouterV1WorkspaceWorkerWorkerChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListWorkerChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListWorkerChannelResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkerChannelResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkerChannelResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkerChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


