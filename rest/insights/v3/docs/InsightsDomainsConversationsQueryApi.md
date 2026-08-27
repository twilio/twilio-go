# InsightsDomainsConversationsQueryApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueryResults**](InsightsDomainsConversationsQueryApi.md#CreateQueryResults) | **Post** /v3/InsightsDomains/Conversations/Query | Execute a synchronous semantic query
[**FetchQueryResults**](InsightsDomainsConversationsQueryApi.md#FetchQueryResults) | **Get** /v3/InsightsDomains/Conversations/Query | Retrieve paginated query results



## CreateQueryResults

> InsightsQueryResponse CreateQueryResults(ctx, optional)

Execute a synchronous semantic query

Execute a semantic query against the Conversations domain. Returns results inline. For long-running queries, use the QueryJobs endpoint. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateQueryResultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page
**InsightsQueryRequest** | [**InsightsQueryRequest**](InsightsQueryRequest.md) | 

### Return type

[**InsightsQueryResponse**](InsightsQueryResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQueryResults

> InsightsQueryResponse FetchQueryResults(ctx, optional)

Retrieve paginated query results

Retrieve subsequent pages of a synchronous query using a page token.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryResultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageToken** | **string** | Pagination token

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

