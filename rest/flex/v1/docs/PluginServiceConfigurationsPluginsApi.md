# PluginServiceConfigurationsPluginsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConfiguredPlugin**](PluginServiceConfigurationsPluginsApi.md#FetchConfiguredPlugin) | **Get** /v1/PluginService/Configurations/{ConfigurationSid}/Plugins/{PluginSid} | 
[**ListConfiguredPlugin**](PluginServiceConfigurationsPluginsApi.md#ListConfiguredPlugin) | **Get** /v1/PluginService/Configurations/{ConfigurationSid}/Plugins | 



## FetchConfiguredPlugin

> FlexV1ConfiguredPlugin FetchConfiguredPlugin(ctx, ConfigurationSidPluginSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationSid** | **string** | The SID of the Flex Plugin Configuration the resource to belongs to.
**PluginSid** | **string** | The unique string that we created to identify the Flex Plugin resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfiguredPluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1ConfiguredPlugin**](FlexV1ConfiguredPlugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfiguredPlugin

> []FlexV1ConfiguredPlugin ListConfiguredPlugin(ctx, ConfigurationSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConfigurationSid** | **string** | The SID of the Flex Plugin Configuration the resource to belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListConfiguredPluginParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1ConfiguredPlugin**](FlexV1ConfiguredPlugin.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

