# AssistantsTasksActionsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskActions**](AssistantsTasksActionsApi.md#FetchTaskActions) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 
[**UpdateTaskActions**](AssistantsTasksActionsApi.md#UpdateTaskActions) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 



## FetchTaskActions

> AutopilotV1TaskActions FetchTaskActions(ctx, AssistantSidTaskSid)



Returns JSON actions for the Task.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to fetch were defined.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to fetch were defined.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskActionsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1TaskActions**](AutopilotV1TaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskActions

> AutopilotV1TaskActions UpdateTaskActions(ctx, AssistantSidTaskSidoptional)



Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to update were defined.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to update were defined.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskActionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Actions** | [**interface{}**](interface{}.md) | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.

### Return type

[**AutopilotV1TaskActions**](AutopilotV1TaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

