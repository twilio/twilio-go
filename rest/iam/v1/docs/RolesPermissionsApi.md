# RolesPermissionsApi

All URIs are relative to *https://iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRolePermission**](RolesPermissionsApi.md#ListRolePermission) | **Get** /v1/Roles/{RoleSid}/Permissions | 



## ListRolePermission

> []IamV1RolePermission ListRolePermission(ctx, RoleSidoptional)



List resolved permissions for a role

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoleSid** | **string** | The SID of the Role for which Permissions will be fetched.

### Other Parameters

Other parameters are passed through a pointer to a ListRolePermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 100.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IamV1RolePermission**](IamV1RolePermission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

