# FlowsExecutionsStepsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchExecutionStep**](FlowsExecutionsStepsApi.md#FetchExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
[**ListExecutionStep**](FlowsExecutionsStepsApi.md#ListExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 



## FetchExecutionStep

> StudioV2ExecutionStep FetchExecutionStep(ctx, FlowSidExecutionSidSid)



Retrieve a Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**ExecutionSid** | **string** | The SID of the Execution resource with the Step to fetch.
**Sid** | **string** | The SID of the ExecutionStep resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionStepParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2ExecutionStep**](StudioV2ExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutionStep

> []StudioV2ExecutionStep ListExecutionStep(ctx, FlowSidExecutionSidoptional)



Retrieve a list of all Steps for an Execution.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Steps to read.
**ExecutionSid** | **string** | The SID of the Execution with the Steps to read.

### Other Parameters

Other parameters are passed through a pointer to a ListExecutionStepParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]StudioV2ExecutionStep**](StudioV2ExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

