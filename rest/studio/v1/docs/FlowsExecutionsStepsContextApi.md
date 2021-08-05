# FlowsExecutionsStepsContextApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchExecutionStepContext**](FlowsExecutionsStepsContextApi.md#FetchExecutionStepContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 



## FetchExecutionStepContext

> StudioV1ExecutionStepContext FetchExecutionStepContext(ctx, FlowSidExecutionSidStepSid)



Retrieve the context for an Execution Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**ExecutionSid** | **string** | The SID of the Execution resource with the Step to fetch.
**StepSid** | **string** | The SID of the Step to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionStepContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1ExecutionStepContext**](StudioV1ExecutionStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

