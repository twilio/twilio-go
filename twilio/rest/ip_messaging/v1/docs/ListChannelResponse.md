# ListChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]IpMessagingV1ServiceChannel**](IpMessagingV1ServiceChannel.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

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

### GetChannels

`func (o *ListChannelResponse) GetChannels() []IpMessagingV1ServiceChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListChannelResponse) GetChannelsOk() (*[]IpMessagingV1ServiceChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListChannelResponse) SetChannels(v []IpMessagingV1ServiceChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListChannelResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListChannelResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListChannelResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


