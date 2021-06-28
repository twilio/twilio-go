# ExportsApi

All URIs are relative to *https://bulkexports.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchExport**](ExportsApi.md#FetchExport) | **Get** /v1/Exports/{ResourceType} | 



## FetchExport

> BulkexportsV1Export FetchExport(ctx, ResourceType)



Fetch a specific Export.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a FetchExportParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**BulkexportsV1Export**](BulkexportsV1Export.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

