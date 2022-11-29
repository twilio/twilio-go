# ServicesConfigurationWebhooksApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchServiceWebhookConfiguration**](ServicesConfigurationWebhooksApi.md#FetchServiceWebhookConfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Webhooks | 
[**UpdateServiceWebhookConfiguration**](ServicesConfigurationWebhooksApi.md#UpdateServiceWebhookConfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Webhooks | 



## FetchServiceWebhookConfiguration

> ConversationsV1ServiceWebhookConfiguration FetchServiceWebhookConfiguration(ctx, ChatServiceSid)



Fetch a specific service webhook configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceWebhookConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceWebhookConfiguration**](ConversationsV1ServiceWebhookConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceWebhookConfiguration

> ConversationsV1ServiceWebhookConfiguration UpdateServiceWebhookConfiguration(ctx, ChatServiceSidoptional)



Update a specific Webhook.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceWebhookConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to.
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to.
**Filters** | **[]string** | The list of events that your configured webhook targets will receive. Events not configured here will not fire. Possible values are `onParticipantAdd`, `onParticipantAdded`, `onDeliveryUpdated`, `onConversationUpdated`, `onConversationRemove`, `onParticipantRemove`, `onConversationUpdate`, `onMessageAdd`, `onMessageRemoved`, `onParticipantUpdated`, `onConversationAdded`, `onMessageAdded`, `onConversationAdd`, `onConversationRemoved`, `onParticipantUpdate`, `onMessageRemove`, `onMessageUpdated`, `onParticipantRemoved`, `onMessageUpdate` or `onConversationStateUpdated`.
**Method** | **string** | The HTTP method to be used when sending a webhook request. One of `GET` or `POST`.

### Return type

[**ConversationsV1ServiceWebhookConfiguration**](ConversationsV1ServiceWebhookConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

