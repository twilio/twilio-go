# PrimitivesApi

All URIs are relative to *https://ai.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPrimitive**](PrimitivesApi.md#FetchPrimitive) | **Get** /v1/Primitives/{Sid} | 
[**ListPrimitive**](PrimitivesApi.md#ListPrimitive) | **Get** /v1/Primitives | 



## FetchPrimitive

> AiV1Primitive FetchPrimitive(ctx, Sid)



Fetch a specific Primitive.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the LIT Primitive.

### Other Parameters

Other parameters are passed through a pointer to a FetchPrimitiveParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AiV1Primitive**](AiV1Primitive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrimitive

> []AiV1Primitive ListPrimitive(ctx, optional)



Retrieve a list of Primitives.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPrimitiveParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AiV1Primitive**](AiV1Primitive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

