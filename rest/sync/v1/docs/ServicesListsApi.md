# ServicesListsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncList**](ServicesListsApi.md#CreateSyncList) | **Post** /v1/Services/{ServiceSid}/Lists | 
[**DeleteSyncList**](ServicesListsApi.md#DeleteSyncList) | **Delete** /v1/Services/{ServiceSid}/Lists/{Sid} | 
[**FetchSyncList**](ServicesListsApi.md#FetchSyncList) | **Get** /v1/Services/{ServiceSid}/Lists/{Sid} | 
[**ListSyncList**](ServicesListsApi.md#ListSyncList) | **Get** /v1/Services/{ServiceSid}/Lists | 
[**UpdateSyncList**](ServicesListsApi.md#UpdateSyncList) | **Post** /v1/Services/{ServiceSid}/Lists/{Sid} | 



## CreateSyncList

> SyncV1SyncList CreateSyncList(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
**Ttl** | **int** | Alias for collection_ttl. If both are provided, this value is ignored.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncList**](SyncV1SyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncList

> DeleteSyncList(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete.
**Sid** | **string** | The SID of the Sync List resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListParams struct


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


## FetchSyncList

> SyncV1SyncList FetchSyncList(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch.
**Sid** | **string** | The SID of the Sync List resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncList**](SyncV1SyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> []SyncV1SyncList ListSyncList(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncList**](SyncV1SyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncList

> SyncV1SyncList UpdateSyncList(ctx, ServiceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update.
**Sid** | **string** | The SID of the Sync List resource to update. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ttl** | **int** | An alias for `collection_ttl`. If both are provided, this value is ignored.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncList**](SyncV1SyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

