# ServicesConversationsWebhooksApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceConversationScopedWebhook**](ServicesConversationsWebhooksApi.md#CreateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**DeleteServiceConversationScopedWebhook**](ServicesConversationsWebhooksApi.md#DeleteServiceConversationScopedWebhook) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchServiceConversationScopedWebhook**](ServicesConversationsWebhooksApi.md#FetchServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**ListServiceConversationScopedWebhook**](ServicesConversationsWebhooksApi.md#ListServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**UpdateServiceConversationScopedWebhook**](ServicesConversationsWebhooksApi.md#UpdateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 



## CreateServiceConversationScopedWebhook

> ConversationsV1ServiceConversationScopedWebhook CreateServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidoptional)



Create a new webhook scoped to the conversation in a specific service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationScopedWebhookParams struct


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

[**ConversationsV1ServiceConversationScopedWebhook**](ConversationsV1ServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceConversationScopedWebhook

> DeleteServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSid)



Remove an existing webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationScopedWebhookParams struct


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


## FetchServiceConversationScopedWebhook

> ConversationsV1ServiceConversationScopedWebhook FetchServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSid)



Fetch the configuration of a conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceConversationScopedWebhook**](ConversationsV1ServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationScopedWebhook

> []ConversationsV1ServiceConversationScopedWebhook ListServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidoptional)



Retrieve a list of all webhooks scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceConversationScopedWebhook**](ConversationsV1ServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationScopedWebhook

> ConversationsV1ServiceConversationScopedWebhook UpdateServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSidoptional)



Update an existing conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**ConfigurationMethod** | **string** | 
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.

### Return type

[**ConversationsV1ServiceConversationScopedWebhook**](ConversationsV1ServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

