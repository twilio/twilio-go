# SupportingDocumentsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportingDocument**](SupportingDocumentsApi.md#CreateSupportingDocument) | **Post** /v1/SupportingDocuments | 
[**DeleteSupportingDocument**](SupportingDocumentsApi.md#DeleteSupportingDocument) | **Delete** /v1/SupportingDocuments/{Sid} | 
[**FetchSupportingDocument**](SupportingDocumentsApi.md#FetchSupportingDocument) | **Get** /v1/SupportingDocuments/{Sid} | 
[**ListSupportingDocument**](SupportingDocumentsApi.md#ListSupportingDocument) | **Get** /v1/SupportingDocuments | 
[**UpdateSupportingDocument**](SupportingDocumentsApi.md#UpdateSupportingDocument) | **Post** /v1/SupportingDocuments/{Sid} | 



## CreateSupportingDocument

> TrusthubV1SupportingDocument CreateSupportingDocument(ctx, optional)



Create a new Supporting Document.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Type** | **string** | The type of the Supporting Document.
**Attributes** | [**interface{}**](interface{}.md) | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.

### Return type

[**TrusthubV1SupportingDocument**](TrusthubV1SupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupportingDocument

> DeleteSupportingDocument(ctx, Sid)



Delete a specific Supporting Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSupportingDocumentParams struct


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


## FetchSupportingDocument

> TrusthubV1SupportingDocument FetchSupportingDocument(ctx, Sid)



Fetch specific Supporting Document Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1SupportingDocument**](TrusthubV1SupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocument

> []TrusthubV1SupportingDocument ListSupportingDocument(ctx, optional)



Retrieve a list of all Supporting Document for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1SupportingDocument**](TrusthubV1SupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportingDocument

> TrusthubV1SupportingDocument UpdateSupportingDocument(ctx, Sidoptional)



Update an existing Supporting Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Attributes** | [**interface{}**](interface{}.md) | The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.

### Return type

[**TrusthubV1SupportingDocument**](TrusthubV1SupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

