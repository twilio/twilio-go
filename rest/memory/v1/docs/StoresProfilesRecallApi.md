# StoresProfilesRecallApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFetchProfileMemory**](StoresProfilesRecallApi.md#CreateFetchProfileMemory) | **Post** /v1/Stores/{storeId}/Profiles/{profileId}/Recall | Retrieve Memories



## CreateFetchProfileMemory

> MemoryRetrievalResponse CreateFetchProfileMemory(ctx, StoreIdProfileIdoptional)

Retrieve Memories

Tailored memory retrieval for agentic workloads. Supports hybrid semantic search, date ranges, and configurable result limits for different memory types. This endpoint is optimized for conversational AI and memory retrieval use cases. If a query is not specified then one is inferred from the conversation context. If neither a query nor a conversationId is provided, results are returned in most-recent order without relevance scores.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a CreateFetchProfileMemoryParams struct


Name | Type | Description
------------- | ------------- | -------------
**AcceptEncoding** | **string** | Compression algorithms supported by the client (e.g., gzip, deflate, br)
**ContentEncoding** | **string** | Compression algorithm used for the request body (e.g., gzip, deflate, br)
**MemoryRetrievalRequest** | [**MemoryRetrievalRequest**](MemoryRetrievalRequest.md) | 

### Return type

[**MemoryRetrievalResponse**](MemoryRetrievalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

