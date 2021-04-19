# ListWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 
**Webhooks** | Pointer to [**[]AutopilotV1AssistantWebhook**](AutopilotV1AssistantWebhook.md) |  | [optional] 

## Methods

### NewListWebhookResponse

`func NewListWebhookResponse() *ListWebhookResponse`

NewListWebhookResponse instantiates a new ListWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookResponseWithDefaults

`func NewListWebhookResponseWithDefaults() *ListWebhookResponse`

NewListWebhookResponseWithDefaults instantiates a new ListWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListWebhookResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWebhookResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWebhookResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWebhookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListWebhookResponse) GetWebhooks() []AutopilotV1AssistantWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListWebhookResponse) GetWebhooksOk() (*[]AutopilotV1AssistantWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListWebhookResponse) SetWebhooks(v []AutopilotV1AssistantWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListWebhookResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


