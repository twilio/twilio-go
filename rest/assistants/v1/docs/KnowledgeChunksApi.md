# KnowledgeChunksApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListKnowledgeChunks**](KnowledgeChunksApi.md#ListKnowledgeChunks) | **Get** /v1/Knowledge/{id}/Chunks | List knowledge chunks



## ListKnowledgeChunks

> []AssistantsV1KnowledgeChunk ListKnowledgeChunks(ctx, Idoptional)

List knowledge chunks

List knowledge chunks

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The knowledge ID.

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeChunksParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1KnowledgeChunk**](AssistantsV1KnowledgeChunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

