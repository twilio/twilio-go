# PluginServiceConfigurationsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePluginConfiguration**](PluginServiceConfigurationsApi.md#CreatePluginConfiguration) | **Post** /v1/PluginService/Configurations | 
[**FetchPluginConfiguration**](PluginServiceConfigurationsApi.md#FetchPluginConfiguration) | **Get** /v1/PluginService/Configurations/{Sid} | 
[**ListPluginConfiguration**](PluginServiceConfigurationsApi.md#ListPluginConfiguration) | **Get** /v1/PluginService/Configurations | 



## CreatePluginConfiguration

> FlexV1PluginConfiguration CreatePluginConfiguration(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePluginConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**Name** | **string** | The Flex Plugin Configuration's name.
**Plugins** | **[]interface{}** | A list of objects that describe the plugin versions included in the configuration. Each object contains the sid of the plugin version.
**Description** | **string** | The Flex Plugin Configuration's description.

### Return type

[**FlexV1PluginConfiguration**](FlexV1PluginConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginConfiguration

> FlexV1PluginConfiguration FetchPluginConfiguration(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Plugin Configuration resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPluginConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1PluginConfiguration**](FlexV1PluginConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPluginConfiguration

> []FlexV1PluginConfiguration ListPluginConfiguration(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPluginConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1PluginConfiguration**](FlexV1PluginConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

