# ServicesListsItemsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncListItem**](ServicesListsItemsApi.md#CreateSyncListItem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**DeleteSyncListItem**](ServicesListsItemsApi.md#DeleteSyncListItem) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**FetchSyncListItem**](ServicesListsItemsApi.md#FetchSyncListItem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**ListSyncListItem**](ServicesListsItemsApi.md#ListSyncListItem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**UpdateSyncListItem**](ServicesListsItemsApi.md#UpdateSyncListItem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 



## CreateSyncListItem

> SyncV1SyncListItem CreateSyncListItem(ctx, ServiceSidListSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
**ListSid** | **string** | The SID of the Sync List to add the new List Item to. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Data** | [**interface{}**](interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**Ttl** | **int** | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted.

### Return type

[**SyncV1SyncListItem**](SyncV1SyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncListItem

> DeleteSyncListItem(ctx, ServiceSidListSidIndexoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.
**Index** | **int** | The index of the Sync List Item resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListItemParams struct


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


## FetchSyncListItem

> SyncV1SyncListItem FetchSyncListItem(ctx, ServiceSidListSidIndex)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.
**Index** | **int** | The index of the Sync List Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1SyncListItem**](SyncV1SyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> []SyncV1SyncListItem ListSyncListItem(ctx, ServiceSidListSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read.
**ListSid** | **string** | The SID of the Sync List with the List Items to read. Can be the Sync List resource's `sid` or its `unique_name`.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | How to order the List Items returned by their `index` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending.
**From** | **string** | The `index` of the first Sync List Item resource to read. See also `bounds`.
**Bounds** | **string** | Whether to include the List Item referenced by the `from` parameter. Can be: `inclusive` to include the List Item referenced by the `from` parameter or `exclusive` to start with the next List Item. The default value is `inclusive`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncListItem**](SyncV1SyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> SyncV1SyncListItem UpdateSyncListItem(ctx, ServiceSidListSidIndexoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource's `sid` or its `unique_name`.
**Index** | **int** | The index of the Sync List Item resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**Data** | [**interface{}**](interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**Ttl** | **int** | An alias for `item_ttl`. If both parameters are provided, this value is ignored.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted. This parameter can only be used when the List Item's `data` or `ttl` is updated in the same request.

### Return type

[**SyncV1SyncListItem**](SyncV1SyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

