# WorkspacesWorkersChannelsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkerChannel**](WorkspacesWorkersChannelsApi.md#FetchWorkerChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 
[**ListWorkerChannel**](WorkspacesWorkersChannelsApi.md#ListWorkerChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels | 
[**UpdateWorkerChannel**](WorkspacesWorkersChannelsApi.md#UpdateWorkerChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 



## FetchWorkerChannel

> TaskrouterV1WorkerChannel FetchWorkerChannel(ctx, WorkspaceSidWorkerSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to fetch.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to fetch.
**Sid** | **string** | The SID of the WorkerChannel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1WorkerChannel**](TaskrouterV1WorkerChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkerChannel

> []TaskrouterV1WorkerChannel ListWorkerChannel(ctx, WorkspaceSidWorkerSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannels to read.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannels to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1WorkerChannel**](TaskrouterV1WorkerChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerChannel

> TaskrouterV1WorkerChannel UpdateWorkerChannel(ctx, WorkspaceSidWorkerSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to update.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to update.
**Sid** | **string** | The SID of the WorkerChannel to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**Capacity** | **int** | The total number of Tasks that the Worker should handle for the TaskChannel type. TaskRouter creates reservations for Tasks of this TaskChannel type up to the specified capacity. If the capacity is 0, no new reservations will be created.
**Available** | **bool** | Whether the WorkerChannel is available. Set to `false` to prevent the Worker from receiving any new Tasks of this TaskChannel type.

### Return type

[**TaskrouterV1WorkerChannel**](TaskrouterV1WorkerChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

