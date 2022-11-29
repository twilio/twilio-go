# ServicesUsersApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceUser**](ServicesUsersApi.md#CreateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users | 
[**DeleteServiceUser**](ServicesUsersApi.md#DeleteServiceUser) | **Delete** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**FetchServiceUser**](ServicesUsersApi.md#FetchServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**ListServiceUser**](ServicesUsersApi.md#ListServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users | 
[**UpdateServiceUser**](ServicesUsersApi.md#UpdateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users/{Sid} | 



## CreateServiceUser

> ConversationsV1ServiceUser CreateServiceUser(ctx, ChatServiceSidoptional)



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
**Identity** | **string** | The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceUser**](ConversationsV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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
**Sid** | **string** | The SID of the User resource to delete. This value can be either the `sid` or the `identity` of the User resource to delete.

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


## FetchServiceUser

> ConversationsV1ServiceUser FetchServiceUser(ctx, ChatServiceSidSid)



Fetch a conversation user from your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the User resource from.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the `sid` or the `identity` of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceUser**](ConversationsV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceUser

> []ConversationsV1ServiceUser ListServiceUser(ctx, ChatServiceSidoptional)



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
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceUser**](ConversationsV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceUser

> ConversationsV1ServiceUser UpdateServiceUser(ctx, ChatServiceSidSidoptional)



Update an existing conversation user in your service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
**Sid** | **string** | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
**RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceUser**](ConversationsV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

