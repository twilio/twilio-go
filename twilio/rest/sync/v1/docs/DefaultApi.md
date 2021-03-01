# DefaultApi

All URIs are relative to *http://localhost*

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

> SyncV1ServiceDocument CreateDocument(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in. | 
 **optional** | ***CreateDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length. | 
**Ttl** | **Int32**| How long, in seconds, before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Document does not expire. The Sync Document will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the Sync Document | 

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> SyncV1Service CreateService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AclEnabled** | **Bool**| Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. | 
**FriendlyName** | **String**| A string that you assign to describe the resource. | 
**ReachabilityDebouncingEnabled** | **Bool**| Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event. | 
**ReachabilityDebouncingWindow** | **Int32**| The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the &#x60;webhook_url&#x60; is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to &#x60;webhook_url&#x60;. | 
**ReachabilityWebhooksEnabled** | **Bool**| Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;. | 
**WebhookUrl** | **String**| The URL we should call when Sync objects are manipulated. | 
**WebhooksFromRestEnabled** | **Bool**| Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;. | 

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamMessage

> SyncV1ServiceSyncStreamStreamMessage CreateStreamMessage(ctx, ServiceSid, StreamSid, optional)



Create a new Stream Message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in. | 
**StreamSid** | **string**| The SID of the Sync Stream to create the new Stream Message resource for. | 
 **optional** | ***CreateStreamMessageRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateStreamMessageRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length. | 

### Return type

[**SyncV1ServiceSyncStreamStreamMessage**](SyncV1ServiceSyncStreamStreamMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncList

> SyncV1ServiceSyncList CreateSyncList(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in. | 
 **optional** | ***CreateSyncListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| Alias for collection_ttl. If both are provided, this value is ignored. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncListItem

> SyncV1ServiceSyncListSyncListItem CreateSyncListItem(ctx, ServiceSid, ListSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in. | 
**ListSid** | **string**| The SID of the Sync List to add the new List Item to. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***CreateSyncListItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length. | 
**ItemTtl** | **Int32**| How long, in seconds, before the List Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the List Item does not expire. The List Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

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


## CreateSyncMap

> SyncV1ServiceSyncMap CreateSyncMap(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in. | 
 **optional** | ***CreateSyncMapRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, ServiceSid, MapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in. | 
**MapSid** | **string**| The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***CreateSyncMapItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length. | 
**ItemTtl** | **Int32**| How long, in seconds, before the Map Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Map Item does not expire.  The Map Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Key** | **String**| The unique, user-defined key for the Map Item. Can be up to 320 characters long. | 
**Ttl** | **Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

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


## CreateSyncStream

> SyncV1ServiceSyncStream CreateSyncStream(ctx, ServiceSid, optional)



Create a new Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in. | 
 **optional** | ***CreateSyncStreamRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncStreamRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Ttl** | **Int32**| How long, in seconds, before the Stream expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Stream does not expire. The Stream will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**UniqueName** | **String**| An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete. | 
**Sid** | **string**| The SID of the Document resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

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

> DeleteDocumentPermission(ctx, ServiceSid, DocumentSid, Identity)



Delete a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete. | 
**DocumentSid** | **string**| The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to delete. | 

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Service resource to delete. | 

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

> DeleteSyncList(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete. | 
**Sid** | **string**| The SID of the Sync List resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

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

> DeleteSyncListItem(ctx, ServiceSid, ListSid, Index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Index** | **int32**| The index of the Sync List Item resource to delete. | 
 **optional** | ***DeleteSyncListItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncListItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **String**| If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). | 

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

> DeleteSyncListPermission(ctx, ServiceSid, ListSid, Identity)



Delete a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to delete. | 

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

> DeleteSyncMap(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete. | 
**Sid** | **string**| The SID of the Sync Map resource to delete. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

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

> DeleteSyncMapItem(ctx, ServiceSid, MapSid, Key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to delete. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to delete. | 
 **optional** | ***DeleteSyncMapItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncMapItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **String**| If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). | 

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

> DeleteSyncMapPermission(ctx, ServiceSid, MapSid, Identity)



Delete a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to delete. | 

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

> DeleteSyncStream(ctx, ServiceSid, Sid)



Delete a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to delete. | 
**Sid** | **string**| The SID of the Stream resource to delete. | 

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

> SyncV1ServiceDocument FetchDocument(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch. | 
**Sid** | **string**| The SID of the Document resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocumentPermission

> SyncV1ServiceDocumentDocumentPermission FetchDocumentPermission(ctx, ServiceSid, DocumentSid, Identity)



Fetch a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch. | 
**DocumentSid** | **string**| The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to fetch. | 

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](SyncV1ServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> SyncV1Service FetchService(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Service resource to fetch. | 

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncList

> SyncV1ServiceSyncList FetchSyncList(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch. | 
**Sid** | **string**| The SID of the Sync List resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListItem

> SyncV1ServiceSyncListSyncListItem FetchSyncListItem(ctx, ServiceSid, ListSid, Index)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Index** | **int32**| The index of the Sync List Item resource to fetch. | 

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


## FetchSyncListPermission

> SyncV1ServiceSyncListSyncListPermission FetchSyncListPermission(ctx, ServiceSid, ListSid, Identity)



Fetch a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to fetch. | 

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](SyncV1ServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMap

> SyncV1ServiceSyncMap FetchSyncMap(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch. | 
**Sid** | **string**| The SID of the Sync Map resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, ServiceSid, MapSid, Key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to fetch. | 

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


## FetchSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, ServiceSid, MapSid, Identity)



Fetch a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to fetch. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](SyncV1ServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncStream

> SyncV1ServiceSyncStream FetchSyncStream(ctx, ServiceSid, Sid)



Fetch a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to fetch. | 
**Sid** | **string**| The SID of the Stream resource to fetch. | 

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> ListDocumentResponse ListDocument(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read. | 
 **optional** | ***ListDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDocumentResponse**](ListDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> ListDocumentPermissionResponse ListDocumentPermission(ctx, ServiceSid, DocumentSid, optional)



Retrieve a list of all Permissions applying to a Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read. | 
**DocumentSid** | **string**| The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListDocumentPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDocumentPermissionResponse**](ListDocumentPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> ListSyncListResponse ListSyncList(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read. | 
 **optional** | ***ListSyncListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncListResponse**](ListSyncListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> ListSyncListItemResponse ListSyncListItem(ctx, ServiceSid, ListSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read. | 
**ListSid** | **string**| The SID of the Sync List with the List Items to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncListItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Order** | **String**| How to order the List Items returned by their &#x60;index&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending. | 
**From** | **String**| The &#x60;index&#x60; of the first Sync List Item resource to read. See also &#x60;bounds&#x60;. | 
**Bounds** | **String**| Whether to include the List Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the List Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next List Item. The default value is &#x60;inclusive&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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


## ListSyncListPermission

> ListSyncListPermissionResponse ListSyncListPermission(ctx, ServiceSid, ListSid, optional)



Retrieve a list of all Permissions applying to a Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncListPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncListPermissionResponse**](ListSyncListPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> ListSyncMapResponse ListSyncMap(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read. | 
 **optional** | ***ListSyncMapRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncMapResponse**](ListSyncMapResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> ListSyncMapItemResponse ListSyncMapItem(ctx, ServiceSid, MapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncMapItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Order** | **String**| How to order the Map Items returned by their &#x60;key&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key. | 
**From** | **String**| The &#x60;key&#x60; of the first Sync Map Item resource to read. See also &#x60;bounds&#x60;. | 
**Bounds** | **String**| Whether to include the Map Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the Map Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next Map Item. The default value is &#x60;inclusive&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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


## ListSyncMapPermission

> ListSyncMapPermissionResponse ListSyncMapPermission(ctx, ServiceSid, MapSid, optional)



Retrieve a list of all Permissions applying to a Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**MapSid** | **string**| The SID of the Sync Map with the Permission resources to read. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncMapPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncMapPermissionResponse**](ListSyncMapPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncStream

> ListSyncStreamResponse ListSyncStream(ctx, ServiceSid, optional)



Retrieve a list of all Streams in a Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Stream resources to read. | 
 **optional** | ***ListSyncStreamRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncStreamRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncStreamResponse**](ListSyncStreamResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> SyncV1ServiceDocument UpdateDocument(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update. | 
**Sid** | **string**| The SID of the Document resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***UpdateDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **String**| The If-Match HTTP request header | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length. | 
**Ttl** | **Int32**| How long, in seconds, before the Sync Document expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Document resource does not expire. The Document resource will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 

### Return type

[**SyncV1ServiceDocument**](SyncV1ServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> SyncV1ServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, ServiceSid, DocumentSid, Identity, optional)



Update an identity's access to a specific Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update. | 
**DocumentSid** | **string**| The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to update. | 
 **optional** | ***UpdateDocumentPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **Bool**| Whether the identity can delete the Sync Document. Default value is &#x60;false&#x60;. | 
**Read** | **Bool**| Whether the identity can read the Sync Document. Default value is &#x60;false&#x60;. | 
**Write** | **Bool**| Whether the identity can update the Sync Document. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](SyncV1ServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> SyncV1Service UpdateService(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Service resource to update. | 
 **optional** | ***UpdateServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AclEnabled** | **Bool**| Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. | 
**FriendlyName** | **String**| A string that you assign to describe the resource. | 
**ReachabilityDebouncingEnabled** | **Bool**| Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event. | 
**ReachabilityDebouncingWindow** | **Int32**| The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called. | 
**ReachabilityWebhooksEnabled** | **Bool**| Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;. | 
**WebhookUrl** | **String**| The URL we should call when Sync objects are manipulated. | 
**WebhooksFromRestEnabled** | **Bool**| Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;. | 

### Return type

[**SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncList

> SyncV1ServiceSyncList UpdateSyncList(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update. | 
**Sid** | **string**| The SID of the Sync List resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***UpdateSyncListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the Sync List expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;collection_ttl&#x60;. If both are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncList**](SyncV1ServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> SyncV1ServiceSyncListSyncListItem UpdateSyncListItem(ctx, ServiceSid, ListSid, Index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Index** | **int32**| The index of the Sync List Item resource to update. | 
 **optional** | ***UpdateSyncListItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **String**| If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). | 
**CollectionTtl** | **Int32**| How long, in seconds, before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length. | 
**ItemTtl** | **Int32**| How long, in seconds, before the List Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the List Item does not expire. The List Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

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


## UpdateSyncListPermission

> SyncV1ServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, ServiceSid, ListSid, Identity, optional)



Update an identity's access to a specific Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update. | 
**ListSid** | **string**| The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to update. | 
 **optional** | ***UpdateSyncListPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **Bool**| Whether the identity can delete the Sync List. Default value is &#x60;false&#x60;. | 
**Read** | **Bool**| Whether the identity can read the Sync List and its Items. Default value is &#x60;false&#x60;. | 
**Write** | **Bool**| Whether the identity can create, update, and delete Items in the Sync List. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](SyncV1ServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMap

> SyncV1ServiceSyncMap UpdateSyncMap(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update. | 
**Sid** | **string**| The SID of the Sync Map resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***UpdateSyncMapRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CollectionTtl** | **Int32**| How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncMap**](SyncV1ServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, ServiceSid, MapSid, Key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to update.  | 
 **optional** | ***UpdateSyncMapItemRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapItemRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **String**| If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). | 
**CollectionTtl** | **Int32**| How long, in seconds, before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync Map does not expire. This parameter can only be used when the Map Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length. | 
**ItemTtl** | **Int32**| How long, in seconds, before the Map Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Map Item does not expire. The Map Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
**Ttl** | **Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

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


## UpdateSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, ServiceSid, MapSid, Identity, optional)



Update an identity's access to a specific Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**MapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**Identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to update. | 
 **optional** | ***UpdateSyncMapPermissionRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapPermissionRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **Bool**| Whether the identity can delete the Sync Map. Default value is &#x60;false&#x60;. | 
**Read** | **Bool**| Whether the identity can read the Sync Map and its Items. Default value is &#x60;false&#x60;. | 
**Write** | **Bool**| Whether the identity can create, update, and delete Items in the Sync Map. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](SyncV1ServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncStream

> SyncV1ServiceSyncStream UpdateSyncStream(ctx, ServiceSid, Sid, optional)



Update a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to update. | 
**Sid** | **string**| The SID of the Stream resource to update. | 
 **optional** | ***UpdateSyncStreamRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncStreamRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Ttl** | **Int32**| How long, in seconds, before the Stream expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Stream does not expire. The Stream will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 

### Return type

[**SyncV1ServiceSyncStream**](SyncV1ServiceSyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

