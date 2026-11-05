# WebhooksOperationsApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperation**](WebhooksOperationsApi.md#FetchOperation) | **Get** /v1/Webhooks/Operations/{operationId} | Get Operation status
[**ListOperation**](WebhooksOperationsApi.md#ListOperation) | **Get** /v1/Webhooks/Operations | List all Operations



## FetchOperation

> Operation FetchOperation(ctx, OperationId)

Get Operation status

Check status of an async Rule operation.  ### Operation Lifecycle  Operations transition from `PENDING` to `COMPLETED` status within 30-60 seconds of a Rule mutation being accepted (202 Accepted with `Operation-Id` header). 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperationId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchOperationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Operation**](Operation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperation

> []Operation ListOperation(ctx, optional)

List all Operations

Returns a paginated list of async Rule operations.  Use this endpoint to recover a lost `Operation-Id` — for example, after a process crash before the ID was persisted. Results are sorted by `createdAt` descending (newest first). 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOperationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**Status** | **string** | Filter by operation state
**ResourceId** | **string** | Filter operations affecting a specific resource ID
**Limit** | **int** | Max number of records to return.

### Return type

[**[]Operation**](Operation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

