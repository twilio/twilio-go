# InsightsDomainsConversationsQueryJobsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueryJob**](InsightsDomainsConversationsQueryJobsApi.md#CreateQueryJob) | **Post** /v3/InsightsDomains/Conversations/QueryJobs | Submit an asynchronous query
[**FetchQueryJobStatus**](InsightsDomainsConversationsQueryJobsApi.md#FetchQueryJobStatus) | **Get** /v3/InsightsDomains/Conversations/QueryJobs/{operationId} | Retrieve asynchronous query status and results
[**ListQueryJobs**](InsightsDomainsConversationsQueryJobsApi.md#ListQueryJobs) | **Get** /v3/InsightsDomains/Conversations/QueryJobs | Retrieve a list of asynchronous queries



## CreateQueryJob

> QueryJobSubmitResponse CreateQueryJob(ctx, optional)

Submit an asynchronous query

Submit a long-running semantic query against the Conversations domain. Returns 202 Accepted with an operationId and a Location header pointing to the QueryJobs status endpoint for polling. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateQueryJobParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | A client-generated UUIDv7 key that ensures idempotent behavior. Submitting the same key with an identical request returns the same operation without re-execution. Scoped to Account + Region. Retained for 24 hours. 
**InsightsQueryRequest** | [**InsightsQueryRequest**](InsightsQueryRequest.md) | 

### Return type

[**QueryJobSubmitResponse**](QueryJobSubmitResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQueryJobStatus

> QueryJobStatusResponse FetchQueryJobStatus(ctx, OperationId)

Retrieve asynchronous query status and results

Poll the operation status. When status is COMPLETED, the response includes query results inline. When FAILED, the response includes error details conforming to the OperationError schema. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperationId** | **string** | The unique identifier for the asynchronous query operation, in TTID format.

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryJobStatusParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**QueryJobStatusResponse**](QueryJobStatusResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueryJobs

> []QueryJobStatusResponse ListQueryJobs(ctx, optional)

Retrieve a list of asynchronous queries

Returns a paginated list of asynchronous query operations for the authenticated account. Optionally filter by status. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListQueryJobsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Status** | [**LongRunningOperationStatus**](LongRunningOperationStatusLongRunningOperationStatus.md) | The operation status to filter results by.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]QueryJobStatusResponse**](QueryJobStatusResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

