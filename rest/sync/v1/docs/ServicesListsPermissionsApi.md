# ServicesListsPermissionsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncListPermission**](ServicesListsPermissionsApi.md#DeleteSyncListPermission) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**FetchSyncListPermission**](ServicesListsPermissionsApi.md#FetchSyncListPermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**ListSyncListPermission**](ServicesListsPermissionsApi.md#ListSyncListPermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions | 
[**UpdateSyncListPermission**](ServicesListsPermissionsApi.md#UpdateSyncListPermission) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 



## DeleteSyncListPermission

> DeleteSyncListPermission(ctx, ServiceSidListSidIdentity)



Delete a specific Sync List Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync List Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListPermissionParams struct


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


## FetchSyncListPermission

> SyncV1SyncListPermission FetchSyncListPermission(ctx, ServiceSidListSidIdentity)



Fetch a specific Sync List Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync List Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncListPermission**](SyncV1SyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> []SyncV1SyncListPermission ListSyncListPermission(ctx, ServiceSidListSidoptional)



Retrieve a list of all Permissions applying to a Sync List.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncListPermission**](SyncV1SyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> SyncV1SyncListPermission UpdateSyncListPermission(ctx, ServiceSidListSidIdentityoptional)



Update an identity's access to a specific Sync List.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync List Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Read** | **bool** | Whether the identity can read the Sync List and its Items. Default value is `false`.
**Write** | **bool** | Whether the identity can create, update, and delete Items in the Sync List. Default value is `false`.
**Manage** | **bool** | Whether the identity can delete the Sync List. Default value is `false`.

### Return type

[**SyncV1SyncListPermission**](SyncV1SyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

