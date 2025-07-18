# OrganizationApi

All URIs are relative to *https://preview-iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOrganization**](OrganizationApi.md#FetchOrganization) | **Get** /Organizations/{OrganizationSid} | List SCIM Users



## FetchOrganization

> FetchOrganization(ctx, OrganizationSid)

List SCIM Users

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchOrganizationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

