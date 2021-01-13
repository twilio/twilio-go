# \DefaultApi

All URIs are relative to *http://localhost*

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

> StudioV1FlowEngagement CreateEngagement(ctx, flowSid, optional)



Triggers a new Engagement for the Flow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow. | 
 **optional** | ***CreateEngagementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEngagementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable &#x60;{{flow.channel.address}}&#x60; | 
 **parameters** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string we will add to your flow&#39;s context and that you can access as variables inside your flow. For example, if you pass in &#x60;Parameters&#x3D;{&#39;name&#39;:&#39;Zeke&#39;}&#x60; then inside a widget you can reference the variable &#x60;{{flow.data.name}}&#x60; which will return the string &#39;Zeke&#39;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string. | 
 **to** | **optional.String**| The Contact phone number to start a Studio Flow Engagement, available as variable &#x60;{{contact.channel.address}}&#x60;. | 

### Return type

[**StudioV1FlowEngagement**](studio.v1.flow.engagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExecution

> StudioV1FlowExecution CreateExecution(ctx, flowSid, optional)



Triggers a new Execution for the Flow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Excecution&#39;s Flow. | 
 **optional** | ***CreateExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;. | 
 **parameters** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string. | 
 **to** | **optional.String**| The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;. | 

### Return type

[**StudioV1FlowExecution**](studio.v1.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngagement

> DeleteEngagement(ctx, flowSid, sid)



Delete this Engagement and all Steps relating to it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow to delete Engagements from. | 
**sid** | **string**| The SID of the Engagement resource to delete. | 

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

> DeleteExecution(ctx, flowSid, sid)



Delete the Execution and all Steps relating to it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Execution resources to delete. | 
**sid** | **string**| The SID of the Execution resource to delete. | 

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

> DeleteFlow(ctx, sid)



Delete a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to delete. | 

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

> StudioV1FlowEngagement FetchEngagement(ctx, flowSid, sid)



Retrieve an Engagement

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow. | 
**sid** | **string**| The SID of the Engagement resource to fetch. | 

### Return type

[**StudioV1FlowEngagement**](studio.v1.flow.engagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEngagementContext

> StudioV1FlowEngagementEngagementContext FetchEngagementContext(ctx, flowSid, engagementSid)



Retrieve the most recent context for an Engagement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow. | 
**engagementSid** | **string**| The SID of the Engagement. | 

### Return type

[**StudioV1FlowEngagementEngagementContext**](studio.v1.flow.engagement.engagement_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecution

> StudioV1FlowExecution FetchExecution(ctx, flowSid, sid)



Retrieve an Execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Execution resource to fetch | 
**sid** | **string**| The SID of the Execution resource to fetch. | 

### Return type

[**StudioV1FlowExecution**](studio.v1.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV1FlowExecutionExecutionContext FetchExecutionContext(ctx, flowSid, executionSid)



Retrieve the most recent context for an Execution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Execution context to fetch. | 
**executionSid** | **string**| The SID of the Execution context to fetch. | 

### Return type

[**StudioV1FlowExecutionExecutionContext**](studio.v1.flow.execution.execution_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV1FlowExecutionExecutionStep FetchExecutionStep(ctx, flowSid, executionSid, sid)



Retrieve a Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**executionSid** | **string**| The SID of the Execution resource with the Step to fetch. | 
**sid** | **string**| The SID of the ExecutionStep resource to fetch. | 

### Return type

[**StudioV1FlowExecutionExecutionStep**](studio.v1.flow.execution.execution_step.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV1FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, flowSid, executionSid, stepSid)



Retrieve the context for an Execution Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**executionSid** | **string**| The SID of the Execution resource with the Step to fetch. | 
**stepSid** | **string**| The SID of the Step to fetch. | 

### Return type

[**StudioV1FlowExecutionExecutionStepExecutionStepContext**](studio.v1.flow.execution.execution_step.execution_step_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlow

> StudioV1Flow FetchFlow(ctx, sid)



Retrieve a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to fetch. | 

### Return type

[**StudioV1Flow**](studio.v1.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStep

> StudioV1FlowEngagementStep FetchStep(ctx, flowSid, engagementSid, sid)



Retrieve a Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**engagementSid** | **string**| The SID of the Engagement with the Step to fetch. | 
**sid** | **string**| The SID of the Step resource to fetch. | 

### Return type

[**StudioV1FlowEngagementStep**](studio.v1.flow.engagement.step.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStepContext

> StudioV1FlowEngagementStepStepContext FetchStepContext(ctx, flowSid, engagementSid, stepSid)



Retrieve the context for an Engagement Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**engagementSid** | **string**| The SID of the Engagement with the Step to fetch. | 
**stepSid** | **string**| The SID of the Step to fetch | 

### Return type

[**StudioV1FlowEngagementStepStepContext**](studio.v1.flow.engagement.step.step_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEngagement

> StudioV1FlowEngagementReadResponse ListEngagement(ctx, flowSid, optional)



Retrieve a list of all Engagements for the Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow to read Engagements from. | 
 **optional** | ***ListEngagementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEngagementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV1FlowEngagementReadResponse**](studio_v1_flow_engagementReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecution

> StudioV1FlowExecutionReadResponse ListExecution(ctx, flowSid, optional)



Retrieve a list of all Executions for the Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Execution resources to read. | 
 **optional** | ***ListExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateCreatedFrom** | **optional.Time**| Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;. | 
 **dateCreatedTo** | **optional.Time**| Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV1FlowExecutionReadResponse**](studio_v1_flow_executionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutionStep

> StudioV1FlowExecutionExecutionStepReadResponse ListExecutionStep(ctx, flowSid, executionSid, optional)



Retrieve a list of all Steps for an Execution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Steps to read. | 
**executionSid** | **string**| The SID of the Execution with the Steps to read. | 
 **optional** | ***ListExecutionStepOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExecutionStepOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV1FlowExecutionExecutionStepReadResponse**](studio_v1_flow_execution_execution_stepReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlow

> StudioV1FlowReadResponse ListFlow(ctx, optional)



Retrieve a list of all Flows.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV1FlowReadResponse**](studio_v1_flowReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStep

> StudioV1FlowEngagementStepReadResponse ListStep(ctx, flowSid, engagementSid, optional)



Retrieve a list of all Steps for an Engagement.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Step to read. | 
**engagementSid** | **string**| The SID of the Engagement with the Step to read. | 
 **optional** | ***ListStepOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListStepOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV1FlowEngagementStepReadResponse**](studio_v1_flow_engagement_stepReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV1FlowExecution UpdateExecution(ctx, flowSid, sid, optional)



Update the status of an Execution to `ended`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSid** | **string**| The SID of the Flow with the Execution resources to update. | 
**sid** | **string**| The SID of the Execution resource to update. | 
 **optional** | ***UpdateExecutionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExecutionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.String**| The status of the Execution. Can only be &#x60;ended&#x60;. | 

### Return type

[**StudioV1FlowExecution**](studio.v1.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

