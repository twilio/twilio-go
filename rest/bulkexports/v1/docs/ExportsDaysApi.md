# ExportsDaysApi

All URIs are relative to *https://bulkexports.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDay**](ExportsDaysApi.md#FetchDay) | **Get** /v1/Exports/{ResourceType}/Days/{Day} | 
[**ListDay**](ExportsDaysApi.md#ListDay) | **Get** /v1/Exports/{ResourceType}/Days | 



## FetchDay

> BulkexportsV1DayInstance FetchDay(ctx, ResourceTypeDay)



Fetch a specific Day.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants
**Day** | **string** | The ISO 8601 format date of the resources in the file, for a UTC day

### Other Parameters

Other parameters are passed through a pointer to a FetchDayParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**BulkexportsV1DayInstance**](BulkexportsV1DayInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> []BulkexportsV1Day ListDay(ctx, ResourceTypeoptional)



Retrieve a list of all Days for a resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a ListDayParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]BulkexportsV1Day**](BulkexportsV1Day.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

