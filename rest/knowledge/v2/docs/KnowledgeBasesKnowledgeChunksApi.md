# KnowledgeBasesKnowledgeChunksApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListKnowledgeChunks**](KnowledgeBasesKnowledgeChunksApi.md#ListKnowledgeChunks) | **Get** /v2/KnowledgeBases/{kbId}/Knowledge/{knowledgeId}/Chunks | List Knowledge Chunks



## ListKnowledgeChunks

> []KnowledgeChunk ListKnowledgeChunks(ctx, KbIdKnowledgeIdoptional)

List Knowledge Chunks

Retrieve a paginated list of all processed content chunks from a specific knowledge source.  Chunks are smaller segments of content that have been extracted and processed from the original  knowledge source. Each chunk contains content text and associated metadata that can be used  for semantic search and retrieval operations.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format
**KnowledgeId** | **string** | A unique Knowledge resource ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeChunksParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**PageToken** | **string** | The page token. This is provided by the API.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]KnowledgeChunk**](KnowledgeChunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

