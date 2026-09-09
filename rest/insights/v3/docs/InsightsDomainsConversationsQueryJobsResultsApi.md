# InsightsDomainsConversationsQueryJobsResultsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchQueryJobResults**](InsightsDomainsConversationsQueryJobsResultsApi.md#FetchQueryJobResults) | **Get** /v3/InsightsDomains/Conversations/QueryJobs/{operationId}/Results | Retrieve asynchronous query results



## FetchQueryJobResults

> InsightsQueryResponse FetchQueryJobResults(ctx, OperationIdoptional)

Retrieve asynchronous query results

Retrieve paginated results for a completed asynchronous query operation. The resultUrl in the status response points to this endpoint. Returns 404 if the operation is not yet COMPLETED or has expired. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OperationId** | **string** | The unique identifier for the asynchronous query operation, in TTID format.

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryJobResultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination

### Return type

[**InsightsQueryResponse**](InsightsQueryResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

