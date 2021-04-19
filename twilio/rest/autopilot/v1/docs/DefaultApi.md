# \DefaultApi

All URIs are relative to *https://autopilot.twilio.com*

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

> AutopilotV1Assistant CreateAssistant(ctx).CallbackEvents(CallbackEvents).CallbackUrl(CallbackUrl).Defaults(Defaults).FriendlyName(FriendlyName).LogQueries(LogQueries).StyleSheet(StyleSheet).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    CallbackEvents := "CallbackEvents_example" // string | Reserved. (optional)
    CallbackUrl := "CallbackUrl_example" // string | Reserved. (optional)
    Defaults := TODO // map[string]interface{} | A JSON object that defines the Assistant's [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. (optional)
    LogQueries := true // bool | Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored. (optional)
    StyleSheet := TODO // map[string]interface{} | The JSON string that defines the Assistant's [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAssistant(context.Background()).CallbackEvents(CallbackEvents).CallbackUrl(CallbackUrl).Defaults(Defaults).FriendlyName(FriendlyName).LogQueries(LogQueries).StyleSheet(StyleSheet).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssistant`: AutopilotV1Assistant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAssistant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CallbackEvents** | **string** | Reserved.
 **CallbackUrl** | **string** | Reserved.
 **Defaults** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
 **LogQueries** | **bool** | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored.
 **StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

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

> AutopilotV1AssistantTaskField CreateField(ctx, AssistantSid, TaskSid).FieldType(FieldType).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the new Field resource.
    FieldType := "FieldType_example" // string | The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the `unique_name`, or the `sid` of a custom Field Type. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateField(context.Background(), AssistantSid, TaskSid).FieldType(FieldType).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateField`: AutopilotV1AssistantTaskField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the new Field resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FieldType** | **string** | The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the &#x60;unique_name&#x60;, or the &#x60;sid&#x60; of a custom Field Type.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

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

> AutopilotV1AssistantFieldType CreateFieldType(ctx, AssistantSid).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFieldType(context.Background(), AssistantSid).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFieldType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFieldType`: AutopilotV1AssistantFieldType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFieldType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

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

> AutopilotV1AssistantFieldTypeFieldValue CreateFieldValue(ctx, AssistantSid, FieldTypeSid).Language(Language).SynonymOf(SynonymOf).Value(Value).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the new resource.
    FieldTypeSid := "FieldTypeSid_example" // string | The SID of the Field Type associated with the Field Value.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US` (optional)
    SynonymOf := "SynonymOf_example" // string | The string value that indicates which word the field value is a synonym of. (optional)
    Value := "Value_example" // string | The Field Value data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFieldValue(context.Background(), AssistantSid, FieldTypeSid).Language(Language).SynonymOf(SynonymOf).Value(Value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFieldValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFieldValue`: AutopilotV1AssistantFieldTypeFieldValue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFieldValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the new resource.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldValueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60;
 **SynonymOf** | **string** | The string value that indicates which word the field value is a synonym of.
 **Value** | **string** | The Field Value data.

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

> AutopilotV1AssistantModelBuild CreateModelBuild(ctx, AssistantSid).StatusCallback(StatusCallback).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
    StatusCallback := "StatusCallback_example" // string | The URL we should call using a POST method to send status information to your application. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateModelBuild(context.Background(), AssistantSid).StatusCallback(StatusCallback).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateModelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateModelBuild`: AutopilotV1AssistantModelBuild
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateModelBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateModelBuildParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **StatusCallback** | **string** | The URL we should call using a POST method to send status information to your application.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

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

> AutopilotV1AssistantQuery CreateQuery(ctx, AssistantSid).Language(Language).ModelBuild(ModelBuild).Query(Query).Tasks(Tasks).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: `en-US`. (optional)
    ModelBuild := "ModelBuild_example" // string | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. (optional)
    Query := "Query_example" // string | The end-user's natural language input. It can be up to 2048 characters long. (optional)
    Tasks := "Tasks_example" // string | The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task `unique_name` values. For example, `task-unique_name-1, task-unique_name-2`. Listing specific tasks is useful to constrain the paths that a user can take. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateQuery(context.Background(), AssistantSid).Language(Language).ModelBuild(ModelBuild).Query(Query).Tasks(Tasks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuery`: AutopilotV1AssistantQuery
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateQueryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: &#x60;en-US&#x60;.
 **ModelBuild** | **string** | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried.
 **Query** | **string** | The end-user&#39;s natural language input. It can be up to 2048 characters long.
 **Tasks** | **string** | The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task &#x60;unique_name&#x60; values. For example, &#x60;task-unique_name-1, task-unique_name-2&#x60;. Listing specific tasks is useful to constrain the paths that a user can take.

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

