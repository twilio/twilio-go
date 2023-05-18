# ConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversation**](ConversationsApi.md#CreateConversation) | **Post** /v1/Conversations | 
[**DeleteConversation**](ConversationsApi.md#DeleteConversation) | **Delete** /v1/Conversations/{Sid} | 
[**FetchConversation**](ConversationsApi.md#FetchConversation) | **Get** /v1/Conversations/{Sid} | 
[**ListConversation**](ConversationsApi.md#ListConversation) | **Get** /v1/Conversations | 
[**UpdateConversation**](ConversationsApi.md#UpdateConversation) | **Post** /v1/Conversations/{Sid} | 



## CreateConversation

> ConversationsV1Conversation CreateConversation(ctx, optional)



Create a new conversation in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**State** | **string** | 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.

### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversation

> DeleteConversation(ctx, Sidoptional)



Remove a conversation from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header

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


## FetchConversation

> ConversationsV1Conversation FetchConversation(ctx, Sid)



Fetch a conversation from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversation

> []ConversationsV1Conversation ListConversation(ctx, optional)



Retrieve a list of conversations in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**StartDate** | **string** | Start date or time in ISO8601 format for filtering list of Conversations. If a date is provided, the start time of the date is used (YYYY-MM-DDT00:00:00Z). Can be combined with other filters.
**EndDate** | **string** | End date or time in ISO8601 format for filtering list of Conversations. If a date is provided, the end time of the date is used (YYYY-MM-DDT23:59:59Z). Can be combined with other filters.
**State** | **string** | State for sorting and filtering list of Conversations. Can be `active`, `inactive` or `closed`
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversation

> ConversationsV1Conversation UpdateConversation(ctx, Sidoptional)



Update an existing conversation in your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**State** | **string** | 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.

### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

