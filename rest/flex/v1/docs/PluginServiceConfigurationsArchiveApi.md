# PluginServiceConfigurationsArchiveApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePluginConfigurationArchive**](PluginServiceConfigurationsArchiveApi.md#UpdatePluginConfigurationArchive) | **Post** /v1/PluginService/Configurations/{Sid}/Archive | 



## UpdatePluginConfigurationArchive

> FlexV1PluginConfigurationArchive UpdatePluginConfigurationArchive(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Plugin Configuration resource to archive.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePluginConfigurationArchiveParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1PluginConfigurationArchive**](FlexV1PluginConfigurationArchive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

