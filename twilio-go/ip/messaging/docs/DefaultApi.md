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

> IpMessagingV2ServiceChannel CreateChannel(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **createdBy** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **friendlyName** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannel**](ip_messaging.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook CreateChannelWebhook(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***CreateChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configurationFilters** | [**optional.Interface of []string**](string.md)|  | 
 **configurationFlowSid** | **optional.String**|  | 
 **configurationMethod** | **optional.String**|  | 
 **configurationRetryCount** | **optional.Int32**|  | 
 **configurationTriggers** | [**optional.Interface of []string**](string.md)|  | 
 **configurationUrl** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](ip_messaging.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> IpMessagingV2Credential CreateCredential(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **optional.String**|  | 
 **certificate** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **privateKey** | **optional.String**|  | 
 **sandbox** | **optional.Bool**|  | 
 **secret** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**IpMessagingV2Credential**](ip_messaging.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> IpMessagingV2ServiceChannelInvite CreateInvite(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***CreateInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInviteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelInvite**](ip_messaging.v2.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> IpMessagingV2ServiceChannelMember CreateMember(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***CreateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **identity** | **optional.String**|  | 
 **lastConsumedMessageIndex** | **optional.Int32**|  | 
 **lastConsumptionTimestamp** | **optional.Time**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelMember**](ip_messaging.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> IpMessagingV2ServiceChannelMessage CreateMessage(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***CreateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **body** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **from** | **optional.String**|  | 
 **lastUpdatedBy** | **optional.String**|  | 
 **mediaSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](ip_messaging.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> IpMessagingV2ServiceRole CreateRole(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**|  | 
 **permission** | [**optional.Interface of []string**](string.md)|  | 
 **type_** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceRole**](ip_messaging.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> IpMessagingV2Service CreateService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**|  | 

### Return type

[**IpMessagingV2Service**](ip_messaging.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> IpMessagingV2ServiceUser CreateUser(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **identity** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceUser**](ip_messaging.v2.service.user.md)

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
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
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
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

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
**sid** | **string**|  | 

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
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

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
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 
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
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 
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
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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
**sid** | **string**|  | 

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
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
**sid** | **string**|  | 

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
**channelSid** | **string**|  | 

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

> IpMessagingV2ServiceBinding FetchBinding(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceBinding**](ip_messaging.v2.service.binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannel

> IpMessagingV2ServiceChannel FetchChannel(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceChannel**](ip_messaging.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook FetchChannelWebhook(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](ip_messaging.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> IpMessagingV2Credential FetchCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

### Return type

[**IpMessagingV2Credential**](ip_messaging.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> IpMessagingV2ServiceChannelInvite FetchInvite(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceChannelInvite**](ip_messaging.v2.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> IpMessagingV2ServiceChannelMember FetchMember(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceChannelMember**](ip_messaging.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> IpMessagingV2ServiceChannelMessage FetchMessage(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](ip_messaging.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> IpMessagingV2ServiceRole FetchRole(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceRole**](ip_messaging.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> IpMessagingV2Service FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

### Return type

[**IpMessagingV2Service**](ip_messaging.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> IpMessagingV2ServiceUser FetchUser(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceUser**](ip_messaging.v2.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBinding

> IpMessagingV2ServiceUserUserBinding FetchUserBinding(ctx, serviceSid, userSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceUserUserBinding**](ip_messaging.v2.service.user.user_binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannel

> IpMessagingV2ServiceUserUserChannel FetchUserChannel(ctx, serviceSid, userSid, channelSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
**channelSid** | **string**|  | 

### Return type

[**IpMessagingV2ServiceUserUserChannel**](ip_messaging.v2.service.user.user_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> IpMessagingV2ServiceBindingReadResponse ListBinding(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindingType** | [**optional.Interface of []string**](string.md)|  | 
 **identity** | [**optional.Interface of []string**](string.md)|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceBindingReadResponse**](ip_messaging_v2_service_bindingReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> IpMessagingV2ServiceChannelReadResponse ListChannel(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceChannelReadResponse**](ip_messaging_v2_service_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhookReadResponse ListChannelWebhook(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***ListChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhookReadResponse**](ip_messaging_v2_service_channel_channel_webhookReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> IpMessagingV2CredentialReadResponse ListCredential(ctx, optional)



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

[**IpMessagingV2CredentialReadResponse**](ip_messaging_v2_credentialReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvite

> IpMessagingV2ServiceChannelInviteReadResponse ListInvite(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***ListInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInviteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | [**optional.Interface of []string**](string.md)|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceChannelInviteReadResponse**](ip_messaging_v2_service_channel_inviteReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> IpMessagingV2ServiceChannelMemberReadResponse ListMember(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***ListMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identity** | [**optional.Interface of []string**](string.md)|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceChannelMemberReadResponse**](ip_messaging_v2_service_channel_memberReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> IpMessagingV2ServiceChannelMessageReadResponse ListMessage(ctx, serviceSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***ListMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceChannelMessageReadResponse**](ip_messaging_v2_service_channel_messageReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> IpMessagingV2ServiceRoleReadResponse ListRole(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceRoleReadResponse**](ip_messaging_v2_service_roleReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> IpMessagingV2ServiceReadResponse ListService(ctx, optional)



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

[**IpMessagingV2ServiceReadResponse**](ip_messaging_v2_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> IpMessagingV2ServiceUserReadResponse ListUser(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceUserReadResponse**](ip_messaging_v2_service_userReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBinding

> IpMessagingV2ServiceUserUserBindingReadResponse ListUserBinding(ctx, serviceSid, userSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
 **optional** | ***ListUserBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bindingType** | [**optional.Interface of []string**](string.md)|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceUserUserBindingReadResponse**](ip_messaging_v2_service_user_user_bindingReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> IpMessagingV2ServiceUserUserChannelReadResponse ListUserChannel(ctx, serviceSid, userSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
 **optional** | ***ListUserChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**IpMessagingV2ServiceUserUserChannelReadResponse**](ip_messaging_v2_service_user_user_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> IpMessagingV2ServiceChannel UpdateChannel(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **createdBy** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **friendlyName** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannel**](ip_messaging.v2.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook UpdateChannelWebhook(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateChannelWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChannelWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **configurationFilters** | [**optional.Interface of []string**](string.md)|  | 
 **configurationFlowSid** | **optional.String**|  | 
 **configurationMethod** | **optional.String**|  | 
 **configurationRetryCount** | **optional.Int32**|  | 
 **configurationTriggers** | [**optional.Interface of []string**](string.md)|  | 
 **configurationUrl** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](ip_messaging.v2.service.channel.channel_webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> IpMessagingV2Credential UpdateCredential(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**|  | 
 **certificate** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **privateKey** | **optional.String**|  | 
 **sandbox** | **optional.Bool**|  | 
 **secret** | **optional.String**|  | 

### Return type

[**IpMessagingV2Credential**](ip_messaging.v2.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> IpMessagingV2ServiceChannelMember UpdateMember(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **lastConsumedMessageIndex** | **optional.Int32**|  | 
 **lastConsumptionTimestamp** | **optional.Time**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelMember**](ip_messaging.v2.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> IpMessagingV2ServiceChannelMessage UpdateMessage(ctx, serviceSid, channelSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **body** | **optional.String**|  | 
 **dateCreated** | **optional.Time**|  | 
 **dateUpdated** | **optional.Time**|  | 
 **from** | **optional.String**|  | 
 **lastUpdatedBy** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](ip_messaging.v2.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> IpMessagingV2ServiceRole UpdateRole(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permission** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**IpMessagingV2ServiceRole**](ip_messaging.v2.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV2Service UpdateService(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumptionReportInterval** | **optional.Int32**|  | 
 **defaultChannelCreatorRoleSid** | **optional.String**|  | 
 **defaultChannelRoleSid** | **optional.String**|  | 
 **defaultServiceRoleSid** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **limitsChannelMembers** | **optional.Int32**|  | 
 **limitsUserChannels** | **optional.Int32**|  | 
 **mediaCompatibilityMessage** | **optional.String**|  | 
 **notificationsAddedToChannelEnabled** | **optional.Bool**|  | 
 **notificationsAddedToChannelSound** | **optional.String**|  | 
 **notificationsAddedToChannelTemplate** | **optional.String**|  | 
 **notificationsInvitedToChannelEnabled** | **optional.Bool**|  | 
 **notificationsInvitedToChannelSound** | **optional.String**|  | 
 **notificationsInvitedToChannelTemplate** | **optional.String**|  | 
 **notificationsLogEnabled** | **optional.Bool**|  | 
 **notificationsNewMessageBadgeCountEnabled** | **optional.Bool**|  | 
 **notificationsNewMessageEnabled** | **optional.Bool**|  | 
 **notificationsNewMessageSound** | **optional.String**|  | 
 **notificationsNewMessageTemplate** | **optional.String**|  | 
 **notificationsRemovedFromChannelEnabled** | **optional.Bool**|  | 
 **notificationsRemovedFromChannelSound** | **optional.String**|  | 
 **notificationsRemovedFromChannelTemplate** | **optional.String**|  | 
 **postWebhookRetryCount** | **optional.Int32**|  | 
 **postWebhookUrl** | **optional.String**|  | 
 **preWebhookRetryCount** | **optional.Int32**|  | 
 **preWebhookUrl** | **optional.String**|  | 
 **reachabilityEnabled** | **optional.Bool**|  | 
 **readStatusEnabled** | **optional.Bool**|  | 
 **typingIndicatorTimeout** | **optional.Int32**|  | 
 **webhookFilters** | [**optional.Interface of []string**](string.md)|  | 
 **webhookMethod** | **optional.String**|  | 

### Return type

[**IpMessagingV2Service**](ip_messaging.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IpMessagingV2ServiceUser UpdateUser(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTwilioWebhookEnabled** | **optional.String**| The X-Twilio-Webhook-Enabled HTTP request header | 
 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceUser**](ip_messaging.v2.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> IpMessagingV2ServiceUserUserChannel UpdateUserChannel(ctx, serviceSid, userSid, channelSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**userSid** | **string**|  | 
**channelSid** | **string**|  | 
 **optional** | ***UpdateUserChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lastConsumedMessageIndex** | **optional.Int32**|  | 
 **lastConsumptionTimestamp** | **optional.Time**|  | 
 **notificationLevel** | **optional.String**|  | 

### Return type

[**IpMessagingV2ServiceUserUserChannel**](ip_messaging.v2.service.user.user_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

