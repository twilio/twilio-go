# DefaultApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversation**](DefaultApi.md#CreateConversation) | **Post** /v1/Conversations | 
[**CreateConversationMessage**](DefaultApi.md#CreateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages | 
[**CreateConversationParticipant**](DefaultApi.md#CreateConversationParticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants | 
[**CreateConversationScopedWebhook**](DefaultApi.md#CreateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v1/Credentials | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v1/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateServiceConversation**](DefaultApi.md#CreateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations | 
[**CreateServiceConversationMessage**](DefaultApi.md#CreateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**CreateServiceConversationParticipant**](DefaultApi.md#CreateServiceConversationParticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
[**CreateServiceConversationScopedWebhook**](DefaultApi.md#CreateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**CreateServiceRole**](DefaultApi.md#CreateServiceRole) | **Post** /v1/Services/{ChatServiceSid}/Roles | 
[**CreateServiceUser**](DefaultApi.md#CreateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v1/Users | 
[**DeleteConversation**](DefaultApi.md#DeleteConversation) | **Delete** /v1/Conversations/{Sid} | 
[**DeleteConversationMessage**](DefaultApi.md#DeleteConversationMessage) | **Delete** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**DeleteConversationParticipant**](DefaultApi.md#DeleteConversationParticipant) | **Delete** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**DeleteConversationScopedWebhook**](DefaultApi.md#DeleteConversationScopedWebhook) | **Delete** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v1/Credentials/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v1/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteServiceBinding**](DefaultApi.md#DeleteServiceBinding) | **Delete** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
[**DeleteServiceConversation**](DefaultApi.md#DeleteServiceConversation) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**DeleteServiceConversationMessage**](DefaultApi.md#DeleteServiceConversationMessage) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**DeleteServiceConversationParticipant**](DefaultApi.md#DeleteServiceConversationParticipant) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**DeleteServiceConversationScopedWebhook**](DefaultApi.md#DeleteServiceConversationScopedWebhook) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**DeleteServiceRole**](DefaultApi.md#DeleteServiceRole) | **Delete** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**DeleteServiceUser**](DefaultApi.md#DeleteServiceUser) | **Delete** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v1/Users/{Sid} | 
[**FetchConfiguration**](DefaultApi.md#FetchConfiguration) | **Get** /v1/Configuration | 
[**FetchConfigurationWebhook**](DefaultApi.md#FetchConfigurationWebhook) | **Get** /v1/Configuration/Webhooks | 
[**FetchConversation**](DefaultApi.md#FetchConversation) | **Get** /v1/Conversations/{Sid} | 
[**FetchConversationMessage**](DefaultApi.md#FetchConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchConversationMessageReceipt**](DefaultApi.md#FetchConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**FetchConversationParticipant**](DefaultApi.md#FetchConversationParticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**FetchConversationScopedWebhook**](DefaultApi.md#FetchConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v1/Credentials/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v1/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchServiceBinding**](DefaultApi.md#FetchServiceBinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
[**FetchServiceConfiguration**](DefaultApi.md#FetchServiceConfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration | 
[**FetchServiceConversation**](DefaultApi.md#FetchServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**FetchServiceConversationMessage**](DefaultApi.md#FetchServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchServiceConversationMessageReceipt**](DefaultApi.md#FetchServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**FetchServiceConversationParticipant**](DefaultApi.md#FetchServiceConversationParticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**FetchServiceConversationScopedWebhook**](DefaultApi.md#FetchServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchServiceNotification**](DefaultApi.md#FetchServiceNotification) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
[**FetchServiceRole**](DefaultApi.md#FetchServiceRole) | **Get** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**FetchServiceUser**](DefaultApi.md#FetchServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v1/Users/{Sid} | 
[**ListConversation**](DefaultApi.md#ListConversation) | **Get** /v1/Conversations | 
[**ListConversationMessage**](DefaultApi.md#ListConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages | 
[**ListConversationMessageReceipt**](DefaultApi.md#ListConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
[**ListConversationParticipant**](DefaultApi.md#ListConversationParticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants | 
[**ListConversationScopedWebhook**](DefaultApi.md#ListConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v1/Credentials | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v1/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListServiceBinding**](DefaultApi.md#ListServiceBinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings | 
[**ListServiceConversation**](DefaultApi.md#ListServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations | 
[**ListServiceConversationMessage**](DefaultApi.md#ListServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**ListServiceConversationMessageReceipt**](DefaultApi.md#ListServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
[**ListServiceConversationParticipant**](DefaultApi.md#ListServiceConversationParticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
[**ListServiceConversationScopedWebhook**](DefaultApi.md#ListServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**ListServiceRole**](DefaultApi.md#ListServiceRole) | **Get** /v1/Services/{ChatServiceSid}/Roles | 
[**ListServiceUser**](DefaultApi.md#ListServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v1/Users | 
[**UpdateConfiguration**](DefaultApi.md#UpdateConfiguration) | **Post** /v1/Configuration | 
[**UpdateConfigurationWebhook**](DefaultApi.md#UpdateConfigurationWebhook) | **Post** /v1/Configuration/Webhooks | 
[**UpdateConversation**](DefaultApi.md#UpdateConversation) | **Post** /v1/Conversations/{Sid} | 
[**UpdateConversationMessage**](DefaultApi.md#UpdateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**UpdateConversationParticipant**](DefaultApi.md#UpdateConversationParticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**UpdateConversationScopedWebhook**](DefaultApi.md#UpdateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v1/Credentials/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v1/Roles/{Sid} | 
[**UpdateServiceConfiguration**](DefaultApi.md#UpdateServiceConfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration | 
[**UpdateServiceConversation**](DefaultApi.md#UpdateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**UpdateServiceConversationMessage**](DefaultApi.md#UpdateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**UpdateServiceConversationParticipant**](DefaultApi.md#UpdateServiceConversationParticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**UpdateServiceConversationScopedWebhook**](DefaultApi.md#UpdateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**UpdateServiceNotification**](DefaultApi.md#UpdateServiceNotification) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
[**UpdateServiceRole**](DefaultApi.md#UpdateServiceRole) | **Post** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**UpdateServiceUser**](DefaultApi.md#UpdateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v1/Users/{Sid} | 



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
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

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


## CreateConversationMessage

> ConversationsV1ConversationConversationMessage CreateConversationMessage(ctx, ConversationSidoptional)



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
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.
**MediaSid** | **string** | The Media SID to be attached to the new Message.

### Return type

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationParticipant

> ConversationsV1ConversationConversationParticipant CreateConversationParticipant(ctx, ConversationSidoptional)



Add a new participant to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
**MessagingBindingAddress** | **string** | The address of the participant&#39;s device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook CreateConversationScopedWebhook(ctx, ConversationSidoptional)



Create a new webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationReplayAfter** | **int** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60;

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> ConversationsV1Credential CreateCredential(ctx, optional)



Add a new push notification credential to your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;.
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;fcm&#x60;, &#x60;gcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ConversationsV1Role CreateRole(ctx, optional)



Create a new user role in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
**Type** | **string** | The type of role. Can be: &#x60;conversation&#x60; for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or &#x60;service&#x60; for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ConversationsV1Service CreateService(ctx, optional)



Create a new conversation service on your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The human-readable name of this service, limited to 256 characters. Optional.

### Return type

[**ConversationsV1Service**](ConversationsV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversation

> ConversationsV1ServiceServiceConversation CreateServiceConversation(ctx, ChatServiceSidoptional)



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
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage CreateServiceConversationMessage(ctx, ChatServiceSidConversationSidoptional)



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
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.
**MediaSid** | **string** | The Media SID to be attached to the new Message.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant CreateServiceConversationParticipant(ctx, ChatServiceSidConversationSidoptional)



Add a new participant to the conversation in a specific service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
**MessagingBindingAddress** | **string** | The address of the participant&#39;s device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook CreateServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidoptional)



Create a new webhook scoped to the conversation in a specific service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationReplayAfter** | **int** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
**Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60;

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceRole

> ConversationsV1ServiceServiceRole CreateServiceRole(ctx, ChatServiceSidoptional)



Create a new user role in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to create the Role resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
**Type** | **string** | The type of role. Can be: &#x60;conversation&#x60; for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or &#x60;service&#x60; for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles.

### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceUser

> ConversationsV1ServiceServiceUser CreateServiceUser(ctx, ChatServiceSidoptional)



Add a new conversation user to your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Identity** | **string** | The application-defined string that uniquely identifies the resource&#39;s User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ConversationsV1User CreateUser(ctx, optional)



Add a new conversation user to your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Identity** | **string** | The application-defined string that uniquely identifies the resource&#39;s User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1User**](ConversationsV1User.md)

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
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

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


## DeleteConversationParticipant

> DeleteConversationParticipant(ctx, ConversationSidSidoptional)



Remove a participant from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationParticipantParams struct


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


## DeleteConversationScopedWebhook

> DeleteConversationScopedWebhook(ctx, ConversationSidSid)



Remove an existing webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteCredential

> DeleteCredential(ctx, Sid)



Remove a push notification credential from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteRole

> DeleteRole(ctx, Sid)



Remove a user role from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteService

> DeleteService(ctx, Sid)



Remove a conversation service with all its nested resources from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteServiceBinding

> DeleteServiceBinding(ctx, ChatServiceSidSid)



Remove a push notification binding from the conversation service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Binding resource from.
**Sid** | **string** | The SID of the Binding resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteServiceConversation

> DeleteServiceConversation(ctx, ChatServiceSidSidoptional)



Remove a conversation from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationParams struct


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


## DeleteServiceConversationParticipant

> DeleteServiceConversationParticipant(ctx, ChatServiceSidConversationSidSidoptional)



Remove a participant from the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationParticipantParams struct


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


## DeleteServiceConversationScopedWebhook

> DeleteServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSid)



Remove an existing webhook scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteServiceRole

> DeleteServiceRole(ctx, ChatServiceSidSid)



Remove a user role from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Role resource from.
**Sid** | **string** | The SID of the Role resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteServiceUser

> DeleteServiceUser(ctx, ChatServiceSidSidoptional)



Remove a conversation user from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the User resource from.
**Sid** | **string** | The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceUserParams struct


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


## DeleteUser

> DeleteUser(ctx, Sidoptional)



Remove a conversation user from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct


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


## FetchConfigurationWebhook

> ConversationsV1ConfigurationConfigurationWebhook FetchConfigurationWebhook(ctx, )



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationWebhookParams struct


### Return type

[**ConversationsV1ConfigurationConfigurationWebhook**](ConversationsV1ConfigurationConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

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


## FetchConversationMessage

> ConversationsV1ConversationConversationMessage FetchConversationMessage(ctx, ConversationSidSid)



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

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationMessageReceipt

> ConversationsV1ConversationConversationMessageConversationMessageReceipt FetchConversationMessageReceipt(ctx, ConversationSidMessageSidSid)



Fetch the delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationConversationMessageConversationMessageReceipt**](ConversationsV1ConversationConversationMessageConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationParticipant

> ConversationsV1ConversationConversationParticipant FetchConversationParticipant(ctx, ConversationSidSid)



Fetch a participant of the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook FetchConversationScopedWebhook(ctx, ConversationSidSid)



Fetch the configuration of a conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ConversationsV1Credential FetchCredential(ctx, Sid)



Fetch a push notification credential from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ConversationsV1Role FetchRole(ctx, Sid)



Fetch a user role from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ConversationsV1Service FetchService(ctx, Sid)



Fetch a conversation service from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1Service**](ConversationsV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceBinding

> ConversationsV1ServiceServiceBinding FetchServiceBinding(ctx, ChatServiceSidSid)



Fetch a push notification binding from the conversation service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceBinding**](ConversationsV1ServiceServiceBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConfiguration

> ConversationsV1ServiceServiceConfiguration FetchServiceConfiguration(ctx, ChatServiceSid)



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

[**ConversationsV1ServiceServiceConfiguration**](ConversationsV1ServiceServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversation

> ConversationsV1ServiceServiceConversation FetchServiceConversation(ctx, ChatServiceSidSid)



Fetch a conversation from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage FetchServiceConversationMessage(ctx, ChatServiceSidConversationSidSid)



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

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationMessageReceipt

> ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt FetchServiceConversationMessageReceipt(ctx, ChatServiceSidConversationSidMessageSidSid)



Fetch the delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt**](ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant FetchServiceConversationParticipant(ctx, ChatServiceSidConversationSidSid)



Fetch a participant of the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook FetchServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSid)



Fetch the configuration of a conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceNotification

> ConversationsV1ServiceServiceConfigurationServiceNotification FetchServiceNotification(ctx, ChatServiceSid)



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

[**ConversationsV1ServiceServiceConfigurationServiceNotification**](ConversationsV1ServiceServiceConfigurationServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceRole

> ConversationsV1ServiceServiceRole FetchServiceRole(ctx, ChatServiceSidSid)



Fetch a user role from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the Role resource from.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceUser

> ConversationsV1ServiceServiceUser FetchServiceUser(ctx, ChatServiceSidSid)



Fetch a conversation user from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the User resource from.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ConversationsV1User FetchUser(ctx, Sid)



Fetch a conversation user from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1User**](ConversationsV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversation

> ListConversationResponse ListConversation(ctx, optional)



Retrieve a list of conversations in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationResponse**](ListConversationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessage

> ListConversationMessageResponse ListConversationMessage(ctx, ConversationSidoptional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationMessageResponse**](ListConversationMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessageReceipt

> ListConversationMessageReceiptResponse ListConversationMessageReceipt(ctx, ConversationSidMessageSidoptional)



Retrieve a list of all delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationMessageReceiptResponse**](ListConversationMessageReceiptResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationParticipant

> ListConversationParticipantResponse ListConversationParticipant(ctx, ConversationSidoptional)



Retrieve a list of all participants of the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationParticipantResponse**](ListConversationParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationScopedWebhook

> ListConversationScopedWebhookResponse ListConversationScopedWebhook(ctx, ConversationSidoptional)



Retrieve a list of all webhooks scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationScopedWebhookResponse**](ListConversationScopedWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> ListCredentialResponse ListCredential(ctx, optional)



Retrieve a list of all push notification credentials on your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialResponse**](ListCredentialResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> ListRoleResponse ListRole(ctx, optional)



Retrieve a list of all user roles in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRoleResponse**](ListRoleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



Retrieve a list of all conversation services on your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceBinding

> ListServiceBindingResponse ListServiceBinding(ctx, ChatServiceSidoptional)



Retrieve a list of all push notification bindings in the conversation service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **[]string** | The push technology used by the Binding resources to read.  Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
**Identity** | **[]string** | The identity of a [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource) this binding belongs to. See [access tokens](https://www.twilio.com/docs/conversations/create-tokens) for more details.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceBindingResponse**](ListServiceBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversation

> ListServiceConversationResponse ListServiceConversation(ctx, ChatServiceSidoptional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationResponse**](ListServiceConversationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessage

> ListServiceConversationMessageResponse ListServiceConversationMessage(ctx, ChatServiceSidConversationSidoptional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationMessageResponse**](ListServiceConversationMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessageReceipt

> ListServiceConversationMessageReceiptResponse ListServiceConversationMessageReceipt(ctx, ChatServiceSidConversationSidMessageSidoptional)



Retrieve a list of all delivery and read receipts of the conversation message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationMessageReceiptParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationMessageReceiptResponse**](ListServiceConversationMessageReceiptResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationParticipant

> ListServiceConversationParticipantResponse ListServiceConversationParticipant(ctx, ChatServiceSidConversationSidoptional)



Retrieve a list of all participants of the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationParticipantResponse**](ListServiceConversationParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationScopedWebhook

> ListServiceConversationScopedWebhookResponse ListServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidoptional)



Retrieve a list of all webhooks scoped to the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationScopedWebhookResponse**](ListServiceConversationScopedWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceRole

> ListServiceRoleResponse ListServiceRole(ctx, ChatServiceSidoptional)



Retrieve a list of all user roles in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the Role resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceRoleResponse**](ListServiceRoleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceUser

> ListServiceUserResponse ListServiceUser(ctx, ChatServiceSidoptional)



Retrieve a list of all conversation users in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the User resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceUserResponse**](ListServiceUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> ListUserResponse ListUser(ctx, optional)



Retrieve a list of all conversation users in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserResponse**](ListUserResponse.md)

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
**DefaultClosedTimer** | **string** | Default ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
**DefaultInactiveTimer** | **string** | Default ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
**DefaultMessagingServiceSid** | **string** | The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) to use when creating a conversation.

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


## UpdateConfigurationWebhook

> ConversationsV1ConfigurationConfigurationWebhook UpdateConfigurationWebhook(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**Filters** | **[]string** | The list of webhook event triggers that are enabled for this Service: &#x60;onMessageAdded&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onParticipantRemoved&#x60;
**Method** | **string** | The HTTP method to be used when sending a webhook request.
**PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to.
**PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to.
**Target** | **string** | The routing target of the webhook.

### Return type

[**ConversationsV1ConfigurationConfigurationWebhook**](ConversationsV1ConfigurationConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
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
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

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


## UpdateConversationMessage

> ConversationsV1ConversationConversationMessage UpdateConversationMessage(ctx, ConversationSidSidoptional)



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
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.

### Return type

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationParticipant

> ConversationsV1ConversationConversationParticipant UpdateConversationParticipant(ctx, ConversationSidSidoptional)



Update an existing participant in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
**LastReadMessageIndex** | **int** | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**LastReadTimestamp** | **string** | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. &#39;null&#39; value will remove it.
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number that the participant is in contact with. &#39;null&#39; value will remove it.
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook UpdateConversationScopedWebhook(ctx, ConversationSidSidoptional)



Update an existing conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ConversationsV1Credential UpdateCredential(ctx, Sidoptional)



Update an existing push notification credential on your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;.
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;fcm&#x60;, &#x60;gcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ConversationsV1Role UpdateRole(ctx, Sidoptional)



Update an existing user role in your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConfiguration

> ConversationsV1ServiceServiceConfiguration UpdateServiceConfiguration(ctx, ChatServiceSidoptional)



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
**DefaultChatServiceRoleSid** | **string** | The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**DefaultConversationCreatorRoleSid** | **string** | The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**DefaultConversationRoleSid** | **string** | The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
**ReachabilityEnabled** | **bool** | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is &#x60;false&#x60;.

### Return type

[**ConversationsV1ServiceServiceConfiguration**](ConversationsV1ServiceServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversation

> ConversationsV1ServiceServiceConversation UpdateServiceConversation(ctx, ChatServiceSidSidoptional)



Update an existing conversation in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
**State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
**TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
**TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage UpdateServiceConversationMessage(ctx, ChatServiceSidConversationSidSidoptional)



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
**Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
**Body** | **string** | The content of the message, can be up to 1,600 characters long.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant UpdateServiceConversationParticipant(ctx, ChatServiceSidConversationSidSidoptional)



Update an existing participant in the conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
**DateCreated** | **time.Time** | The date that this resource was created.
**DateUpdated** | **time.Time** | The date that this resource was last updated.
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
**LastReadMessageIndex** | **int** | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**LastReadTimestamp** | **string** | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. &#39;null&#39; value will remove it.
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number that the participant is in contact with. &#39;null&#39; value will remove it.
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook UpdateServiceConversationScopedWebhook(ctx, ChatServiceSidConversationSidSidoptional)



Update an existing conversation-scoped webhook

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationScopedWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
**ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
**ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
**ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
**ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceNotification

> ConversationsV1ServiceServiceConfigurationServiceNotification UpdateServiceNotification(ctx, ChatServiceSidoptional)



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
**AddedToConversationEnabled** | **bool** | Whether to send a notification when a participant is added to a conversation. The default is &#x60;false&#x60;.
**AddedToConversationSound** | **string** | The name of the sound to play when a participant is added to a conversation and &#x60;added_to_conversation.enabled&#x60; is &#x60;true&#x60;.
**AddedToConversationTemplate** | **string** | The template to use to create the notification text displayed when a participant is added to a conversation and &#x60;added_to_conversation.enabled&#x60; is &#x60;true&#x60;.
**LogEnabled** | **bool** | Weather the notification logging is enabled.
**NewMessageBadgeCountEnabled** | **bool** | Whether the new message badge is enabled. The default is &#x60;false&#x60;.
**NewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a conversation. The default is &#x60;false&#x60;.
**NewMessageSound** | **string** | The name of the sound to play when a new message is added to a conversation and &#x60;new_message.enabled&#x60; is &#x60;true&#x60;.
**NewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a conversation and &#x60;new_message.enabled&#x60; is &#x60;true&#x60;.
**RemovedFromConversationEnabled** | **bool** | Whether to send a notification to a user when they are removed from a conversation. The default is &#x60;false&#x60;.
**RemovedFromConversationSound** | **string** | The name of the sound to play to a user when they are removed from a conversation and &#x60;removed_from_conversation.enabled&#x60; is &#x60;true&#x60;.
**RemovedFromConversationTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a conversation and &#x60;removed_from_conversation.enabled&#x60; is &#x60;true&#x60;.

### Return type

[**ConversationsV1ServiceServiceConfigurationServiceNotification**](ConversationsV1ServiceServiceConfigurationServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceRole

> ConversationsV1ServiceServiceRole UpdateServiceRole(ctx, ChatServiceSidSidoptional)



Update an existing user role in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to update the Role resource in.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.

### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceUser

> ConversationsV1ServiceServiceUser UpdateServiceUser(ctx, ChatServiceSidSidoptional)



Update an existing conversation user in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ConversationsV1User UpdateUser(ctx, Sidoptional)



Update an existing conversation user in your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1User**](ConversationsV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

