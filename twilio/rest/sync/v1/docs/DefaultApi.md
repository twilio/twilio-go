# DefaultApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DefaultApi.md#CreateDocument) | **Post** /v1/Services/{ServiceSid}/Documents | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateStreamMessage**](DefaultApi.md#CreateStreamMessage) | **Post** /v1/Services/{ServiceSid}/Streams/{StreamSid}/Messages | 
[**CreateSyncList**](DefaultApi.md#CreateSyncList) | **Post** /v1/Services/{ServiceSid}/Lists | 
[**CreateSyncListItem**](DefaultApi.md#CreateSyncListItem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**CreateSyncMap**](DefaultApi.md#CreateSyncMap) | **Post** /v1/Services/{ServiceSid}/Maps | 
[**CreateSyncMapItem**](DefaultApi.md#CreateSyncMapItem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**CreateSyncStream**](DefaultApi.md#CreateSyncStream) | **Post** /v1/Services/{ServiceSid}/Streams | 
[**DeleteDocument**](DefaultApi.md#DeleteDocument) | **Delete** /v1/Services/{ServiceSid}/Documents/{Sid} | 
[**DeleteDocumentPermission**](DefaultApi.md#DeleteDocumentPermission) | **Delete** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteSyncList**](DefaultApi.md#DeleteSyncList) | **Delete** /v1/Services/{ServiceSid}/Lists/{Sid} | 
[**DeleteSyncListItem**](DefaultApi.md#DeleteSyncListItem) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**DeleteSyncListPermission**](DefaultApi.md#DeleteSyncListPermission) | **Delete** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**DeleteSyncMap**](DefaultApi.md#DeleteSyncMap) | **Delete** /v1/Services/{ServiceSid}/Maps/{Sid} | 
[**DeleteSyncMapItem**](DefaultApi.md#DeleteSyncMapItem) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**DeleteSyncMapPermission**](DefaultApi.md#DeleteSyncMapPermission) | **Delete** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**DeleteSyncStream**](DefaultApi.md#DeleteSyncStream) | **Delete** /v1/Services/{ServiceSid}/Streams/{Sid} | 
[**FetchDocument**](DefaultApi.md#FetchDocument) | **Get** /v1/Services/{ServiceSid}/Documents/{Sid} | 
[**FetchDocumentPermission**](DefaultApi.md#FetchDocumentPermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchSyncList**](DefaultApi.md#FetchSyncList) | **Get** /v1/Services/{ServiceSid}/Lists/{Sid} | 
[**FetchSyncListItem**](DefaultApi.md#FetchSyncListItem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**FetchSyncListPermission**](DefaultApi.md#FetchSyncListPermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**FetchSyncMap**](DefaultApi.md#FetchSyncMap) | **Get** /v1/Services/{ServiceSid}/Maps/{Sid} | 
[**FetchSyncMapItem**](DefaultApi.md#FetchSyncMapItem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**FetchSyncMapPermission**](DefaultApi.md#FetchSyncMapPermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**FetchSyncStream**](DefaultApi.md#FetchSyncStream) | **Get** /v1/Services/{ServiceSid}/Streams/{Sid} | 
[**ListDocument**](DefaultApi.md#ListDocument) | **Get** /v1/Services/{ServiceSid}/Documents | 
[**ListDocumentPermission**](DefaultApi.md#ListDocumentPermission) | **Get** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListSyncList**](DefaultApi.md#ListSyncList) | **Get** /v1/Services/{ServiceSid}/Lists | 
[**ListSyncListItem**](DefaultApi.md#ListSyncListItem) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**ListSyncListPermission**](DefaultApi.md#ListSyncListPermission) | **Get** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions | 
[**ListSyncMap**](DefaultApi.md#ListSyncMap) | **Get** /v1/Services/{ServiceSid}/Maps | 
[**ListSyncMapItem**](DefaultApi.md#ListSyncMapItem) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**ListSyncMapPermission**](DefaultApi.md#ListSyncMapPermission) | **Get** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions | 
[**ListSyncStream**](DefaultApi.md#ListSyncStream) | **Get** /v1/Services/{ServiceSid}/Streams | 
[**UpdateDocument**](DefaultApi.md#UpdateDocument) | **Post** /v1/Services/{ServiceSid}/Documents/{Sid} | 
[**UpdateDocumentPermission**](DefaultApi.md#UpdateDocumentPermission) | **Post** /v1/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 
[**UpdateSyncList**](DefaultApi.md#UpdateSyncList) | **Post** /v1/Services/{ServiceSid}/Lists/{Sid} | 
[**UpdateSyncListItem**](DefaultApi.md#UpdateSyncListItem) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**UpdateSyncListPermission**](DefaultApi.md#UpdateSyncListPermission) | **Post** /v1/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**UpdateSyncMap**](DefaultApi.md#UpdateSyncMap) | **Post** /v1/Services/{ServiceSid}/Maps/{Sid} | 
[**UpdateSyncMapItem**](DefaultApi.md#UpdateSyncMapItem) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**UpdateSyncMapPermission**](DefaultApi.md#UpdateSyncMapPermission) | **Post** /v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**UpdateSyncStream**](DefaultApi.md#UpdateSyncStream) | **Post** /v1/Services/{ServiceSid}/Streams/{Sid} | 



## CreateDocument

> SyncV1ServiceDocument CreateDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.

### Other Parameters

Other parameters are passed through a pointer to a CreateDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live).
**UniqueName** | **string** | An application-defined string that uniquely identifies the Sync Document

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> SyncV1Service CreateService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**AclEnabled** | **bool** | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
**FriendlyName** | **string** | A string that you assign to describe the resource.
**ReachabilityDebouncingEnabled** | **bool** | Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event.
**ReachabilityDebouncingWindow** | **int32** | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the &#x60;webhook_url&#x60; is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to &#x60;webhook_url&#x60;.
**ReachabilityWebhooksEnabled** | **bool** | Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;.
**WebhookUrl** | **string** | The URL we should call when Sync objects are manipulated.
**WebhooksFromRestEnabled** | **bool** | Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;.

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamMessage

> SyncV1ServiceSyncStreamStreamMessage CreateStreamMessage(ctx, ServiceSidStreamSidoptional)



Create a new Stream Message.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in.
**StreamSid** | **string** | The SID of the Sync Stream to create the new Stream Message resource for.

### Other Parameters

Other parameters are passed through a pointer to a CreateStreamMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.

### Return type

[**SyncV1ServiceSyncStreamStreamMessage**](SyncV1ServiceSyncStreamStreamMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncList

> SyncV1ServiceSyncList CreateSyncList(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
**Ttl** | **int32** | Alias for collection_ttl. If both are provided, this value is ignored.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMap

> SyncV1ServiceSyncMap CreateSyncMap(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**Key** | **string** | The unique, user-defined key for the Map Item. Can be up to 320 characters long.
**Ttl** | **int32** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncStream

> SyncV1ServiceSyncStream CreateSyncStream(ctx, ServiceSidoptional)



Create a new Stream.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncStreamParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete.
**Sid** | **string** | The SID of the Document resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentParams struct


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


## DeleteDocumentPermission

> DeleteDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Delete a specific Sync Document Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to delete.

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


## DeleteService

> DeleteService(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


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


## DeleteSyncList

> DeleteSyncList(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete.
**Sid** | **string** | The SID of the Sync List resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

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


## DeleteSyncListItem

> DeleteSyncListItem(ctx, ServiceSidListSidIndexoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int32** | The index of the Sync List Item resource to delete.

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


## DeleteSyncListPermission

> DeleteSyncListPermission(ctx, ServiceSidListSidIdentity)



Delete a specific Sync List Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to delete.

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


## DeleteSyncMap

> DeleteSyncMap(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete.
**Sid** | **string** | The SID of the Sync Map resource to delete. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

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


## DeleteSyncMapPermission

> DeleteSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Delete a specific Sync Map Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to delete.

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


## DeleteSyncStream

> DeleteSyncStream(ctx, ServiceSidSid)



Delete a specific Stream.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to delete.
**Sid** | **string** | The SID of the Stream resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncStreamParams struct


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


## FetchDocument

> SyncV1ServiceDocument FetchDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch.
**Sid** | **string** | The SID of the Document resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocumentPermission

> SyncV1ServiceDocumentDocumentPermission FetchDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Fetch a specific Sync Document Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](SyncV1ServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> SyncV1Service FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncList

> SyncV1ServiceSyncList FetchSyncList(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch.
**Sid** | **string** | The SID of the Sync List resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

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
**Index** | **int32** | The index of the Sync List Item resource to fetch.

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
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListPermission

> SyncV1ServiceSyncListSyncListPermission FetchSyncListPermission(ctx, ServiceSidListSidIdentity)



Fetch a specific Sync List Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](SyncV1ServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMap

> SyncV1ServiceSyncMap FetchSyncMap(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch.
**Sid** | **string** | The SID of the Sync Map resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

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
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Fetch a specific Sync Map Permission.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](SyncV1ServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncStream

> SyncV1ServiceSyncStream FetchSyncStream(ctx, ServiceSidSid)



Fetch a specific Stream.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to fetch.
**Sid** | **string** | The SID of the Stream resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncStreamParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> ListDocumentResponse ListDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDocumentResponse**](ListDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> ListDocumentPermissionResponse ListDocumentPermission(ctx, ServiceSidDocumentSidoptional)



Retrieve a list of all Permissions applying to a Sync Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDocumentPermissionResponse**](ListDocumentPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> ListSyncListResponse ListSyncList(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncListResponse**](ListSyncListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncListItemResponse**](ListSyncListItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> ListSyncListPermissionResponse ListSyncListPermission(ctx, ServiceSidListSidoptional)



Retrieve a list of all Permissions applying to a Sync List.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncListPermissionResponse**](ListSyncListPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> ListSyncMapResponse ListSyncMap(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncMapResponse**](ListSyncMapResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncMapItemResponse**](ListSyncMapItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> ListSyncMapPermissionResponse ListSyncMapPermission(ctx, ServiceSidMapSidoptional)



Retrieve a list of all Permissions applying to a Sync Map.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Permission resources to read. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncMapPermissionResponse**](ListSyncMapPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncStream

> ListSyncStreamResponse ListSyncStream(ctx, ServiceSidoptional)



Retrieve a list of all Streams in a Service Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Stream resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncStreamParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSyncStreamResponse**](ListSyncStreamResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> SyncV1ServiceDocument UpdateDocument(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update.
**Sid** | **string** | The SID of the Document resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
**Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (time-to-live).

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> SyncV1ServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, ServiceSidDocumentSidIdentityoptional)



Update an identity's access to a specific Sync Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Manage** | **bool** | Whether the identity can delete the Sync Document. Default value is &#x60;false&#x60;.
**Read** | **bool** | Whether the identity can read the Sync Document. Default value is &#x60;false&#x60;.
**Write** | **bool** | Whether the identity can update the Sync Document. Default value is &#x60;false&#x60;.

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](SyncV1ServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> SyncV1Service UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**AclEnabled** | **bool** | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource.
**FriendlyName** | **string** | A string that you assign to describe the resource.
**ReachabilityDebouncingEnabled** | **bool** | Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event.
**ReachabilityDebouncingWindow** | **int32** | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called.
**ReachabilityWebhooksEnabled** | **bool** | Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;.
**WebhookUrl** | **string** | The URL we should call when Sync objects are manipulated.
**WebhooksFromRestEnabled** | **bool** | Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;.

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncList

> SyncV1ServiceSyncList UpdateSyncList(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update.
**Sid** | **string** | The SID of the Sync List resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;collection_ttl&#x60;. If both are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

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
**Index** | **int32** | The index of the Sync List Item resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListItemParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted. This parameter can only be used when the List Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncListSyncListItem**](SyncV1ServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> SyncV1ServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, ServiceSidListSidIdentityoptional)



Update an identity's access to a specific Sync List.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Manage** | **bool** | Whether the identity can delete the Sync List. Default value is &#x60;false&#x60;.
**Read** | **bool** | Whether the identity can read the Sync List and its Items. Default value is &#x60;false&#x60;.
**Write** | **bool** | Whether the identity can create, update, and delete Items in the Sync List. Default value is &#x60;false&#x60;.

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](SyncV1ServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMap

> SyncV1ServiceSyncMap UpdateSyncMap(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update.
**Sid** | **string** | The SID of the Sync Map resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapParams struct


Name | Type | Description
------------- | ------------- | -------------
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

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
**CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted. This parameter can only be used when the Map Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
**ItemTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
**Ttl** | **int32** | An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored.

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](SyncV1ServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, ServiceSidMapSidIdentityoptional)



Update an identity's access to a specific Sync Map.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapPermissionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Manage** | **bool** | Whether the identity can delete the Sync Map. Default value is &#x60;false&#x60;.
**Read** | **bool** | Whether the identity can read the Sync Map and its Items. Default value is &#x60;false&#x60;.
**Write** | **bool** | Whether the identity can create, update, and delete Items in the Sync Map. Default value is &#x60;false&#x60;.

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](SyncV1ServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncStream

> SyncV1ServiceSyncStream UpdateSyncStream(ctx, ServiceSidSidoptional)



Update a specific Stream.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to update.
**Sid** | **string** | The SID of the Stream resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncStreamParams struct


Name | Type | Description
------------- | ------------- | -------------
**Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

