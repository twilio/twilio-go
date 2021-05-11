# DefaultApi

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

> ChatV2ServiceChannel CreateChannel(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Channel resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**CreatedBy** | **string** | The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used in cases where a Channel is being recreated from a backup/separate source  and where a Message was previously updated.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**Type** | **string** | The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the Channel resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service.

### Return type

[**ChatV2ServiceChannel**](ChatV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelWebhook

> ChatV2ServiceChannelChannelWebhook CreateChannelWebhook(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to create the Webhook resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Channel Webhook resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
**ConfigurationFlowSid** | **string** | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; is &#x60;studio&#x60;.
**ConfigurationMethod** | **string** | The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**ConfigurationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.
**ConfigurationTriggers** | **[]string** | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;.
**ConfigurationUrl** | **string** | The URL of the webhook to call using the &#x60;configuration.method&#x60;.
**Type** | **string** | The type of webhook. Can be: &#x60;webhook&#x60;, &#x60;studio&#x60;, or &#x60;trigger&#x60;.

### Return type

[**ChatV2ServiceChannelChannelWebhook**](ChatV2ServiceChannelChannelWebhook.md)

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
**Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ChatV2Credential**](ChatV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> ChatV2ServiceChannelInvite CreateInvite(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Invite resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Invite resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) assigned to the new member.

### Return type

[**ChatV2ServiceChannelInvite**](ChatV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> ChatV2ServiceChannelMember CreateMember(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Member resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Member resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is &#x60;null&#x60;. Note that this parameter should only be used when a Member is being recreated from a backup/separate source and where a Member was previously updated.
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**LastConsumedMessageIndex** | **int32** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. This parameter should only be used when recreating a Member from a backup/separate source.
**LastConsumptionTimestamp** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource).

### Return type

[**ChatV2ServiceChannelMember**](ChatV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ChatV2ServiceChannelMessage CreateMessage(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Message resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the new Message resource belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**Body** | **string** | The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
**From** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;.
**LastUpdatedBy** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
**MediaSid** | **string** | The SID of the [Media](https://www.twilio.com/docs/chat/rest/media) to attach to the new Message.

### Return type

[**ChatV2ServiceChannelMessage**](ChatV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ChatV2ServiceRole CreateRole(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the Role resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
**Type** | **string** | The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/rest/service-resource) roles.

### Return type

[**ChatV2ServiceRole**](ChatV2ServiceRole.md)

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource.

### Return type

[**ChatV2Service**](ChatV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ChatV2ServiceUser CreateUser(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the User resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. This value is often used for display purposes.
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). This value is often a username or email address. See the Identity documentation for more info.
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the new User.

### Return type

[**ChatV2ServiceUser**](ChatV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinding

> DeleteBinding(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Binding resource from.
**Sid** | **string** | The SID of the Binding resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBindingParams struct


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


## DeleteChannel

> DeleteChannel(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the Channel resource to delete.  This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct


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


## DeleteChannelWebhook

> DeleteChannelWebhook(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to delete the Webhook resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Channel Webhook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelWebhookParams struct


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



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Credential resource to delete.

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


## DeleteInvite

> DeleteInvite(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Invite resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Invite resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInviteParams struct


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


## DeleteMember

> DeleteMember(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Member resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Member resource to delete. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMemberParams struct


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


## DeleteMessage

> DeleteMessage(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Message resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to delete belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Message resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct


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


## DeleteRole

> DeleteRole(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the Role resource from.
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



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to delete.

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


## DeleteUser

> DeleteUser(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the User resource from.
**Sid** | **string** | The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct


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


## DeleteUserBinding

> DeleteUserBinding(ctx, ServiceSidUserSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the User Binding resource from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resources to delete.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
**Sid** | **string** | The SID of the User Binding resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserBindingParams struct


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


## DeleteUserChannel

> DeleteUserChannel(ctx, ServiceSidUserSidChannelSid)



Removes User from selected Channel.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource belongs to.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserChannelParams struct


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


## FetchBinding

> ChatV2ServiceBinding FetchBinding(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Binding resource from.
**Sid** | **string** | The SID of the Binding resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceBinding**](ChatV2ServiceBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannel

> ChatV2ServiceChannel FetchChannel(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Channel resource from.
**Sid** | **string** | The SID of the Channel resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceChannel**](ChatV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelWebhook

> ChatV2ServiceChannelChannelWebhook FetchChannelWebhook(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to fetch the Webhook resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Channel Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceChannelChannelWebhook**](ChatV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ChatV2Credential FetchCredential(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Credential resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2Credential**](ChatV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> ChatV2ServiceChannelInvite FetchInvite(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Invite resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Invite resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInviteParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceChannelInvite**](ChatV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ChatV2ServiceChannelMember FetchMember(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Member resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Member resource to fetch. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value.

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceChannelMember**](ChatV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ChatV2ServiceChannelMessage FetchMessage(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Message resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to fetch belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Message resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceChannelMessage**](ChatV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ChatV2ServiceRole FetchRole(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the Role resource from.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceRole**](ChatV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ChatV2Service FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2Service**](ChatV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ChatV2ServiceUser FetchUser(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User resource from.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceUser**](ChatV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBinding

> ChatV2ServiceUserUserBinding FetchUserBinding(ctx, ServiceSidUserSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User Binding resource from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resource to fetch.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
**Sid** | **string** | The SID of the User Binding resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceUserUserBinding**](ChatV2ServiceUserUserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannel

> ChatV2ServiceUserUserChannel FetchUserChannel(ctx, ServiceSidUserSidChannelSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the User Channel resource from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to fetch the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) that has the User Channel to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV2ServiceUserUserChannel**](ChatV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> ListBindingResponse ListBinding(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Binding resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **[]string** | The push technology used by the Binding resources to read.  Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
**Identity** | **[]string** | The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBindingResponse**](ListBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ListChannelResponse ListChannel(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Channel resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **[]string** | The visibility of the Channels to read. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListChannelResponse**](ListChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelWebhook

> ListChannelWebhookResponse ListChannelWebhook(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel to read the resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListChannelWebhookResponse**](ListChannelWebhookResponse.md)

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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


## ListInvite

> ListInviteResponse ListInvite(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Invite resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **[]string** | The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListInviteResponse**](ListInviteResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ListMemberResponse ListMember(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Member resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resources to read belong to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **[]string** | The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the Member resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ListMessageResponse ListMessage(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Message resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to read belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the returned messages. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;asc&#x60; as the default.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessageResponse**](ListMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> ListRoleResponse ListRole(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the Role resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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


## ListUser

> ListUserResponse ListUser(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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


## ListUserBinding

> ListUserBindingResponse ListUserBinding(ctx, ServiceSidUserSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User Binding resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) with the User Binding resources to read.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.

### Other Parameters

Other parameters are passed through a pointer to a ListUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **[]string** | The push technology used by the User Binding resources to read. Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserBindingResponse**](ListUserBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> ListUserChannelResponse ListUserChannel(ctx, ServiceSidUserSidoptional)



List all Channels for a given User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the User Channel resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to read the User Channel resources from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserChannelResponse**](ListUserChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> ChatV2ServiceChannel UpdateChannel(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Channel resource in.
**Sid** | **string** | The SID of the Channel resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;unique_name&#x60; of the Channel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**CreatedBy** | **string** | The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 256 characters long.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 256 characters or less in length and unique within the Service.

### Return type

[**ChatV2ServiceChannel**](ChatV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> ChatV2ServiceChannelChannelWebhook UpdateChannelWebhook(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) with the Channel that has the Webhook resource to update.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Channel Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | The events that cause us to call the Channel Webhook. Used when &#x60;type&#x60; is &#x60;webhook&#x60;. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
**ConfigurationFlowSid** | **string** | The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in &#x60;configuration.filters&#x60; occurs. Used only when &#x60;type&#x60; &#x3D; &#x60;studio&#x60;.
**ConfigurationMethod** | **string** | The HTTP method used to call &#x60;configuration.url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
**ConfigurationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.
**ConfigurationTriggers** | **[]string** | A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when &#x60;type&#x60; &#x3D; &#x60;trigger&#x60;.
**ConfigurationUrl** | **string** | The URL of the webhook to call using the &#x60;configuration.method&#x60;.

### Return type

[**ChatV2ServiceChannelChannelWebhook**](ChatV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ChatV2Credential UpdateCredential(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Credential resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.

### Return type

[**ChatV2Credential**](ChatV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ChatV2ServiceChannelMember UpdateMember(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Member resource in.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Member resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Member resource to update. This value can be either the Member&#39;s &#x60;sid&#x60; or its &#x60;identity&#x60; value.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
**LastConsumedMessageIndex** | **int32** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) that the Member has read within the [Channel](https://www.twilio.com/docs/chat/channels).
**LastConsumptionTimestamp** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource).

### Return type

[**ChatV2ServiceChannelMember**](ChatV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ChatV2ServiceChannelMessage UpdateMessage(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Message resource in.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource to update belongs to. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The SID of the Message resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**Body** | **string** | The message to send to the channel. Can be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
**DateCreated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat&#39;s history is being recreated from a backup/separate source.
**DateUpdated** | **time.Time** | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
**From** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the message&#39;s author.
**LastUpdatedBy** | **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.

### Return type

[**ChatV2ServiceChannelMessage**](ChatV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ChatV2ServiceRole UpdateRole(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the Role resource in.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.

### Return type

[**ChatV2ServiceRole**](ChatV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ChatV2Service UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConsumptionReportInterval** | **int32** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource.
**LimitsChannelMembers** | **int32** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
**LimitsUserChannels** | **int32** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
**MediaCompatibilityMessage** | **string** | The message to send when a media message has no text. Can be used as placeholder message.
**NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. The default is &#x60;false&#x60;.
**NotificationsAddedToChannelSound** | **string** | The name of the sound to play when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. The default is &#x60;false&#x60;.
**NotificationsInvitedToChannelSound** | **string** | The name of the sound to play when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsLogEnabled** | **bool** | Whether to log notifications. The default is &#x60;false&#x60;.
**NotificationsNewMessageBadgeCountEnabled** | **bool** | Whether the new message badge is enabled. The default is &#x60;false&#x60;.
**NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. The default is &#x60;false&#x60;.
**NotificationsNewMessageSound** | **string** | The name of the sound to play when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
**NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. The default is &#x60;false&#x60;.
**NotificationsRemovedFromChannelSound** | **string** | The name of the sound to play to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
**PostWebhookRetryCount** | **int32** | The number of times to retry a call to the &#x60;post_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won&#39;t be retried.
**PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**PreWebhookRetryCount** | **int32** | The number of times to retry a call to the &#x60;pre_webhook_url&#x60; if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won&#39;t be retried.
**PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;.
**ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;.
**TypingIndicatorTimeout** | **int32** | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds.
**WebhookFilters** | **[]string** | The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
**WebhookMethod** | **string** | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.

### Return type

[**ChatV2Service**](ChatV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ChatV2ServiceUser UpdateUser(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the User resource in.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is often used for display purposes.
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the User.

### Return type

[**ChatV2ServiceUser**](ChatV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> ChatV2ServiceUserUserChannel UpdateUserChannel(ctx, ServiceSidUserSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to update the User Channel resource in.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) to update the User Channel resource from. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) with the User Channel resource to update. This value can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**LastConsumedMessageIndex** | **int32** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read.
**LastConsumptionTimestamp** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
**NotificationLevel** | **string** | The push notification level to assign to the User Channel. Can be: &#x60;default&#x60; or &#x60;muted&#x60;.

### Return type

[**ChatV2ServiceUserUserChannel**](ChatV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

