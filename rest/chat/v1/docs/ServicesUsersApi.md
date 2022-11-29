# ServicesUsersApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](ServicesUsersApi.md#CreateUser) | **Post** /v1/Services/{ServiceSid}/Users | 
[**DeleteUser**](ServicesUsersApi.md#DeleteUser) | **Delete** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**FetchUser**](ServicesUsersApi.md#FetchUser) | **Get** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**ListUser**](ServicesUsersApi.md#ListUser) | **Get** /v1/Services/{ServiceSid}/Users | 
[**UpdateUser**](ServicesUsersApi.md#UpdateUser) | **Post** /v1/Services/{ServiceSid}/Users/{Sid} | 



## CreateUser

> ChatV1User CreateUser(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. This value is often used for display purposes.

### Return type

[**ChatV1User**](ChatV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to delete.

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


## FetchUser

> ChatV1User FetchUser(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ChatV1User**](ChatV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> []ChatV1User ListUser(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ChatV1User**](ChatV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ChatV1User UpdateUser(ctx, ServiceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user.
**Attributes** | **string** | A valid JSON string that contains application-specific data.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is often used for display purposes.

### Return type

[**ChatV1User**](ChatV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

