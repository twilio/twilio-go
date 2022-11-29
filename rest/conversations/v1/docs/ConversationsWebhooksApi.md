# ConversationsWebhooksApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationScopedWebhook**](ConversationsWebhooksApi.md#CreateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks | 
[**DeleteConversationScopedWebhook**](ConversationsWebhooksApi.md#DeleteConversationScopedWebhook) | **Delete** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchConversationScopedWebhook**](ConversationsWebhooksApi.md#FetchConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**ListConversationScopedWebhook**](ConversationsWebhooksApi.md#ListConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks | 
[**UpdateConversationScopedWebhook**](ConversationsWebhooksApi.md#UpdateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 



## CreateConversationScopedWebhook

> ConversationsV1ConversationScopedWebhook CreateConversationScopedWebhook(ctx, ConversationSidoptional)



Create a new webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Target** | **string** | 
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**ConfigurationMethod** | **string** | 
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationReplayAfter** | **int** | The message index for which and it's successors the webhook will be replayed. Not set by default

### Return type

[**ConversationsV1ConversationScopedWebhook**](ConversationsV1ConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversationScopedWebhook

> DeleteConversationScopedWebhook(ctx, ConversationSidSid)



Remove an existing webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationScopedWebhook

> ConversationsV1ConversationScopedWebhook FetchConversationScopedWebhook(ctx, ConversationSidSid)



Fetch the configuration of a conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationScopedWebhook**](ConversationsV1ConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationScopedWebhook

> []ConversationsV1ConversationScopedWebhook ListConversationScopedWebhook(ctx, ConversationSidoptional)



Retrieve a list of all webhooks scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ConversationScopedWebhook**](ConversationsV1ConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationScopedWebhook

> ConversationsV1ConversationScopedWebhook UpdateConversationScopedWebhook(ctx, ConversationSidSidoptional)



Update an existing conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**ConfigurationMethod** | **string** | 
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.

### Return type

[**ConversationsV1ConversationScopedWebhook**](ConversationsV1ConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

