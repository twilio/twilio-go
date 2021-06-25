# UsersApi

All URIs are relative to *https://frontline-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUser**](UsersApi.md#FetchUser) | **Get** /v1/Users/{Sid} | 
[**UpdateUser**](UsersApi.md#UpdateUser) | **Post** /v1/Users/{Sid} | 



## FetchUser

> FrontlineV1User FetchUser(ctx, Sid)



Fetch a frontline user

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

[**FrontlineV1User**](FrontlineV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> FrontlineV1User UpdateUser(ctx, Sidoptional)



Update an existing frontline user

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Avatar** | **string** | The avatar URL which will be shown in Frontline application.
**FriendlyName** | **string** | The string that you assigned to describe the User.
**State** | **string** | Current state of this user. Can be either &#x60;active&#x60; or &#x60;deactivated&#x60; and defaults to &#x60;active&#x60;

### Return type

[**FrontlineV1User**](FrontlineV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

