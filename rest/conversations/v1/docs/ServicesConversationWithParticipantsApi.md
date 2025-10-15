# ServicesConversationWithParticipantsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceConversationWithParticipants**](ServicesConversationWithParticipantsApi.md#CreateServiceConversationWithParticipants) | **Post** /v1/Services/{ChatServiceSid}/ConversationWithParticipants | Create a new conversation with the list of participants in your account&#39;s default service



## CreateServiceConversationWithParticipants

> ConversationsV1ServiceConversationWithParticipants CreateServiceConversationWithParticipants(ctx, ChatServiceSidoptional)

Create a new conversation with the list of participants in your account's default service

Create a new conversation with the list of participants in your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationWithParticipantsParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | [**string**](stringstring.md) | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/messaging/api/service-resource) this conversation belongs to.
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
**State** | [**string**](string.md) | 
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
**BindingsEmailAddress** | **string** | The default email address that will be used when sending outbound emails in this conversation.
**BindingsEmailName** | **string** | The default name that will be used when sending outbound emails in this conversation.
**Participant** | **[]string** | The participant to be added to the conversation in JSON format. The JSON object attributes are as parameters in [Participant Resource](https://www.twilio.com/docs/conversations/api/conversation-participant-resource). The maximum number of participants that can be added in a single request is 10.

### Return type

[**ConversationsV1ServiceConversationWithParticipants**](ConversationsV1ServiceConversationWithParticipants.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

