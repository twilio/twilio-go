# ServicesMapsItemsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncMapItem**](ServicesMapsItemsApi.md#CreateSyncMapItem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**DeleteSyncMapItem**](ServicesMapsItemsApi.md#DeleteSyncMapItem) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**FetchSyncMapItem**](ServicesMapsItemsApi.md#FetchSyncMapItem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**ListSyncMapItem**](ServicesMapsItemsApi.md#ListSyncMapItem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**UpdateSyncMapItem**](ServicesMapsItemsApi.md#UpdateSyncMapItem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 



## CreateSyncMapItem

> SyncV1SyncMapItem CreateSyncMapItem(ctx, ServiceSidMapSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
**MapSid** | **string** | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | The unique, user-defined key for the Map Item. Can be up to 320 characters long.
**Data** | [**interface{}**](interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**Ttl** | **int** | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncMapItem**](SyncV1SyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncMapItem

> DeleteSyncMapItem(ctx, ServiceSidMapSidKeyoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to delete.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map resource's `sid` or its `unique_name`.
**Key** | **string** | The `key` value of the Sync Map Item resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).

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


## FetchSyncMapItem

> SyncV1SyncMapItem FetchSyncMapItem(ctx, ServiceSidMapSidKey)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.
**Key** | **string** | The `key` value of the Sync Map Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncMapItem**](SyncV1SyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> []SyncV1SyncMapItem ListSyncMapItem(ctx, ServiceSidMapSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | How to order the Map Items returned by their `key` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key.
**From** | **string** | The `key` of the first Sync Map Item resource to read. See also `bounds`.
**Bounds** | **string** | Whether to include the Map Item referenced by the `from` parameter. Can be: `inclusive` to include the Map Item referenced by the `from` parameter or `exclusive` to start with the next Map Item. The default value is `inclusive`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncMapItem**](SyncV1SyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> SyncV1SyncMapItem UpdateSyncMapItem(ctx, ServiceSidMapSidKeyoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map resource's `sid` or its `unique_name`.
**Key** | **string** | The `key` value of the Sync Map Item resource to update. 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**Data** | [**interface{}**](interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**Ttl** | **int** | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted. This parameter can only be used when the Map Item's `data` or `ttl` is updated in the same request.

### Return type

[**SyncV1SyncMapItem**](SyncV1SyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

