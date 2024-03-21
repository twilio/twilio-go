# PluginServiceReleasesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePluginRelease**](PluginServiceReleasesApi.md#CreatePluginRelease) | **Post** /v1/PluginService/Releases | 
[**FetchPluginRelease**](PluginServiceReleasesApi.md#FetchPluginRelease) | **Get** /v1/PluginService/Releases/{Sid} | 
[**ListPluginRelease**](PluginServiceReleasesApi.md#ListPluginRelease) | **Get** /v1/PluginService/Releases | 



## CreatePluginRelease

> FlexV1PluginRelease CreatePluginRelease(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePluginReleaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**ConfigurationId** | **string** | The SID or the Version of the Flex Plugin Configuration to release.

### Return type

[**FlexV1PluginRelease**](FlexV1PluginRelease.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPluginRelease

> FlexV1PluginRelease FetchPluginRelease(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Plugin Release resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPluginReleaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header

### Return type

[**FlexV1PluginRelease**](FlexV1PluginRelease.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPluginRelease

> []FlexV1PluginRelease ListPluginRelease(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPluginReleaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**FlexMetadata** | **string** | The Flex-Metadata HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1PluginRelease**](FlexV1PluginRelease.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

