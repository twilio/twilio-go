# ServicesMapsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncMap**](ServicesMapsApi.md#CreateSyncMap) | **Post** /v1/Services/{ServiceSid}/Maps | 
[**DeleteSyncMap**](ServicesMapsApi.md#DeleteSyncMap) | **Delete** /v1/Services/{ServiceSid}/Maps/{Sid} | 
[**FetchSyncMap**](ServicesMapsApi.md#FetchSyncMap) | **Get** /v1/Services/{ServiceSid}/Maps/{Sid} | 
[**ListSyncMap**](ServicesMapsApi.md#ListSyncMap) | **Get** /v1/Services/{ServiceSid}/Maps | 
[**UpdateSyncMap**](ServicesMapsApi.md#UpdateSyncMap) | **Post** /v1/Services/{ServiceSid}/Maps/{Sid} | 



## CreateSyncMap

> SyncV1SyncMap CreateSyncMap(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource.
**Ttl** | **int** | An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncMap**](SyncV1SyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncMap

> DeleteSyncMap(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete.
**Sid** | **string** | The SID of the Sync Map resource to delete. Can be the Sync Map's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapParams struct


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


## FetchSyncMap

> SyncV1SyncMap FetchSyncMap(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch.
**Sid** | **string** | The SID of the Sync Map resource to fetch. Can be the Sync Map's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncMap**](SyncV1SyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> []SyncV1SyncMap ListSyncMap(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncMap**](SyncV1SyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMap

> SyncV1SyncMap UpdateSyncMap(ctx, ServiceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update.
**Sid** | **string** | The SID of the Sync Map resource to update. Can be the Sync Map's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ttl** | **int** | An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncMap**](SyncV1SyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

