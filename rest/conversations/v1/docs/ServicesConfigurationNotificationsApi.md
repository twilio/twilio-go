# ServicesConfigurationNotificationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchServiceNotification**](ServicesConfigurationNotificationsApi.md#FetchServiceNotification) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
[**UpdateServiceNotification**](ServicesConfigurationNotificationsApi.md#UpdateServiceNotification) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 



## FetchServiceNotification

> ConversationsV1ServiceNotification FetchServiceNotification(ctx, ChatServiceSid)



Fetch push notification service settings

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceNotificationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceNotification**](ConversationsV1ServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceNotification

> ConversationsV1ServiceNotification UpdateServiceNotification(ctx, ChatServiceSidoptional)



Update push notification service settings

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceNotificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**LogEnabled** | **bool** | Weather the notification logging is enabled.
**NewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a conversation. The default is `false`.
**NewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a conversation and `new_message.enabled` is `true`.
**NewMessageSound** | **string** | The name of the sound to play when a new message is added to a conversation and `new_message.enabled` is `true`.
**NewMessageBadgeCountEnabled** | **bool** | Whether the new message badge is enabled. The default is `false`.
**AddedToConversationEnabled** | **bool** | Whether to send a notification when a participant is added to a conversation. The default is `false`.
**AddedToConversationTemplate** | **string** | The template to use to create the notification text displayed when a participant is added to a conversation and `added_to_conversation.enabled` is `true`.
**AddedToConversationSound** | **string** | The name of the sound to play when a participant is added to a conversation and `added_to_conversation.enabled` is `true`.
**RemovedFromConversationEnabled** | **bool** | Whether to send a notification to a user when they are removed from a conversation. The default is `false`.
**RemovedFromConversationTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`.
**RemovedFromConversationSound** | **string** | The name of the sound to play to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`.
**NewMessageWithMediaEnabled** | **bool** | Whether to send a notification when a new message with media/file attachments is added to a conversation. The default is `false`.
**NewMessageWithMediaTemplate** | **string** | The template to use to create the notification text displayed when a new message with media/file attachments is added to a conversation and `new_message.attachments.enabled` is `true`.

### Return type

[**ConversationsV1ServiceNotification**](ConversationsV1ServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

