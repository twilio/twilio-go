# ConfigurationApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConfiguration**](ConfigurationApi.md#FetchConfiguration) | **Get** /v1/Configuration | 
[**UpdateConfiguration**](ConfigurationApi.md#UpdateConfiguration) | **Post** /v1/Configuration | 



## FetchConfiguration

> ConversationsV1Configuration FetchConfiguration(ctx, )



Fetch the global configuration of conversations on your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


### Return type

[**ConversationsV1Configuration**](ConversationsV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ConversationsV1Configuration UpdateConfiguration(ctx, optional)



Update the global configuration of conversations on your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**DefaultChatServiceSid** | **string** | The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to use when creating a conversation.
**DefaultMessagingServiceSid** | **string** | The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) to use when creating a conversation.
**DefaultInactiveTimer** | **string** | Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**DefaultClosedTimer** | **string** | Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.

### Return type

[**ConversationsV1Configuration**](ConversationsV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

