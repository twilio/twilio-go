# AccountsApi

All URIs are relative to *https://preview-iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOrganizationAccount**](AccountsApi.md#FetchOrganizationAccount) | **Get** /Organizations/{OrganizationSid}/Accounts/{AccountSid} | Get details of organization account
[**ListOrganizationAccounts**](AccountsApi.md#ListOrganizationAccounts) | **Get** /Organizations/{OrganizationSid}/Accounts | List organization accounts



## FetchOrganizationAccount

> PublicApiAccountResponse FetchOrganizationAccount(ctx, OrganizationSidoptional)

Get details of organization account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchOrganizationAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | 

### Return type

[**PublicApiAccountResponse**](PublicApiAccountResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationAccounts

> []PublicApiAccountResponse ListOrganizationAccounts(ctx, OrganizationSidoptional)

List organization accounts

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListOrganizationAccountsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | 
**Limit** | **int** | Max number of records to return.

### Return type

[**[]PublicApiAccountResponse**](PublicApiAccountResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

