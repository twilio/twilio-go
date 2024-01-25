# AssistantsTasksFieldsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateField**](AssistantsTasksFieldsApi.md#CreateField) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 
[**DeleteField**](AssistantsTasksFieldsApi.md#DeleteField) | **Delete** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**FetchField**](AssistantsTasksFieldsApi.md#FetchField) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**ListField**](AssistantsTasksFieldsApi.md#ListField) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 



## CreateField

> AutopilotV1Field CreateField(ctx, AssistantSidTaskSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the new Field resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldParams struct


Name | Type | Description
------------- | ------------- | -------------
**FieldType** | **string** | The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the `unique_name`, or the `sid` of a custom Field Type.
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.

### Return type

[**AutopilotV1Field**](AutopilotV1Field.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteField

> DeleteField(ctx, AssistantSidTaskSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Field resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldParams struct


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


## FetchField

> AutopilotV1Field FetchField(ctx, AssistantSidTaskSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Field resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Field**](AutopilotV1Field.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListField

> []AutopilotV1Field ListField(ctx, AssistantSidTaskSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1Field**](AutopilotV1Field.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

