# AccountsSigningKeysApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSigningKey**](AccountsSigningKeysApi.md#CreateNewSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**DeleteSigningKey**](AccountsSigningKeysApi.md#DeleteSigningKey) | **Delete** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**FetchSigningKey**](AccountsSigningKeysApi.md#FetchSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**ListSigningKey**](AccountsSigningKeysApi.md#ListSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**UpdateSigningKey**](AccountsSigningKeysApi.md#UpdateSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 



## CreateNewSigningKey

> ApiV2010NewSigningKey CreateNewSigningKey(ctx, optional)



Create a new Signing Key for the account making the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewSigningKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010NewSigningKey**](ApiV2010NewSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSigningKey

> DeleteSigningKey(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSigningKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | 

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


## FetchSigningKey

> ApiV2010SigningKey FetchSigningKey(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSigningKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | 

### Return type

[**ApiV2010SigningKey**](ApiV2010SigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSigningKey

> []ApiV2010SigningKey ListSigningKey(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSigningKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SigningKey**](ApiV2010SigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSigningKey

> ApiV2010SigningKey UpdateSigningKey(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSigningKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | 
**FriendlyName** | **string** | 

### Return type

[**ApiV2010SigningKey**](ApiV2010SigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

