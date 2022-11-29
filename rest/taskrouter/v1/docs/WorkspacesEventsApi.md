# WorkspacesEventsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEvent**](WorkspacesEventsApi.md#FetchEvent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events/{Sid} | 
[**ListEvent**](WorkspacesEventsApi.md#ListEvent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events | 



## FetchEvent

> TaskrouterV1Event FetchEvent(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Event to fetch.
**Sid** | **string** | The SID of the Event resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1Event**](TaskrouterV1Event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> []TaskrouterV1Event ListEvent(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Events to read. Returns only the Events that pertain to the specified Workspace.

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndDate** | **time.Time** | Only include Events that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**EventType** | **string** | The type of Events to read. Returns only Events of the type specified.
**Minutes** | **int** | The period of events to read in minutes. Returns only Events that occurred since this many minutes in the past. The default is `15` minutes. Task Attributes for Events occuring more 43,200 minutes ago will be redacted.
**ReservationSid** | **string** | The SID of the Reservation with the Events to read. Returns only Events that pertain to the specified Reservation.
**StartDate** | **time.Time** | Only include Events from on or after this date and time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Task Attributes for Events older than 30 days will be redacted.
**TaskQueueSid** | **string** | The SID of the TaskQueue with the Events to read. Returns only the Events that pertain to the specified TaskQueue.
**TaskSid** | **string** | The SID of the Task with the Events to read. Returns only the Events that pertain to the specified Task.
**WorkerSid** | **string** | The SID of the Worker with the Events to read. Returns only the Events that pertain to the specified Worker.
**WorkflowSid** | **string** | The SID of the Workflow with the Events to read. Returns only the Events that pertain to the specified Workflow.
**TaskChannel** | **string** | The TaskChannel with the Events to read. Returns only the Events that pertain to the specified TaskChannel.
**Sid** | **string** | The SID of the Event resource to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1Event**](TaskrouterV1Event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

