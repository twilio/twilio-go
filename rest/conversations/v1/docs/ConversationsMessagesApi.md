# ConversationsMessagesApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationMessage**](ConversationsMessagesApi.md#CreateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages | 
[**DeleteConversationMessage**](ConversationsMessagesApi.md#DeleteConversationMessage) | **Delete** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchConversationMessage**](ConversationsMessagesApi.md#FetchConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**ListConversationMessage**](ConversationsMessagesApi.md#ListConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages | 
[**UpdateConversationMessage**](ConversationsMessagesApi.md#UpdateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 



## CreateConversationMessage

> ConversationsV1ConversationMessage CreateConversationMessage(ctx, ConversationSidoptional)



Add a new message to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Author** | **string** | The channel specific identifier of the message's author. Defaults to `system`.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. `null` if the message has not been edited.
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**MediaSid** | **string** | The Media SID to be attached to the new Message.
**ContentSid** | **string** | The unique ID of the multi-channel [Rich Content](https://www.twilio.com/docs/content-api) template, required for template-generated messages.  **Note** that if this field is set, `Body` and `MediaSid` parameters are ignored.
**ContentVariables** | **string** | A structurally valid JSON string that contains values to resolve Rich Content template variables.

### Return type

[**ConversationsV1ConversationMessage**](ConversationsV1ConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversationMessage

> DeleteConversationMessage(ctx, ConversationSidSidoptional)



Remove a message from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationMessageParams struct


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


## FetchConversationMessage

> ConversationsV1ConversationMessage FetchConversationMessage(ctx, ConversationSidSid)



Fetch a message from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationMessage**](ConversationsV1ConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessage

> []ConversationsV1ConversationMessage ListConversationMessage(ctx, ConversationSidoptional)



Retrieve a list of all messages in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending), with `asc` as the default.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ConversationMessage**](ConversationsV1ConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationMessage

> ConversationsV1ConversationMessage UpdateConversationMessage(ctx, ConversationSidSidoptional)



Update an existing message in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Author** | **string** | The channel specific identifier of the message's author. Defaults to `system`.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. `null` if the message has not been edited.
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.

### Return type

[**ConversationsV1ConversationMessage**](ConversationsV1ConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

