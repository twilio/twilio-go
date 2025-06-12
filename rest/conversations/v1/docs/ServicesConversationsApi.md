# ServicesConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceConversation**](ServicesConversationsApi.md#CreateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations | Create a new conversation in your service
[**DeleteServiceConversation**](ServicesConversationsApi.md#DeleteServiceConversation) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | Remove a conversation from your service
[**FetchServiceConversation**](ServicesConversationsApi.md#FetchServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | Fetch a conversation from your service
[**ListServiceConversation**](ServicesConversationsApi.md#ListServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations | Retrieve a list of conversations in your service
[**UpdateServiceConversation**](ServicesConversationsApi.md#UpdateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | Update an existing conversation in your service



## CreateServiceConversation

> ConversationsV1ServiceConversation CreateServiceConversation(ctx, ChatServiceSidoptional)

Create a new conversation in your service

Create a new conversation in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | [**string**](stringstring.md) | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) this conversation belongs to.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**State** | [**string**](string.md) | 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**BindingsEmailAddress** | **string** | The default email address that will be used when sending outbound emails in this conversation.
**BindingsEmailName** | **string** | The default name that will be used when sending outbound emails in this conversation.

### Return type

[**ConversationsV1ServiceConversation**](ConversationsV1ServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceConversation

> DeleteServiceConversation(ctx, ChatServiceSidSidoptional)

Remove a conversation from your service

Remove a conversation from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | [**string**](stringstring.md) | The X-Twilio-Webhook-Enabled HTTP request header

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


## FetchServiceConversation

> ConversationsV1ServiceConversation FetchServiceConversation(ctx, ChatServiceSidSid)

Fetch a conversation from your service

Fetch a conversation from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceConversation**](ConversationsV1ServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversation

> []ConversationsV1ServiceConversation ListServiceConversation(ctx, ChatServiceSidoptional)

Retrieve a list of conversations in your service

Retrieve a list of conversations in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**StartDate** | **string** | Specifies the beginning of the date range for filtering Conversations based on their creation date. Conversations that were created on or after this date will be included in the results. The date must be in ISO8601 format, specifically starting at the beginning of the specified date (YYYY-MM-DDT00:00:00Z), for precise filtering. This parameter can be combined with other filters. If this filter is used, the returned list is sorted by latest conversation creation date in descending order.
**EndDate** | **string** | Defines the end of the date range for filtering conversations by their creation date. Only conversations that were created on or before this date will appear in the results.  The date must be in ISO8601 format, specifically capturing up to the end of the specified date (YYYY-MM-DDT23:59:59Z), to ensure that conversations from the entire end day are included. This parameter can be combined with other filters. If this filter is used, the returned list is sorted by latest conversation creation date in descending order.
**State** | [**string**](stringstring.md) | State for sorting and filtering list of Conversations. Can be `active`, `inactive` or `closed`
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 100.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceConversation**](ConversationsV1ServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversation

> ConversationsV1ServiceConversation UpdateServiceConversation(ctx, ChatServiceSidSidoptional)

Update an existing conversation in your service

Update an existing conversation in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | [**string**](stringstring.md) | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) this conversation belongs to.
**State** | [**string**](string.md) | 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**BindingsEmailAddress** | **string** | The default email address that will be used when sending outbound emails in this conversation.
**BindingsEmailName** | **string** | The default name that will be used when sending outbound emails in this conversation.

### Return type

[**ConversationsV1ServiceConversation**](ConversationsV1ServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

