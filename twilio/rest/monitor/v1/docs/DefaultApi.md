# DefaultApi

All URIs are relative to *https://monitor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAlert**](DefaultApi.md#FetchAlert) | **Get** /v1/Alerts/{Sid} | 
[**FetchEvent**](DefaultApi.md#FetchEvent) | **Get** /v1/Events/{Sid} | 
[**ListAlert**](DefaultApi.md#ListAlert) | **Get** /v1/Alerts | 
[**ListEvent**](DefaultApi.md#ListEvent) | **Get** /v1/Events | 



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


## FetchEvent

> MonitorV1Event FetchEvent(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Event resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventParams struct

Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MonitorV1Event**](MonitorV1Event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlert

> ListAlertResponse ListAlert(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAlertParams struct

Name | Type | Description
------------- | ------------- | -------------
**LogLevel** | **string** | Only show alerts for this log-level.  Can be: &#x60;error&#x60;, &#x60;warning&#x60;, &#x60;notice&#x60;, or &#x60;debug&#x60;.
**StartDate** | **time.Time** | Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as &#x60;YYYY-MM-DD&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;. Queries for alerts older than 30 days are not supported.
**EndDate** | **time.Time** | Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as &#x60;YYYY-MM-DD&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;. Queries for alerts older than 30 days are not supported.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAlertResponse**](ListAlertResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> ListEventResponse ListEvent(ctx, optional)



Returns a list of events in the account, sorted by event-date.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct

Name | Type | Description
------------- | ------------- | -------------
**ActorSid** | **string** | Only include events initiated by this Actor. Useful for auditing actions taken by specific users or API credentials.
**EventType** | **string** | Only include events of this [Event Type](https://www.twilio.com/docs/usage/monitor-events#event-types).
**ResourceSid** | **string** | Only include events that refer to this resource. Useful for discovering the history of a specific resource.
**SourceIpAddress** | **string** | Only include events that originated from this IP address. Useful for tracking suspicious activity originating from the API or the Twilio Console.
**StartDate** | **time.Time** | Only include events that occurred on or after this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**EndDate** | **time.Time** | Only include events that occurred on or before this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventResponse**](ListEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

