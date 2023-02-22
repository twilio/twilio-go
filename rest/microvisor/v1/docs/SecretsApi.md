# SecretsApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSecret**](SecretsApi.md#CreateAccountSecret) | **Post** /v1/Secrets | 
[**DeleteAccountSecret**](SecretsApi.md#DeleteAccountSecret) | **Delete** /v1/Secrets/{Key} | 
[**FetchAccountSecret**](SecretsApi.md#FetchAccountSecret) | **Get** /v1/Secrets/{Key} | 
[**ListAccountSecret**](SecretsApi.md#ListAccountSecret) | **Get** /v1/Secrets | 
[**UpdateAccountSecret**](SecretsApi.md#UpdateAccountSecret) | **Post** /v1/Secrets/{Key} | 



## CreateAccountSecret

> MicrovisorV1AccountSecret CreateAccountSecret(ctx, optional)



Create a secret for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | The secret key; up to 100 characters.
**Value** | **string** | The secret value; up to 4096 characters.

### Return type

[**MicrovisorV1AccountSecret**](MicrovisorV1AccountSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountSecret

> DeleteAccountSecret(ctx, Key)



Delete a secret for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAccountSecretParams struct


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


## FetchAccountSecret

> MicrovisorV1AccountSecret FetchAccountSecret(ctx, Key)



Retrieve a Secret for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountSecretParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1AccountSecret**](MicrovisorV1AccountSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountSecret

> []MicrovisorV1AccountSecret ListAccountSecret(ctx, optional)



Retrieve a list of all Secrets for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAccountSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1AccountSecret**](MicrovisorV1AccountSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSecret

> MicrovisorV1AccountSecret UpdateAccountSecret(ctx, Keyoptional)



Update a secret for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**Value** | **string** | The secret value; up to 4096 characters.

### Return type

[**MicrovisorV1AccountSecret**](MicrovisorV1AccountSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

