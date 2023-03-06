# RegulatoryComplianceIdentitiesDocumentsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDocument**](RegulatoryComplianceIdentitiesDocumentsApi.md#ListDocument) | **Get** /v1/RegulatoryCompliance/Identities/{IdentitySid}/Documents | 



## ListDocument

> []NumbersV1Document ListDocument(ctx, IdentitySidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdentitySid** | **string** | Show only the Document resources that support the Identity with this SID.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1Document**](NumbersV1Document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

