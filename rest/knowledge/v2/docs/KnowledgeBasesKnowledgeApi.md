# KnowledgeBasesKnowledgeApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKnowledge**](KnowledgeBasesKnowledgeApi.md#CreateKnowledge) | **Post** /v2/KnowledgeBases/{kbId}/Knowledge | Create Knowledge Source
[**DeleteKnowledge**](KnowledgeBasesKnowledgeApi.md#DeleteKnowledge) | **Delete** /v2/KnowledgeBases/{kbId}/Knowledge/{knowledgeId} | Delete Knowledge Source
[**FetchKnowledge**](KnowledgeBasesKnowledgeApi.md#FetchKnowledge) | **Get** /v2/KnowledgeBases/{kbId}/Knowledge/{knowledgeId} | Retrieve Knowledge Source
[**ListKnowledge**](KnowledgeBasesKnowledgeApi.md#ListKnowledge) | **Get** /v2/KnowledgeBases/{kbId}/Knowledge | List Knowledge Sources
[**UpdateKnowledge**](KnowledgeBasesKnowledgeApi.md#UpdateKnowledge) | **Patch** /v2/KnowledgeBases/{kbId}/Knowledge/{knowledgeId} | Update Knowledge Source



## CreateKnowledge

> Knowledge CreateKnowledge(ctx, KbIdoptional)

Create Knowledge Source

Create a new knowledge source from various data sources such as web content,  files, or raw text. The knowledge source will be processed and indexed to enable semantic search and retrieval.  ## Best practices  To maximize the effectiveness of Knowledge, consider the following best practices:  **Assess and optimize content:** Regularly evaluate your existing Knowledge sources for accuracy, relevance, and coverage. Identify any gaps or outdated information that could hinder the Assistant's performance.  **Simplify and structure content:** Ensure that the content is clear and concise. Use headings, bullet points, and metadata to make information straightforward to navigate for both the AI Assistant and your users.  **Prioritize high-impact content:** Focus on updating and maintaining content that is frequently accessed or critical to customer interactions. Consider using analytics to determine which Knowledge entries are most valuable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**KnowledgeCore** | [**KnowledgeCore**](KnowledgeCore.md) | 

### Return type

[**Knowledge**](Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledge

> DeleteKnowledge(ctx, KbIdKnowledgeId)

Delete Knowledge Source

Permanently delete knowledge source and all its associated data, including  processed chunks and embeddings. This action cannot be undone. The knowledge  resource will no longer be available for search or retrieval operations.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format
**KnowledgeId** | **string** | A unique Knowledge resource ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a DeleteKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKnowledge

> Knowledge FetchKnowledge(ctx, KbIdKnowledgeId)

Retrieve Knowledge Source

Fetch detailed information about a specific knowledge source by its ID.  This returns the complete knowledge source object including  processing status, source details, and configuration information.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format
**KnowledgeId** | **string** | A unique Knowledge resource ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a FetchKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Knowledge**](Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKnowledge

> []Knowledge ListKnowledge(ctx, KbIdoptional)

List Knowledge Sources

Retrieve a paginated list of all knowledge sources for a specific knowledge base.  Knowledge sources represent unstructured data sources such as documents,  websites, or text content that can be used for context and information retrieval.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**Page** | **int** | The page index. This value is simply for client state.
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
**PageToken** | **string** | The token for the page of results to retrieve.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]Knowledge**](Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKnowledge

> Knowledge UpdateKnowledge(ctx, KbIdKnowledgeIdoptional)

Update Knowledge Source

Partially update mutable fields of an existing knowledge source such as name, description, tags, or source-specific configuration. Only the fields provided in the request body will be updated. Some changes (e.g., KnowledgeSourceTypes) may trigger asynchronous reprocessing of the underlying content. Fields omitted from the request remain unchanged. Immutable fields (id, type, status, url, createdAt, updatedAt) cannot be modified directly.  ## Refresh ##  To request reprocessing without changing fields, pass the query parameter `refresh=true`. When `refresh=true` is provided, the server will re-queue processing for this knowledge resource (transitioning the persisted status to `QUEUED`) and return 202 Accepted. This query parameter is idempotent while the resource is already QUEUED or PROCESSING.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**KbId** | **string** | A unique Knowledge Base ID using Twilio Type ID (TTID) format
**KnowledgeId** | **string** | A unique Knowledge resource ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a PatchKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**Refresh** | **bool** | When true, re-queues processing for this knowledge resource. Idempotent while the resource is already QUEUED or PROCESSING.
**KnowledgeCore** | [**KnowledgeCore**](KnowledgeCore.md) | 

### Return type

[**Knowledge**](Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

