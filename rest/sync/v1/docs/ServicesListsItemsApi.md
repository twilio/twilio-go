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

> SyncV1ServiceSyncListSyncListItem CreateSyncListItem(ctx, ServiceSidListSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
**ListSid** | **string** | The SID of the Sync List to add the new List Item to. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**Ttl** | **int** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md)

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
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
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

> SyncV1ServiceSyncListSyncListItem FetchSyncListItem(ctx, ServiceSidListSidIndex)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int** | The index of the Sync List Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> ListSyncListItemResponse ListSyncListItem(ctx, ServiceSidListSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read.
**ListSid** | **string** | The SID of the Sync List with the List Items to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | How to order the List Items returned by their &#x60;index&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending.
**From** | **string** | The &#x60;index&#x60; of the first Sync List Item resource to read. See also &#x60;bounds&#x60;.
**Bounds** | **string** | Whether to include the List Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the List Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next List Item. The default value is &#x60;inclusive&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncListItemResponse**](ListSyncListItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> SyncV1ServiceSyncListSyncListItem UpdateSyncListItem(ctx, ServiceSidListSidIndexoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int** | The index of the Sync List Item resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted. This parameter can only be used when the List Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**Ttl** | **int** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

