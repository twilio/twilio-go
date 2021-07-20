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

> SyncV1ServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, ServiceSidMapSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
**MapSid** | **string** | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**Key** | **string** | The unique, user-defined key for the Map Item. Can be up to 320 characters long.
**Ttl** | **int** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md)

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
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to delete.

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

> SyncV1ServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, ServiceSidMapSidKey)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> ListSyncMapItemResponse ListSyncMapItem(ctx, ServiceSidMapSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | How to order the Map Items returned by their &#x60;key&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key.
**From** | **string** | The &#x60;key&#x60; of the first Sync Map Item resource to read. See also &#x60;bounds&#x60;.
**Bounds** | **string** | Whether to include the Map Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the Map Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next Map Item. The default value is &#x60;inclusive&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListSyncMapItemResponse**](ListSyncMapItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, ServiceSidMapSidKeyoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to update. 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**CollectionTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted. This parameter can only be used when the Map Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**Ttl** | **int** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

