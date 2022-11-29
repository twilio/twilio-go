# ConfigurationWebhooksApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConfigurationWebhook**](ConfigurationWebhooksApi.md#FetchConfigurationWebhook) | **Get** /v1/Configuration/Webhooks | 
[**UpdateConfigurationWebhook**](ConfigurationWebhooksApi.md#UpdateConfigurationWebhook) | **Post** /v1/Configuration/Webhooks | 



## FetchConfigurationWebhook

> ConversationsV1ConfigurationWebhook FetchConfigurationWebhook(ctx, )





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationWebhookParams struct


### Return type

[**ConversationsV1ConfigurationWebhook**](ConversationsV1ConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationWebhook

> ConversationsV1ConfigurationWebhook UpdateConfigurationWebhook(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Method** | **string** | The HTTP method to be used when sending a webhook request.
**Filters** | **[]string** | The list of webhook event triggers that are enabled for this Service: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to.
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to.
**Target** | **string** | 

### Return type

[**ConversationsV1ConfigurationWebhook**](ConversationsV1ConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

