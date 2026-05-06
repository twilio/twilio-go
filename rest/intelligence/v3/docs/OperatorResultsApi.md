# OperatorResultsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOperatorResult**](OperatorResultsApi.md#DeleteOperatorResult) | **Delete** /v3/OperatorResults/{operatorResultId} | Delete an OperatorResult
[**FetchOperatorResult**](OperatorResultsApi.md#FetchOperatorResult) | **Get** /v3/OperatorResults/{operatorResultId} | Retrieve an OperatorResult
[**ListOperatorResults**](OperatorResultsApi.md#ListOperatorResults) | **Get** /v3/OperatorResults | Retrieve a list of OperatorResults



## DeleteOperatorResult

> DeleteOperatorResult(ctx, OperatorResultId)

Delete an OperatorResult

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperatorResultId** | **string** | Operator Result id (TTID)

### Other Parameters

Other parameters are passed through a pointer to a DeleteOperatorResultParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOperatorResult

> OperatorResultsResponseV1 FetchOperatorResult(ctx, OperatorResultId)

Retrieve an OperatorResult

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperatorResultId** | **string** | Operator Result id (TTID)

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorResultParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**OperatorResultsResponseV1**](OperatorResultsResponseV1.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperatorResults

> []OperatorResultsResponseV1 ListOperatorResults(ctx, optional)

Retrieve a list of OperatorResults

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorResultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConversationId** | **string** | Filter Operator Results by attached Conversation `id`.
**IntelligenceConfigurationId** | **string** | Filter Operator Results by Intelligence Configuration `id` used to generate them.
**OperatorId** | **string** | Filter Operator Results by Language Operator `id`.
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]OperatorResultsResponseV1**](OperatorResultsResponseV1.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

