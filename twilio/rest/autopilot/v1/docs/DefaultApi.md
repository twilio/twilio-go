# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistant**](DefaultApi.md#CreateAssistant) | **Post** /v1/Assistants | 
[**CreateField**](DefaultApi.md#CreateField) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 
[**CreateFieldType**](DefaultApi.md#CreateFieldType) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes | 
[**CreateFieldValue**](DefaultApi.md#CreateFieldValue) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 
[**CreateModelBuild**](DefaultApi.md#CreateModelBuild) | **Post** /v1/Assistants/{AssistantSid}/ModelBuilds | 
[**CreateQuery**](DefaultApi.md#CreateQuery) | **Post** /v1/Assistants/{AssistantSid}/Queries | 
[**CreateSample**](DefaultApi.md#CreateSample) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /v1/Assistants/{AssistantSid}/Tasks | 
[**CreateWebhook**](DefaultApi.md#CreateWebhook) | **Post** /v1/Assistants/{AssistantSid}/Webhooks | 
[**DeleteAssistant**](DefaultApi.md#DeleteAssistant) | **Delete** /v1/Assistants/{Sid} | 
[**DeleteField**](DefaultApi.md#DeleteField) | **Delete** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**DeleteFieldType**](DefaultApi.md#DeleteFieldType) | **Delete** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**DeleteFieldValue**](DefaultApi.md#DeleteFieldValue) | **Delete** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**DeleteModelBuild**](DefaultApi.md#DeleteModelBuild) | **Delete** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**DeleteQuery**](DefaultApi.md#DeleteQuery) | **Delete** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 
[**DeleteSample**](DefaultApi.md#DeleteSample) | **Delete** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /v1/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 
[**FetchAssistant**](DefaultApi.md#FetchAssistant) | **Get** /v1/Assistants/{Sid} | 
[**FetchDefaults**](DefaultApi.md#FetchDefaults) | **Get** /v1/Assistants/{AssistantSid}/Defaults | 
[**FetchDialogue**](DefaultApi.md#FetchDialogue) | **Get** /v1/Assistants/{AssistantSid}/Dialogues/{Sid} | 
[**FetchField**](DefaultApi.md#FetchField) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**FetchFieldType**](DefaultApi.md#FetchFieldType) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**FetchFieldValue**](DefaultApi.md#FetchFieldValue) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**FetchModelBuild**](DefaultApi.md#FetchModelBuild) | **Get** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**FetchQuery**](DefaultApi.md#FetchQuery) | **Get** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 
[**FetchSample**](DefaultApi.md#FetchSample) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**FetchStyleSheet**](DefaultApi.md#FetchStyleSheet) | **Get** /v1/Assistants/{AssistantSid}/StyleSheet | 
[**FetchTask**](DefaultApi.md#FetchTask) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**FetchTaskActions**](DefaultApi.md#FetchTaskActions) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 
[**FetchTaskStatistics**](DefaultApi.md#FetchTaskStatistics) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Statistics | 
[**FetchWebhook**](DefaultApi.md#FetchWebhook) | **Get** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 
[**ListAssistant**](DefaultApi.md#ListAssistant) | **Get** /v1/Assistants | 
[**ListField**](DefaultApi.md#ListField) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 
[**ListFieldType**](DefaultApi.md#ListFieldType) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes | 
[**ListFieldValue**](DefaultApi.md#ListFieldValue) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 
[**ListModelBuild**](DefaultApi.md#ListModelBuild) | **Get** /v1/Assistants/{AssistantSid}/ModelBuilds | 
[**ListQuery**](DefaultApi.md#ListQuery) | **Get** /v1/Assistants/{AssistantSid}/Queries | 
[**ListSample**](DefaultApi.md#ListSample) | **Get** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**ListTask**](DefaultApi.md#ListTask) | **Get** /v1/Assistants/{AssistantSid}/Tasks | 
[**ListWebhook**](DefaultApi.md#ListWebhook) | **Get** /v1/Assistants/{AssistantSid}/Webhooks | 
[**UpdateAssistant**](DefaultApi.md#UpdateAssistant) | **Post** /v1/Assistants/{Sid} | 
[**UpdateDefaults**](DefaultApi.md#UpdateDefaults) | **Post** /v1/Assistants/{AssistantSid}/Defaults | 
[**UpdateFieldType**](DefaultApi.md#UpdateFieldType) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**UpdateModelBuild**](DefaultApi.md#UpdateModelBuild) | **Post** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**UpdateQuery**](DefaultApi.md#UpdateQuery) | **Post** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 
[**UpdateRestoreAssistant**](DefaultApi.md#UpdateRestoreAssistant) | **Post** /v1/Assistants/Restore | 
[**UpdateSample**](DefaultApi.md#UpdateSample) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**UpdateStyleSheet**](DefaultApi.md#UpdateStyleSheet) | **Post** /v1/Assistants/{AssistantSid}/StyleSheet | 
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**UpdateTaskActions**](DefaultApi.md#UpdateTaskActions) | **Post** /v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 
[**UpdateWebhook**](DefaultApi.md#UpdateWebhook) | **Post** /v1/Assistants/{AssistantSid}/Webhooks/{Sid} | 



## CreateAssistant

> AutopilotV1Assistant CreateAssistant(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAssistantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAssistantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackEvents** | **String**| Reserved. | 
**CallbackUrl** | **String**| Reserved. | 
**Defaults** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
**LogQueries** | **Bool**| Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored. | 
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateField

> AutopilotV1AssistantTaskField CreateField(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the new Field resource. | 
 **optional** | ***CreateFieldRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FieldType** | **String**| The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the &#x60;unique_name&#x60;, or the &#x60;sid&#x60; of a custom Field Type. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantTaskField**](AutopilotV1AssistantTaskField.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldType

> AutopilotV1AssistantFieldType CreateFieldType(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource. | 
 **optional** | ***CreateFieldTypeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldTypeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1AssistantFieldType**](AutopilotV1AssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldValue

> AutopilotV1AssistantFieldTypeFieldValue CreateFieldValue(ctx, AssistantSid, FieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the new resource. | 
**FieldTypeSid** | **string**| The SID of the Field Type associated with the Field Value. | 
 **optional** | ***CreateFieldValueRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldValueRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60; | 
**SynonymOf** | **String**| The string value that indicates which word the field value is a synonym of. | 
**Value** | **String**| The Field Value data. | 

### Return type

[**AutopilotV1AssistantFieldTypeFieldValue**](AutopilotV1AssistantFieldTypeFieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelBuild

> AutopilotV1AssistantModelBuild CreateModelBuild(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource. | 
 **optional** | ***CreateModelBuildRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateModelBuildRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**StatusCallback** | **String**| The URL we should call using a POST method to send status information to your application. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantModelBuild**](AutopilotV1AssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuery

> AutopilotV1AssistantQuery CreateQuery(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource. | 
 **optional** | ***CreateQueryRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: &#x60;en-US&#x60;. | 
**ModelBuild** | **String**| The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. | 
**Query** | **String**| The end-user&#39;s natural language input. It can be up to 2048 characters long. | 
**Tasks** | **String**| The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task &#x60;unique_name&#x60; values. For example, &#x60;task-unique_name-1, task-unique_name-2&#x60;. Listing specific tasks is useful to constrain the paths that a user can take. | 

### Return type

[**AutopilotV1AssistantQuery**](AutopilotV1AssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSample

> AutopilotV1AssistantTaskSample CreateSample(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create. | 
 **optional** | ***CreateSampleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSampleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: &#x60;en-US&#x60;. | 
**SourceChannel** | **String**| The communication channel from which the new sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included. | 
**TaggedText** | **String**| The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). | 

### Return type

[**AutopilotV1AssistantTaskSample**](AutopilotV1AssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> AutopilotV1AssistantTask CreateTask(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource. | 
 **optional** | ***CreateTaskRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique. | 
**ActionsUrl** | **String**| The URL from which the Assistant can fetch actions. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 

### Return type

[**AutopilotV1AssistantTask**](AutopilotV1AssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> AutopilotV1AssistantWebhook CreateWebhook(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource. | 
 **optional** | ***CreateWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Events** | **String**| The list of space-separated events that this Webhook will subscribe to. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 
**WebhookMethod** | **String**| The method to be used when calling the webhook&#39;s URL. | 
**WebhookUrl** | **String**| The URL associated with this Webhook. | 

### Return type

[**AutopilotV1AssistantWebhook**](AutopilotV1AssistantWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistant

> DeleteAssistant(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Assistant resource to delete. | 

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


## DeleteField

> DeleteField(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Field resource to delete. | 

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


## DeleteFieldType

> DeleteFieldType(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the FieldType resource to delete. | 

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


## DeleteFieldValue

> DeleteFieldValue(ctx, AssistantSid, FieldTypeSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to delete. | 
**FieldTypeSid** | **string**| The SID of the Field Type associated with the Field Value to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the FieldValue resource to delete. | 

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


## DeleteModelBuild

> DeleteModelBuild(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ModelBuild resource to delete. | 

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


## DeleteQuery

> DeleteQuery(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Query resource to delete. | 

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


## DeleteSample

> DeleteSample(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Sample resource to delete. | 

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


## DeleteTask

> DeleteTask(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Task resource to delete. | 

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


## DeleteWebhook

> DeleteWebhook(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to delete. | 

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


## FetchAssistant

> AutopilotV1Assistant FetchAssistant(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Assistant resource to fetch. | 

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDefaults

> AutopilotV1AssistantDefaults FetchDefaults(ctx, AssistantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 

### Return type

[**AutopilotV1AssistantDefaults**](AutopilotV1AssistantDefaults.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialogue

> AutopilotV1AssistantDialogue FetchDialogue(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Dialogue resource to fetch. | 

### Return type

[**AutopilotV1AssistantDialogue**](AutopilotV1AssistantDialogue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchField

> AutopilotV1AssistantTaskField FetchField(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Field resource to fetch. | 

### Return type

[**AutopilotV1AssistantTaskField**](AutopilotV1AssistantTaskField.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldType

> AutopilotV1AssistantFieldType FetchFieldType(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the FieldType resource to fetch. | 

### Return type

[**AutopilotV1AssistantFieldType**](AutopilotV1AssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldValue

> AutopilotV1AssistantFieldTypeFieldValue FetchFieldValue(ctx, AssistantSid, FieldTypeSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource to fetch. | 
**FieldTypeSid** | **string**| The SID of the Field Type associated with the Field Value to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the FieldValue resource to fetch. | 

### Return type

[**AutopilotV1AssistantFieldTypeFieldValue**](AutopilotV1AssistantFieldTypeFieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchModelBuild

> AutopilotV1AssistantModelBuild FetchModelBuild(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ModelBuild resource to fetch. | 

### Return type

[**AutopilotV1AssistantModelBuild**](AutopilotV1AssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQuery

> AutopilotV1AssistantQuery FetchQuery(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Query resource to fetch. | 

### Return type

[**AutopilotV1AssistantQuery**](AutopilotV1AssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSample

> AutopilotV1AssistantTaskSample FetchSample(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Sample resource to fetch. | 

### Return type

[**AutopilotV1AssistantTaskSample**](AutopilotV1AssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStyleSheet

> AutopilotV1AssistantStyleSheet FetchStyleSheet(ctx, AssistantSid)



Returns Style sheet JSON object for the Assistant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 

### Return type

[**AutopilotV1AssistantStyleSheet**](AutopilotV1AssistantStyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTask

> AutopilotV1AssistantTask FetchTask(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Task resource to fetch. | 

### Return type

[**AutopilotV1AssistantTask**](AutopilotV1AssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskActions

> AutopilotV1AssistantTaskTaskActions FetchTaskActions(ctx, AssistantSid, TaskSid)



Returns JSON actions for the Task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to fetch were defined. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to fetch were defined. | 

### Return type

[**AutopilotV1AssistantTaskTaskActions**](AutopilotV1AssistantTaskTaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskStatistics

> AutopilotV1AssistantTaskTaskStatistics FetchTaskStatistics(ctx, AssistantSid, TaskSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) that is associated with the resource to fetch. | 

### Return type

[**AutopilotV1AssistantTaskTaskStatistics**](AutopilotV1AssistantTaskTaskStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebhook

> AutopilotV1AssistantWebhook FetchWebhook(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to fetch. | 

### Return type

[**AutopilotV1AssistantWebhook**](AutopilotV1AssistantWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistant

> ListAssistantResponse ListAssistant(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAssistantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssistantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAssistantResponse**](ListAssistantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListField

> ListFieldResponse ListField(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resources to read. | 
 **optional** | ***ListFieldRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldResponse**](ListFieldResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldType

> ListFieldTypeResponse ListFieldType(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read. | 
 **optional** | ***ListFieldTypeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldTypeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldTypeResponse**](ListFieldTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldValue

> ListFieldValueResponse ListFieldValue(ctx, AssistantSid, FieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to read. | 
**FieldTypeSid** | **string**| The SID of the Field Type associated with the Field Value to read. | 
 **optional** | ***ListFieldValueRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldValueRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60; | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldValueResponse**](ListFieldValueResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelBuild

> ListModelBuildResponse ListModelBuild(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read. | 
 **optional** | ***ListModelBuildRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListModelBuildRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListModelBuildResponse**](ListModelBuildResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuery

> ListQueryResponse ListQuery(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read. | 
 **optional** | ***ListQueryRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQueryRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query resources to read. For example: &#x60;en-US&#x60;. | 
**ModelBuild** | **String**| The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. | 
**Status** | **String**| The status of the resources to read. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60; | 
**DialogueSid** | **String**| The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListQueryResponse**](ListQueryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSample

> ListSampleResponse ListSample(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resources to read. | 
 **optional** | ***ListSampleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSampleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSampleResponse**](ListSampleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> ListTaskResponse ListTask(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read. | 
 **optional** | ***ListTaskRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListTaskResponse**](ListTaskResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhook

> ListWebhookResponse ListWebhook(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read. | 
 **optional** | ***ListWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListWebhookResponse**](ListWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistant

> AutopilotV1Assistant UpdateAssistant(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Assistant resource to update. | 
 **optional** | ***UpdateAssistantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackEvents** | **String**| Reserved. | 
**CallbackUrl** | **String**| Reserved. | 
**Defaults** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. | 
**DevelopmentStage** | **String**| A string describing the state of the assistant. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**LogQueries** | **Bool**| Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored. | 
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaults

> AutopilotV1AssistantDefaults UpdateDefaults(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
 **optional** | ***UpdateDefaultsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDefaultsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Defaults** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that describes the default task links for the &#x60;assistant_initiation&#x60;, &#x60;collect&#x60;, and &#x60;fallback&#x60; situations. | 

### Return type

[**AutopilotV1AssistantDefaults**](AutopilotV1AssistantDefaults.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldType

> AutopilotV1AssistantFieldType UpdateFieldType(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the FieldType resource to update. | 
 **optional** | ***UpdateFieldTypeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFieldTypeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1AssistantFieldType**](AutopilotV1AssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelBuild

> AutopilotV1AssistantModelBuild UpdateModelBuild(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ModelBuild resource to update. | 
 **optional** | ***UpdateModelBuildRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateModelBuildRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantModelBuild**](AutopilotV1AssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuery

> AutopilotV1AssistantQuery UpdateQuery(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Query resource to update. | 
 **optional** | ***UpdateQueryRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQueryRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**SampleSid** | **String**| The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query. | 
**Status** | **String**| The new status of the resource. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60; | 

### Return type

[**AutopilotV1AssistantQuery**](AutopilotV1AssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRestoreAssistant

> AutopilotV1RestoreAssistant UpdateRestoreAssistant(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateRestoreAssistantRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRestoreAssistantRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Assistant** | **String**| The Twilio-provided string that uniquely identifies the Assistant resource to restore. | 

### Return type

[**AutopilotV1RestoreAssistant**](AutopilotV1RestoreAssistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSample

> AutopilotV1AssistantTaskSample UpdateSample(ctx, AssistantSid, TaskSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to update. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Sample resource to update. | 
 **optional** | ***UpdateSampleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSampleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;. | 
**SourceChannel** | **String**| The communication channel from which the sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included. | 
**TaggedText** | **String**| The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). | 

### Return type

[**AutopilotV1AssistantTaskSample**](AutopilotV1AssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStyleSheet

> AutopilotV1AssistantStyleSheet UpdateStyleSheet(ctx, AssistantSid, optional)



Updates the style sheet for an Assistant identified by `assistant_sid`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
 **optional** | ***UpdateStyleSheetRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateStyleSheetRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that describes the style sheet object. | 

### Return type

[**AutopilotV1AssistantStyleSheet**](AutopilotV1AssistantStyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> AutopilotV1AssistantTask UpdateTask(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Task resource to update. | 
 **optional** | ***UpdateTaskRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. | 
**ActionsUrl** | **String**| The URL from which the Assistant can fetch actions. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantTask**](AutopilotV1AssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskActions

> AutopilotV1AssistantTaskTaskActions UpdateTaskActions(ctx, AssistantSid, TaskSid, optional)



Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to update were defined. | 
**TaskSid** | **string**| The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to update were defined. | 
 **optional** | ***UpdateTaskActionsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskActionsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. | 

### Return type

[**AutopilotV1AssistantTaskTaskActions**](AutopilotV1AssistantTaskTaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> AutopilotV1AssistantWebhook UpdateWebhook(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to update. | 
 **optional** | ***UpdateWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Events** | **String**| The list of space-separated events that this Webhook will subscribe to. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 
**WebhookMethod** | **String**| The method to be used when calling the webhook&#39;s URL. | 
**WebhookUrl** | **String**| The URL associated with this Webhook. | 

### Return type

[**AutopilotV1AssistantWebhook**](AutopilotV1AssistantWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

