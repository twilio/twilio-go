# ConversationsMessagesReceiptsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConversationMessageReceipt**](ConversationsMessagesReceiptsApi.md#FetchConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**ListConversationMessageReceipt**](ConversationsMessagesReceiptsApi.md#ListConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 



## FetchConversationMessageReceipt

> ConversationsV1ConversationMessageReceipt FetchConversationMessageReceipt(ctx, ConversationSidMessageSidSid)



Fetch the delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationMessageReceipt**](ConversationsV1ConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessageReceipt

> []ConversationsV1ConversationMessageReceipt ListConversationMessageReceipt(ctx, ConversationSidMessageSidoptional)



Retrieve a list of all delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ConversationMessageReceipt**](ConversationsV1ConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

