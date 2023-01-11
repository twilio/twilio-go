# ServicesConversationsMessagesApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceConversationMessage**](ServicesConversationsMessagesApi.md#CreateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**DeleteServiceConversationMessage**](ServicesConversationsMessagesApi.md#DeleteServiceConversationMessage) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchServiceConversationMessage**](ServicesConversationsMessagesApi.md#FetchServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**ListServiceConversationMessage**](ServicesConversationsMessagesApi.md#ListServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**UpdateServiceConversationMessage**](ServicesConversationsMessagesApi.md#UpdateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 



## CreateServiceConversationMessage

> ConversationsV1ServiceConversationMessage CreateServiceConversationMessage(ctx, ChatServiceSidConversationSidoptional)



Add a new message to the conversation in a specific service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationMessageParams struct


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

[**ConversationsV1ServiceConversationMessage**](ConversationsV1ServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceConversationMessage

> DeleteServiceConversationMessage(ctx, ChatServiceSidConversationSidSidoptional)



Remove a message from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationMessageParams struct


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


## FetchServiceConversationMessage

> ConversationsV1ServiceConversationMessage FetchServiceConversationMessage(ctx, ChatServiceSidConversationSidSid)



Fetch a message from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceConversationMessage**](ConversationsV1ServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessage

> []ConversationsV1ServiceConversationMessage ListServiceConversationMessage(ctx, ChatServiceSidConversationSidoptional)



Retrieve a list of all messages in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending), with `asc` as the default.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceConversationMessage**](ConversationsV1ServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationMessage

> ConversationsV1ServiceConversationMessage UpdateServiceConversationMessage(ctx, ChatServiceSidConversationSidSidoptional)



Update an existing message in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Author** | **string** | The channel specific identifier of the message's author. Defaults to `system`.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. `null` if the message has not been edited.
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.

### Return type

[**ConversationsV1ServiceConversationMessage**](ConversationsV1ServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

