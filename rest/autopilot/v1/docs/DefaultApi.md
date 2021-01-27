# \DefaultApi

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
 **optional** | ***CreateAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CallbackEvents** | **optional.String**| Reserved. | 
 **CallbackUrl** | **optional.String**| Reserved. | 
 **Defaults** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
 **LogQueries** | **optional.Bool**| Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored. | 
 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1Assistant**](autopilot.v1.assistant.md)

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
 **optional** | ***CreateFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FieldType** | **optional.String**| The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the &#x60;unique_name&#x60;, or the &#x60;sid&#x60; of a custom Field Type. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantTaskField**](autopilot.v1.assistant.task.field.md)

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
 **optional** | ***CreateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1AssistantFieldType**](autopilot.v1.assistant.field_type.md)

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
 **optional** | ***CreateFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldValueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60; | 
 **SynonymOf** | **optional.String**| The string value that indicates which word the field value is a synonym of. | 
 **Value** | **optional.String**| The Field Value data. | 

### Return type

[**AutopilotV1AssistantFieldTypeFieldValue**](autopilot.v1.assistant.field_type.field_value.md)

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
 **optional** | ***CreateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **StatusCallback** | **optional.String**| The URL we should call using a POST method to send status information to your application. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantModelBuild**](autopilot.v1.assistant.model_build.md)

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
 **optional** | ***CreateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: &#x60;en-US&#x60;. | 
 **ModelBuild** | **optional.String**| The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. | 
 **Query** | **optional.String**| The end-user&#39;s natural language input. It can be up to 2048 characters long. | 
 **Tasks** | **optional.String**| The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task &#x60;unique_name&#x60; values. For example, &#x60;task-unique_name-1, task-unique_name-2&#x60;. Listing specific tasks is useful to constrain the paths that a user can take. | 

### Return type

[**AutopilotV1AssistantQuery**](autopilot.v1.assistant.query.md)

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
 **optional** | ***CreateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: &#x60;en-US&#x60;. | 
 **SourceChannel** | **optional.String**| The communication channel from which the new sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included. | 
 **TaggedText** | **optional.String**| The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). | 

### Return type

[**AutopilotV1AssistantTaskSample**](autopilot.v1.assistant.task.sample.md)

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
 **optional** | ***CreateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique. | 
 **ActionsUrl** | **optional.String**| The URL from which the Assistant can fetch actions. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 

### Return type

[**AutopilotV1AssistantTask**](autopilot.v1.assistant.task.md)

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
 **optional** | ***CreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Events** | **optional.String**| The list of space-separated events that this Webhook will subscribe to. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 
 **WebhookMethod** | **optional.String**| The method to be used when calling the webhook&#39;s URL. | 
 **WebhookUrl** | **optional.String**| The URL associated with this Webhook. | 

### Return type

[**AutopilotV1AssistantWebhook**](autopilot.v1.assistant.webhook.md)

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

[**AutopilotV1Assistant**](autopilot.v1.assistant.md)

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

[**AutopilotV1AssistantDefaults**](autopilot.v1.assistant.defaults.md)

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

[**AutopilotV1AssistantDialogue**](autopilot.v1.assistant.dialogue.md)

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

[**AutopilotV1AssistantTaskField**](autopilot.v1.assistant.task.field.md)

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

[**AutopilotV1AssistantFieldType**](autopilot.v1.assistant.field_type.md)

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

[**AutopilotV1AssistantFieldTypeFieldValue**](autopilot.v1.assistant.field_type.field_value.md)

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

[**AutopilotV1AssistantModelBuild**](autopilot.v1.assistant.model_build.md)

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

[**AutopilotV1AssistantQuery**](autopilot.v1.assistant.query.md)

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

[**AutopilotV1AssistantTaskSample**](autopilot.v1.assistant.task.sample.md)

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

[**AutopilotV1AssistantStyleSheet**](autopilot.v1.assistant.style_sheet.md)

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

[**AutopilotV1AssistantTask**](autopilot.v1.assistant.task.md)

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

[**AutopilotV1AssistantTaskTaskActions**](autopilot.v1.assistant.task.task_actions.md)

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

[**AutopilotV1AssistantTaskTaskStatistics**](autopilot.v1.assistant.task.task_statistics.md)

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

[**AutopilotV1AssistantWebhook**](autopilot.v1.assistant.webhook.md)

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
 **optional** | ***ListAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldValueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60; | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query resources to read. For example: &#x60;en-US&#x60;. | 
 **ModelBuild** | **optional.String**| The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. | 
 **Status** | **optional.String**| The status of the resources to read. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60; | 
 **DialogueSid** | **optional.String**| The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***ListWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWebhookOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **optional** | ***UpdateAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallbackEvents** | **optional.String**| Reserved. | 
 **CallbackUrl** | **optional.String**| Reserved. | 
 **Defaults** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. | 
 **DevelopmentStage** | **optional.String**| A string describing the state of the assistant. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **LogQueries** | **optional.Bool**| Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored. | 
 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1Assistant**](autopilot.v1.assistant.md)

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
 **optional** | ***UpdateDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDefaultsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Defaults** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that describes the default task links for the &#x60;assistant_initiation&#x60;, &#x60;collect&#x60;, and &#x60;fallback&#x60; situations. | 

### Return type

[**AutopilotV1AssistantDefaults**](autopilot.v1.assistant.defaults.md)

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
 **optional** | ***UpdateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique. | 

### Return type

[**AutopilotV1AssistantFieldType**](autopilot.v1.assistant.field_type.md)

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
 **optional** | ***UpdateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantModelBuild**](autopilot.v1.assistant.model_build.md)

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
 **optional** | ***UpdateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **SampleSid** | **optional.String**| The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query. | 
 **Status** | **optional.String**| The new status of the resource. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60; | 

### Return type

[**AutopilotV1AssistantQuery**](autopilot.v1.assistant.query.md)

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
 **optional** | ***UpdateRestoreAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRestoreAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Assistant** | **optional.String**| The Twilio-provided string that uniquely identifies the Assistant resource to restore. | 

### Return type

[**AutopilotV1RestoreAssistant**](autopilot.v1.restore_assistant.md)

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
 **optional** | ***UpdateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Language** | **optional.String**| The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;. | 
 **SourceChannel** | **optional.String**| The communication channel from which the sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included. | 
 **TaggedText** | **optional.String**| The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). | 

### Return type

[**AutopilotV1AssistantTaskSample**](autopilot.v1.assistant.task.sample.md)

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
 **optional** | ***UpdateStyleSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateStyleSheetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that describes the style sheet object. | 

### Return type

[**AutopilotV1AssistantStyleSheet**](autopilot.v1.assistant.style_sheet.md)

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
 **optional** | ***UpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. | 
 **ActionsUrl** | **optional.String**| The URL from which the Assistant can fetch actions. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**AutopilotV1AssistantTask**](autopilot.v1.assistant.task.md)

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
 **optional** | ***UpdateTaskActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskActionsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. | 

### Return type

[**AutopilotV1AssistantTaskTaskActions**](autopilot.v1.assistant.task.task_actions.md)

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
 **optional** | ***UpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebhookOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Events** | **optional.String**| The list of space-separated events that this Webhook will subscribe to. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length. | 
 **WebhookMethod** | **optional.String**| The method to be used when calling the webhook&#39;s URL. | 
 **WebhookUrl** | **optional.String**| The URL associated with this Webhook. | 

### Return type

[**AutopilotV1AssistantWebhook**](autopilot.v1.assistant.webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
