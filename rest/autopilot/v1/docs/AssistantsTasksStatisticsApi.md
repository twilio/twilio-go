# AssistantsTasksStatisticsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskStatistics**](AssistantsTasksStatisticsApi.md#FetchTaskStatistics) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Statistics | 



## FetchTaskStatistics

> AutopilotV1TaskStatistics FetchTaskStatistics(ctx, AssistantSidTaskSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) that is associated with the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1TaskStatistics**](AutopilotV1TaskStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

