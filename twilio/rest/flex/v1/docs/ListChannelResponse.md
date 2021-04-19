# ListChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlexChatChannels** | Pointer to [**[]FlexV1Channel**](FlexV1Channel.md) |  | [optional] 
**Meta** | Pointer to [**ListChannelResponseMeta**](ListChannelResponse_meta.md) |  | [optional] 

## Methods

### NewListChannelResponse

`func NewListChannelResponse() *ListChannelResponse`

NewListChannelResponse instantiates a new ListChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannelResponseWithDefaults

`func NewListChannelResponseWithDefaults() *ListChannelResponse`

NewListChannelResponseWithDefaults instantiates a new ListChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlexChatChannels

`func (o *ListChannelResponse) GetFlexChatChannels() []FlexV1Channel`

GetFlexChatChannels returns the FlexChatChannels field if non-nil, zero value otherwise.

### GetFlexChatChannelsOk

`func (o *ListChannelResponse) GetFlexChatChannelsOk() (*[]FlexV1Channel, bool)`

GetFlexChatChannelsOk returns a tuple with the FlexChatChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatChannels

`func (o *ListChannelResponse) SetFlexChatChannels(v []FlexV1Channel)`

SetFlexChatChannels sets FlexChatChannels field to given value.

### HasFlexChatChannels

`func (o *ListChannelResponse) HasFlexChatChannels() bool`

HasFlexChatChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListChannelResponse) GetMeta() ListChannelResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListChannelResponse) GetMetaOk() (*ListChannelResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListChannelResponse) SetMeta(v ListChannelResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


