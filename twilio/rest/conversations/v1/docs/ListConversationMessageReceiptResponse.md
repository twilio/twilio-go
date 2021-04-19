# ListConversationMessageReceiptResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DeliveryReceipts** | Pointer to [**[]ConversationsV1ConversationConversationMessageConversationMessageReceipt**](ConversationsV1ConversationConversationMessageConversationMessageReceipt.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListConversationMessageReceiptResponse

`func NewListConversationMessageReceiptResponse() *ListConversationMessageReceiptResponse`

NewListConversationMessageReceiptResponse instantiates a new ListConversationMessageReceiptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationMessageReceiptResponseWithDefaults

`func NewListConversationMessageReceiptResponseWithDefaults() *ListConversationMessageReceiptResponse`

NewListConversationMessageReceiptResponseWithDefaults instantiates a new ListConversationMessageReceiptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryReceipts

`func (o *ListConversationMessageReceiptResponse) GetDeliveryReceipts() []ConversationsV1ConversationConversationMessageConversationMessageReceipt`

GetDeliveryReceipts returns the DeliveryReceipts field if non-nil, zero value otherwise.

### GetDeliveryReceiptsOk

`func (o *ListConversationMessageReceiptResponse) GetDeliveryReceiptsOk() (*[]ConversationsV1ConversationConversationMessageConversationMessageReceipt, bool)`

GetDeliveryReceiptsOk returns a tuple with the DeliveryReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryReceipts

`func (o *ListConversationMessageReceiptResponse) SetDeliveryReceipts(v []ConversationsV1ConversationConversationMessageConversationMessageReceipt)`

SetDeliveryReceipts sets DeliveryReceipts field to given value.

### HasDeliveryReceipts

`func (o *ListConversationMessageReceiptResponse) HasDeliveryReceipts() bool`

HasDeliveryReceipts returns a boolean if a field has been set.

### GetMeta

`func (o *ListConversationMessageReceiptResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConversationMessageReceiptResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConversationMessageReceiptResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConversationMessageReceiptResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


