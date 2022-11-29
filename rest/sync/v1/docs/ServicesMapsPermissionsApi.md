# ServicesMapsPermissionsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncMapPermission**](ServicesMapsPermissionsApi.md#DeleteSyncMapPermission) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**FetchSyncMapPermission**](ServicesMapsPermissionsApi.md#FetchSyncMapPermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**ListSyncMapPermission**](ServicesMapsPermissionsApi.md#ListSyncMapPermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions | 
[**UpdateSyncMapPermission**](ServicesMapsPermissionsApi.md#UpdateSyncMapPermission) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 



## DeleteSyncMapPermission

> DeleteSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Delete a specific Sync Map Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service's `sid` value or `default`.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync Map Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapPermissionParams struct


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


## FetchSyncMapPermission

> SyncV1SyncMapPermission FetchSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Fetch a specific Sync Map Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service's `sid` value or `default`.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync Map Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncMapPermission**](SyncV1SyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> []SyncV1SyncMapPermission ListSyncMapPermission(ctx, ServiceSidMapSidoptional)



Retrieve a list of all Permissions applying to a Sync Map.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service's `sid` value or `default`.
**MapSid** | **string** | The SID of the Sync Map with the Permission resources to read. Can be the Sync Map resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncMapPermission**](SyncV1SyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> SyncV1SyncMapPermission UpdateSyncMapPermission(ctx, ServiceSidMapSidIdentityoptional)



Update an identity's access to a specific Sync Map.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service's `sid` value or `default`.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Sync Map Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Read** | **bool** | Whether the identity can read the Sync Map and its Items. Default value is `false`.
**Write** | **bool** | Whether the identity can create, update, and delete Items in the Sync Map. Default value is `false`.
**Manage** | **bool** | Whether the identity can delete the Sync Map. Default value is `false`.

### Return type

[**SyncV1SyncMapPermission**](SyncV1SyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

