# AssistantsToolsApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistantToolAttachment**](AssistantsToolsApi.md#CreateAssistantToolAttachment) | **Post** /v1/Assistants/{assistantId}/Tools/{id} | Attach Tool to Assistant
[**DeleteAssistantToolAttachment**](AssistantsToolsApi.md#DeleteAssistantToolAttachment) | **Delete** /v1/Assistants/{assistantId}/Tools/{id} | Detach Tool to Assistant
[**ListToolsByAssistant**](AssistantsToolsApi.md#ListToolsByAssistant) | **Get** /v1/Assistants/{assistantId}/Tools | List tools for an Assistant



## CreateAssistantToolAttachment

> CreateAssistantToolAttachment(ctx, AssistantIdId)

Attach Tool to Assistant

Attach Tool to Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.
**Id** | **string** | The tool ID.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantToolAttachmentParams struct


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


## DeleteAssistantToolAttachment

> DeleteAssistantToolAttachment(ctx, AssistantIdId)

Detach Tool to Assistant

Detach Tool to Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.
**Id** | **string** | The tool ID.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantToolAttachmentParams struct


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


## ListToolsByAssistant

> []AssistantsV1Tool ListToolsByAssistant(ctx, AssistantIdoptional)

List tools for an Assistant

List tools for an Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.

### Other Parameters

Other parameters are passed through a pointer to a ListToolsByAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
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

