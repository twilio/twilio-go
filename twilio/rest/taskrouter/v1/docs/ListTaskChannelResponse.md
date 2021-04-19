# ListTaskChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]TaskrouterV1WorkspaceTaskChannel**](TaskrouterV1WorkspaceTaskChannel.md) |  | [optional] 
**Meta** | Pointer to [**ListWorkspaceResponseMeta**](ListWorkspaceResponse_meta.md) |  | [optional] 

## Methods

### NewListTaskChannelResponse

`func NewListTaskChannelResponse() *ListTaskChannelResponse`

NewListTaskChannelResponse instantiates a new ListTaskChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaskChannelResponseWithDefaults

`func NewListTaskChannelResponseWithDefaults() *ListTaskChannelResponse`

NewListTaskChannelResponseWithDefaults instantiates a new ListTaskChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListTaskChannelResponse) GetChannels() []TaskrouterV1WorkspaceTaskChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListTaskChannelResponse) GetChannelsOk() (*[]TaskrouterV1WorkspaceTaskChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListTaskChannelResponse) SetChannels(v []TaskrouterV1WorkspaceTaskChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListTaskChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListTaskChannelResponse) GetMeta() ListWorkspaceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTaskChannelResponse) GetMetaOk() (*ListWorkspaceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTaskChannelResponse) SetMeta(v ListWorkspaceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTaskChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


