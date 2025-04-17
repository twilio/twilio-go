# ScimUsersApi

All URIs are relative to *https://preview-iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationUser**](ScimUsersApi.md#CreateOrganizationUser) | **Post** /Organizations/{OrganizationSid}/scim/Users | Create SCIM User
[**DeleteOrganizationUser**](ScimUsersApi.md#DeleteOrganizationUser) | **Delete** /Organizations/{OrganizationSid}/scim/Users/{Id} | Delete SCIM User
[**FetchOrganizationUser**](ScimUsersApi.md#FetchOrganizationUser) | **Get** /Organizations/{OrganizationSid}/scim/Users/{Id} | Get SCIM User
[**ListOrganizationUsers**](ScimUsersApi.md#ListOrganizationUsers) | **Get** /Organizations/{OrganizationSid}/scim/Users | List SCIM Users
[**UpdateOrganizationUser**](ScimUsersApi.md#UpdateOrganizationUser) | **Put** /Organizations/{OrganizationSid}/scim/Users/{Id} | Update SCIM User



## CreateOrganizationUser

> ScimUser CreateOrganizationUser(ctx, OrganizationSidoptional)

Create SCIM User

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateOrganizationUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**ScimUser** | [**ScimUser**](ScimUser.md) | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/scim+json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationUser

> DeleteOrganizationUser(ctx, OrganizationSidId)

Delete SCIM User

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteOrganizationUserParams struct


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


## FetchOrganizationUser

> ScimUser FetchOrganizationUser(ctx, OrganizationSidId)

Get SCIM User

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchOrganizationUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUsers

> []ScimUser ListOrganizationUsers(ctx, OrganizationSidoptional)

List SCIM Users

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListOrganizationUsersParams struct


Name | Type | Description
------------- | ------------- | -------------
**Filter** | **string** | 
**Limit** | **int** | Max number of records to return.
**PageSize** | **int** | Max number of records to return in a page

### Return type

[**[]ScimUser**](ScimUser.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationUser

> ScimUser UpdateOrganizationUser(ctx, OrganizationSidIdoptional)

Update SCIM User

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 
**Id** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateOrganizationUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | 
**ScimUser** | [**ScimUser**](ScimUser.md) | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/scim+json
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

