# ListUserChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]ChatV2ServiceUserUserChannel**](ChatV2ServiceUserUserChannel.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListUserChannelResponse

`func NewListUserChannelResponse() *ListUserChannelResponse`

NewListUserChannelResponse instantiates a new ListUserChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserChannelResponseWithDefaults

`func NewListUserChannelResponseWithDefaults() *ListUserChannelResponse`

NewListUserChannelResponseWithDefaults instantiates a new ListUserChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListUserChannelResponse) GetChannels() []ChatV2ServiceUserUserChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListUserChannelResponse) GetChannelsOk() (*[]ChatV2ServiceUserUserChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListUserChannelResponse) SetChannels(v []ChatV2ServiceUserUserChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListUserChannelResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetMeta

`func (o *ListUserChannelResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUserChannelResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUserChannelResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUserChannelResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


