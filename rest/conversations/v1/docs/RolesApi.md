# RolesApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /v1/Roles | Create a new user role in your account&#39;s default service
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /v1/Roles/{Sid} | Remove a user role from your account&#39;s default service
[**FetchRole**](RolesApi.md#FetchRole) | **Get** /v1/Roles/{Sid} | Fetch a user role from your account&#39;s default service
[**ListRole**](RolesApi.md#ListRole) | **Get** /v1/Roles | Retrieve a list of all user roles in your account&#39;s default service
[**UpdateRole**](RolesApi.md#UpdateRole) | **Post** /v1/Roles/{Sid} | Update an existing user role in your account&#39;s default service



## CreateRole

> ConversationsV1Role CreateRole(ctx, optional)

Create a new user role in your account's default service

Create a new user role in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
**Type** | [**string**](string.md) | 
**Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, Sid)

Remove a user role from your account's default service

Remove a user role from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoleParams struct


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


## FetchRole

> ConversationsV1Role FetchRole(ctx, Sid)

Fetch a user role from your account's default service

Fetch a user role from your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> []ConversationsV1Role ListRole(ctx, optional)

Retrieve a list of all user roles in your account's default service

Retrieve a list of all user roles in your account's default service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 50.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ConversationsV1Role UpdateRole(ctx, Sidoptional)

Update an existing user role in your account's default service

Update an existing user role in your account's default service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

