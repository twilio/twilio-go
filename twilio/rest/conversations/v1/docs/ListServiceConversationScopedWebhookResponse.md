# ListServiceConversationScopedWebhookResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Webhooks** | Pointer to [**[]ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md) |  | [optional] 

## Methods

### NewListServiceConversationScopedWebhookResponse

`func NewListServiceConversationScopedWebhookResponse() *ListServiceConversationScopedWebhookResponse`

NewListServiceConversationScopedWebhookResponse instantiates a new ListServiceConversationScopedWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceConversationScopedWebhookResponseWithDefaults

`func NewListServiceConversationScopedWebhookResponseWithDefaults() *ListServiceConversationScopedWebhookResponse`

NewListServiceConversationScopedWebhookResponseWithDefaults instantiates a new ListServiceConversationScopedWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListServiceConversationScopedWebhookResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceConversationScopedWebhookResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceConversationScopedWebhookResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceConversationScopedWebhookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListServiceConversationScopedWebhookResponse) GetWebhooks() []ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListServiceConversationScopedWebhookResponse) GetWebhooksOk() (*[]ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListServiceConversationScopedWebhookResponse) SetWebhooks(v []ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListServiceConversationScopedWebhookResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


