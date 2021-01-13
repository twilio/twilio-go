# \DefaultApi

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

> SyncV1ServiceDocument CreateDocument(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in. | 
 **optional** | ***CreateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16KB in length. | 
 **ttl** | **optional.Int32**| How long, in seconds, before the Sync Document expires and is deleted (the Sync Document&#39;s time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Document does not expire. The Sync Document will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the Sync Document | 

### Return type

[**SyncV1ServiceDocument**](sync.v1.service.document.md)

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
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclEnabled** | **optional.Bool**| Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. | 
 **friendlyName** | **optional.String**| A string that you assign to describe the resource. | 
 **reachabilityDebouncingEnabled** | **optional.Bool**| Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event. | 
 **reachabilityDebouncingWindow** | **optional.Int32**| The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the &#x60;webhook_url&#x60; is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to &#x60;webhook_url&#x60;. | 
 **reachabilityWebhooksEnabled** | **optional.Bool**| Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;. | 
 **webhookUrl** | **optional.String**| The URL we should call when Sync objects are manipulated. | 
 **webhooksFromRestEnabled** | **optional.Bool**| Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;. | 

### Return type

[**SyncV1Service**](sync.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamMessage

> SyncV1ServiceSyncStreamStreamMessage CreateStreamMessage(ctx, serviceSid, streamSid, optional)



Create a new Stream Message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in. | 
**streamSid** | **string**| The SID of the Sync Stream to create the new Stream Message resource for. | 
 **optional** | ***CreateStreamMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateStreamMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4KB in length. | 

### Return type

[**SyncV1ServiceSyncStreamStreamMessage**](sync.v1.service.sync_stream.stream_message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncList

> SyncV1ServiceSyncList CreateSyncList(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in. | 
 **optional** | ***CreateSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| Alias for collection_ttl. If both are provided, this value is ignored. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncList**](sync.v1.service.sync_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncListItem

> SyncV1ServiceSyncListSyncListItem CreateSyncListItem(ctx, serviceSid, listSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in. | 
**listSid** | **string**| The SID of the Sync List to add the new List Item to. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***CreateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionTtl** | **optional.Int32**| How long, in seconds, before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16KB in length. | 
 **itemTtl** | **optional.Int32**| How long, in seconds, before the List Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the List Item does not expire. The List Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncListSyncListItem**](sync.v1.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMap

> SyncV1ServiceSyncMap CreateSyncMap(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in. | 
 **optional** | ***CreateSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncMap**](sync.v1.service.sync_map.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, serviceSid, mapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in. | 
**mapSid** | **string**| The SID of the Sync Map to add the new Map Item to. | 
 **optional** | ***CreateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16KB in length. | 
 **itemTtl** | **optional.Int32**| How long, in seconds, before the Map Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Map Item does not expire.  The Map Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **key** | **optional.String**| The unique, user-defined key for the Map Item. Can be up to 320 characters long. | 
 **ttl** | **optional.Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](sync.v1.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncStream

> SyncV1ServiceSyncStream CreateSyncStream(ctx, serviceSid, optional)



Create a new Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in. | 
 **optional** | ***CreateSyncStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ttl** | **optional.Int32**| How long, in seconds, before the Stream expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Stream does not expire. The Stream will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. | 

### Return type

[**SyncV1ServiceSyncStream**](sync.v1.service.sync_stream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete. | 
**sid** | **string**| The SID of the Document resource to delete. | 
 **optional** | ***DeleteDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteDocumentPermission(ctx, serviceSid, documentSid, identity)



Delete a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete. | 
**documentSid** | **string**| The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to delete. | 

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

> DeleteService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to delete. | 

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

> DeleteSyncList(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete. | 
**sid** | **string**| The SID of the Sync List resource to delete. | 

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

> DeleteSyncListItem(ctx, serviceSid, listSid, index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**index** | **int32**| The index of the Sync List Item resource to delete. | 
 **optional** | ***DeleteSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteSyncListPermission(ctx, serviceSid, listSid, identity)



Delete a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to delete. | 

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

> DeleteSyncMap(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete. | 
**sid** | **string**| The SID of the Sync Map resource to delete. | 

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

> DeleteSyncMapItem(ctx, serviceSid, mapSid, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to delete. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to delete. | 
 **optional** | ***DeleteSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteSyncMapPermission(ctx, serviceSid, mapSid, identity)



Delete a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to delete. | 

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

> DeleteSyncStream(ctx, serviceSid, sid)



Delete a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to delete. | 
**sid** | **string**| The SID of the Stream resource to delete. | 

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

> SyncV1ServiceDocument FetchDocument(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch. | 
**sid** | **string**| The SID of the Document resource to fetch. | 

### Return type

[**SyncV1ServiceDocument**](sync.v1.service.document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocumentPermission

> SyncV1ServiceDocumentDocumentPermission FetchDocumentPermission(ctx, serviceSid, documentSid, identity)



Fetch a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch. | 
**documentSid** | **string**| The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to fetch. | 

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](sync.v1.service.document.document_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> SyncV1Service FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to fetch. | 

### Return type

[**SyncV1Service**](sync.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncList

> SyncV1ServiceSyncList FetchSyncList(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch. | 
**sid** | **string**| The SID of the Sync List resource to fetch. | 

### Return type

[**SyncV1ServiceSyncList**](sync.v1.service.sync_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListItem

> SyncV1ServiceSyncListSyncListItem FetchSyncListItem(ctx, serviceSid, listSid, index)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**index** | **int32**| The index of the Sync List Item resource to fetch. | 

### Return type

[**SyncV1ServiceSyncListSyncListItem**](sync.v1.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListPermission

> SyncV1ServiceSyncListSyncListPermission FetchSyncListPermission(ctx, serviceSid, listSid, identity)



Fetch a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to fetch. | 

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](sync.v1.service.sync_list.sync_list_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMap

> SyncV1ServiceSyncMap FetchSyncMap(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch. | 
**sid** | **string**| The SID of the Sync Map resource to fetch. | 

### Return type

[**SyncV1ServiceSyncMap**](sync.v1.service.sync_map.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, serviceSid, mapSid, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to fetch. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](sync.v1.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, serviceSid, mapSid, identity)



Fetch a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to fetch. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](sync.v1.service.sync_map.sync_map_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncStream

> SyncV1ServiceSyncStream FetchSyncStream(ctx, serviceSid, sid)



Fetch a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to fetch. | 
**sid** | **string**| The SID of the Stream resource to fetch. | 

### Return type

[**SyncV1ServiceSyncStream**](sync.v1.service.sync_stream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> SyncV1ServiceDocumentReadResponse ListDocument(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read. | 
 **optional** | ***ListDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hideExpired** | **optional.String**| The default list of Sync Documents will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceDocumentReadResponse**](sync_v1_service_documentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> SyncV1ServiceDocumentDocumentPermissionReadResponse ListDocumentPermission(ctx, serviceSid, documentSid, optional)



Retrieve a list of all Permissions applying to a Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read. | 
**documentSid** | **string**| The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceDocumentDocumentPermissionReadResponse**](sync_v1_service_document_document_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> SyncV1ServiceReadResponse ListService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceReadResponse**](sync_v1_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> SyncV1ServiceSyncListReadResponse ListSyncList(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read. | 
 **optional** | ***ListSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hideExpired** | **optional.String**| The default list of Sync Lists will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncListReadResponse**](sync_v1_service_sync_listReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> SyncV1ServiceSyncListSyncListItemReadResponse ListSyncListItem(ctx, serviceSid, listSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read. | 
**listSid** | **string**| The SID of the Sync List with the List Items to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**| How to order the List Items returned by their &#x60;index&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending. | 
 **from** | **optional.String**| The &#x60;index&#x60; of the first Sync List Item resource to read. See also &#x60;bounds&#x60;. | 
 **bounds** | **optional.String**| Whether to include the List Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the List Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next List Item. The default value is &#x60;inclusive&#x60;. | 
 **hideExpired** | **optional.String**| The default list of Sync List items will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncListSyncListItemReadResponse**](sync_v1_service_sync_list_sync_list_itemReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> SyncV1ServiceSyncListSyncListPermissionReadResponse ListSyncListPermission(ctx, serviceSid, listSid, optional)



Retrieve a list of all Permissions applying to a Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncListSyncListPermissionReadResponse**](sync_v1_service_sync_list_sync_list_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> SyncV1ServiceSyncMapReadResponse ListSyncMap(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read. | 
 **optional** | ***ListSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hideExpired** | **optional.String**| The default list of Sync Maps will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncMapReadResponse**](sync_v1_service_sync_mapReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> SyncV1ServiceSyncMapSyncMapItemReadResponse ListSyncMapItem(ctx, serviceSid, mapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**| How to order the Map Items returned by their &#x60;key&#x60; value. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key. | 
 **from** | **optional.String**| The &#x60;key&#x60; of the first Sync Map Item resource to read. See also &#x60;bounds&#x60;. | 
 **bounds** | **optional.String**| Whether to include the Map Item referenced by the &#x60;from&#x60; parameter. Can be: &#x60;inclusive&#x60; to include the Map Item referenced by the &#x60;from&#x60; parameter or &#x60;exclusive&#x60; to start with the next Map Item. The default value is &#x60;inclusive&#x60;. | 
 **hideExpired** | **optional.String**| The default list of Sync Map items will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapItemReadResponse**](sync_v1_service_sync_map_sync_map_itemReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermissionReadResponse ListSyncMapPermission(ctx, serviceSid, mapSid, optional)



Retrieve a list of all Permissions applying to a Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**mapSid** | **string**| The SID of the Sync Map with the Permission resources to read. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
 **optional** | ***ListSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapPermissionReadResponse**](sync_v1_service_sync_map_sync_map_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncStream

> SyncV1ServiceSyncStreamReadResponse ListSyncStream(ctx, serviceSid, optional)



Retrieve a list of all Streams in a Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Stream resources to read. | 
 **optional** | ***ListSyncStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hideExpired** | **optional.String**| The default list of Sync Streams will show both active and expired items. It is possible to filter only the active ones by hiding the expired ones. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SyncV1ServiceSyncStreamReadResponse**](sync_v1_service_sync_streamReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> SyncV1ServiceDocument UpdateDocument(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update. | 
**sid** | **string**| The SID of the Document resource to update. | 
 **optional** | ***UpdateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16KB in length. | 
 **ttl** | **optional.Int32**| How long, in seconds, before the Sync Document expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Document resource does not expire. The Document resource will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 

### Return type

[**SyncV1ServiceDocument**](sync.v1.service.document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> SyncV1ServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, serviceSid, documentSid, identity, optional)



Update an identity's access to a specific Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update. | 
**documentSid** | **string**| The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Document Permission resource to update. | 
 **optional** | ***UpdateDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Whether the identity can delete the Sync Document. Default value is &#x60;false&#x60;. | 
 **read** | **optional.Bool**| Whether the identity can read the Sync Document. Default value is &#x60;false&#x60;. | 
 **write** | **optional.Bool**| Whether the identity can update the Sync Document. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceDocumentDocumentPermission**](sync.v1.service.document.document_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> SyncV1Service UpdateService(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Service resource to update. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aclEnabled** | **optional.Bool**| Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. | 
 **friendlyName** | **optional.String**| A string that you assign to describe the resource. | 
 **reachabilityDebouncingEnabled** | **optional.Bool**| Whether every &#x60;endpoint_disconnected&#x60; event should occur after a configurable delay. The default is &#x60;false&#x60;, where the &#x60;endpoint_disconnected&#x60; event occurs immediately after disconnection. When &#x60;true&#x60;, intervening reconnections can prevent the &#x60;endpoint_disconnected&#x60; event. | 
 **reachabilityDebouncingWindow** | **optional.Int32**| The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called. | 
 **reachabilityWebhooksEnabled** | **optional.Bool**| Whether the service instance should call &#x60;webhook_url&#x60; when client endpoints connect to Sync. The default is &#x60;false&#x60;. | 
 **webhookUrl** | **optional.String**| The URL we should call when Sync objects are manipulated. | 
 **webhooksFromRestEnabled** | **optional.Bool**| Whether the Service instance should call &#x60;webhook_url&#x60; when the REST API is used to update Sync objects. The default is &#x60;false&#x60;. | 

### Return type

[**SyncV1Service**](sync.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncList

> SyncV1ServiceSyncList UpdateSyncList(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update. | 
**sid** | **string**| The SID of the Sync List resource to update. | 
 **optional** | ***UpdateSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Sync List expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;collection_ttl&#x60;. If both are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncList**](sync.v1.service.sync_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> SyncV1ServiceSyncListSyncListItem UpdateSyncListItem(ctx, serviceSid, listSid, index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**index** | **int32**| The index of the Sync List Item resource to update. | 
 **optional** | ***UpdateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **collectionTtl** | **optional.Int32**| How long, in seconds, before the List Item&#39;s parent Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16KB in length. | 
 **itemTtl** | **optional.Int32**| How long, in seconds, before the List Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the List Item does not expire. The List Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncListSyncListItem**](sync.v1.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> SyncV1ServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, serviceSid, listSid, identity, optional)



Update an identity's access to a specific Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update. | 
**listSid** | **string**| The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to update. | 
 **optional** | ***UpdateSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Whether the identity can delete the Sync List. Default value is &#x60;false&#x60;. | 
 **read** | **optional.Bool**| Whether the identity can read the Sync List and its Items. Default value is &#x60;false&#x60;. | 
 **write** | **optional.Bool**| Whether the identity can create, update, and delete Items in the Sync List. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceSyncListSyncListPermission**](sync.v1.service.sync_list.sync_list_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMap

> SyncV1ServiceSyncMap UpdateSyncMap(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update. | 
**sid** | **string**| The SID of the Sync Map resource to update. | 
 **optional** | ***UpdateSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncMap**](sync.v1.service.sync_map.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, serviceSid, mapSid, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**key** | **string**| The &#x60;key&#x60; value of the Sync Map Item resource to update.  | 
 **optional** | ***UpdateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **collectionTtl** | **optional.Int32**| How long, in seconds, before the Map Item&#39;s parent Sync Map expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the parent Sync Map does not expire. This parameter can only be used when the Map Item&#39;s &#x60;data&#x60; or &#x60;ttl&#x60; is updated in the same request. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16KB in length. | 
 **itemTtl** | **optional.Int32**| How long, in seconds, before the Map Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Map Item does not expire. The Map Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 
 **ttl** | **optional.Int32**| An alias for &#x60;item_ttl&#x60;. If both parameters are provided, this value is ignored. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapItem**](sync.v1.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, serviceSid, mapSid, identity, optional)



Update an identity's access to a specific Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;. | 
**mapSid** | **string**| The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;. | 
**identity** | **string**| The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to update. | 
 **optional** | ***UpdateSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Whether the identity can delete the Sync Map. Default value is &#x60;false&#x60;. | 
 **read** | **optional.Bool**| Whether the identity can read the Sync Map and its Items. Default value is &#x60;false&#x60;. | 
 **write** | **optional.Bool**| Whether the identity can create, update, and delete Items in the Sync Map. Default value is &#x60;false&#x60;. | 

### Return type

[**SyncV1ServiceSyncMapSyncMapPermission**](sync.v1.service.sync_map.sync_map_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncStream

> SyncV1ServiceSyncStream UpdateSyncStream(ctx, serviceSid, sid, optional)



Update a specific Stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to update. | 
**sid** | **string**| The SID of the Stream resource to update. | 
 **optional** | ***UpdateSyncStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ttl** | **optional.Int32**| How long, in seconds, before the Stream expires and is deleted (time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is &#x60;0&#x60;, which means the Stream does not expire. The Stream will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources&#39;s deletion. | 

### Return type

[**SyncV1ServiceSyncStream**](sync.v1.service.sync_stream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

