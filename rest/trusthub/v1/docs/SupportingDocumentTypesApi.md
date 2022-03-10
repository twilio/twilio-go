# SupportingDocumentTypesApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSupportingDocumentType**](SupportingDocumentTypesApi.md#FetchSupportingDocumentType) | **Get** /v1/SupportingDocumentTypes/{Sid} | 
[**ListSupportingDocumentType**](SupportingDocumentTypesApi.md#ListSupportingDocumentType) | **Get** /v1/SupportingDocumentTypes | 



## FetchSupportingDocumentType

> TrusthubV1SupportingDocumentType FetchSupportingDocumentType(ctx, Sid)



Fetch a specific Supporting Document Type Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Supporting Document Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1SupportingDocumentType**](TrusthubV1SupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocumentType

> []TrusthubV1SupportingDocumentType ListSupportingDocumentType(ctx, optional)



Retrieve a list of all Supporting Document Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1SupportingDocumentType**](TrusthubV1SupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

