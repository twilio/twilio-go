# FlowsExecutionsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExecution**](FlowsExecutionsApi.md#CreateExecution) | **Post** /v2/Flows/{FlowSid}/Executions | 
[**DeleteExecution**](FlowsExecutionsApi.md#DeleteExecution) | **Delete** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**FetchExecution**](FlowsExecutionsApi.md#FetchExecution) | **Get** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**ListExecution**](FlowsExecutionsApi.md#ListExecution) | **Get** /v2/Flows/{FlowSid}/Executions | 
[**UpdateExecution**](FlowsExecutionsApi.md#UpdateExecution) | **Post** /v2/Flows/{FlowSid}/Executions/{Sid} | 



## CreateExecution

> StudioV2Execution CreateExecution(ctx, FlowSidoptional)



Triggers a new Execution for the Flow

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Excecution's Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------
**To** | **string** | The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. For SMS, this can also be a Messaging Service SID.
**Parameters** | [**interface{}**](interface{}.md) | JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.

### Return type

[**StudioV2Execution**](StudioV2Execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExecution

> DeleteExecution(ctx, FlowSidSid)



Delete the Execution and all Steps relating to it.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to delete.
**Sid** | **string** | The SID of the Execution resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteExecutionParams struct


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


## FetchExecution

> StudioV2Execution FetchExecution(ctx, FlowSidSid)



Retrieve an Execution

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resource to fetch
**Sid** | **string** | The SID of the Execution resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2Execution**](StudioV2Execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecution

> []StudioV2Execution ListExecution(ctx, FlowSidoptional)



Retrieve a list of all Executions for the Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateCreatedFrom** | **time.Time** | Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
**DateCreatedTo** | **time.Time** | Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]StudioV2Execution**](StudioV2Execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV2Execution UpdateExecution(ctx, FlowSidSidoptional)



Update the status of an Execution to `ended`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to update.
**Sid** | **string** | The SID of the Execution resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 

### Return type

[**StudioV2Execution**](StudioV2Execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

