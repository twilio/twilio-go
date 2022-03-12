# ServicesConversationsMessagesReceiptsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchServiceConversationMessageReceipt**](ServicesConversationsMessagesReceiptsApi.md#FetchServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**ListServiceConversationMessageReceipt**](ServicesConversationsMessagesReceiptsApi.md#ListServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 



## FetchServiceConversationMessageReceipt

> ConversationsV1ServiceConversationMessageReceipt FetchServiceConversationMessageReceipt(ctx, ChatServiceSidConversationSidMessageSidSid)



Fetch the delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceConversationMessageReceipt**](ConversationsV1ServiceConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessageReceipt

> []ConversationsV1ServiceConversationMessageReceipt ListServiceConversationMessageReceipt(ctx, ChatServiceSidConversationSidMessageSidoptional)



Retrieve a list of all delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceConversationMessageReceipt**](ConversationsV1ServiceConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

