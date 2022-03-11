# TypesApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEventType**](TypesApi.md#FetchEventType) | **Get** /v1/Types/{Type} | 
[**ListEventType**](TypesApi.md#ListEventType) | **Get** /v1/Types | 



## FetchEventType

> EventsV1EventType FetchEventType(ctx, Type)



Fetch a specific Event Type.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | A string that uniquely identifies this Event Type.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1EventType**](EventsV1EventType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventType

> []EventsV1EventType ListEventType(ctx, optional)



Retrieve a paginated list of all the available Event Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEventTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**SchemaId** | **string** | A string parameter filtering the results to return only the Event Types using a given schema.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EventsV1EventType**](EventsV1EventType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

