# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v2/Services/{ServiceSid}/Channels | 
[**CreateChannelWebhook**](DefaultApi.md#CreateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v2/Credentials | 
[**CreateInvite**](DefaultApi.md#CreateInvite) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**CreateMember**](DefaultApi.md#CreateMember) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v2/Services/{ServiceSid}/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v2/Services | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v2/Services/{ServiceSid}/Users | 
[**DeleteBinding**](DefaultApi.md#DeleteBinding) | **Delete** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**DeleteChannelWebhook**](DefaultApi.md#DeleteChannelWebhook) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v2/Credentials/{Sid} | 
[**DeleteInvite**](DefaultApi.md#DeleteInvite) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**DeleteMember**](DefaultApi.md#DeleteMember) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**DeleteUserBinding**](DefaultApi.md#DeleteUserBinding) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid} | 
[**DeleteUserChannel**](DefaultApi.md#DeleteUserChannel) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**FetchBinding**](DefaultApi.md#FetchBinding) | **Get** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**FetchChannelWebhook**](DefaultApi.md#FetchChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v2/Credentials/{Sid} | 
[**FetchInvite**](DefaultApi.md#FetchInvite) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**FetchUserBinding**](DefaultApi.md#FetchUserBinding) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid} | 
[**FetchUserChannel**](DefaultApi.md#FetchUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**ListBinding**](DefaultApi.md#ListBinding) | **Get** /v2/Services/{ServiceSid}/Bindings | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v2/Services/{ServiceSid}/Channels | 
[**ListChannelWebhook**](DefaultApi.md#ListChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v2/Credentials | 
[**ListInvite**](DefaultApi.md#ListInvite) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v2/Services/{ServiceSid}/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v2/Services | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v2/Services/{ServiceSid}/Users | 
[**ListUserBinding**](DefaultApi.md#ListUserBinding) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings | 
[**ListUserChannel**](DefaultApi.md#ListUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateChannel**](DefaultApi.md#UpdateChannel) | **Post** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**UpdateChannelWebhook**](DefaultApi.md#UpdateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v2/Credentials/{Sid} | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**UpdateUserChannel**](DefaultApi.md#UpdateUserChannel) | **Post** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 



## CreateChannel

> ChatV2ServiceChannel CreateChannel(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Channel resource under. | 
 **optional** | ***CreateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **createdBy** | **optional.String**| The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used in cases where a Channel is being recreated from a backup/separate source  and where a Message was previously updated. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
 **type_** | **optional.String**| The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the Channel resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service. | 

### Return type

[**ChatV2ServiceChannel**](chat.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelWebhook

> ChatV2ServiceChannelChannelWebhook CreateChannelWebhook(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to create the Webhook resource under. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Channel Webhook resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***CreateChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configurationFilters** | [**optional.Interface of []string**](string.md)| The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger). | 
 **configurationFlowSid** | **optional.String**| The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; is &#x60;studio&#x60;. | 
 **configurationMethod** | **optional.String**| The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **configurationRetryCount** | **optional.Int32**| The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0. | 
 **configurationTriggers** | [**optional.Interface of []string**](string.md)| A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;. | 
 **configurationUrl** | **optional.String**| The URL of the webhook to call using the &#x60;configuration.method&#x60;. | 
 **type_** | **optional.String**| The type of webhook. Can be: &#x60;webhook&#x60;, &#x60;studio&#x60;, or &#x60;trigger&#x60;. | 

### Return type

[**ChatV2ServiceChannelChannelWebhook**](chat.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> ChatV2Credential CreateCredential(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **optional.String**| [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. | 
 **certificate** | **optional.String**| [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60; | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
 **privateKey** | **optional.String**| [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60; | 
 **sandbox** | **optional.Bool**| [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
 **secret** | **optional.String**| [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. | 
 **type_** | **optional.String**| The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;. | 

### Return type

[**ChatV2Credential**](chat.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> ChatV2ServiceChannelInvite CreateInvite(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Invite resource under. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Invite resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***CreateInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInviteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **optional.String**| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info. | 
 **roleSid** | **optional.String**| The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) assigned to the new member. | 

### Return type

[**ChatV2ServiceChannelInvite**](chat.v2.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> ChatV2ServiceChannelMember CreateMember(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Member resource under. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Member resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***CreateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used when a Member is being recreated from a backup/separate source and where a Member was previously updated. | 
 **identity** | **optional.String**| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info. | 
 **lastConsumedMessageIndex** | **optional.Int32**| The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. This parameter should only be used when recreating a Member from a backup/separate source. | 
 **lastConsumptionTimestamp** | **optional.Time**| The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels). | 
 **roleSid** | **optional.String**| The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource). | 

### Return type

[**ChatV2ServiceChannelMember**](chat.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ChatV2ServiceChannelMessage CreateMessage(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Message resource under. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Message resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***CreateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **body** | **optional.String**| The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | 
 **from** | **optional.String**| The [Identity](https://www.twilio.com/docs/chat/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;. | 
 **lastUpdatedBy** | **optional.String**| The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable. | 
 **mediaSid** | **optional.String**| The SID of the [Media](https://www.twilio.com/docs/chat/rest/media) to attach to the new Message. | 

### Return type

[**ChatV2ServiceChannelMessage**](chat.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ChatV2ServiceRole CreateRole(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Role resource under. | 
 **optional** | ***CreateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
 **permission** | [**optional.Interface of []string**](string.md)| A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;. | 
 **type_** | **optional.String**| The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/rest/service-resource) roles. | 

### Return type

[**ChatV2ServiceRole**](chat.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ChatV2Service CreateService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. | 

### Return type

[**ChatV2Service**](chat.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ChatV2ServiceUser CreateUser(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the User resource under. | 
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. This value is often used for display purposes. | 
 **identity** | **optional.String**| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). This value is often a username or email address. See the Identity documentation for more info. | 
 **roleSid** | **optional.String**| The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the new User. | 

### Return type

[**ChatV2ServiceUser**](chat.v2.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinding

> DeleteBinding(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Binding resource from. | 
**sid** | **string**| The SID of the Binding resource to delete. | 

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


## DeleteChannel

> DeleteChannel(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from. | 
**sid** | **string**| The SID of the Channel resource to delete.  This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to delete. | 
 **optional** | ***DeleteChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 

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


## DeleteChannelWebhook

> DeleteChannelWebhook(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to delete the Webhook resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Channel Webhook resource to delete. | 

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

> DeleteCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Credential resource to delete. | 

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


## DeleteInvite

> DeleteInvite(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Invite resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Invite resource to delete. | 

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


## DeleteMember

> DeleteMember(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Member resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Member resource to delete. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value. | 
 **optional** | ***DeleteMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 

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


## DeleteMessage

> DeleteMessage(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Message resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Message resource to delete. | 
 **optional** | ***DeleteMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 

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

> DeleteRole(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Role resource from. | 
**sid** | **string**| The SID of the Role resource to delete. | 

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

> DeleteService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to delete. | 

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

> DeleteUser(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the User resource from. | 
**sid** | **string**| The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete. | 

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


## DeleteUserBinding

> DeleteUserBinding(ctx, serviceSid, userSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the User Binding resource from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resources to delete.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. | 
**sid** | **string**| The SID of the User Binding resource to delete. | 

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


## DeleteUserChannel

> DeleteUserChannel(ctx, serviceSid, userSid, channelSid)



Removes User from selected Channel.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource belongs to. | 

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


## FetchBinding

> ChatV2ServiceBinding FetchBinding(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Binding resource from. | 
**sid** | **string**| The SID of the Binding resource to fetch. | 

### Return type

[**ChatV2ServiceBinding**](chat.v2.service.binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannel

> ChatV2ServiceChannel FetchChannel(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Channel resource from. | 
**sid** | **string**| The SID of the Channel resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to fetch. | 

### Return type

[**ChatV2ServiceChannel**](chat.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelWebhook

> ChatV2ServiceChannelChannelWebhook FetchChannelWebhook(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to fetch the Webhook resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Channel Webhook resource to fetch. | 

### Return type

[**ChatV2ServiceChannelChannelWebhook**](chat.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ChatV2Credential FetchCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Credential resource to fetch. | 

### Return type

[**ChatV2Credential**](chat.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> ChatV2ServiceChannelInvite FetchInvite(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Invite resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Invite resource to fetch. | 

### Return type

[**ChatV2ServiceChannelInvite**](chat.v2.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ChatV2ServiceChannelMember FetchMember(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Member resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Member resource to fetch. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value. | 

### Return type

[**ChatV2ServiceChannelMember**](chat.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ChatV2ServiceChannelMessage FetchMessage(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Message resource from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Message resource to fetch. | 

### Return type

[**ChatV2ServiceChannelMessage**](chat.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ChatV2ServiceRole FetchRole(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Role resource from. | 
**sid** | **string**| The SID of the Role resource to fetch. | 

### Return type

[**ChatV2ServiceRole**](chat.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ChatV2Service FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to fetch. | 

### Return type

[**ChatV2Service**](chat.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ChatV2ServiceUser FetchUser(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User resource from. | 
**sid** | **string**| The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch. | 

### Return type

[**ChatV2ServiceUser**](chat.v2.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBinding

> ChatV2ServiceUserUserBinding FetchUserBinding(ctx, serviceSid, userSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User Binding resource from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resource to fetch.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. | 
**sid** | **string**| The SID of the User Binding resource to fetch. | 

### Return type

[**ChatV2ServiceUserUserBinding**](chat.v2.service.user.user_binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannel

> ChatV2ServiceUserUserChannel FetchUserChannel(ctx, serviceSid, userSid, channelSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User Channel resource from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to fetch the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) that has the User Channel to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel to fetch. | 

### Return type

[**ChatV2ServiceUserUserChannel**](chat.v2.service.user.user_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> ChatV2ServiceBindingReadResponse ListBinding(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Binding resources from. | 
 **optional** | ***ListBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindingType** | [**optional.Interface of []string**](string.md)| The push technology used by the Binding resources to read.  Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. | 
 **identity** | [**optional.Interface of []string**](string.md)| The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceBindingReadResponse**](chat_v2_service_bindingReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ChatV2ServiceChannelReadResponse ListChannel(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Channel resources from. | 
 **optional** | ***ListChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of []string**](string.md)| The visibility of the Channels to read. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceChannelReadResponse**](chat_v2_service_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelWebhook

> ChatV2ServiceChannelChannelWebhookReadResponse ListChannelWebhook(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to read the resources from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***ListChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceChannelChannelWebhookReadResponse**](chat_v2_service_channel_channel_webhookReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> ChatV2CredentialReadResponse ListCredential(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2CredentialReadResponse**](chat_v2_credentialReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvite

> ChatV2ServiceChannelInviteReadResponse ListInvite(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Invite resources from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***ListInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInviteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | [**optional.Interface of []string**](string.md)| The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceChannelInviteReadResponse**](chat_v2_service_channel_inviteReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ChatV2ServiceChannelMemberReadResponse ListMember(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Member resources from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***ListMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | [**optional.Interface of []string**](string.md)| The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the Member resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceChannelMemberReadResponse**](chat_v2_service_channel_memberReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ChatV2ServiceChannelMessageReadResponse ListMessage(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Message resources from. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to read belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***ListMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**| The sort order of the returned messages. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;asc&#x60; as the default. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceChannelMessageReadResponse**](chat_v2_service_channel_messageReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> ChatV2ServiceRoleReadResponse ListRole(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Role resources from. | 
 **optional** | ***ListRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceRoleReadResponse**](chat_v2_service_roleReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ChatV2ServiceReadResponse ListService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceReadResponse**](chat_v2_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> ChatV2ServiceUserReadResponse ListUser(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User resources from. | 
 **optional** | ***ListUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceUserReadResponse**](chat_v2_service_userReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBinding

> ChatV2ServiceUserUserBindingReadResponse ListUserBinding(ctx, serviceSid, userSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User Binding resources from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resources to read.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. | 
 **optional** | ***ListUserBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bindingType** | [**optional.Interface of []string**](string.md)| The push technology used by the User Binding resources to read. Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceUserUserBindingReadResponse**](chat_v2_service_user_user_bindingReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> ChatV2ServiceUserUserChannelReadResponse ListUserChannel(ctx, serviceSid, userSid, optional)



List all Channels for a given User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User Channel resources from. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to read the User Channel resources from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource. | 
 **optional** | ***ListUserChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ChatV2ServiceUserUserChannelReadResponse**](chat_v2_service_user_user_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> ChatV2ServiceChannel UpdateChannel(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Channel resource in. | 
**sid** | **string**| The SID of the Channel resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to update. | 
 **optional** | ***UpdateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **createdBy** | **optional.String**| The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 256 characters long. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 256 characters or less in length and unique within the Service. | 

### Return type

[**ChatV2ServiceChannel**](chat.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> ChatV2ServiceChannelChannelWebhook UpdateChannelWebhook(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel that has the Webhook resource to update. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Channel Webhook resource to update. | 
 **optional** | ***UpdateChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configurationFilters** | [**optional.Interface of []string**](string.md)| The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger). | 
 **configurationFlowSid** | **optional.String**| The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; &#x3D; &#x60;studio&#x60;. | 
 **configurationMethod** | **optional.String**| The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **configurationRetryCount** | **optional.Int32**| The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0. | 
 **configurationTriggers** | [**optional.Interface of []string**](string.md)| A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;. | 
 **configurationUrl** | **optional.String**| The URL of the webhook to call using the &#x60;configuration.method&#x60;. | 

### Return type

[**ChatV2ServiceChannelChannelWebhook**](chat.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ChatV2Credential UpdateCredential(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Credential resource to update. | 
 **optional** | ***UpdateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. | 
 **certificate** | **optional.String**| [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60; | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **privateKey** | **optional.String**| [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60; | 
 **sandbox** | **optional.Bool**| [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
 **secret** | **optional.String**| [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. | 

### Return type

[**ChatV2Credential**](chat.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ChatV2ServiceChannelMember UpdateMember(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Member resource in. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Member resource to update. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value. | 
 **optional** | ***UpdateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | 
 **lastConsumedMessageIndex** | **optional.Int32**| The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) that the Member has read within the [Channel](https://www.twilio.com/docs/chat/channels). | 
 **lastConsumptionTimestamp** | **optional.Time**| The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels). | 
 **roleSid** | **optional.String**| The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource). | 

### Return type

[**ChatV2ServiceChannelMember**](chat.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ChatV2ServiceChannelMessage UpdateMessage(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Message resource in. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**sid** | **string**| The SID of the Message resource to update. | 
 **optional** | ***UpdateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **body** | **optional.String**| The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. | 
 **dateCreated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source. | 
 **dateUpdated** | **optional.Time**| The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | 
 **from** | **optional.String**| The [Identity](https://www.twilio.com/docs/chat/identity) of the message&#39;s author. | 
 **lastUpdatedBy** | **optional.String**| The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable. | 

### Return type

[**ChatV2ServiceChannelMessage**](chat.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ChatV2ServiceRole UpdateRole(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Role resource in. | 
**sid** | **string**| The SID of the Role resource to update. | 
 **optional** | ***UpdateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permission** | [**optional.Interface of []string**](string.md)| A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;. | 

### Return type

[**ChatV2ServiceRole**](chat.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ChatV2Service UpdateService(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to update. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReportInterval** | **optional.Int32**| DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints. | 
 **defaultChannelCreatorRoleSid** | **optional.String**| The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | 
 **defaultChannelRoleSid** | **optional.String**| The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | 
 **defaultServiceRoleSid** | **optional.String**| The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. | 
 **limitsChannelMembers** | **optional.Int32**| The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000. | 
 **limitsUserChannels** | **optional.Int32**| The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000. | 
 **mediaCompatibilityMessage** | **optional.String**| The message to send when a media message has no text. Can be used as placeholder message. | 
 **notificationsAddedToChannelEnabled** | **optional.Bool**| Whether to send a notification when a member is added to a channel. The default is &#x60;false&#x60;. | 
 **notificationsAddedToChannelSound** | **optional.String**| The name of the sound to play when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsAddedToChannelTemplate** | **optional.String**| The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsInvitedToChannelEnabled** | **optional.Bool**| Whether to send a notification when a user is invited to a channel. The default is &#x60;false&#x60;. | 
 **notificationsInvitedToChannelSound** | **optional.String**| The name of the sound to play when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsInvitedToChannelTemplate** | **optional.String**| The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsLogEnabled** | **optional.Bool**| Whether to log notifications. The default is &#x60;false&#x60;. | 
 **notificationsNewMessageBadgeCountEnabled** | **optional.Bool**| Whether the new message badge is enabled. The default is &#x60;false&#x60;. | 
 **notificationsNewMessageEnabled** | **optional.Bool**| Whether to send a notification when a new message is added to a channel. The default is &#x60;false&#x60;. | 
 **notificationsNewMessageSound** | **optional.String**| The name of the sound to play when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsNewMessageTemplate** | **optional.String**| The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsRemovedFromChannelEnabled** | **optional.Bool**| Whether to send a notification to a user when they are removed from a channel. The default is &#x60;false&#x60;. | 
 **notificationsRemovedFromChannelSound** | **optional.String**| The name of the sound to play to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **notificationsRemovedFromChannelTemplate** | **optional.String**| The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;. | 
 **postWebhookRetryCount** | **optional.Int32**| The number of times to retry a call to the &#x60;post_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won&#39;t be retried. | 
 **postWebhookUrl** | **optional.String**| The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 
 **preWebhookRetryCount** | **optional.Int32**| The number of times to retry a call to the &#x60;pre_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won&#39;t be retried. | 
 **preWebhookUrl** | **optional.String**| The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 
 **reachabilityEnabled** | **optional.Bool**| Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;. | 
 **readStatusEnabled** | **optional.Bool**| Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;. | 
 **typingIndicatorTimeout** | **optional.Int32**| How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds. | 
 **webhookFilters** | [**optional.Interface of []string**](string.md)| The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 
 **webhookMethod** | **optional.String**| The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 

### Return type

[**ChatV2Service**](chat.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ChatV2ServiceUser UpdateUser(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the User resource in. | 
**sid** | **string**| The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update. | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**| A valid JSON string that contains application-specific data. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is often used for display purposes. | 
 **roleSid** | **optional.String**| The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the User. | 

### Return type

[**ChatV2ServiceUser**](chat.v2.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> ChatV2ServiceUserUserChannel UpdateUserChannel(ctx, serviceSid, userSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the User Channel resource in. | 
**userSid** | **string**| The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to update the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource. | 
**channelSid** | **string**| The SID of the [Channel](https://www.twilio.com/docs/chat/channels) with the User Channel resource to update. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
 **optional** | ***UpdateUserChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lastConsumedMessageIndex** | **optional.Int32**| The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. | 
 **lastConsumptionTimestamp** | **optional.Time**| The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels). | 
 **notificationLevel** | **optional.String**| The push notification level to assign to the User Channel. Can be: &#x60;default&#x60; or &#x60;muted&#x60;. | 

### Return type

[**ChatV2ServiceUserUserChannel**](chat.v2.service.user.user_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

