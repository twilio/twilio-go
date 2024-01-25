# AssistantsTasksSamplesApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSample**](AssistantsTasksSamplesApi.md#CreateSample) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**DeleteSample**](AssistantsTasksSamplesApi.md#DeleteSample) | **Delete** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**FetchSample**](AssistantsTasksSamplesApi.md#FetchSample) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**ListSample**](AssistantsTasksSamplesApi.md#ListSample) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**UpdateSample**](AssistantsTasksSamplesApi.md#UpdateSample) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 



## CreateSample

> AutopilotV1Sample CreateSample(ctx, AssistantSidTaskSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.

### Other Parameters

Other parameters are passed through a pointer to a CreateSampleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: `en-US`.
**TaggedText** | **string** | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
**SourceChannel** | **string** | The communication channel from which the new sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.

### Return type

[**AutopilotV1Sample**](AutopilotV1Sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSample

> DeleteSample(ctx, AssistantSidTaskSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSampleParams struct


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


## FetchSample

> AutopilotV1Sample FetchSample(ctx, AssistantSidTaskSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSampleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Sample**](AutopilotV1Sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSample

> []AutopilotV1Sample ListSample(ctx, AssistantSidTaskSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSampleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1Sample**](AutopilotV1Sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSample

> AutopilotV1Sample UpdateSample(ctx, AssistantSidTaskSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to update.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSampleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
**TaggedText** | **string** | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
**SourceChannel** | **string** | The communication channel from which the sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.

### Return type

[**AutopilotV1Sample**](AutopilotV1Sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

