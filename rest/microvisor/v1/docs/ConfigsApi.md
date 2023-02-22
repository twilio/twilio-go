# ConfigsApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountConfig**](ConfigsApi.md#CreateAccountConfig) | **Post** /v1/Configs | 
[**DeleteAccountConfig**](ConfigsApi.md#DeleteAccountConfig) | **Delete** /v1/Configs/{Key} | 
[**FetchAccountConfig**](ConfigsApi.md#FetchAccountConfig) | **Get** /v1/Configs/{Key} | 
[**ListAccountConfig**](ConfigsApi.md#ListAccountConfig) | **Get** /v1/Configs | 
[**UpdateAccountConfig**](ConfigsApi.md#UpdateAccountConfig) | **Post** /v1/Configs/{Key} | 



## CreateAccountConfig

> MicrovisorV1AccountConfig CreateAccountConfig(ctx, optional)



Create a config for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | The config key; up to 100 characters.
**Value** | **string** | The config value; up to 4096 characters.

### Return type

[**MicrovisorV1AccountConfig**](MicrovisorV1AccountConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountConfig

> DeleteAccountConfig(ctx, Key)



Delete a config for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAccountConfigParams struct


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


## FetchAccountConfig

> MicrovisorV1AccountConfig FetchAccountConfig(ctx, Key)



Retrieve a Config for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountConfigParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1AccountConfig**](MicrovisorV1AccountConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountConfig

> []MicrovisorV1AccountConfig ListAccountConfig(ctx, optional)



Retrieve a list of all Configs for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAccountConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1AccountConfig**](MicrovisorV1AccountConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountConfig

> MicrovisorV1AccountConfig UpdateAccountConfig(ctx, Keyoptional)



Update a config for an Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**Value** | **string** | The config value; up to 4096 characters.

### Return type

[**MicrovisorV1AccountConfig**](MicrovisorV1AccountConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

