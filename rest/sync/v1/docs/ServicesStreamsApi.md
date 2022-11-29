# ServicesStreamsApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncStream**](ServicesStreamsApi.md#CreateSyncStream) | **Post** /v1/Services/{ServiceSid}/Streams | 
[**DeleteSyncStream**](ServicesStreamsApi.md#DeleteSyncStream) | **Delete** /v1/Services/{ServiceSid}/Streams/{Sid} | 
[**FetchSyncStream**](ServicesStreamsApi.md#FetchSyncStream) | **Get** /v1/Services/{ServiceSid}/Streams/{Sid} | 
[**ListSyncStream**](ServicesStreamsApi.md#ListSyncStream) | **Get** /v1/Services/{ServiceSid}/Streams | 
[**UpdateSyncStream**](ServicesStreamsApi.md#UpdateSyncStream) | **Post** /v1/Services/{ServiceSid}/Streams/{Sid} | 



## CreateSyncStream

> SyncV1SyncStream CreateSyncStream(ctx, ServiceSidoptional)



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
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
**Ttl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).

### Return type

[**SyncV1SyncStream**](SyncV1SyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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


## FetchSyncStream

> SyncV1SyncStream FetchSyncStream(ctx, ServiceSidSid)



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

[**SyncV1SyncStream**](SyncV1SyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncStream

> []SyncV1SyncStream ListSyncStream(ctx, ServiceSidoptional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SyncV1SyncStream**](SyncV1SyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncStream

> SyncV1SyncStream UpdateSyncStream(ctx, ServiceSidSidoptional)



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
**Ttl** | **int** | How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).

### Return type

[**SyncV1SyncStream**](SyncV1SyncStream.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

