# DefaultApi

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

> ChatV1ServiceChannel CreateChannel(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
**Type** | **string** | The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service. | 

### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> ChatV1Credential CreateCredential(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. | 
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60; | 
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----&#x60; | 
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. | 
**Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;. | 

### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> ChatV1ServiceChannelInvite CreateInvite(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateInviteParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info. | 
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member. | 

### Return type

[**ChatV1ServiceChannelInvite**](ChatV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> ChatV1ServiceChannelMember CreateMember(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new member belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMemberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. | 
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services). | 

### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ChatV1ServiceChannelMessage CreateMessage(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**Body** | **string** | The message to send to the channel. Can also be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. | 
**From** | **string** | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;. | 

### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ChatV1ServiceRole CreateRole(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. | 
**Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60; and are described in the documentation. | 
**Type** | **string** | The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/api/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/api/services) roles. | 

### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ChatV1Service CreateService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ChatV1ServiceUser CreateUser(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. This value is often used for display purposes. | 
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details. | 
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User. | 

### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to delete belongs to. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Invite resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteInviteParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteMember(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to.  Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteMemberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteMessage(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to delete belongs to.  Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> ChatV1ServiceChannel FetchChannel(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ChatV1Credential FetchCredential(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> ChatV1ServiceChannelInvite FetchInvite(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to fetch belongs to. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Invite resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchInviteParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceChannelInvite**](ChatV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ChatV1ServiceChannelMember FetchMember(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to fetch belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60; value. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ChatV1ServiceChannelMessage FetchMessage(ctx, ServiceSidChannelSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to fetch belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ChatV1ServiceRole FetchRole(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ChatV1Service FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ChatV1ServiceUser FetchUser(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ListChannelResponse ListChannel(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Type** | **[]string** | The visibility of the Channels to read. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListChannelResponse**](ListChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> ListCredentialResponse ListCredential(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCredentialResponse**](ListCredentialResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvite

> ListInviteResponse ListInvite(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resources to read belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a ListInviteParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Identity** | **[]string** | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListInviteResponse**](ListInviteResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ListMemberResponse ListMember(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the members to read belong to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60; value. | 

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Identity** | **[]string** | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ListMessageResponse ListMessage(ctx, ServiceSidChannelSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to read belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Order** | **string** | The sort order of the returned messages. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;asc&#x60; as the default. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMessageResponse**](ListMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> ListRoleResponse ListRole(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListRoleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRoleResponse**](ListRoleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> ListUserResponse ListUser(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUserResponse**](ListUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> ListUserChannelResponse ListUserChannel(ctx, ServiceSidUserSidoptional)



List all Channels for a given User.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from. | 
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from. | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUserChannelResponse**](ListUserChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> ChatV1ServiceChannel UpdateChannel(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service. | 

### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ChatV1Credential UpdateCredential(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. | 
**Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60; | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----&#x60; | 
**Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
**Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. | 

### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ChatV1ServiceChannelMember UpdateMember(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to update belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**LastConsumedMessageIndex** | **int32** | The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels). | 
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services). | 

### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ChatV1ServiceChannelMessage UpdateMessage(ctx, ServiceSidChannelSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from. | 
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**Body** | **string** | The message to send to the channel. Can also be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. | 

### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ChatV1ServiceRole UpdateRole(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60; and are described in the documentation. | 

### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ChatV1Service UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ConsumptionReportInterval** | **int32** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints. | 
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. | 
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. | 
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**LimitsChannelMembers** | **int32** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000. | 
**LimitsUserChannels** | **int32** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000. | 
**NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
**NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
**NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
**NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;. | 
**NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
**NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;. | 
**NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
**NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;. | 
**PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. | 
**PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. | 
**ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;. | 
**ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;. | 
**TypingIndicatorTimeout** | **int32** | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds. | 
**WebhookFilters** | **[]string** | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 
**WebhookMethod** | **string** | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. | 
**WebhooksOnChannelAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_add.url&#x60;. | 
**WebhooksOnChannelAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_add&#x60; event using the &#x60;webhooks.on_channel_add.method&#x60; HTTP method. | 
**WebhooksOnChannelAddedMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event&#x60;. | 
**WebhooksOnChannelAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_added.method&#x60; HTTP method. | 
**WebhooksOnChannelDestroyMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroy.url&#x60;. | 
**WebhooksOnChannelDestroyUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_destroy&#x60; event using the &#x60;webhooks.on_channel_destroy.method&#x60; HTTP method. | 
**WebhooksOnChannelDestroyedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroyed.url&#x60;. | 
**WebhooksOnChannelDestroyedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_destroyed.method&#x60; HTTP method. | 
**WebhooksOnChannelUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_update.url&#x60;. | 
**WebhooksOnChannelUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_update&#x60; event using the &#x60;webhooks.on_channel_update.method&#x60; HTTP method. | 
**WebhooksOnChannelUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;. | 
**WebhooksOnChannelUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method. | 
**WebhooksOnMemberAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_add.url&#x60;. | 
**WebhooksOnMemberAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_add&#x60; event using the &#x60;webhooks.on_member_add.method&#x60; HTTP method. | 
**WebhooksOnMemberAddedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;. | 
**WebhooksOnMemberAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method. | 
**WebhooksOnMemberRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_remove.url&#x60;. | 
**WebhooksOnMemberRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_remove&#x60; event using the &#x60;webhooks.on_member_remove.method&#x60; HTTP method. | 
**WebhooksOnMemberRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_removed.url&#x60;. | 
**WebhooksOnMemberRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_removed&#x60; event using the &#x60;webhooks.on_member_removed.method&#x60; HTTP method. | 
**WebhooksOnMessageRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_remove.url&#x60;. | 
**WebhooksOnMessageRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_remove&#x60; event using the &#x60;webhooks.on_message_remove.method&#x60; HTTP method. | 
**WebhooksOnMessageRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_removed.url&#x60;. | 
**WebhooksOnMessageRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_removed&#x60; event using the &#x60;webhooks.on_message_removed.method&#x60; HTTP method. | 
**WebhooksOnMessageSendMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_send.url&#x60;. | 
**WebhooksOnMessageSendUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_send&#x60; event using the &#x60;webhooks.on_message_send.method&#x60; HTTP method. | 
**WebhooksOnMessageSentMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event&#x60;. | 
**WebhooksOnMessageSentUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event using the &#x60;webhooks.on_message_sent.method&#x60; HTTP method. | 
**WebhooksOnMessageUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_update.url&#x60;. | 
**WebhooksOnMessageUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_update&#x60; event using the &#x60;webhooks.on_message_update.method&#x60; HTTP method. | 
**WebhooksOnMessageUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_updated.url&#x60;. | 
**WebhooksOnMessageUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_updated&#x60; event using the &#x60;webhooks.on_message_updated.method&#x60; HTTP method. | 

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ChatV1ServiceUser UpdateUser(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from. | 
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is often used for display purposes. | 
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user. | 

### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

