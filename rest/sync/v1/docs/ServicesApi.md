# ServicesApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v1/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



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
**ReachabilityDebouncingWindow** | **int** | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the &#x60;webhook_url&#x60; is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the call to &#x60;webhook_url&#x60;.
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> []SyncV1Service ListService(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1Service**](SyncV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**ReachabilityDebouncingWindow** | **int** | The reachability event delay in milliseconds if &#x60;reachability_debouncing_enabled&#x60; &#x3D; &#x60;true&#x60;.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before the webhook is called if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the webhook from being called.
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

