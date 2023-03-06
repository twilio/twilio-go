# HostedNumberAuthorizationDocumentsDependentOrdersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDependentOrder**](HostedNumberAuthorizationDocumentsDependentOrdersApi.md#ListDependentOrder) | **Get** /v1/HostedNumber/AuthorizationDocuments/{SigningDocumentSid}/DependentOrders | 



## ListDependentOrder

> []NumbersV1DependentOrder ListDependentOrder(ctx, SigningDocumentSidoptional)



Retrieve a list of HostedNumberOrder resources that belong to the account initiating the request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SigningDocumentSid** | **string** | The SID of the AuthorizationDocument resource that refers to the DependentOrder resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDependentOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that identifies the DependentOrder resources to read.
**IncomingPhoneNumberSid** | **string** | The SID of the IncomingPhoneNumber resource of the DependentOrder resources to read.
**PhoneNumber** | **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone numbers of the DependentOrder resources to read.
**Status** | **string** | The status of the DependentOrder resources to read. Can be: `twilio-processing`, `received`, `pending-verification`, `verified`, `pending-loa`, `carrier-processing`, `testing`, `completed`, `failed`, or `action-required`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1DependentOrder**](NumbersV1DependentOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

