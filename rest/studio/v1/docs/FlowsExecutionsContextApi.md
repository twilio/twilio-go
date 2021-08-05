# FlowsExecutionsContextApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchExecutionContext**](FlowsExecutionsContextApi.md#FetchExecutionContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 



## FetchExecutionContext

> StudioV1ExecutionContext FetchExecutionContext(ctx, FlowSidExecutionSid)



Retrieve the most recent context for an Execution.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution context to fetch.
**ExecutionSid** | **string** | The SID of the Execution context to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1ExecutionContext**](StudioV1ExecutionContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

