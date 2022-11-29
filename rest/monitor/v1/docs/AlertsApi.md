# AlertsApi

All URIs are relative to *https://monitor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAlert**](AlertsApi.md#FetchAlert) | **Get** /v1/Alerts/{Sid} | 
[**ListAlert**](AlertsApi.md#ListAlert) | **Get** /v1/Alerts | 



## FetchAlert

> MonitorV1AlertInstance FetchAlert(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Alert resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAlertParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MonitorV1AlertInstance**](MonitorV1AlertInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlert

> []MonitorV1Alert ListAlert(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAlertParams struct


Name | Type | Description
------------- | ------------- | -------------
**LogLevel** | **string** | Only show alerts for this log-level.  Can be: `error`, `warning`, `notice`, or `debug`.
**StartDate** | **time.Time** | Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
**EndDate** | **time.Time** | Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MonitorV1Alert**](MonitorV1Alert.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

