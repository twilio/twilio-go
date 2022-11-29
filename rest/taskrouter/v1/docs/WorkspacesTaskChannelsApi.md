# WorkspacesTaskChannelsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskChannel**](WorkspacesTaskChannelsApi.md#CreateTaskChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
[**DeleteTaskChannel**](WorkspacesTaskChannelsApi.md#DeleteTaskChannel) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
[**FetchTaskChannel**](WorkspacesTaskChannelsApi.md#FetchTaskChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
[**ListTaskChannel**](WorkspacesTaskChannelsApi.md#ListTaskChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
[**UpdateTaskChannel**](WorkspacesTaskChannelsApi.md#UpdateTaskChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 



## CreateTaskChannel

> TaskrouterV1TaskChannel CreateTaskChannel(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Task Channel belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
**UniqueName** | **string** | An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`.
**ChannelOptimizedRouting** | **bool** | Whether the Task Channel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized.

### Return type

[**TaskrouterV1TaskChannel**](TaskrouterV1TaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskChannel

> DeleteTaskChannel(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to delete.
**Sid** | **string** | The SID of the Task Channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskChannelParams struct


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


## FetchTaskChannel

> TaskrouterV1TaskChannel FetchTaskChannel(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to fetch.
**Sid** | **string** | The SID of the Task Channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1TaskChannel**](TaskrouterV1TaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskChannel

> []TaskrouterV1TaskChannel ListTaskChannel(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1TaskChannel**](TaskrouterV1TaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskChannel

> TaskrouterV1TaskChannel UpdateTaskChannel(ctx, WorkspaceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to update.
**Sid** | **string** | The SID of the Task Channel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
**ChannelOptimizedRouting** | **bool** | Whether the TaskChannel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized.

### Return type

[**TaskrouterV1TaskChannel**](TaskrouterV1TaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

