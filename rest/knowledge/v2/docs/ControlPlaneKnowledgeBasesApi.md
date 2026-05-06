# ControlPlaneKnowledgeBasesApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKnowledgeBase**](ControlPlaneKnowledgeBasesApi.md#CreateKnowledgeBase) | **Post** /v2/ControlPlane/KnowledgeBases | Create a Knowledge Base
[**DeleteKnowledgeBase**](ControlPlaneKnowledgeBasesApi.md#DeleteKnowledgeBase) | **Delete** /v2/ControlPlane/KnowledgeBases/{kbId} | Delete a Knowledge Base
[**FetchKnowledgeBase**](ControlPlaneKnowledgeBasesApi.md#FetchKnowledgeBase) | **Get** /v2/ControlPlane/KnowledgeBases/{kbId} | Retrieve a Knowledge Base
[**ListKnowledgeBases**](ControlPlaneKnowledgeBasesApi.md#ListKnowledgeBases) | **Get** /v2/ControlPlane/KnowledgeBases | List Knowledge Bases
[**UpdateKnowledgeBase**](ControlPlaneKnowledgeBasesApi.md#UpdateKnowledgeBase) | **Patch** /v2/ControlPlane/KnowledgeBases/{kbId} | Update a Knowledge Base



## CreateKnowledgeBase

> CreateKnowledgeBaseResponse CreateKnowledgeBase(ctx, optional)

Create a Knowledge Base

Create a new Knowledge Base for the Twilio account. Accounts can have multiple knowledge bases. Each knowledge base can contain multiple knowledge resources such as documents, websites, or text content that can be used for context and information retrieval.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateKnowledgeBaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**KnowledgeBaseCore** | [**KnowledgeBaseCore**](KnowledgeBaseCore.md) | 

### Return type

[**CreateKnowledgeBaseResponse**](CreateKnowledgeBase202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledgeBase

> DeleteKnowledgeBaseResponse DeleteKnowledgeBase(ctx, KbId)

Delete a Knowledge Base

Delete a Knowledge Base and all associated knowledge resources. This action cannot be undone.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a DeleteKnowledgeBaseParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteKnowledgeBaseResponse**](DeleteKnowledgeBase202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKnowledgeBase

> KnowledgeBase FetchKnowledgeBase(ctx, KbId)

Retrieve a Knowledge Base

Retrieve the details of a specific Knowledge Base by its unique ID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a FetchKnowledgeBaseParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**KnowledgeBase**](KnowledgeBase.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKnowledgeBases

> []KnowledgeBase ListKnowledgeBases(ctx, optional)

List Knowledge Bases

Get a list of knowledge bases for the Twilio account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeBasesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 100.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]KnowledgeBase**](KnowledgeBase.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKnowledgeBase

> UpdateKnowledgeBaseResponse UpdateKnowledgeBase(ctx, KbIdoptional)

Update a Knowledge Base

Partially update a Knowledge Base. Only the fields provided in the request body will be updated. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a PatchKnowledgeBaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Allows for optimistic concurrency control by making the request conditional. Server will only act if the resource's current Entity Tag (ETag) matches the one provided, preventing accidental overwrites.
**PatchKnowledgeBaseRequest** | [**PatchKnowledgeBaseRequest**](PatchKnowledgeBaseRequest.md) | 

### Return type

[**UpdateKnowledgeBaseResponse**](UpdateKnowledgeBase202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

