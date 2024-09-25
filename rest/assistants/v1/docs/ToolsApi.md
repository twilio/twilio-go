# ToolsApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTool**](ToolsApi.md#CreateTool) | **Post** /v1/Tools | Create tool
[**DeleteTool**](ToolsApi.md#DeleteTool) | **Delete** /v1/Tools/{id} | Delete tool
[**FetchTool**](ToolsApi.md#FetchTool) | **Get** /v1/Tools/{id} | Get tool
[**ListTools**](ToolsApi.md#ListTools) | **Get** /v1/Tools | List tools
[**UpdateTool**](ToolsApi.md#UpdateTool) | **Put** /v1/Tools/{id} | Update tool



## CreateTool

> AssistantsV1Tool CreateTool(ctx, optional)

Create tool

Create tool

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateToolParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1CreateToolRequest** | [**AssistantsV1CreateToolRequest**](AssistantsV1CreateToolRequest.md) | 

### Return type

[**AssistantsV1Tool**](AssistantsV1Tool.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTool

> DeleteTool(ctx, Id)

Delete tool

delete a tool

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The tool ID.

### Other Parameters

Other parameters are passed through a pointer to a DeleteToolParams struct


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


## FetchTool

> AssistantsV1ToolWithPolicies FetchTool(ctx, Id)

Get tool

Get tool

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchToolParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AssistantsV1ToolWithPolicies**](AssistantsV1ToolWithPolicies.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTools

> []AssistantsV1Tool ListTools(ctx, optional)

List tools

List tools

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListToolsParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantId** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Tool**](AssistantsV1Tool.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTool

> AssistantsV1Tool UpdateTool(ctx, Idoptional)

Update tool

Update tool

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateToolParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1UpdateToolRequest** | [**AssistantsV1UpdateToolRequest**](AssistantsV1UpdateToolRequest.md) | 

### Return type

[**AssistantsV1Tool**](AssistantsV1Tool.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

