# ListConversationScopedWebhookResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Webhooks** | Pointer to [**[]ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md) |  | [optional] 

## Methods

### NewListConversationScopedWebhookResponse

`func NewListConversationScopedWebhookResponse() *ListConversationScopedWebhookResponse`

NewListConversationScopedWebhookResponse instantiates a new ListConversationScopedWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationScopedWebhookResponseWithDefaults

`func NewListConversationScopedWebhookResponseWithDefaults() *ListConversationScopedWebhookResponse`

NewListConversationScopedWebhookResponseWithDefaults instantiates a new ListConversationScopedWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListConversationScopedWebhookResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConversationScopedWebhookResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConversationScopedWebhookResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConversationScopedWebhookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListConversationScopedWebhookResponse) GetWebhooks() []ConversationsV1ConversationConversationScopedWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListConversationScopedWebhookResponse) GetWebhooksOk() (*[]ConversationsV1ConversationConversationScopedWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListConversationScopedWebhookResponse) SetWebhooks(v []ConversationsV1ConversationConversationScopedWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListConversationScopedWebhookResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


