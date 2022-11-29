# ServicesConfigurationApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchServiceConfiguration**](ServicesConfigurationApi.md#FetchServiceConfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration | 
[**UpdateServiceConfiguration**](ServicesConfigurationApi.md#UpdateServiceConfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration | 



## FetchServiceConfiguration

> ConversationsV1ServiceConfiguration FetchServiceConfiguration(ctx, ChatServiceSid)



Fetch the configuration of a conversation service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the Service configuration resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceConfiguration**](ConversationsV1ServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConfiguration

> ConversationsV1ServiceConfiguration UpdateServiceConfiguration(ctx, ChatServiceSidoptional)



Update configuration settings of a conversation service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the Service configuration resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**DefaultConversationCreatorRoleSid** | **string** | The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**DefaultConversationRoleSid** | **string** | The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**DefaultChatServiceRoleSid** | **string** | The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**ReachabilityEnabled** | **bool** | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is `false`.

### Return type

[**ConversationsV1ServiceConfiguration**](ConversationsV1ServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

