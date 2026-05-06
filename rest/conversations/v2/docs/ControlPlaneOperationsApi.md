# ControlPlaneOperationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperationStatus**](ControlPlaneOperationsApi.md#FetchOperationStatus) | **Get** /v2/ControlPlane/Operations/{Sid} | Get Operation Status



## FetchOperationStatus

> FetchOperationStatusResponse FetchOperationStatus(ctx, Sid)

Get Operation Status

Retrieve the current status of a long-running operation. Operations progress through: PENDING -> RUNNING -> COMPLETED or FAILED. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchOperationStatusParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FetchOperationStatusResponse**](FetchOperationStatus200Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

