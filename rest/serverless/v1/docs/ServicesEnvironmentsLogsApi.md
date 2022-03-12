# ServicesEnvironmentsLogsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchLog**](ServicesEnvironmentsLogsApi.md#FetchLog) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs/{Sid} | 
[**ListLog**](ServicesEnvironmentsLogsApi.md#ListLog) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs | 



## FetchLog

> ServerlessV1Log FetchLog(ctx, ServiceSidEnvironmentSidSid)



Retrieve a specific log.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resource to fetch.
**Sid** | **string** | The SID of the Log resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchLogParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Log**](ServerlessV1Log.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLog

> []ServerlessV1Log ListLog(ctx, ServiceSidEnvironmentSidoptional)



Retrieve a list of all logs.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListLogParams struct


Name | Type | Description
------------- | ------------- | -------------
**FunctionSid** | **string** | The SID of the function whose invocation produced the Log resources to read.
**StartDate** | **time.Time** | The date/time (in GMT, ISO 8601) after which the Log resources must have been created. Defaults to 1 day prior to current date/time.
**EndDate** | **time.Time** | The date/time (in GMT, ISO 8601) before which the Log resources must have been created. Defaults to current date/time.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Log**](ServerlessV1Log.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

