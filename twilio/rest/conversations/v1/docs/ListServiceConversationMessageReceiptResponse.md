# ListServiceConversationMessageReceiptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryReceipts** | Pointer to [**[]ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt**](ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListServiceConversationMessageReceiptResponse

`func NewListServiceConversationMessageReceiptResponse() *ListServiceConversationMessageReceiptResponse`

NewListServiceConversationMessageReceiptResponse instantiates a new ListServiceConversationMessageReceiptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceConversationMessageReceiptResponseWithDefaults

`func NewListServiceConversationMessageReceiptResponseWithDefaults() *ListServiceConversationMessageReceiptResponse`

NewListServiceConversationMessageReceiptResponseWithDefaults instantiates a new ListServiceConversationMessageReceiptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryReceipts

`func (o *ListServiceConversationMessageReceiptResponse) GetDeliveryReceipts() []ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt`

GetDeliveryReceipts returns the DeliveryReceipts field if non-nil, zero value otherwise.

### GetDeliveryReceiptsOk

`func (o *ListServiceConversationMessageReceiptResponse) GetDeliveryReceiptsOk() (*[]ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt, bool)`

GetDeliveryReceiptsOk returns a tuple with the DeliveryReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryReceipts

`func (o *ListServiceConversationMessageReceiptResponse) SetDeliveryReceipts(v []ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt)`

SetDeliveryReceipts sets DeliveryReceipts field to given value.

### HasDeliveryReceipts

`func (o *ListServiceConversationMessageReceiptResponse) HasDeliveryReceipts() bool`

HasDeliveryReceipts returns a boolean if a field has been set.

### GetMeta

`func (o *ListServiceConversationMessageReceiptResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceConversationMessageReceiptResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceConversationMessageReceiptResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceConversationMessageReceiptResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


