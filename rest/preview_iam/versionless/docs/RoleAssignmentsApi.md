# RoleAssignmentsApi

All URIs are relative to *https://preview-iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleAssignment**](RoleAssignmentsApi.md#CreateRoleAssignment) | **Post** /Organizations/{OrganizationSid}/RoleAssignments | Create a role assignment
[**DeleteRoleAssignment**](RoleAssignmentsApi.md#DeleteRoleAssignment) | **Delete** /Organizations/{OrganizationSid}/RoleAssignments/{Sid} | Delete a role assignment
[**ListRoleAssignments**](RoleAssignmentsApi.md#ListRoleAssignments) | **Get** /Organizations/{OrganizationSid}/RoleAssignments | List role assignments



## CreateRoleAssignment

> PublicApiRoleAssignmentResponse CreateRoleAssignment(ctx, OrganizationSidoptional)

Create a role assignment

Create a role assignment for the given organization

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PublicApiCreateRoleAssignmentRequest** | [**PublicApiCreateRoleAssignmentRequest**](PublicApiCreateRoleAssignmentRequest.md) | 

### Return type

[**PublicApiRoleAssignmentResponse**](PublicApiRoleAssignmentResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleAssignment

> DeleteRoleAssignment(ctx, OrganizationSidSid)

Delete a role assignment

Delete a role assignment for the given organization

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoleAssignmentParams struct


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


## ListRoleAssignments

> []PublicApiRoleAssignmentResponse ListRoleAssignments(ctx, OrganizationSidoptional)

List role assignments

List role assignments for the given organization

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**OrganizationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListRoleAssignmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | 
**Identity** | **string** | 
**Scope** | **string** | 
**Limit** | **int** | Max number of records to return.

### Return type

[**[]PublicApiRoleAssignmentResponse**](PublicApiRoleAssignmentResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

