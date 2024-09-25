# KnowledgeStatusApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchKnowledgeStatus**](KnowledgeStatusApi.md#FetchKnowledgeStatus) | **Get** /v1/Knowledge/{id}/Status | Get knowledge status



## FetchKnowledgeStatus

> AssistantsV1KnowledgeStatus FetchKnowledgeStatus(ctx, Id)

Get knowledge status

Get knowledge status

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | the Knowledge ID.

### Other Parameters

Other parameters are passed through a pointer to a FetchKnowledgeStatusParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AssistantsV1KnowledgeStatus**](AssistantsV1KnowledgeStatus.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

