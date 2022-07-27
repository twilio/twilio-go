# FlowsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlow**](FlowsApi.md#CreateFlow) | **Post** /v2/Flows | 
[**DeleteFlow**](FlowsApi.md#DeleteFlow) | **Delete** /v2/Flows/{Sid} | 
[**FetchFlow**](FlowsApi.md#FetchFlow) | **Get** /v2/Flows/{Sid} | 
[**ListFlow**](FlowsApi.md#ListFlow) | **Get** /v2/Flows | 
[**UpdateFlow**](FlowsApi.md#UpdateFlow) | **Post** /v2/Flows/{Sid} | 



## CreateFlow

> StudioV2Flow CreateFlow(ctx, optional)



Create a Flow.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Status** | **string** | 
**Definition** | [**interface{}**](interface{}.md) | JSON representation of flow definition.
**CommitMessage** | **string** | Description of change made in the revision.

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


## ListFlow

> []StudioV2Flow ListFlow(ctx, optional)



Retrieve a list of all Flows.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]StudioV2Flow**](StudioV2Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
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
**Status** | **string** | 
**FriendlyName** | **string** | The string that you assigned to describe the Flow.
**Definition** | [**interface{}**](interface{}.md) | JSON representation of flow definition.
**CommitMessage** | **string** | Description of change made in the revision.

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

