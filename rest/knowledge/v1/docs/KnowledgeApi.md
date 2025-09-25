# KnowledgeApi

All URIs are relative to *https://knowledge.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKnowledge**](KnowledgeApi.md#CreateKnowledge) | **Post** /v1/Knowledge | Create knowledge
[**DeleteKnowledge**](KnowledgeApi.md#DeleteKnowledge) | **Delete** /v1/Knowledge/{id} | Delete knowledge
[**FetchKnowledge**](KnowledgeApi.md#FetchKnowledge) | **Get** /v1/Knowledge/{id} | Get knowledge
[**ListKnowledge**](KnowledgeApi.md#ListKnowledge) | **Get** /v1/Knowledge | List all knowledge
[**UpdateKnowledge**](KnowledgeApi.md#UpdateKnowledge) | **Put** /v1/Knowledge/{id} | Update knowledge



## CreateKnowledge

> KnowledgeV1Knowledge CreateKnowledge(ctx, optional)

Create knowledge

Create knowledge

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**KnowledgeV1CreateKnowledgeRequest** | [**KnowledgeV1CreateKnowledgeRequest**](KnowledgeV1CreateKnowledgeRequest.md) | 

### Return type

[**KnowledgeV1Knowledge**](KnowledgeV1Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledge

> DeleteKnowledge(ctx, Id)

Delete knowledge

Delete knowledge

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | the Knowledge ID.

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKnowledge

> KnowledgeV1Knowledge FetchKnowledge(ctx, Id)

Get knowledge

Get knowledge

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**KnowledgeV1Knowledge**](KnowledgeV1Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKnowledge

> []ListKnowledgeResponseKnowledge ListKnowledge(ctx, optional)

List all knowledge

List all knowledge

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Tags** | **string** | Json array of tag and value pairs for tag filtering.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListKnowledgeResponseKnowledge**](ListKnowledgeResponseKnowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKnowledge

> KnowledgeV1Knowledge UpdateKnowledge(ctx, Idoptional)

Update knowledge

Update knowledge

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateKnowledgeParams struct


Name | Type | Description
------------- | ------------- | -------------
**KnowledgeV1UpdateKnowledgeRequest** | [**KnowledgeV1UpdateKnowledgeRequest**](KnowledgeV1UpdateKnowledgeRequest.md) | 

### Return type

[**KnowledgeV1Knowledge**](KnowledgeV1Knowledge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

