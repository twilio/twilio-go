# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExecution**](DefaultApi.md#CreateExecution) | **Post** /v2/Flows/{FlowSid}/Executions | 
[**CreateFlow**](DefaultApi.md#CreateFlow) | **Post** /v2/Flows | 
[**DeleteExecution**](DefaultApi.md#DeleteExecution) | **Delete** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**DeleteFlow**](DefaultApi.md#DeleteFlow) | **Delete** /v2/Flows/{Sid} | 
[**FetchExecution**](DefaultApi.md#FetchExecution) | **Get** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**FetchExecutionContext**](DefaultApi.md#FetchExecutionContext) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 
[**FetchExecutionStep**](DefaultApi.md#FetchExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
[**FetchExecutionStepContext**](DefaultApi.md#FetchExecutionStepContext) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 
[**FetchFlow**](DefaultApi.md#FetchFlow) | **Get** /v2/Flows/{Sid} | 
[**FetchFlowRevision**](DefaultApi.md#FetchFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions/{Revision} | 
[**FetchTestUser**](DefaultApi.md#FetchTestUser) | **Get** /v2/Flows/{Sid}/TestUsers | 
[**ListExecution**](DefaultApi.md#ListExecution) | **Get** /v2/Flows/{FlowSid}/Executions | 
[**ListExecutionStep**](DefaultApi.md#ListExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 
[**ListFlow**](DefaultApi.md#ListFlow) | **Get** /v2/Flows | 
[**ListFlowRevision**](DefaultApi.md#ListFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions | 
[**UpdateExecution**](DefaultApi.md#UpdateExecution) | **Post** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**UpdateFlow**](DefaultApi.md#UpdateFlow) | **Post** /v2/Flows/{Sid} | 
[**UpdateFlowValidate**](DefaultApi.md#UpdateFlowValidate) | **Post** /v2/Flows/Validate | 
[**UpdateTestUser**](DefaultApi.md#UpdateTestUser) | **Post** /v2/Flows/{Sid}/TestUsers | 



## CreateExecution

> StudioV2FlowExecution CreateExecution(ctx, FlowSid, optional)



Triggers a new Execution for the Flow

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Excecution&#39;s Flow. | 
 **optional** | ***CreateExecutionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExecutionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**From** | **String**| The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;. | 
**Parameters** | [**map[string]interface{}**](map[string]interface{}.md)| JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string. | 
**To** | **String**| The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;. | 

### Return type

[**StudioV2FlowExecution**](studio.v2.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlow

> StudioV2Flow CreateFlow(ctx, optional)



Create a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFlowRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFlowRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommitMessage** | **String**| Description of change made in the revision. | 
**Definition** | [**map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
**FriendlyName** | **String**| The string that you assigned to describe the Flow. | 
**Status** | **String**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExecution

> DeleteExecution(ctx, FlowSid, Sid)



Delete the Execution and all Steps relating to it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Execution resources to delete. | 
**Sid** | **string**| The SID of the Execution resource to delete. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flow resource to delete. | 

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

> StudioV2FlowExecution FetchExecution(ctx, FlowSid, Sid)



Retrieve an Execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Execution resource to fetch | 
**Sid** | **string**| The SID of the Execution resource to fetch. | 

### Return type

[**StudioV2FlowExecution**](studio.v2.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV2FlowExecutionExecutionContext FetchExecutionContext(ctx, FlowSid, ExecutionSid)



Retrieve the most recent context for an Execution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Execution context to fetch. | 
**ExecutionSid** | **string**| The SID of the Execution context to fetch. | 

### Return type

[**StudioV2FlowExecutionExecutionContext**](studio.v2.flow.execution.execution_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV2FlowExecutionExecutionStep FetchExecutionStep(ctx, FlowSid, ExecutionSid, Sid)



Retrieve a Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**ExecutionSid** | **string**| The SID of the Execution resource with the Step to fetch. | 
**Sid** | **string**| The SID of the ExecutionStep resource to fetch. | 

### Return type

[**StudioV2FlowExecutionExecutionStep**](studio.v2.flow.execution.execution_step.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV2FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, FlowSid, ExecutionSid, StepSid)



Retrieve the context for an Execution Step.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Step to fetch. | 
**ExecutionSid** | **string**| The SID of the Execution resource with the Step to fetch. | 
**StepSid** | **string**| The SID of the Step to fetch. | 

### Return type

[**StudioV2FlowExecutionExecutionStepExecutionStepContext**](studio.v2.flow.execution.execution_step.execution_step_context.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlow

> StudioV2Flow FetchFlow(ctx, Sid)



Retrieve a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flow resource to fetch. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlowRevision

> StudioV2FlowFlowRevision FetchFlowRevision(ctx, Sid, Revision)



Retrieve a specific Flow revision.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flow resource to fetch. | 
**Revision** | **string**| Specific Revision number or can be &#x60;LatestPublished&#x60; and &#x60;LatestRevision&#x60;. | 

### Return type

[**StudioV2FlowFlowRevision**](studio.v2.flow.flow_revision.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTestUser

> StudioV2FlowTestUser FetchTestUser(ctx, Sid)



Fetch flow test users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| Unique identifier of the flow. | 

### Return type

[**StudioV2FlowTestUser**](studio.v2.flow.test_user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecution

> ListExecutionResponse ListExecution(ctx, FlowSid, optional)



Retrieve a list of all Executions for the Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Execution resources to read. | 
 **optional** | ***ListExecutionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExecutionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DateCreatedFrom** | **Time**| Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;. | 
**DateCreatedTo** | **Time**| Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListExecutionStepResponse ListExecutionStep(ctx, FlowSid, ExecutionSid, optional)



Retrieve a list of all Steps for an Execution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Steps to read. | 
**ExecutionSid** | **string**| The SID of the Execution with the Steps to read. | 
 **optional** | ***ListExecutionStepRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExecutionStepRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFlowRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlowRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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


## ListFlowRevision

> ListFlowRevisionResponse ListFlowRevision(ctx, Sid, optional)



Retrieve a list of all Flows revisions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flow resource to fetch. | 
 **optional** | ***ListFlowRevisionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlowRevisionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFlowRevisionResponse**](ListFlowRevisionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV2FlowExecution UpdateExecution(ctx, FlowSid, Sid, optional)



Update the status of an Execution to `ended`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string**| The SID of the Flow with the Execution resources to update. | 
**Sid** | **string**| The SID of the Execution resource to update. | 
 **optional** | ***UpdateExecutionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExecutionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| The status of the Execution. Can only be &#x60;ended&#x60;. | 

### Return type

[**StudioV2FlowExecution**](studio.v2.flow.execution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlow

> StudioV2Flow UpdateFlow(ctx, Sid, optional)



Update a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flow resource to fetch. | 
 **optional** | ***UpdateFlowRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFlowRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommitMessage** | **String**| Description of change made in the revision. | 
**Definition** | [**map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
**FriendlyName** | **String**| The string that you assigned to describe the Flow. | 
**Status** | **String**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlowValidate

> StudioV2FlowValidate UpdateFlowValidate(ctx, optional)



Validate flow JSON definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateFlowValidateRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFlowValidateRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommitMessage** | **String**| Description of change made in the revision. | 
**Definition** | [**map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
**FriendlyName** | **String**| The string that you assigned to describe the Flow. | 
**Status** | **String**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

### Return type

[**StudioV2FlowValidate**](studio.v2.flow_validate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTestUser

> StudioV2FlowTestUser UpdateTestUser(ctx, Sid, optional)



Update flow test users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| Unique identifier of the flow. | 
 **optional** | ***UpdateTestUserRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTestUserRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**TestUsers** | [**[]string**](string.md)| List of test user identities that can test draft versions of the flow. | 

### Return type

[**StudioV2FlowTestUser**](studio.v2.flow.test_user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

