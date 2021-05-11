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

> IpMessagingV2ServiceChannel CreateChannel(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**CreatedBy** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**FriendlyName** | **string** | 
**Type** | **string** | 
**UniqueName** | **string** | 

### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook CreateChannelWebhook(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | 
**ConfigurationFlowSid** | **string** | 
**ConfigurationMethod** | **string** | 
**ConfigurationRetryCount** | **int32** | 
**ConfigurationTriggers** | **[]string** | 
**ConfigurationUrl** | **string** | 
**Type** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | 
**Certificate** | **string** | 
**FriendlyName** | **string** | 
**PrivateKey** | **string** | 
**Sandbox** | **bool** | 
**Secret** | **string** | 
**Type** | **string** | 

### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> IpMessagingV2ServiceChannelInvite CreateInvite(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | 
**RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelInvite**](IpMessagingV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> IpMessagingV2ServiceChannelMember CreateMember(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**Identity** | **string** | 
**LastConsumedMessageIndex** | **int32** | 
**LastConsumptionTimestamp** | **time.Time** | 
**RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> IpMessagingV2ServiceChannelMessage CreateMessage(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**Body** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**From** | **string** | 
**LastUpdatedBy** | **string** | 
**MediaSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> IpMessagingV2ServiceRole CreateRole(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | 
**Permission** | **[]string** | 
**Type** | **string** | 

### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> IpMessagingV2ServiceUser CreateUser(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**FriendlyName** | **string** | 
**Identity** | **string** | 
**RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

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
**ServiceSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

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
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**Sid** | **string** | 

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
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**Sid** | **string** | 

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
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**Sid** | **string** | 

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



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

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

> IpMessagingV2ServiceBinding FetchBinding(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceBinding**](IpMessagingV2ServiceBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannel

> IpMessagingV2ServiceChannel FetchChannel(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook FetchChannelWebhook(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> IpMessagingV2Credential FetchCredential(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> IpMessagingV2ServiceChannelInvite FetchInvite(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchInviteParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceChannelInvite**](IpMessagingV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> IpMessagingV2ServiceChannelMember FetchMember(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> IpMessagingV2ServiceChannelMessage FetchMessage(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> IpMessagingV2ServiceRole FetchRole(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> IpMessagingV2Service FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> IpMessagingV2ServiceUser FetchUser(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBinding

> IpMessagingV2ServiceUserUserBinding FetchUserBinding(ctx, ServiceSidUserSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceUserUserBinding**](IpMessagingV2ServiceUserUserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannel

> IpMessagingV2ServiceUserUserChannel FetchUserChannel(ctx, ServiceSidUserSidChannelSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2ServiceUserUserChannel**](IpMessagingV2ServiceUserUserChannel.md)

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
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **[]string** | 
**Identity** | **[]string** | 
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
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **[]string** | 
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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **[]string** | 
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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **[]string** | 
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
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | 
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
**ServiceSid** | **string** | 

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
**ServiceSid** | **string** | 

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
**ServiceSid** | **string** | 
**UserSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **[]string** | 
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



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 

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

> IpMessagingV2ServiceChannel UpdateChannel(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**CreatedBy** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**FriendlyName** | **string** | 
**UniqueName** | **string** | 

### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook UpdateChannelWebhook(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelWebhookParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConfigurationFilters** | **[]string** | 
**ConfigurationFlowSid** | **string** | 
**ConfigurationMethod** | **string** | 
**ConfigurationRetryCount** | **int32** | 
**ConfigurationTriggers** | **[]string** | 
**ConfigurationUrl** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> IpMessagingV2Credential UpdateCredential(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApiKey** | **string** | 
**Certificate** | **string** | 
**FriendlyName** | **string** | 
**PrivateKey** | **string** | 
**Sandbox** | **bool** | 
**Secret** | **string** | 

### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> IpMessagingV2ServiceChannelMember UpdateMember(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**LastConsumedMessageIndex** | **int32** | 
**LastConsumptionTimestamp** | **time.Time** | 
**RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> IpMessagingV2ServiceChannelMessage UpdateMessage(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**Body** | **string** | 
**DateCreated** | **time.Time** | 
**DateUpdated** | **time.Time** | 
**From** | **string** | 
**LastUpdatedBy** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> IpMessagingV2ServiceRole UpdateRole(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permission** | **[]string** | 

### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV2Service UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConsumptionReportInterval** | **int32** | 
**DefaultChannelCreatorRoleSid** | **string** | 
**DefaultChannelRoleSid** | **string** | 
**DefaultServiceRoleSid** | **string** | 
**FriendlyName** | **string** | 
**LimitsChannelMembers** | **int32** | 
**LimitsUserChannels** | **int32** | 
**MediaCompatibilityMessage** | **string** | 
**NotificationsAddedToChannelEnabled** | **bool** | 
**NotificationsAddedToChannelSound** | **string** | 
**NotificationsAddedToChannelTemplate** | **string** | 
**NotificationsInvitedToChannelEnabled** | **bool** | 
**NotificationsInvitedToChannelSound** | **string** | 
**NotificationsInvitedToChannelTemplate** | **string** | 
**NotificationsLogEnabled** | **bool** | 
**NotificationsNewMessageBadgeCountEnabled** | **bool** | 
**NotificationsNewMessageEnabled** | **bool** | 
**NotificationsNewMessageSound** | **string** | 
**NotificationsNewMessageTemplate** | **string** | 
**NotificationsRemovedFromChannelEnabled** | **bool** | 
**NotificationsRemovedFromChannelSound** | **string** | 
**NotificationsRemovedFromChannelTemplate** | **string** | 
**PostWebhookRetryCount** | **int32** | 
**PostWebhookUrl** | **string** | 
**PreWebhookRetryCount** | **int32** | 
**PreWebhookUrl** | **string** | 
**ReachabilityEnabled** | **bool** | 
**ReadStatusEnabled** | **bool** | 
**TypingIndicatorTimeout** | **int32** | 
**WebhookFilters** | **[]string** | 
**WebhookMethod** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IpMessagingV2ServiceUser UpdateUser(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Attributes** | **string** | 
**FriendlyName** | **string** | 
**RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> IpMessagingV2ServiceUserUserChannel UpdateUserChannel(ctx, ServiceSidUserSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**LastConsumedMessageIndex** | **int32** | 
**LastConsumptionTimestamp** | **time.Time** | 
**NotificationLevel** | **string** | 

### Return type

[**IpMessagingV2ServiceUserUserChannel**](IpMessagingV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

