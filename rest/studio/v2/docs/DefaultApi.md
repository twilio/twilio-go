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

> StudioV2FlowExecution CreateExecution(ctx, FlowSidoptional)



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
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;. For SMS, this can also be a Messaging Service SID.
**Parameters** | [**map[string]interface{}**](map[string]interface{}.md) | JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
**To** | **string** | The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;.

### Return type

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

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

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**CommitMessage** | **string** | Description of change made in the revision.
**Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2Flow**](StudioV2Flow.md)

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


## FetchExecution

> StudioV2FlowExecution FetchExecution(ctx, FlowSidSid)



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

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV2FlowExecutionExecutionContext FetchExecutionContext(ctx, FlowSidExecutionSid)



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

[**StudioV2FlowExecutionExecutionContext**](StudioV2FlowExecutionExecutionContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV2FlowExecutionExecutionStep FetchExecutionStep(ctx, FlowSidExecutionSidSid)



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

[**StudioV2FlowExecutionExecutionStep**](StudioV2FlowExecutionExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV2FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, FlowSidExecutionSidStepSid)



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

[**StudioV2FlowExecutionExecutionStepExecutionStepContext**](StudioV2FlowExecutionExecutionStepExecutionStepContext.md)

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

[**StudioV2Flow**](StudioV2Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlowRevision

> StudioV2FlowFlowRevision FetchFlowRevision(ctx, SidRevision)



Retrieve a specific Flow revision.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.
**Revision** | **string** | Specific Revision number or can be &#x60;LatestPublished&#x60; and &#x60;LatestRevision&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlowRevisionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2FlowFlowRevision**](StudioV2FlowFlowRevision.md)

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

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a FetchTestUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2FlowTestUser**](StudioV2FlowTestUser.md)

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


## ListFlowRevision

> ListFlowRevisionResponse ListFlowRevision(ctx, Sidoptional)



Retrieve a list of all Flows revisions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowRevisionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> StudioV2FlowExecution UpdateExecution(ctx, FlowSidSidoptional)



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

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlow

> StudioV2Flow UpdateFlow(ctx, Sidoptional)



Update a Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**CommitMessage** | **string** | Description of change made in the revision.
**Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2Flow**](StudioV2Flow.md)

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

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlowValidateParams struct


Name | Type | Description
------------- | ------------- | -------------
**CommitMessage** | **string** | Description of change made in the revision.
**Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2FlowValidate**](StudioV2FlowValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTestUser

> StudioV2FlowTestUser UpdateTestUser(ctx, Sidoptional)



Update flow test users

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTestUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**TestUsers** | **[]string** | List of test user identities that can test draft versions of the flow.

### Return type

[**StudioV2FlowTestUser**](StudioV2FlowTestUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

