# WebChannelsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebChannel**](WebChannelsApi.md#CreateWebChannel) | **Post** /v1/WebChannels | 
[**DeleteWebChannel**](WebChannelsApi.md#DeleteWebChannel) | **Delete** /v1/WebChannels/{Sid} | 
[**FetchWebChannel**](WebChannelsApi.md#FetchWebChannel) | **Get** /v1/WebChannels/{Sid} | 
[**ListWebChannel**](WebChannelsApi.md#ListWebChannel) | **Get** /v1/WebChannels | 
[**UpdateWebChannel**](WebChannelsApi.md#UpdateWebChannel) | **Post** /v1/WebChannels/{Sid} | 



## CreateWebChannel

> FlexV1WebChannel CreateWebChannel(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexFlowSid** | **string** | The SID of the Flex Flow.
**Identity** | **string** | The chat identity.
**CustomerFriendlyName** | **string** | The chat participant's friendly name.
**ChatFriendlyName** | **string** | The chat channel's friendly name.
**ChatUniqueName** | **string** | The chat channel's unique name.
**PreEngagementData** | **string** | The pre-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebChannel

> DeleteWebChannel(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebChannelParams struct


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


## FetchWebChannel

> FlexV1WebChannel FetchWebChannel(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebChannel

> []FlexV1WebChannel ListWebChannel(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebChannel

> FlexV1WebChannel UpdateWebChannel(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChatStatus** | **string** | 
**PostEngagementData** | **string** | The post-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

