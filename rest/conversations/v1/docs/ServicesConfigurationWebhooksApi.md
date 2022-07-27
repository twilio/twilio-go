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
**Filters** | **[]string** | The list of events that your configured webhook targets will receive. Events not configured here will not fire. Possible values are &#x60;onParticipantAdd&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onDeliveryUpdated&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationRemove&#x60;, &#x60;onParticipantRemove&#x60;, &#x60;onConversationUpdate&#x60;, &#x60;onMessageAdd&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onConversationAdded&#x60;, &#x60;onMessageAdded&#x60;, &#x60;onConversationAdd&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantUpdate&#x60;, &#x60;onMessageRemove&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onParticipantRemoved&#x60;, &#x60;onMessageUpdate&#x60; or &#x60;onConversationStateUpdated&#x60;.
**Method** | **string** | The HTTP method to be used when sending a webhook request. One of &#x60;GET&#x60; or &#x60;POST&#x60;.

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

