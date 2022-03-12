# RegulatoryComplianceSupportingDocumentTypesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSupportingDocumentType**](RegulatoryComplianceSupportingDocumentTypesApi.md#FetchSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid} | 
[**ListSupportingDocumentType**](RegulatoryComplianceSupportingDocumentTypesApi.md#ListSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes | 



## FetchSupportingDocumentType

> NumbersV2SupportingDocumentType FetchSupportingDocumentType(ctx, Sid)



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

[**NumbersV2SupportingDocumentType**](NumbersV2SupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocumentType

> []NumbersV2SupportingDocumentType ListSupportingDocumentType(ctx, optional)



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

[**[]NumbersV2SupportingDocumentType**](NumbersV2SupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

