# ListWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 
**Webhooks** | Pointer to [**[]VerifyV2ServiceWebhook**](VerifyV2ServiceWebhook.md) |  | [optional] 

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

`func (o *ListWebhookResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWebhookResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWebhookResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWebhookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWebhooks

`func (o *ListWebhookResponse) GetWebhooks() []VerifyV2ServiceWebhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ListWebhookResponse) GetWebhooksOk() (*[]VerifyV2ServiceWebhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ListWebhookResponse) SetWebhooks(v []VerifyV2ServiceWebhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ListWebhookResponse) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


