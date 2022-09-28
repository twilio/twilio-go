# EventsApi

All URIs are relative to *https://monitor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEvent**](EventsApi.md#FetchEvent) | **Get** /v1/Events/{Sid} | 
[**ListEvent**](EventsApi.md#ListEvent) | **Get** /v1/Events | 



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


## ListEvent

> []MonitorV1Event ListEvent(ctx, optional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MonitorV1Event**](MonitorV1Event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

