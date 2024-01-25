# AssistantsQueriesApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuery**](AssistantsQueriesApi.md#CreateQuery) | **Post** /v1/Assistants/{AssistantSid}/Queries | 
[**DeleteQuery**](AssistantsQueriesApi.md#DeleteQuery) | **Delete** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 
[**FetchQuery**](AssistantsQueriesApi.md#FetchQuery) | **Get** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 
[**ListQuery**](AssistantsQueriesApi.md#ListQuery) | **Get** /v1/Assistants/{AssistantSid}/Queries | 
[**UpdateQuery**](AssistantsQueriesApi.md#UpdateQuery) | **Post** /v1/Assistants/{AssistantSid}/Queries/{Sid} | 



## CreateQuery

> AutopilotV1Query CreateQuery(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateQueryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new query. For example: `en-US`.
**Query** | **string** | The end-user's natural language input. It can be up to 2048 characters long.
**Tasks** | **string** | The list of tasks to limit the new query to. Tasks are expressed as a comma-separated list of task `unique_name` values. For example, `task-unique_name-1, task-unique_name-2`. Listing specific tasks is useful to constrain the paths that a user can take.
**ModelBuild** | **string** | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried.

### Return type

[**AutopilotV1Query**](AutopilotV1Query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuery

> DeleteQuery(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteQueryParams struct


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


## FetchQuery

> AutopilotV1Query FetchQuery(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Query**](AutopilotV1Query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuery

> []AutopilotV1Query ListQuery(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListQueryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used by the Query resources to read. For example: `en-US`.
**ModelBuild** | **string** | The SID or unique name of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) to be queried.
**Status** | **string** | The status of the resources to read. Can be: `pending-review`, `reviewed`, or `discarded`
**DialogueSid** | **string** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue).
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1Query**](AutopilotV1Query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuery

> AutopilotV1Query UpdateQuery(ctx, AssistantSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Query resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateQueryParams struct


Name | Type | Description
------------- | ------------- | -------------
**SampleSid** | **string** | The SID of an optional reference to the [Sample](https://www.twilio.com/docs/autopilot/api/task-sample) created from the query.
**Status** | **string** | The new status of the resource. Can be: `pending-review`, `reviewed`, or `discarded`

### Return type

[**AutopilotV1Query**](AutopilotV1Query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