> AutopilotV1AssistantTaskSample CreateSample(ctx, AssistantSid, TaskSid).Language(Language).SourceChannel(SourceChannel).TaggedText(TaggedText).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: `en-US`. (optional)
    SourceChannel := "SourceChannel_example" // string | The communication channel from which the new sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included. (optional)
    TaggedText := "TaggedText_example" // string | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSample(context.Background(), AssistantSid, TaskSid).Language(Language).SourceChannel(SourceChannel).TaggedText(TaggedText).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSample`: AutopilotV1AssistantTaskSample
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the new resource.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.

### Other Parameters

Other parameters are passed through a pointer to a CreateSampleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: &#x60;en-US&#x60;.
 **SourceChannel** | **string** | The communication channel from which the new sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included.
 **TaggedText** | **string** | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).

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

> AutopilotV1AssistantTask CreateTask(ctx, AssistantSid).Actions(Actions).ActionsUrl(ActionsUrl).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
    Actions := TODO // map[string]interface{} | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique. (optional)
    ActionsUrl := "ActionsUrl_example" // string | The URL from which the Assistant can fetch actions. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTask(context.Background(), AssistantSid).Actions(Actions).ActionsUrl(ActionsUrl).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTask`: AutopilotV1AssistantTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Actions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique.
 **ActionsUrl** | **string** | The URL from which the Assistant can fetch actions.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length.

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

> AutopilotV1AssistantWebhook CreateWebhook(ctx, AssistantSid).Events(Events).UniqueName(UniqueName).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.
    Events := "Events_example" // string | The list of space-separated events that this Webhook will subscribe to. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length. (optional)
    WebhookMethod := "WebhookMethod_example" // string | The method to be used when calling the webhook's URL. (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL associated with this Webhook. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWebhook(context.Background(), AssistantSid).Events(Events).UniqueName(UniqueName).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: AutopilotV1AssistantWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Events** | **string** | The list of space-separated events that this Webhook will subscribe to.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length.
 **WebhookMethod** | **string** | The method to be used when calling the webhook&#39;s URL.
 **WebhookUrl** | **string** | The URL associated with this Webhook.

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

