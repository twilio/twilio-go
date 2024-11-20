# AssistantsKnowledgeApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistantKnowledgeAttachment**](AssistantsKnowledgeApi.md#CreateAssistantKnowledgeAttachment) | **Post** /v1/Assistants/{assistantId}/Knowledge/{id} | Attach Knowledge to Assistant
[**DeleteAssistantKnowledgeAttachment**](AssistantsKnowledgeApi.md#DeleteAssistantKnowledgeAttachment) | **Delete** /v1/Assistants/{assistantId}/Knowledge/{id} | Detach Knowledge to Assistant
[**ListKnowledgeByAssistant**](AssistantsKnowledgeApi.md#ListKnowledgeByAssistant) | **Get** /v1/Assistants/{assistantId}/Knowledge | List all knowledge for an Assistant



## CreateAssistantKnowledgeAttachment

> CreateAssistantKnowledgeAttachment(ctx, AssistantIdId)

Attach Knowledge to Assistant

Attach Knowledge to Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.
**Id** | **string** | The knowledge ID.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantKnowledgeAttachmentParams struct


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


## DeleteAssistantKnowledgeAttachment

> DeleteAssistantKnowledgeAttachment(ctx, AssistantIdId)

Detach Knowledge to Assistant

Detach Knowledge to Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.
**Id** | **string** | The knowledge ID.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantKnowledgeAttachmentParams struct


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


## ListKnowledgeByAssistant

> []AssistantsV1Knowledge ListKnowledgeByAssistant(ctx, AssistantIdoptional)

List all knowledge for an Assistant

List all knowledge for an Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantId** | **string** | The assistant ID.

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeByAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Knowledge**](AssistantsV1Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

