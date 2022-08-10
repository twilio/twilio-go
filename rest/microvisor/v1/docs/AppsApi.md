# AppsApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApp**](AppsApi.md#DeleteApp) | **Delete** /v1/Apps/{Sid} | 
[**FetchApp**](AppsApi.md#FetchApp) | **Get** /v1/Apps/{Sid} | 
[**ListApp**](AppsApi.md#ListApp) | **Get** /v1/Apps | 



## DeleteApp

> DeleteApp(ctx, Sid)



Delete a specific App.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34-character string that uniquely identifies this App.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAppParams struct


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


## FetchApp

> MicrovisorV1App FetchApp(ctx, Sid)



Fetch a specific App.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34-character string that uniquely identifies this App.

### Other Parameters

Other parameters are passed through a pointer to a FetchAppParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1App**](MicrovisorV1App.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApp

> []MicrovisorV1App ListApp(ctx, optional)



Retrieve a list of all Apps for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1App**](MicrovisorV1App.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

