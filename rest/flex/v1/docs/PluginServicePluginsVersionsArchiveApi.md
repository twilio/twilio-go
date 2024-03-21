# PluginServicePluginsVersionsArchiveApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePluginVersionArchive**](PluginServicePluginsVersionsArchiveApi.md#UpdatePluginVersionArchive) | **Post** /v1/PluginService/Plugins/{PluginSid}/Versions/{Sid}/Archive | 



## UpdatePluginVersionArchive

> FlexV1PluginVersionArchive UpdatePluginVersionArchive(ctx, PluginSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PluginSid** | **string** | The SID of the Flex Plugin the resource to belongs to.
**Sid** | **string** | The SID of the Flex Plugin Version resource to archive.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePluginVersionArchiveParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1PluginVersionArchive**](FlexV1PluginVersionArchive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

