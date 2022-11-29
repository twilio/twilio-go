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
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the `sid` or the `identity` of the User resource to fetch.

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
**Sid** | **string** | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the User.
**Avatar** | **string** | The avatar URL which will be shown in Frontline application.
**State** | **string** | 
**IsAvailable** | **bool** | Whether the User is available for new conversations. Set to `false` to prevent User from receiving new inbound conversations if you are using [Pool Routing](https://www.twilio.com/docs/frontline/handle-incoming-conversations#3-pool-routing).

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

