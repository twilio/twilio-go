# \V2Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2FlowsCreate**](V2Api.md#V2FlowsCreate) | **Post** /v2/Flows | 
[**V2FlowsDelete**](V2Api.md#V2FlowsDelete) | **Delete** /v2/Flows/{Sid} | 
[**V2FlowsList**](V2Api.md#V2FlowsList) | **Get** /v2/Flows | 
[**V2FlowsRead**](V2Api.md#V2FlowsRead) | **Get** /v2/Flows/{Sid} | 
[**V2FlowsUpdate**](V2Api.md#V2FlowsUpdate) | **Post** /v2/Flows/{Sid} | 



## V2FlowsCreate

> StudioV2Flow V2FlowsCreate(ctx, optional)



Create a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V2FlowsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2FlowsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitMessage** | **optional.**| Description on change made in the revision. | 
 **definition** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
 **friendlyName** | **optional.**| The string that you assigned to describe the Flow. | 
 **status** | **optional.**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

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


## V2FlowsDelete

> V2FlowsDelete(ctx, sid)



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


## V2FlowsList

> InlineResponse2001 V2FlowsList(ctx, optional)



Retrieve a list of all Flows.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V2FlowsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2FlowsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2FlowsRead

> StudioV2Flow V2FlowsRead(ctx, sid)



Retrieve a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to fetch. | 

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


## V2FlowsUpdate

> StudioV2Flow V2FlowsUpdate(ctx, sid, optional)



Update a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to fetch. | 
 **optional** | ***V2FlowsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2FlowsUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitMessage** | **optional.**| Description on change made in the revision. | 
 **definition** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| JSON representation of flow definition. | 
 **friendlyName** | **optional.**| The string that you assigned to describe the Flow. | 
 **status** | **optional.**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 

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

