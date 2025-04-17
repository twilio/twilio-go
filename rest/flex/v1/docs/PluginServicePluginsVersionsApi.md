# PluginServicePluginsVersionsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePluginVersion**](PluginServicePluginsVersionsApi.md#CreatePluginVersion) | **Post** /v1/PluginService/Plugins/{PluginSid}/Versions | 
[**FetchPluginVersion**](PluginServicePluginsVersionsApi.md#FetchPluginVersion) | **Get** /v1/PluginService/Plugins/{PluginSid}/Versions/{Sid} | 
[**ListPluginVersion**](PluginServicePluginsVersionsApi.md#ListPluginVersion) | **Get** /v1/PluginService/Plugins/{PluginSid}/Versions | 



## CreatePluginVersion

> FlexV1PluginVersion CreatePluginVersion(ctx, PluginSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PluginSid** | **string** | The SID of the Flex Plugin the resource to belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreatePluginVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**Version** | **string** | The Flex Plugin Version's version.
**PluginUrl** | **string** | The URL of the Flex Plugin Version bundle
**Changelog** | **string** | The changelog of the Flex Plugin Version.
**Private** | **bool** | Whether this Flex Plugin Version requires authorization.
**CliVersion** | **string** | The version of Flex Plugins CLI used to create this plugin
**ValidateStatus** | **string** | The validation status of the plugin, indicating whether it has been validated

### Return type

[**FlexV1PluginVersion**](FlexV1PluginVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginVersion

> FlexV1PluginVersion FetchPluginVersion(ctx, PluginSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PluginSid** | **string** | The SID of the Flex Plugin the resource to belongs to.
**Sid** | **string** | The SID of the Flex Plugin Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPluginVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1PluginVersion**](FlexV1PluginVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPluginVersion

> []FlexV1PluginVersion ListPluginVersion(ctx, PluginSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PluginSid** | **string** | The SID of the Flex Plugin the resource to belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListPluginVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1PluginVersion**](FlexV1PluginVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

