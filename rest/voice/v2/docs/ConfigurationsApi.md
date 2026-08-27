# ConfigurationsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ConfigurationsApi.md#CreateConfiguration) | **Post** /v2/Configurations/{Type} | Create a Configuration
[**DeleteConfiguration**](ConfigurationsApi.md#DeleteConfiguration) | **Delete** /v2/Configurations/{Type}/{IdOrUniqueName} | Delete a Configuration by Id
[**FetchConfiguration**](ConfigurationsApi.md#FetchConfiguration) | **Get** /v2/Configurations/{Type}/{IdOrUniqueName} | Get a Configuration by Id or UniqueName
[**UpdateConfiguration**](ConfigurationsApi.md#UpdateConfiguration) | **Put** /v2/Configurations/{Type}/{IdOrUniqueName} | Update a Configuration by Id



## CreateConfiguration

> VoiceV2Response CreateConfiguration(ctx, Typeoptional)

Create a Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**VoiceV2Request** | [**VoiceV2Request**](VoiceV2Request.md) | 

### Return type

[**VoiceV2Response**](VoiceV2Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, TypeIdOrUniqueName)

Delete a Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> VoiceV2Response FetchConfiguration(ctx, TypeIdOrUniqueName)

Get a Configuration by Id or UniqueName

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 
**IdOrUniqueName** | **string** | Config id or unique name

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV2Response**](VoiceV2Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> VoiceV2Response UpdateConfiguration(ctx, TypeIdOrUniqueNameoptional)

Update a Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**VoiceV2Request** | [**VoiceV2Request**](VoiceV2Request.md) | 

### Return type

[**VoiceV2Response**](VoiceV2Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

