# SinksApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSink**](SinksApi.md#CreateSink) | **Post** /v1/Sinks | 
[**DeleteSink**](SinksApi.md#DeleteSink) | **Delete** /v1/Sinks/{Sid} | 
[**FetchSink**](SinksApi.md#FetchSink) | **Get** /v1/Sinks/{Sid} | 
[**ListSink**](SinksApi.md#ListSink) | **Get** /v1/Sinks | 
[**UpdateSink**](SinksApi.md#UpdateSink) | **Post** /v1/Sinks/{Sid} | 



## CreateSink

> EventsV1Sink CreateSink(ctx, optional)



Create a new Sink

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Sink **This value should not contain PII.**
**SinkConfiguration** | [**interface{}**](interface{}.md) | The information required for Twilio to connect to the provided Sink encoded as JSON.
**SinkType** | **string** | 

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSink

> DeleteSink(ctx, Sid)



Delete a specific Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSinkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSink

> EventsV1Sink FetchSink(ctx, Sid)



Fetch a specific Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a FetchSinkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSink

> []EventsV1Sink ListSink(ctx, optional)



Retrieve a paginated list of Sinks belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**InUse** | **bool** | A boolean query parameter filtering the results to return sinks used/not used by a subscription.
**Status** | **string** | A String query parameter filtering the results by status `initialized`, `validating`, `active` or `failed`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSink

> EventsV1Sink UpdateSink(ctx, Sidoptional)



Update a specific Sink

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Sink **This value should not contain PII.**

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

