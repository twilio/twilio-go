# ListWebChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlexChatChannels** | Pointer to [**[]FlexV1WebChannel**](FlexV1WebChannel.md) |  | [optional] 
**Meta** | Pointer to [**ListChannelResponseMeta**](ListChannelResponse_meta.md) |  | [optional] 

## Methods

### NewListWebChannelResponse

`func NewListWebChannelResponse() *ListWebChannelResponse`

NewListWebChannelResponse instantiates a new ListWebChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebChannelResponseWithDefaults

`func NewListWebChannelResponseWithDefaults() *ListWebChannelResponse`

NewListWebChannelResponseWithDefaults instantiates a new ListWebChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlexChatChannels

`func (o *ListWebChannelResponse) GetFlexChatChannels() []FlexV1WebChannel`

GetFlexChatChannels returns the FlexChatChannels field if non-nil, zero value otherwise.

### GetFlexChatChannelsOk

`func (o *ListWebChannelResponse) GetFlexChatChannelsOk() (*[]FlexV1WebChannel, bool)`

GetFlexChatChannelsOk returns a tuple with the FlexChatChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatChannels

`func (o *ListWebChannelResponse) SetFlexChatChannels(v []FlexV1WebChannel)`

SetFlexChatChannels sets FlexChatChannels field to given value.

### HasFlexChatChannels

`func (o *ListWebChannelResponse) HasFlexChatChannels() bool`

HasFlexChatChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListWebChannelResponse) GetMeta() ListChannelResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWebChannelResponse) GetMetaOk() (*ListChannelResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWebChannelResponse) SetMeta(v ListChannelResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWebChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


