# KnowledgeBasesSearchApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKnowledgeSearch**](KnowledgeBasesSearchApi.md#CreateKnowledgeSearch) | **Post** /v2/KnowledgeBases/{kbId}/Search | Search Knowledge Chunks



## CreateKnowledgeSearch

> CreateKnowledgeSearchResponse CreateKnowledgeSearch(ctx, KbIdoptional)

Search Knowledge Chunks

Perform semantic search across knowledge sources within a knowledge base to find the most relevant content chunks based on a natural language query.  Returns ranked chunks with similarity scores, allowing you to retrieve contextually relevant information for AI applications, chatbots, or  information retrieval systems. You can filter results by specific knowledge sources.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateKnowledgeSearchParams struct


Name | Type | Description
------------- | ------------- | -------------
**KnowledgeSearch** | [**KnowledgeSearch**](KnowledgeSearch.md) | 

### Return type

[**CreateKnowledgeSearchResponse**](CreateKnowledgeSearch200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

