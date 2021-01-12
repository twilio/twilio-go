# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v1/Services/{ServiceSid}/Channels | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v1/Credentials | 
[**CreateInvite**](DefaultApi.md#CreateInvite) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**CreateMember**](DefaultApi.md#CreateMember) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v1/Services/{ServiceSid}/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v1/Services/{ServiceSid}/Users | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v1/Credentials/{Sid} | 
[**DeleteInvite**](DefaultApi.md#DeleteInvite) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**DeleteMember**](DefaultApi.md#DeleteMember) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v1/Credentials/{Sid} | 
[**FetchInvite**](DefaultApi.md#FetchInvite) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v1/Services/{ServiceSid}/Channels | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v1/Credentials | 
[**ListInvite**](DefaultApi.md#ListInvite) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v1/Services/{ServiceSid}/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v1/Services/{ServiceSid}/Users | 
[**ListUserChannel**](DefaultApi.md#ListUserChannel) | **Get** /v1/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateChannel**](DefaultApi.md#UpdateChannel) | **Post** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v1/Credentials/{Sid} | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v1/Services/{ServiceSid}/Users/{Sid} | 



## CreateChannel

> IpMessagingV1ServiceChannel CreateChannel(ctx, serviceSid, optional)



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

 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannel**](ip_messaging.v1.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> IpMessagingV1Credential CreateCredential(ctx, optional)



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

[**IpMessagingV1Credential**](ip_messaging.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> IpMessagingV1ServiceChannelInvite CreateInvite(ctx, serviceSid, channelSid, optional)



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

[**IpMessagingV1ServiceChannelInvite**](ip_messaging.v1.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> IpMessagingV1ServiceChannelMember CreateMember(ctx, serviceSid, channelSid, optional)



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


 **identity** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannelMember**](ip_messaging.v1.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> IpMessagingV1ServiceChannelMessage CreateMessage(ctx, serviceSid, channelSid, optional)



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


 **attributes** | **optional.String**|  | 
 **body** | **optional.String**|  | 
 **from** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannelMessage**](ip_messaging.v1.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> IpMessagingV1ServiceRole CreateRole(ctx, serviceSid, optional)



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

[**IpMessagingV1ServiceRole**](ip_messaging.v1.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> IpMessagingV1Service CreateService(ctx, optional)



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

[**IpMessagingV1Service**](ip_messaging.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> IpMessagingV1ServiceUser CreateUser(ctx, serviceSid, optional)



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

 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **identity** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceUser**](ip_messaging.v1.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, serviceSid, sid)



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

> DeleteMember(ctx, serviceSid, channelSid, sid)



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


## DeleteMessage

> DeleteMessage(ctx, serviceSid, channelSid, sid)



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


## FetchChannel

> IpMessagingV1ServiceChannel FetchChannel(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceChannel**](ip_messaging.v1.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> IpMessagingV1Credential FetchCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

### Return type

[**IpMessagingV1Credential**](ip_messaging.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> IpMessagingV1ServiceChannelInvite FetchInvite(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceChannelInvite**](ip_messaging.v1.service.channel.invite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> IpMessagingV1ServiceChannelMember FetchMember(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceChannelMember**](ip_messaging.v1.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> IpMessagingV1ServiceChannelMessage FetchMessage(ctx, serviceSid, channelSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**channelSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceChannelMessage**](ip_messaging.v1.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> IpMessagingV1ServiceRole FetchRole(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceRole**](ip_messaging.v1.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> IpMessagingV1Service FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

### Return type

[**IpMessagingV1Service**](ip_messaging.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> IpMessagingV1ServiceUser FetchUser(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

### Return type

[**IpMessagingV1ServiceUser**](ip_messaging.v1.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> InlineResponse2002 ListChannel(ctx, serviceSid, optional)



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

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> InlineResponse200 ListCredential(ctx, optional)



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

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvite

> InlineResponse2003 ListInvite(ctx, serviceSid, channelSid, optional)



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

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> InlineResponse2004 ListMember(ctx, serviceSid, channelSid, optional)



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

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> InlineResponse2005 ListMessage(ctx, serviceSid, channelSid, optional)



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

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> InlineResponse2006 ListRole(ctx, serviceSid, optional)



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

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> InlineResponse2001 ListService(ctx, optional)



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

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> InlineResponse2007 ListUser(ctx, serviceSid, optional)



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

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> InlineResponse2008 ListUserChannel(ctx, serviceSid, userSid, optional)



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

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> IpMessagingV1ServiceChannel UpdateChannel(ctx, serviceSid, sid, optional)



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


 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannel**](ip_messaging.v1.service.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> IpMessagingV1Credential UpdateCredential(ctx, sid, optional)



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

[**IpMessagingV1Credential**](ip_messaging.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> IpMessagingV1ServiceChannelMember UpdateMember(ctx, serviceSid, channelSid, sid, optional)



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



 **lastConsumedMessageIndex** | **optional.Int32**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannelMember**](ip_messaging.v1.service.channel.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> IpMessagingV1ServiceChannelMessage UpdateMessage(ctx, serviceSid, channelSid, sid, optional)



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



 **attributes** | **optional.String**|  | 
 **body** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceChannelMessage**](ip_messaging.v1.service.channel.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> IpMessagingV1ServiceRole UpdateRole(ctx, serviceSid, sid, optional)



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

[**IpMessagingV1ServiceRole**](ip_messaging.v1.service.role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV1Service UpdateService(ctx, sid, optional)



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
 **notificationsAddedToChannelEnabled** | **optional.Bool**|  | 
 **notificationsAddedToChannelTemplate** | **optional.String**|  | 
 **notificationsInvitedToChannelEnabled** | **optional.Bool**|  | 
 **notificationsInvitedToChannelTemplate** | **optional.String**|  | 
 **notificationsNewMessageEnabled** | **optional.Bool**|  | 
 **notificationsNewMessageTemplate** | **optional.String**|  | 
 **notificationsRemovedFromChannelEnabled** | **optional.Bool**|  | 
 **notificationsRemovedFromChannelTemplate** | **optional.String**|  | 
 **postWebhookUrl** | **optional.String**|  | 
 **preWebhookUrl** | **optional.String**|  | 
 **reachabilityEnabled** | **optional.Bool**|  | 
 **readStatusEnabled** | **optional.Bool**|  | 
 **typingIndicatorTimeout** | **optional.Int32**|  | 
 **webhookFilters** | [**optional.Interface of []string**](string.md)|  | 
 **webhookMethod** | **optional.String**|  | 
 **webhooksOnChannelAddMethod** | **optional.String**|  | 
 **webhooksOnChannelAddUrl** | **optional.String**|  | 
 **webhooksOnChannelAddedMethod** | **optional.String**|  | 
 **webhooksOnChannelAddedUrl** | **optional.String**|  | 
 **webhooksOnChannelDestroyMethod** | **optional.String**|  | 
 **webhooksOnChannelDestroyUrl** | **optional.String**|  | 
 **webhooksOnChannelDestroyedMethod** | **optional.String**|  | 
 **webhooksOnChannelDestroyedUrl** | **optional.String**|  | 
 **webhooksOnChannelUpdateMethod** | **optional.String**|  | 
 **webhooksOnChannelUpdateUrl** | **optional.String**|  | 
 **webhooksOnChannelUpdatedMethod** | **optional.String**|  | 
 **webhooksOnChannelUpdatedUrl** | **optional.String**|  | 
 **webhooksOnMemberAddMethod** | **optional.String**|  | 
 **webhooksOnMemberAddUrl** | **optional.String**|  | 
 **webhooksOnMemberAddedMethod** | **optional.String**|  | 
 **webhooksOnMemberAddedUrl** | **optional.String**|  | 
 **webhooksOnMemberRemoveMethod** | **optional.String**|  | 
 **webhooksOnMemberRemoveUrl** | **optional.String**|  | 
 **webhooksOnMemberRemovedMethod** | **optional.String**|  | 
 **webhooksOnMemberRemovedUrl** | **optional.String**|  | 
 **webhooksOnMessageRemoveMethod** | **optional.String**|  | 
 **webhooksOnMessageRemoveUrl** | **optional.String**|  | 
 **webhooksOnMessageRemovedMethod** | **optional.String**|  | 
 **webhooksOnMessageRemovedUrl** | **optional.String**|  | 
 **webhooksOnMessageSendMethod** | **optional.String**|  | 
 **webhooksOnMessageSendUrl** | **optional.String**|  | 
 **webhooksOnMessageSentMethod** | **optional.String**|  | 
 **webhooksOnMessageSentUrl** | **optional.String**|  | 
 **webhooksOnMessageUpdateMethod** | **optional.String**|  | 
 **webhooksOnMessageUpdateUrl** | **optional.String**|  | 
 **webhooksOnMessageUpdatedMethod** | **optional.String**|  | 
 **webhooksOnMessageUpdatedUrl** | **optional.String**|  | 

### Return type

[**IpMessagingV1Service**](ip_messaging.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IpMessagingV1ServiceUser UpdateUser(ctx, serviceSid, sid, optional)



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


 **attributes** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **roleSid** | **optional.String**|  | 

### Return type

[**IpMessagingV1ServiceUser**](ip_messaging.v1.service.user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

