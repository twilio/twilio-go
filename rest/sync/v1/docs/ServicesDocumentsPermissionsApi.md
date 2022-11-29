# ServicesDocumentsPermissionsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDocumentPermission**](ServicesDocumentsPermissionsApi.md#DeleteDocumentPermission) | **Delete** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**FetchDocumentPermission**](ServicesDocumentsPermissionsApi.md#FetchDocumentPermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**ListDocumentPermission**](ServicesDocumentsPermissionsApi.md#ListDocumentPermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions | 
[**UpdateDocumentPermission**](ServicesDocumentsPermissionsApi.md#UpdateDocumentPermission) | **Post** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 



## DeleteDocumentPermission

> DeleteDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Delete a specific Sync Document Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Document Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentPermissionParams struct


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


## FetchDocumentPermission

> SyncV1DocumentPermission FetchDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Fetch a specific Sync Document Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Document Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1DocumentPermission**](SyncV1DocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> []SyncV1DocumentPermission ListDocumentPermission(ctx, ServiceSidDocumentSidoptional)



Retrieve a list of all Permissions applying to a Sync Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1DocumentPermission**](SyncV1DocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> SyncV1DocumentPermission UpdateDocumentPermission(ctx, ServiceSidDocumentSidIdentityoptional)



Update an identity's access to a specific Sync Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource's `sid` or its `unique_name`.
**Identity** | **string** | The application-defined string that uniquely identifies the User's Document Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Read** | **bool** | Whether the identity can read the Sync Document. Default value is `false`.
**Write** | **bool** | Whether the identity can update the Sync Document. Default value is `false`.
**Manage** | **bool** | Whether the identity can delete the Sync Document. Default value is `false`.

### Return type

[**SyncV1DocumentPermission**](SyncV1DocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

