# DefaultApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngagement**](DefaultApi.md#CreateEngagement) | **Post** /v1/Flows/{FlowSid}/Engagements | 
[**CreateExecution**](DefaultApi.md#CreateExecution) | **Post** /v1/Flows/{FlowSid}/Executions | 
[**DeleteEngagement**](DefaultApi.md#DeleteEngagement) | **Delete** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**DeleteExecution**](DefaultApi.md#DeleteExecution) | **Delete** /v1/Flows/{FlowSid}/Executions/{Sid} | 
[**DeleteFlow**](DefaultApi.md#DeleteFlow) | **Delete** /v1/Flows/{Sid} | 
[**FetchEngagement**](DefaultApi.md#FetchEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**FetchEngagementContext**](DefaultApi.md#FetchEngagementContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Context | 
[**FetchExecution**](DefaultApi.md#FetchExecution) | **Get** /v1/Flows/{FlowSid}/Executions/{Sid} | 
[**FetchExecutionContext**](DefaultApi.md#FetchExecutionContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 
[**FetchExecutionStep**](DefaultApi.md#FetchExecutionStep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
[**FetchExecutionStepContext**](DefaultApi.md#FetchExecutionStepContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 
[**FetchFlow**](DefaultApi.md#FetchFlow) | **Get** /v1/Flows/{Sid} | 
[**FetchStep**](DefaultApi.md#FetchStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid} | 
[**FetchStepContext**](DefaultApi.md#FetchStepContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{StepSid}/Context | 
[**ListEngagement**](DefaultApi.md#ListEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements | 
[**ListExecution**](DefaultApi.md#ListExecution) | **Get** /v1/Flows/{FlowSid}/Executions | 
[**ListExecutionStep**](DefaultApi.md#ListExecutionStep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 
[**ListFlow**](DefaultApi.md#ListFlow) | **Get** /v1/Flows | 
[**ListStep**](DefaultApi.md#ListStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps | 
[**UpdateExecution**](DefaultApi.md#UpdateExecution) | **Post** /v1/Flows/{FlowSid}/Executions/{Sid} | 



## CreateEngagement

> StudioV1FlowEngagement CreateEngagement(ctx, FlowSidoptional)



Triggers a new Engagement for the Flow

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable &#x60;{{flow.channel.address}}&#x60;
**Parameters** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string we will add to your flow&#39;s context and that you can access as variables inside your flow. For example, if you pass in &#x60;Parameters&#x3D;{&#39;name&#39;:&#39;Zeke&#39;}&#x60; then inside a widget you can reference the variable &#x60;{{flow.data.name}}&#x60; which will return the string &#39;Zeke&#39;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
**To** | **string** | The Contact phone number to start a Studio Flow Engagement, available as variable &#x60;{{contact.channel.address}}&#x60;.

### Return type

[**StudioV1FlowEngagement**](StudioV1FlowEngagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExecution

> StudioV1FlowExecution CreateExecution(ctx, FlowSidoptional)



Triggers a new Execution for the Flow

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Excecution&#39;s Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;.
**Parameters** | [**map[string]interface{}**](map[string]interface{}.md) | JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
**To** | **string** | The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;.

### Return type

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngagement

> DeleteEngagement(ctx, FlowSidSid)



Delete this Engagement and all Steps relating to it.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to delete Engagements from.
**Sid** | **string** | The SID of the Engagement resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEngagementParams struct


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


## DeleteFlow

> DeleteFlow(ctx, Sid)



Delete a specific Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFlowParams struct


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


## FetchEngagement

> StudioV1FlowEngagement FetchEngagement(ctx, FlowSidSid)



Retrieve an Engagement

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**Sid** | **string** | The SID of the Engagement resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1FlowEngagement**](StudioV1FlowEngagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEngagementContext

> StudioV1FlowEngagementEngagementContext FetchEngagementContext(ctx, FlowSidEngagementSid)



Retrieve the most recent context for an Engagement.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**EngagementSid** | **string** | The SID of the Engagement.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1FlowEngagementEngagementContext**](StudioV1FlowEngagementEngagementContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecution

> StudioV1FlowExecution FetchExecution(ctx, FlowSidSid)



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

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV1FlowExecutionExecutionContext FetchExecutionContext(ctx, FlowSidExecutionSid)



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

[**StudioV1FlowExecutionExecutionContext**](StudioV1FlowExecutionExecutionContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV1FlowExecutionExecutionStep FetchExecutionStep(ctx, FlowSidExecutionSidSid)



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

[**StudioV1FlowExecutionExecutionStep**](StudioV1FlowExecutionExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV1FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, FlowSidExecutionSidStepSid)



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

[**StudioV1FlowExecutionExecutionStepExecutionStepContext**](StudioV1FlowExecutionExecutionStepExecutionStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlow

> StudioV1Flow FetchFlow(ctx, Sid)



Retrieve a specific Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlowParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1Flow**](StudioV1Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStep

> StudioV1FlowEngagementStep FetchStep(ctx, FlowSidEngagementSidSid)



Retrieve a Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**Sid** | **string** | The SID of the Step resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchStepParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1FlowEngagementStep**](StudioV1FlowEngagementStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStepContext

> StudioV1FlowEngagementStepStepContext FetchStepContext(ctx, FlowSidEngagementSidStepSid)



Retrieve the context for an Engagement Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**StepSid** | **string** | The SID of the Step to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchStepContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1FlowEngagementStepStepContext**](StudioV1FlowEngagementStepStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEngagement

> ListEngagementResponse ListEngagement(ctx, FlowSidoptional)



Retrieve a list of all Engagements for the Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to read Engagements from.

### Other Parameters

Other parameters are passed through a pointer to a ListEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEngagementResponse**](ListEngagementResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecution

> ListExecutionResponse ListExecution(ctx, FlowSidoptional)



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
**DateCreatedFrom** | **time.Time** | Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;.
**DateCreatedTo** | **time.Time** | Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListExecutionResponse**](ListExecutionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutionStep

> ListExecutionStepResponse ListExecutionStep(ctx, FlowSidExecutionSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListExecutionStepResponse**](ListExecutionStepResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlow

> ListFlowResponse ListFlow(ctx, optional)



Retrieve a list of all Flows.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFlowResponse**](ListFlowResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStep

> ListStepResponse ListStep(ctx, FlowSidEngagementSidoptional)



Retrieve a list of all Steps for an Engagement.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to read.
**EngagementSid** | **string** | The SID of the Engagement with the Step to read.

### Other Parameters

Other parameters are passed through a pointer to a ListStepParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListStepResponse**](ListStepResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV1FlowExecution UpdateExecution(ctx, FlowSidSidoptional)



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
**Status** | **string** | The status of the Execution. Can only be &#x60;ended&#x60;.

### Return type

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

