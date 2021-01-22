# \FlowsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlow**](FlowsApi.md#CreateFlow) | **Post** /v2/Flows | 
[**ListFlow**](FlowsApi.md#ListFlow) | **Get** /v2/Flows | 



## CreateFlow

> StudioV2Flow CreateFlow(ctx, optional)



Create a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFlowOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CommitMessage** | **optional.String**| Description on change made in the revision. | 
 **Definition** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
 **FriendlyName** | **optional.String**| The string that you assigned to describe the Flow. | 
 **Status** | **optional.String**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

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


## ListFlow

> StudioV2FlowReadResponse ListFlow(ctx, optional)



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
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**StudioV2FlowReadResponse**](studio_v2_flowReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

