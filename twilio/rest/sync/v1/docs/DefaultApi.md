# \DefaultApi

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

> SyncV1ServiceDocument CreateDocument(ctx, ServiceSid).Data(Data).Ttl(Ttl).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length. (optional)
    Ttl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document's time-to-live). (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the Sync Document (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDocument(context.Background(), ServiceSid).Data(Data).Ttl(Ttl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: SyncV1ServiceDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Document resource in.

### Other Parameters

Other parameters are passed through a pointer to a CreateDocumentParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> SyncV1Service CreateService(ctx).AclEnabled(AclEnabled).FriendlyName(FriendlyName).ReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled).ReachabilityDebouncingWindow(ReachabilityDebouncingWindow).ReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled).WebhookUrl(WebhookUrl).WebhooksFromRestEnabled(WebhooksFromRestEnabled).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AclEnabled := true // bool | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that you assign to describe the resource. (optional)
    ReachabilityDebouncingEnabled := true // bool | Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event. (optional)
    ReachabilityDebouncingWindow := int32(56) // int32 | The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the `webhook_url` is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to `webhook_url`. (optional)
    ReachabilityWebhooksEnabled := true // bool | Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`. (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL we should call when Sync objects are manipulated. (optional)
    WebhooksFromRestEnabled := true // bool | Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).AclEnabled(AclEnabled).FriendlyName(FriendlyName).ReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled).ReachabilityDebouncingWindow(ReachabilityDebouncingWindow).ReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled).WebhookUrl(WebhookUrl).WebhooksFromRestEnabled(WebhooksFromRestEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: SyncV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamMessage

> SyncV1ServiceSyncStreamStreamMessage CreateStreamMessage(ctx, ServiceSid, StreamSid).Data(Data).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in.
    StreamSid := "StreamSid_example" // string | The SID of the Sync Stream to create the new Stream Message resource for.
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateStreamMessage(context.Background(), ServiceSid, StreamSid).Data(Data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStreamMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamMessage`: SyncV1ServiceSyncStreamStreamMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStreamMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in.
**StreamSid** | **string** | The SID of the Sync Stream to create the new Stream Message resource for.

### Other Parameters

Other parameters are passed through a pointer to a CreateStreamMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Data** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.

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

> SyncV1ServiceSyncList CreateSyncList(ctx, ServiceSid).CollectionTtl(CollectionTtl).Ttl(Ttl).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | Alias for collection_ttl. If both are provided, this value is ignored. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSyncList(context.Background(), ServiceSid).CollectionTtl(CollectionTtl).Ttl(Ttl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncList`: SyncV1ServiceSyncList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSyncList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Sync List in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncListItem

> SyncV1ServiceSyncListSyncListItem CreateSyncListItem(ctx, ServiceSid, ListSid).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
    ListSid := "ListSid_example" // string | The SID of the Sync List to add the new List Item to. Can be the Sync List resource's `sid` or its `unique_name`.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted. (optional)
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length. (optional)
    ItemTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `item_ttl`. If both parameters are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSyncListItem(context.Background(), ServiceSid, ListSid).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSyncListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncListItem`: SyncV1ServiceSyncListSyncListItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSyncListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new List Item in.
**ListSid** | **string** | The SID of the Sync List to add the new List Item to. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListItemParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMap

> SyncV1ServiceSyncMap CreateSyncMap(ctx, ServiceSid).CollectionTtl(CollectionTtl).Ttl(Ttl).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `collection_ttl`. If both parameters are provided, this value is ignored. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSyncMap(context.Background(), ServiceSid).CollectionTtl(CollectionTtl).Ttl(Ttl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSyncMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncMap`: SyncV1ServiceSyncMap
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSyncMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Sync Map in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMapItem

> SyncV1ServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, ServiceSid, MapSid).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Key(Key).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
    MapSid := "MapSid_example" // string | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource's `sid` or its `unique_name`.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted. (optional)
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length. (optional)
    ItemTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted. (optional)
    Key := "Key_example" // string | The unique, user-defined key for the Map Item. Can be up to 320 characters long. (optional)
    Ttl := int32(56) // int32 | An alias for `item_ttl`. If both parameters are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSyncMapItem(context.Background(), ServiceSid, MapSid).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Key(Key).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSyncMapItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncMapItem`: SyncV1ServiceSyncMapSyncMapItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSyncMapItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the Map Item in.
**MapSid** | **string** | The SID of the Sync Map to add the new Map Item to. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapItemParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncStream

