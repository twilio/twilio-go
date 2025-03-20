# TrustProductsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustProduct**](TrustProductsApi.md#CreateTrustProduct) | **Post** /v1/TrustProducts | Create a new Trust Product.
[**DeleteTrustProduct**](TrustProductsApi.md#DeleteTrustProduct) | **Delete** /v1/TrustProducts/{Sid} | Delete a specific Trust Product.
[**FetchTrustProduct**](TrustProductsApi.md#FetchTrustProduct) | **Get** /v1/TrustProducts/{Sid} | Fetch a specific Trust Product instance.
[**ListTrustProduct**](TrustProductsApi.md#ListTrustProduct) | **Get** /v1/TrustProducts | Retrieve a list of all Trust Products for an account.
[**UpdateTrustProduct**](TrustProductsApi.md#UpdateTrustProduct) | **Post** /v1/TrustProducts/{Sid} | Updates a Trust Product in an account.



## CreateTrustProduct

> TrusthubV1TrustProduct CreateTrustProduct(ctx, optional)

Create a new Trust Product.

Create a new Trust Product.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrustProductParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Trust Product resource changes status.
**PolicySid** | **string** | The unique string of a policy that is associated to the Trust Product resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**TrusthubV1TrustProduct**](TrusthubV1TrustProduct.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustProduct

> DeleteTrustProduct(ctx, Sid)

Delete a specific Trust Product.

Delete a specific Trust Product.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trust Product resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrustProductParams struct


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


## FetchTrustProduct

> TrusthubV1TrustProduct FetchTrustProduct(ctx, Sid)

Fetch a specific Trust Product instance.

Fetch a specific Trust Product instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trust Product resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrustProductParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProduct**](TrusthubV1TrustProduct.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustProduct

> []TrusthubV1TrustProduct ListTrustProduct(ctx, optional)

Retrieve a list of all Trust Products for an account.

Retrieve a list of all Trust Products for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTrustProductParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | [**string**](stringstring.md) | The verification status of the Trust Product resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**PolicySid** | **string** | The unique string of a policy that is associated to the Trust Product resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1TrustProduct**](TrusthubV1TrustProduct.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrustProduct

> TrusthubV1TrustProduct UpdateTrustProduct(ctx, Sidoptional)

Updates a Trust Product in an account.

Updates a Trust Product in an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trust Product resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTrustProductParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | [**string**](string.md) | 
**StatusCallback** | **string** | The URL we call to inform your application of status changes.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Trust Product resource changes status.

### Return type

[**TrusthubV1TrustProduct**](TrusthubV1TrustProduct.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

