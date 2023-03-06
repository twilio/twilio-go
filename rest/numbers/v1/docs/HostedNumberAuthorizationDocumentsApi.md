# HostedNumberAuthorizationDocumentsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#CreateAuthorizationDocument) | **Post** /v1/HostedNumber/AuthorizationDocuments | 
[**DeleteAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#DeleteAuthorizationDocument) | **Delete** /v1/HostedNumber/AuthorizationDocuments/{Sid} | 
[**FetchAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#FetchAuthorizationDocument) | **Get** /v1/HostedNumber/AuthorizationDocuments/{Sid} | 
[**ListAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#ListAuthorizationDocument) | **Get** /v1/HostedNumber/AuthorizationDocuments | 
[**UpdateAuthorizationDocument**](HostedNumberAuthorizationDocumentsApi.md#UpdateAuthorizationDocument) | **Post** /v1/HostedNumber/AuthorizationDocuments/{Sid} | 



## CreateAuthorizationDocument

> NumbersV1AuthorizationDocument CreateAuthorizationDocument(ctx, optional)



Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on our platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the AuthorizationDocument resource to update.
**ContactPhoneNumber** | **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number of the person who will sign the new authorization document.
**ContactTitle** | **string** | The formal title of the person who will sign the authorization document. For example, `Mr`, `Product Manager`, or `Marketing Director`.
**Email** | **string** | The email address that the new authorization document should be sent to for signing.
**HostedNumberOrderSids** | **[]string** | The HostedNumberOrder SIDs that the new authorization document will authorize for hosting phone number capabilities on our platform.
**CcEmails** | **[]string** | The email addresses that the new authorization document should be copied to. This parameter takes only one email address. To provide more than one email address, repeat this parameter for each address.

### Return type

[**NumbersV1AuthorizationDocument**](NumbersV1AuthorizationDocument.md)

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



Delete the AuthorizationDocument and remove any associations with HostedNumberOrders.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the AuthorizationDocument resource to delete.

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

> NumbersV1AuthorizationDocument FetchAuthorizationDocument(ctx, Sid)



Fetch a specific AuthorizationDocument.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the AuthorizationDocument resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1AuthorizationDocument**](NumbersV1AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationDocument

> []NumbersV1AuthorizationDocument ListAuthorizationDocument(ctx, optional)



Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | The email address of the AuthorizationDocument resources to read.
**Status** | **string** | The status of the AuthorizationDocument resources to read. Can be: `opened`, `signing`, `signed`, `canceled`, or `failed`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1AuthorizationDocument**](NumbersV1AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDocument

> NumbersV1AuthorizationDocument UpdateAuthorizationDocument(ctx, Sidoptional)



Updates a specific AuthorizationDocument.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the AuthorizationDocument resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAuthorizationDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the AuthorizationDocument to update.
**CcEmails** | **[]string** | The email addresses that the new authorization document should be copied to. This parameter takes only one email address. To provide more than one email address, repeat this parameter for each address.
**Email** | **string** | The email address that the authorization document should be sent to for signing.
**HostedNumberOrderSids** | **[]string** | The HostedNumberOrder SIDs that the authorization document will authorize for hosting phone number capabilities on our platform.
**Status** | **string** | 

### Return type

[**NumbersV1AuthorizationDocument**](NumbersV1AuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

