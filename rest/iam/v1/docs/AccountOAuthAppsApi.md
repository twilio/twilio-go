# AccountOAuthAppsApi

All URIs are relative to *https://iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthAppAccount**](AccountOAuthAppsApi.md#CreateOauthAppAccount) | **Post** /v1/Account/OAuthApps | Creates an oauth app for the given AccountSid
[**DeleteOauthAppAccount**](AccountOAuthAppsApi.md#DeleteOauthAppAccount) | **Delete** /v1/Account/OAuthApps/{sid} | Deletes an oauth app for the given AccountSid
[**UpdateOauthAppAccount**](AccountOAuthAppsApi.md#UpdateOauthAppAccount) | **Put** /v1/Account/OAuthApps/{sid} | Updates an existing OAuth app for the given OauthAppSid



## CreateOauthAppAccount

> IamV1VendorOauthAppCreateUpdateResponse CreateOauthAppAccount(ctx, optional)

Creates an oauth app for the given AccountSid

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateOauthAppAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**IamV1VendorOauthAppCreateRequest** | [**IamV1VendorOauthAppCreateRequest**](IamV1VendorOauthAppCreateRequest.md) | 

### Return type

[**IamV1VendorOauthAppCreateUpdateResponse**](IamV1VendorOauthAppCreateUpdateResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthAppAccount

> DeleteOauthAppAccount(ctx, Sid)

Deletes an oauth app for the given AccountSid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique ID (sid) of the OAuth app

### Other Parameters

Other parameters are passed through a pointer to a DeleteOauthAppAccountParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthAppAccount

> IamV1VendorOauthAppCreateUpdateResponse UpdateOauthAppAccount(ctx, Sidoptional)

Updates an existing OAuth app for the given OauthAppSid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique ID (sid) of the OAuth app

### Other Parameters

Other parameters are passed through a pointer to a UpdateOauthAppAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**IamV1VendorOauthAppUpdateRequest** | [**IamV1VendorOauthAppUpdateRequest**](IamV1VendorOauthAppUpdateRequest.md) | 

### Return type

[**IamV1VendorOauthAppCreateUpdateResponse**](IamV1VendorOauthAppCreateUpdateResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

