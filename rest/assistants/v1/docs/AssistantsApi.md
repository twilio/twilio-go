# AssistantsApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistant**](AssistantsApi.md#CreateAssistant) | **Post** /v1/Assistants | Create a new assistant
[**DeleteAssistant**](AssistantsApi.md#DeleteAssistant) | **Delete** /v1/Assistants/{id} | Delete an assistant by ID
[**FetchAssistant**](AssistantsApi.md#FetchAssistant) | **Get** /v1/Assistants/{id} | Get an assistant by ID
[**ListAssistants**](AssistantsApi.md#ListAssistants) | **Get** /v1/Assistants | List all assistants
[**UpdateAssistant**](AssistantsApi.md#UpdateAssistant) | **Put** /v1/Assistants/{id} | Update an assistant by ID



## CreateAssistant

> AssistantsV1Assistant CreateAssistant(ctx, optional)

Create a new assistant

create an assistant

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1CreateAssistantRequest** | [**AssistantsV1CreateAssistantRequest**](AssistantsV1CreateAssistantRequest.md) | 

### Return type

[**AssistantsV1Assistant**](AssistantsV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistant

> DeleteAssistant(ctx, Id)

Delete an assistant by ID

delete an assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantParams struct


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


## FetchAssistant

> AssistantsV1AssistantWithToolsAndKnowledge FetchAssistant(ctx, Id)

Get an assistant by ID

get an assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AssistantsV1AssistantWithToolsAndKnowledge**](AssistantsV1AssistantWithToolsAndKnowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistants

> []AssistantsV1Assistant ListAssistants(ctx, optional)

List all assistants

list assistants

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAssistantsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Assistant**](AssistantsV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistant

> AssistantsV1Assistant UpdateAssistant(ctx, Idoptional)

Update an assistant by ID

update an assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1UpdateAssistantRequest** | [**AssistantsV1UpdateAssistantRequest**](AssistantsV1UpdateAssistantRequest.md) | 

### Return type

[**AssistantsV1Assistant**](AssistantsV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

