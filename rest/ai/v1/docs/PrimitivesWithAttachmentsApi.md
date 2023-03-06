# PrimitivesWithAttachmentsApi

All URIs are relative to *https://ai.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrimitiveWithAttachments**](PrimitivesWithAttachmentsApi.md#ListPrimitiveWithAttachments) | **Get** /v1/PrimitivesWithAttachments | 



## ListPrimitiveWithAttachments

> []AiV1PrimitiveWithAttachments ListPrimitiveWithAttachments(ctx, optional)



Retrieve a list of Primitives and the list of services attached to each primitive.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPrimitiveWithAttachmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AiV1PrimitiveWithAttachments**](AiV1PrimitiveWithAttachments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

