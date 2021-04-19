# ListChannelWebhookResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 
**Webhooks** | Pointer to [**[]ChatV2ServiceChannelChannelWebhook**](ChatV2ServiceChannelChannelWebhook.md) |  | [optional] 

## Methods

### NewListChannelWebhookResponse

`func NewListChannelWebhookResponse() *ListChannelWebhookResponse`

NewListChannelWebhookResponse instantiates a new ListChannelWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannelWebhookResponseWithDefaults

`func NewListChannelWebhookResponseWithDefaults() *ListChannelWebhookResponse`

NewListChannelWebhookResponseWithDefaults instantiates a new ListChannelWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListChannelWebhookResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListChannelWebhookResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListChannelWebhookResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListChannelWebhookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListChannelWebhookResponse) GetWebhooks() []ChatV2ServiceChannelChannelWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListChannelWebhookResponse) GetWebhooksOk() (*[]ChatV2ServiceChannelChannelWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListChannelWebhookResponse) SetWebhooks(v []ChatV2ServiceChannelChannelWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListChannelWebhookResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


