# HostedNumberAuthorizationDocumentsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#CreateAuthorizationDocument) | **Post** /v2/HostedNumber/AuthorizationDocuments | Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio&#39;s platform.
[**DeleteAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#DeleteAuthorizationDocument) | **Delete** /v2/HostedNumber/AuthorizationDocuments/{Sid} | Cancel the AuthorizationDocument request.
[**FetchAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#FetchAuthorizationDocument) | **Get** /v2/HostedNumber/AuthorizationDocuments/{Sid} | Fetch a specific AuthorizationDocument.
[**ListAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#ListAuthorizationDocument) | **Get** /v2/HostedNumber/AuthorizationDocuments | Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.



## CreateAuthorizationDocument

> NumbersV2AuthorizationDocument CreateAuthorizationDocument(ctx, optional)

Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio's platform.

Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio's platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument.
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing.
**ContactPhoneNumber** | **string** | The contact phone number of the person authorized to sign the Authorization Document.
**HostedNumberOrderSids** | **[]string** | A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio's platform.
**ContactTitle** | **string** | The title of the person authorized to sign the Authorization Document for this phone number.
**CcEmails** | **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed.

### Return type

[**NumbersV2AuthorizationDocument**](NumbersV2AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorizationDocument

> DeleteAuthorizationDocument(ctx, Sid)

Cancel the AuthorizationDocument request.

Cancel the AuthorizationDocument request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this AuthorizationDocument.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAuthorizationDocumentParams struct


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


## FetchAuthorizationDocument

> NumbersV2AuthorizationDocument FetchAuthorizationDocument(ctx, Sid)

Fetch a specific AuthorizationDocument.

Fetch a specific AuthorizationDocument.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this AuthorizationDocument.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2AuthorizationDocument**](NumbersV2AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationDocument

> []NumbersV2AuthorizationDocument ListAuthorizationDocument(ctx, optional)

Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.

Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing.
**Status** | [**string**](stringstring.md) | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource#status-values) for more information on each of these statuses.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2AuthorizationDocument**](NumbersV2AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

