# PluginServicePluginsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlugin**](PluginServicePluginsApi.md#CreatePlugin) | **Post** /v1/PluginService/Plugins | 
[**FetchPlugin**](PluginServicePluginsApi.md#FetchPlugin) | **Get** /v1/PluginService/Plugins/{Sid} | 
[**ListPlugin**](PluginServicePluginsApi.md#ListPlugin) | **Get** /v1/PluginService/Plugins | 
[**UpdatePlugin**](PluginServicePluginsApi.md#UpdatePlugin) | **Post** /v1/PluginService/Plugins/{Sid} | 



## CreatePlugin

> FlexV1Plugin CreatePlugin(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**UniqueName** | **string** | The Flex Plugin's unique name.
**FriendlyName** | **string** | The Flex Plugin's friendly name.
**Description** | **string** | A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long

### Return type

[**FlexV1Plugin**](FlexV1Plugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPlugin

> FlexV1Plugin FetchPlugin(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Plugin resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1Plugin**](FlexV1Plugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlugin

> []FlexV1Plugin ListPlugin(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1Plugin**](FlexV1Plugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlugin

> FlexV1Plugin UpdatePlugin(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Plugin resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**FriendlyName** | **string** | The Flex Plugin's friendly name.
**Description** | **string** | A descriptive string that you update to describe the plugin resource. It can be up to 500 characters long

### Return type

[**FlexV1Plugin**](FlexV1Plugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

