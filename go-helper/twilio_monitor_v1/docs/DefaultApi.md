# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAlert**](DefaultApi.md#FetchAlert) | **Get** /v1/Alerts/{Sid} | 
[**FetchEvent**](DefaultApi.md#FetchEvent) | **Get** /v1/Events/{Sid} | 
[**ListAlert**](DefaultApi.md#ListAlert) | **Get** /v1/Alerts | 
[**ListEvent**](DefaultApi.md#ListEvent) | **Get** /v1/Events | 



## FetchAlert

> MonitorV1AlertInstance FetchAlert(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Alert resource to fetch. | 

### Return type

[**MonitorV1AlertInstance**](monitor.v1.alert-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvent

> MonitorV1Event FetchEvent(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Event resource to fetch. | 

### Return type

[**MonitorV1Event**](monitor.v1.event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlert

> InlineResponse200 ListAlert(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAlertOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logLevel** | **optional.String**| Only show alerts for this log-level.  Can be: &#x60;error&#x60;, &#x60;warning&#x60;, &#x60;notice&#x60;, or &#x60;debug&#x60;. | 
 **startDate** | **optional.Time**| Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as &#x60;YYYY-MM-DD&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;. Queries for alerts older than 30 days are not supported. | 
 **endDate** | **optional.Time**| Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as &#x60;YYYY-MM-DD&#x60; or &#x60;YYYY-MM-DDThh:mm:ssZ&#x60;. Queries for alerts older than 30 days are not supported. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> InlineResponse2001 ListEvent(ctx, optional)



Returns a list of events in the account, sorted by event-date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorSid** | **optional.String**| Only include events initiated by this Actor. Useful for auditing actions taken by specific users or API credentials. | 
 **eventType** | **optional.String**| Only include events of this [Event Type](https://www.twilio.com/docs/usage/monitor-events#event-types). | 
 **resourceSid** | **optional.String**| Only include events that refer to this resource. Useful for discovering the history of a specific resource. | 
 **sourceIpAddress** | **optional.String**| Only include events that originated from this IP address. Useful for tracking suspicious activity originating from the API or the Twilio Console. | 
 **startDate** | **optional.Time**| Only include events that occurred on or after this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **endDate** | **optional.Time**| Only include events that occurred on or before this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

