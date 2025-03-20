# HostedNumberAuthorizationDocumentsDependentHostedNumberOrdersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDependentHostedNumberOrder**](HostedNumberAuthorizationDocumentsDependentHostedNumberOrdersApi.md#ListDependentHostedNumberOrder) | **Get** /v2/HostedNumber/AuthorizationDocuments/{SigningDocumentSid}/DependentHostedNumberOrders | Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.



## ListDependentHostedNumberOrder

> []NumbersV2DependentHostedNumberOrder ListDependentHostedNumberOrder(ctx, SigningDocumentSidoptional)

Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.

Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SigningDocumentSid** | **string** | A 34 character string that uniquely identifies the LOA document associated with this HostedNumberOrder.

### Other Parameters

Other parameters are passed through a pointer to a ListDependentHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | [**string**](stringstring.md) | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource#status-values) for more information on each of these statuses.
**PhoneNumber** | **string** | An E164 formatted phone number hosted by this HostedNumberOrder.
**IncomingPhoneNumberSid** | **string** | A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder.
**FriendlyName** | **string** | A human readable description of this resource, up to 128 characters.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2DependentHostedNumberOrder**](NumbersV2DependentHostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