> SyncV1ServiceSyncStream CreateSyncStream(ctx, ServiceSid).Ttl(Ttl).UniqueName(UniqueName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in.
    Ttl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live). (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSyncStream(context.Background(), ServiceSid).Ttl(Ttl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSyncStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncStream`: SyncV1ServiceSyncStream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSyncStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream in.

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncStreamParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The &#x60;unique_name&#x60; value can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource.

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

> DeleteDocument(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete.
    Sid := "Sid_example" // string | The SID of the Document resource to delete. Can be the Document resource's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteDocument(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to delete.
**Sid** | **string** | The SID of the Document resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentParams struct via the builder pattern


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

> DeleteDocumentPermission(ctx, ServiceSid, DocumentSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete.
    DocumentSid := "DocumentSid_example" // string | The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Document Permission resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteDocumentPermission(context.Background(), ServiceSid, DocumentSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDocumentPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to delete.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to delete. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentPermissionParams struct via the builder pattern


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

> DeleteService(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Service resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct via the builder pattern


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

> DeleteSyncList(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete.
    Sid := "Sid_example" // string | The SID of the Sync List resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncList(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to delete.
**Sid** | **string** | The SID of the Sync List resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListParams struct via the builder pattern


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

> DeleteSyncListItem(ctx, ServiceSid, ListSid, Index).IfMatch(IfMatch).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.
    Index := int32(56) // int32 | The index of the Sync List Item resource to delete.
    IfMatch := "IfMatch_example" // string | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncListItem(context.Background(), ServiceSid, ListSid, Index).IfMatch(IfMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int32** | The index of the Sync List Item resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListItemParams struct via the builder pattern


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

> DeleteSyncListPermission(ctx, ServiceSid, ListSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync List Permission resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncListPermission(context.Background(), ServiceSid, ListSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to delete.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to delete. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListPermissionParams struct via the builder pattern


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

> DeleteSyncMap(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete.
    Sid := "Sid_example" // string | The SID of the Sync Map resource to delete. Can be the Sync Map's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncMap(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to delete.
**Sid** | **string** | The SID of the Sync Map resource to delete. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapParams struct via the builder pattern


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

> DeleteSyncMapItem(ctx, ServiceSid, MapSid, Key).IfMatch(IfMatch).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to delete.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map resource's `sid` or its `unique_name`.
    Key := "Key_example" // string | The `key` value of the Sync Map Item resource to delete.
    IfMatch := "IfMatch_example" // string | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncMapItem(context.Background(), ServiceSid, MapSid, Key).IfMatch(IfMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncMapItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to delete.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapItemParams struct via the builder pattern


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

> DeleteSyncMapPermission(ctx, ServiceSid, MapSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service's `sid` value or `default`.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync Map Permission resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncMapPermission(context.Background(), ServiceSid, MapSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncMapPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to delete. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to delete. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapPermissionParams struct via the builder pattern


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

> DeleteSyncStream(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to delete.
    Sid := "Sid_example" // string | The SID of the Stream resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSyncStream(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSyncStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to delete.
**Sid** | **string** | The SID of the Stream resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncStreamParams struct via the builder pattern


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

> SyncV1ServiceDocument FetchDocument(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch.
    Sid := "Sid_example" // string | The SID of the Document resource to fetch. Can be the Document resource's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDocument(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDocument`: SyncV1ServiceDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to fetch.
**Sid** | **string** | The SID of the Document resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> SyncV1ServiceDocumentDocumentPermission FetchDocumentPermission(ctx, ServiceSid, DocumentSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch.
    DocumentSid := "DocumentSid_example" // string | The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Document Permission resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDocumentPermission(context.Background(), ServiceSid, DocumentSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDocumentPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDocumentPermission`: SyncV1ServiceDocumentDocumentPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDocumentPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to fetch.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to fetch. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> SyncV1Service FetchService(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Service resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: SyncV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> SyncV1ServiceSyncList FetchSyncList(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch.
    Sid := "Sid_example" // string | The SID of the Sync List resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncList(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncList`: SyncV1ServiceSyncList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to fetch.
**Sid** | **string** | The SID of the Sync List resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> SyncV1ServiceSyncListSyncListItem FetchSyncListItem(ctx, ServiceSid, ListSid, Index).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.
    Index := int32(56) // int32 | The index of the Sync List Item resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncListItem(context.Background(), ServiceSid, ListSid, Index).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncListItem`: SyncV1ServiceSyncListSyncListItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int32** | The index of the Sync List Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListItemParams struct via the builder pattern


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


## FetchSyncListPermission

> SyncV1ServiceSyncListSyncListPermission FetchSyncListPermission(ctx, ServiceSid, ListSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync List Permission resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncListPermission(context.Background(), ServiceSid, ListSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncListPermission`: SyncV1ServiceSyncListSyncListPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncListPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to fetch.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to fetch. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> SyncV1ServiceSyncMap FetchSyncMap(ctx, ServiceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch.
    Sid := "Sid_example" // string | The SID of the Sync Map resource to fetch. Can be the Sync Map's `sid` or its `unique_name`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncMap(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncMap`: SyncV1ServiceSyncMap
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to fetch.
**Sid** | **string** | The SID of the Sync Map resource to fetch. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> SyncV1ServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, ServiceSid, MapSid, Key).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.
    Key := "Key_example" // string | The `key` value of the Sync Map Item resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncMapItem(context.Background(), ServiceSid, MapSid, Key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncMapItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncMapItem`: SyncV1ServiceSyncMapSyncMapItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncMapItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to fetch.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapItemParams struct via the builder pattern


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


## FetchSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, ServiceSid, MapSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service's `sid` value or `default`.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync Map Permission resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncMapPermission(context.Background(), ServiceSid, MapSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncMapPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncMapPermission`: SyncV1ServiceSyncMapSyncMapPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncMapPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to fetch. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> SyncV1ServiceSyncStream FetchSyncStream(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to fetch.
    Sid := "Sid_example" // string | The SID of the Stream resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSyncStream(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSyncStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSyncStream`: SyncV1ServiceSyncStream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSyncStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to fetch.
**Sid** | **string** | The SID of the Stream resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncStreamParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> ListDocumentResponse ListDocument(ctx, ServiceSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDocument(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocument`: ListDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListDocumentPermissionResponse ListDocumentPermission(ctx, ServiceSid, DocumentSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read.
    DocumentSid := "DocumentSid_example" // string | The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource's `sid` or its `unique_name`.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDocumentPermission(context.Background(), ServiceSid, DocumentSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDocumentPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocumentPermission`: ListDocumentPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDocumentPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resources to read.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resources to read. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListServiceResponse ListService(ctx).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListService(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListService`: ListServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSyncListResponse ListSyncList(ctx, ServiceSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncList(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncList`: ListSyncListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSyncListItemResponse ListSyncListItem(ctx, ServiceSid, ListSid).Order(Order).From(From).Bounds(Bounds).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the List Items to read. Can be the Sync List resource's `sid` or its `unique_name`.
    Order := "Order_example" // string | How to order the List Items returned by their `index` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending. (optional)
    From := "From_example" // string | The `index` of the first Sync List Item resource to read. See also `bounds`. (optional)
    Bounds := "Bounds_example" // string | Whether to include the List Item referenced by the `from` parameter. Can be: `inclusive` to include the List Item referenced by the `from` parameter or `exclusive` to start with the next List Item. The default value is `inclusive`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncListItem(context.Background(), ServiceSid, ListSid).Order(Order).From(From).Bounds(Bounds).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncListItem`: ListSyncListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the List Item resources to read.
**ListSid** | **string** | The SID of the Sync List with the List Items to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListItemParams struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> ListSyncListPermissionResponse ListSyncListPermission(ctx, ServiceSid, ListSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource's `sid` or its `unique_name`.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncListPermission(context.Background(), ServiceSid, ListSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncListPermission`: ListSyncListPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncListPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resources to read.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resources to read. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSyncMapResponse ListSyncMap(ctx, ServiceSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncMap(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncMap`: ListSyncMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSyncMapItemResponse ListSyncMapItem(ctx, ServiceSid, MapSid).Order(Order).From(From).Bounds(Bounds).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource's `sid` or its `unique_name`.
    Order := "Order_example" // string | How to order the Map Items returned by their `key` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key. (optional)
    From := "From_example" // string | The `key` of the first Sync Map Item resource to read. See also `bounds`. (optional)
    Bounds := "Bounds_example" // string | Whether to include the Map Item referenced by the `from` parameter. Can be: `inclusive` to include the Map Item referenced by the `from` parameter or `exclusive` to start with the next Map Item. The default value is `inclusive`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncMapItem(context.Background(), ServiceSid, MapSid).Order(Order).From(From).Bounds(Bounds).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncMapItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncMapItem`: ListSyncMapItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncMapItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Map Item resources to read.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to fetch. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapItemParams struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> ListSyncMapPermissionResponse ListSyncMapPermission(ctx, ServiceSid, MapSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service's `sid` value or `default`.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Permission resources to read. Can be the Sync Map resource's `sid` or its `unique_name`.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncMapPermission(context.Background(), ServiceSid, MapSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncMapPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncMapPermission`: ListSyncMapPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncMapPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resources to read. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Permission resources to read. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapPermissionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSyncStreamResponse ListSyncStream(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Stream resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSyncStream(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSyncStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSyncStream`: ListSyncStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSyncStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Stream resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSyncStreamParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> SyncV1ServiceDocument UpdateDocument(ctx, ServiceSid, Sid).IfMatch(IfMatch).Data(Data).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update.
    Sid := "Sid_example" // string | The SID of the Document resource to update. Can be the Document resource's `sid` or its `unique_name`.
    IfMatch := "IfMatch_example" // string | The If-Match HTTP request header (optional)
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length. (optional)
    Ttl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (time-to-live). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateDocument(context.Background(), ServiceSid, Sid).IfMatch(IfMatch).Data(Data).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: SyncV1ServiceDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document resource to update.
**Sid** | **string** | The SID of the Document resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> SyncV1ServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, ServiceSid, DocumentSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update.
    DocumentSid := "DocumentSid_example" // string | The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Document Permission resource to update.
    Manage := true // bool | Whether the identity can delete the Sync Document. Default value is `false`. (optional)
    Read := true // bool | Whether the identity can read the Sync Document. Default value is `false`. (optional)
    Write := true // bool | Whether the identity can update the Sync Document. Default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateDocumentPermission(context.Background(), ServiceSid, DocumentSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDocumentPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocumentPermission`: SyncV1ServiceDocumentDocumentPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateDocumentPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Document Permission resource to update.
**DocumentSid** | **string** | The SID of the Sync Document with the Document Permission resource to update. Can be the Document resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Document Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentPermissionParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> SyncV1Service UpdateService(ctx, Sid).AclEnabled(AclEnabled).FriendlyName(FriendlyName).ReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled).ReachabilityDebouncingWindow(ReachabilityDebouncingWindow).ReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled).WebhookUrl(WebhookUrl).WebhooksFromRestEnabled(WebhooksFromRestEnabled).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Service resource to update.
    AclEnabled := true // bool | Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that you assign to describe the resource. (optional)
    ReachabilityDebouncingEnabled := true // bool | Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event. (optional)
    ReachabilityDebouncingWindow := int32(56) // int32 | The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called. (optional)
    ReachabilityWebhooksEnabled := true // bool | Whether the service instance should call `webhook_url` when client endpoints connect to Sync. The default is `false`. (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL we should call when Sync objects are manipulated. (optional)
    WebhooksFromRestEnabled := true // bool | Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).AclEnabled(AclEnabled).FriendlyName(FriendlyName).ReachabilityDebouncingEnabled(ReachabilityDebouncingEnabled).ReachabilityDebouncingWindow(ReachabilityDebouncingWindow).ReachabilityWebhooksEnabled(ReachabilityWebhooksEnabled).WebhookUrl(WebhookUrl).WebhooksFromRestEnabled(WebhooksFromRestEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: SyncV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncList

> SyncV1ServiceSyncList UpdateSyncList(ctx, ServiceSid, Sid).CollectionTtl(CollectionTtl).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update.
    Sid := "Sid_example" // string | The SID of the Sync List resource to update. Can be the Sync List resource's `sid` or its `unique_name`.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `collection_ttl`. If both are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncList(context.Background(), ServiceSid, Sid).CollectionTtl(CollectionTtl).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncList`: SyncV1ServiceSyncList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List resource to update.
**Sid** | **string** | The SID of the Sync List resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
 **Ttl** | **int32** | An alias for &#x60;collection_ttl&#x60;. If both are provided, this value is ignored.

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

> SyncV1ServiceSyncListSyncListItem UpdateSyncListItem(ctx, ServiceSid, ListSid, Index).IfMatch(IfMatch).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource's `sid` or its `unique_name`.
    Index := int32(56) // int32 | The index of the Sync List Item resource to update.
    IfMatch := "IfMatch_example" // string | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). (optional)
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted. This parameter can only be used when the List Item's `data` or `ttl` is updated in the same request. (optional)
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length. (optional)
    ItemTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `item_ttl`. If both parameters are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncListItem(context.Background(), ServiceSid, ListSid, Index).IfMatch(IfMatch).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncListItem`: SyncV1ServiceSyncListSyncListItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Item resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Item resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Index** | **int32** | The index of the Sync List Item resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListItemParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> SyncV1ServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, ServiceSid, ListSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update.
    ListSid := "ListSid_example" // string | The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync List Permission resource to update.
    Manage := true // bool | Whether the identity can delete the Sync List. Default value is `false`. (optional)
    Read := true // bool | Whether the identity can read the Sync List and its Items. Default value is `false`. (optional)
    Write := true // bool | Whether the identity can create, update, and delete Items in the Sync List. Default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncListPermission(context.Background(), ServiceSid, ListSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncListPermission`: SyncV1ServiceSyncListSyncListPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncListPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync List Permission resource to update.
**ListSid** | **string** | The SID of the Sync List with the Sync List Permission resource to update. Can be the Sync List resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync List Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListPermissionParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMap

> SyncV1ServiceSyncMap UpdateSyncMap(ctx, ServiceSid, Sid).CollectionTtl(CollectionTtl).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update.
    Sid := "Sid_example" // string | The SID of the Sync Map resource to update. Can be the Sync Map's `sid` or its `unique_name`.
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `collection_ttl`. If both parameters are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncMap(context.Background(), ServiceSid, Sid).CollectionTtl(CollectionTtl).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncMap`: SyncV1ServiceSyncMap
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map resource to update.
**Sid** | **string** | The SID of the Sync Map resource to update. Can be the Sync Map&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CollectionTtl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
 **Ttl** | **int32** | An alias for &#x60;collection_ttl&#x60;. If both parameters are provided, this value is ignored.

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

> SyncV1ServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, ServiceSid, MapSid, Key).IfMatch(IfMatch).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map resource's `sid` or its `unique_name`.
    Key := "Key_example" // string | The `key` value of the Sync Map Item resource to update. 
    IfMatch := "IfMatch_example" // string | If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). (optional)
    CollectionTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted. This parameter can only be used when the Map Item's `data` or `ttl` is updated in the same request. (optional)
    Data := TODO // map[string]interface{} | A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length. (optional)
    ItemTtl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted. (optional)
    Ttl := int32(56) // int32 | An alias for `item_ttl`. If both parameters are provided, this value is ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncMapItem(context.Background(), ServiceSid, MapSid, Key).IfMatch(IfMatch).CollectionTtl(CollectionTtl).Data(Data).ItemTtl(ItemTtl).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncMapItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncMapItem`: SyncV1ServiceSyncMapSyncMapItem
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncMapItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Item resource to update.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Item resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Key** | **string** | The &#x60;key&#x60; value of the Sync Map Item resource to update. 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapItemParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> SyncV1ServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, ServiceSid, MapSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service's `sid` value or `default`.
    MapSid := "MapSid_example" // string | The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map resource's `sid` or its `unique_name`.
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the User's Sync Map Permission resource to update.
    Manage := true // bool | Whether the identity can delete the Sync Map. Default value is `false`. (optional)
    Read := true // bool | Whether the identity can read the Sync Map and its Items. Default value is `false`. (optional)
    Write := true // bool | Whether the identity can create, update, and delete Items in the Sync Map. Default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncMapPermission(context.Background(), ServiceSid, MapSid, Identity).Manage(Manage).Read(Read).Write(Write).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncMapPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncMapPermission`: SyncV1ServiceSyncMapSyncMapPermission
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncMapPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Map Permission resource to update. Can be the Service&#39;s &#x60;sid&#x60; value or &#x60;default&#x60;.
**MapSid** | **string** | The SID of the Sync Map with the Sync Map Permission resource to update. Can be the Sync Map resource&#39;s &#x60;sid&#x60; or its &#x60;unique_name&#x60;.
**Identity** | **string** | The application-defined string that uniquely identifies the User&#39;s Sync Map Permission resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapPermissionParams struct via the builder pattern


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

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncStream

> SyncV1ServiceSyncStream UpdateSyncStream(ctx, ServiceSid, Sid).Ttl(Ttl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to update.
    Sid := "Sid_example" // string | The SID of the Stream resource to update.
    Ttl := int32(56) // int32 | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSyncStream(context.Background(), ServiceSid, Sid).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSyncStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyncStream`: SyncV1ServiceSyncStream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSyncStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) with the Sync Stream resource to update.
**Sid** | **string** | The SID of the Stream resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncStreamParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Ttl** | **int32** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).

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

