# ControlPlaneOperationsApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperation**](ControlPlaneOperationsApi.md#FetchOperation) | **Get** /v2/ControlPlane/Operations/{operationId} | Get Operation Status



## FetchOperation

> OperationStatus FetchOperation(ctx, OperationId)

Get Operation Status

Retrieve the status and result of an asynchronous operation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperationId** | **string** | The operation ID returned from a write endpoint.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**OperationStatus**](OperationStatus.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