> DeleteAssistant(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Assistant resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAssistant(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantParams struct via the builder pattern


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


## DeleteField

> DeleteField(ctx, AssistantSid, TaskSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Field resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteField(context.Background(), AssistantSid, TaskSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Field resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldParams struct via the builder pattern


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


## DeleteFieldType

> DeleteFieldType(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the FieldType resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFieldType(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFieldType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldTypeParams struct via the builder pattern


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


## DeleteFieldValue

> DeleteFieldValue(ctx, AssistantSid, FieldTypeSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to delete.
    FieldTypeSid := "FieldTypeSid_example" // string | The SID of the Field Type associated with the Field Value to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the FieldValue resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFieldValue(context.Background(), AssistantSid, FieldTypeSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFieldValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to delete.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldValue resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldValueParams struct via the builder pattern


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


## DeleteModelBuild

> DeleteModelBuild(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ModelBuild resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteModelBuild(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteModelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteModelBuildParams struct via the builder pattern


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


## DeleteQuery

> DeleteQuery(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Query resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteQuery(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteQueryParams struct via the builder pattern


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


## DeleteSample

> DeleteSample(ctx, AssistantSid, TaskSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Sample resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSample(context.Background(), AssistantSid, TaskSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to delete.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSampleParams struct via the builder pattern


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


## DeleteTask

> DeleteTask(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Task resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTask(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Task resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskParams struct via the builder pattern


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


## DeleteWebhook

> DeleteWebhook(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhook(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebhookParams struct via the builder pattern


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


## FetchAssistant

> AutopilotV1Assistant FetchAssistant(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Assistant resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAssistant(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAssistant`: AutopilotV1Assistant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAssistant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> AutopilotV1AssistantDefaults FetchDefaults(ctx, AssistantSid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDefaults(context.Background(), AssistantSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDefaults`: AutopilotV1AssistantDefaults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDefaultsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> AutopilotV1AssistantDialogue FetchDialogue(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Dialogue resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDialogue(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDialogue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDialogue`: AutopilotV1AssistantDialogue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDialogue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Dialogue resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDialogueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantTaskField FetchField(ctx, AssistantSid, TaskSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Field resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchField(context.Background(), AssistantSid, TaskSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchField`: AutopilotV1AssistantTaskField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Field resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> AutopilotV1AssistantFieldType FetchFieldType(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the FieldType resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFieldType(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFieldType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFieldType`: AutopilotV1AssistantFieldType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFieldType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantFieldTypeFieldValue FetchFieldValue(ctx, AssistantSid, FieldTypeSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource to fetch.
    FieldTypeSid := "FieldTypeSid_example" // string | The SID of the Field Type associated with the Field Value to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the FieldValue resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFieldValue(context.Background(), AssistantSid, FieldTypeSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFieldValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFieldValue`: AutopilotV1AssistantFieldTypeFieldValue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFieldValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource to fetch.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldValue resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldValueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> AutopilotV1AssistantModelBuild FetchModelBuild(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ModelBuild resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchModelBuild(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchModelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchModelBuild`: AutopilotV1AssistantModelBuild
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchModelBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchModelBuildParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantQuery FetchQuery(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Query resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchQuery(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchQuery`: AutopilotV1AssistantQuery
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantTaskSample FetchSample(ctx, AssistantSid, TaskSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Sample resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSample(context.Background(), AssistantSid, TaskSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSample`: AutopilotV1AssistantTaskSample
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to create.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSampleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> AutopilotV1AssistantStyleSheet FetchStyleSheet(ctx, AssistantSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchStyleSheet(context.Background(), AssistantSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchStyleSheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStyleSheet`: AutopilotV1AssistantStyleSheet
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchStyleSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchStyleSheetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> AutopilotV1AssistantTask FetchTask(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Task resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTask(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTask`: AutopilotV1AssistantTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Task resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantTaskTaskActions FetchTaskActions(ctx, AssistantSid, TaskSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to fetch were defined.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to fetch were defined.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskActions(context.Background(), AssistantSid, TaskSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskActions`: AutopilotV1AssistantTaskTaskActions
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to fetch were defined.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to fetch were defined.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskActionsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantTaskTaskStatistics FetchTaskStatistics(ctx, AssistantSid, TaskSid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) that is associated with the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskStatistics(context.Background(), AssistantSid, TaskSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskStatistics`: AutopilotV1AssistantTaskTaskStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) that is associated with the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> AutopilotV1AssistantWebhook FetchWebhook(ctx, AssistantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWebhook(context.Background(), AssistantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWebhook`: AutopilotV1AssistantWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> ListAssistantResponse ListAssistant(ctx).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAssistant(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssistant`: ListAssistantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAssistant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAssistantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListFieldResponse ListField(ctx, AssistantSid, TaskSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListField(context.Background(), AssistantSid, TaskSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListField`: ListFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with the Field resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListFieldTypeResponse ListFieldType(ctx, AssistantSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFieldType(context.Background(), AssistantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFieldType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFieldType`: ListFieldTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFieldType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListFieldValueResponse ListFieldValue(ctx, AssistantSid, FieldTypeSid).Language(Language).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to read.
    FieldTypeSid := "FieldTypeSid_example" // string | The SID of the Field Type associated with the Field Value to read.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US` (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFieldValue(context.Background(), AssistantSid, FieldTypeSid).Language(Language).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFieldValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFieldValue`: ListFieldValueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFieldValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to read.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldValueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: &#x60;en-US&#x60;
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListModelBuildResponse ListModelBuild(ctx, AssistantSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListModelBuild(context.Background(), AssistantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListModelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListModelBuild`: ListModelBuildResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListModelBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListModelBuildParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListQueryResponse ListQuery(ctx, AssistantSid).Language(Language).ModelBuild(ModelBuild).Status(Status).DialogueSid(DialogueSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query resources to read. For example: `en-US`. (optional)
    ModelBuild := "ModelBuild_example" // string | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried. (optional)
    Status := "Status_example" // string | The status of the resources to read. Can be: `pending-review`, `reviewed`, or `discarded` (optional)
    DialogueSid := "DialogueSid_example" // string | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListQuery(context.Background(), AssistantSid).Language(Language).ModelBuild(ModelBuild).Status(Status).DialogueSid(DialogueSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQuery`: ListQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListQueryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query resources to read. For example: &#x60;en-US&#x60;.
 **ModelBuild** | **string** | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried.
 **Status** | **string** | The status of the resources to read. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60;
 **DialogueSid** | **string** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSampleResponse ListSample(ctx, AssistantSid, TaskSid).Language(Language).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resources to read.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSample(context.Background(), AssistantSid, TaskSid).Language(Language).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSample`: ListSampleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resources to read.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSampleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListTaskResponse ListTask(ctx, AssistantSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTask(context.Background(), AssistantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTask`: ListTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListWebhookResponse ListWebhook(ctx, AssistantSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWebhook(context.Background(), AssistantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhook`: ListWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> AutopilotV1Assistant UpdateAssistant(ctx, Sid).CallbackEvents(CallbackEvents).CallbackUrl(CallbackUrl).Defaults(Defaults).DevelopmentStage(DevelopmentStage).FriendlyName(FriendlyName).LogQueries(LogQueries).StyleSheet(StyleSheet).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Assistant resource to update.
    CallbackEvents := "CallbackEvents_example" // string | Reserved. (optional)
    CallbackUrl := "CallbackUrl_example" // string | Reserved. (optional)
    Defaults := TODO // map[string]interface{} | A JSON object that defines the Assistant's [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks. (optional)
    DevelopmentStage := "DevelopmentStage_example" // string | A string describing the state of the assistant. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    LogQueries := true // bool | Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored. (optional)
    StyleSheet := TODO // map[string]interface{} | The JSON string that defines the Assistant's [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet) (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAssistant(context.Background(), Sid).CallbackEvents(CallbackEvents).CallbackUrl(CallbackUrl).Defaults(Defaults).DevelopmentStage(DevelopmentStage).FriendlyName(FriendlyName).LogQueries(LogQueries).StyleSheet(StyleSheet).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssistant`: AutopilotV1Assistant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAssistant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CallbackEvents** | **string** | Reserved.
 **CallbackUrl** | **string** | Reserved.
 **Defaults** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
 **DevelopmentStage** | **string** | A string describing the state of the assistant.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **LogQueries** | **bool** | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored.
 **StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

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

> AutopilotV1AssistantDefaults UpdateDefaults(ctx, AssistantSid).Defaults(Defaults).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    Defaults := TODO // map[string]interface{} | A JSON string that describes the default task links for the `assistant_initiation`, `collect`, and `fallback` situations. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateDefaults(context.Background(), AssistantSid).Defaults(Defaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaults`: AutopilotV1AssistantDefaults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDefaultsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Defaults** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that describes the default task links for the &#x60;assistant_initiation&#x60;, &#x60;collect&#x60;, and &#x60;fallback&#x60; situations.

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

> AutopilotV1AssistantFieldType UpdateFieldType(ctx, AssistantSid, Sid).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the FieldType resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFieldType(context.Background(), AssistantSid, Sid).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFieldType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFieldType`: AutopilotV1AssistantFieldType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFieldType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFieldTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

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

> AutopilotV1AssistantModelBuild UpdateModelBuild(ctx, AssistantSid, Sid).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ModelBuild resource to update.
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateModelBuild(context.Background(), AssistantSid, Sid).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateModelBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateModelBuild`: AutopilotV1AssistantModelBuild
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateModelBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateModelBuildParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

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

> AutopilotV1AssistantQuery UpdateQuery(ctx, AssistantSid, Sid).SampleSid(SampleSid).Status(Status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Query resource to update.
    SampleSid := "SampleSid_example" // string | The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query. (optional)
    Status := "Status_example" // string | The new status of the resource. Can be: `pending-review`, `reviewed`, or `discarded` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateQuery(context.Background(), AssistantSid, Sid).SampleSid(SampleSid).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuery`: AutopilotV1AssistantQuery
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateQueryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **SampleSid** | **string** | The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query.
 **Status** | **string** | The new status of the resource. Can be: &#x60;pending-review&#x60;, &#x60;reviewed&#x60;, or &#x60;discarded&#x60;

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

> AutopilotV1RestoreAssistant UpdateRestoreAssistant(ctx).Assistant(Assistant).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Assistant := "Assistant_example" // string | The Twilio-provided string that uniquely identifies the Assistant resource to restore. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRestoreAssistant(context.Background()).Assistant(Assistant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRestoreAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRestoreAssistant`: AutopilotV1RestoreAssistant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRestoreAssistant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRestoreAssistantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Assistant** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to restore.

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

> AutopilotV1AssistantTaskSample UpdateSample(ctx, AssistantSid, TaskSid, Sid).Language(Language).SourceChannel(SourceChannel).TaggedText(TaggedText).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to update.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Sample resource to update.
    Language := "Language_example" // string | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`. (optional)
    SourceChannel := "SourceChannel_example" // string | The communication channel from which the sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included. (optional)
    TaggedText := "TaggedText_example" // string | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSample(context.Background(), AssistantSid, TaskSid, Sid).Language(Language).SourceChannel(SourceChannel).TaggedText(TaggedText).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSample``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSample`: AutopilotV1AssistantTaskSample
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSample`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task associated with the resource to update.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) associated with the Sample resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Sample resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSampleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: &#x60;en-US&#x60;.
 **SourceChannel** | **string** | The communication channel from which the sample was captured. Can be: &#x60;voice&#x60;, &#x60;sms&#x60;, &#x60;chat&#x60;, &#x60;alexa&#x60;, &#x60;google-assistant&#x60;, &#x60;slack&#x60;, or null if not included.
 **TaggedText** | **string** | The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).

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

> AutopilotV1AssistantStyleSheet UpdateStyleSheet(ctx, AssistantSid).StyleSheet(StyleSheet).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    StyleSheet := TODO // map[string]interface{} | The JSON string that describes the style sheet object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateStyleSheet(context.Background(), AssistantSid).StyleSheet(StyleSheet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStyleSheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStyleSheet`: AutopilotV1AssistantStyleSheet
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateStyleSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateStyleSheetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that describes the style sheet object.

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

> AutopilotV1AssistantTask UpdateTask(ctx, AssistantSid, Sid).Actions(Actions).ActionsUrl(ActionsUrl).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Task resource to update.
    Actions := TODO // map[string]interface{} | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. (optional)
    ActionsUrl := "ActionsUrl_example" // string | The URL from which the Assistant can fetch actions. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTask(context.Background(), AssistantSid, Sid).Actions(Actions).ActionsUrl(ActionsUrl).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: AutopilotV1AssistantTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Task resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Actions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.
 **ActionsUrl** | **string** | The URL from which the Assistant can fetch actions.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

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

> AutopilotV1AssistantTaskTaskActions UpdateTaskActions(ctx, AssistantSid, TaskSid).Actions(Actions).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to update were defined.
    TaskSid := "TaskSid_example" // string | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to update were defined.
    Actions := TODO // map[string]interface{} | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTaskActions(context.Background(), AssistantSid, TaskSid).Actions(Actions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTaskActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskActions`: AutopilotV1AssistantTaskTaskActions
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTaskActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the Task for which the task actions to update were defined.
**TaskSid** | **string** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) for which the task actions to update were defined.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskActionsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Actions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.

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

> AutopilotV1AssistantWebhook UpdateWebhook(ctx, AssistantSid, Sid).Events(Events).UniqueName(UniqueName).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AssistantSid := "AssistantSid_example" // string | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to update.
    Events := "Events_example" // string | The list of space-separated events that this Webhook will subscribe to. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length. (optional)
    WebhookMethod := "WebhookMethod_example" // string | The method to be used when calling the webhook's URL. (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL associated with this Webhook. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebhook(context.Background(), AssistantSid, Sid).Events(Events).UniqueName(UniqueName).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: AutopilotV1AssistantWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Events** | **string** | The list of space-separated events that this Webhook will subscribe to.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. This value must be unique and 64 characters or less in length.
 **WebhookMethod** | **string** | The method to be used when calling the webhook&#39;s URL.
 **WebhookUrl** | **string** | The URL associated with this Webhook.

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

