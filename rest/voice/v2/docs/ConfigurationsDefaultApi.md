# ConfigurationsDefaultApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateDefaultConfiguration**](ConfigurationsDefaultApi.md#CreateOrUpdateDefaultConfiguration) | **Post** /v2/Configurations/{Type}/Default | Create/Update a Default Configuration
[**FetchDefaultConfiguration**](ConfigurationsDefaultApi.md#FetchDefaultConfiguration) | **Get** /v2/Configurations/{Type}/Default | Get the Default Configuration



## CreateOrUpdateDefaultConfiguration

> CreateOrUpdateDefaultConfiguration(ctx, Typeoptional)

Create/Update a Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateOrUpdateDefaultConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**DefaultConfigurationRequest** | [**DefaultConfigurationRequest**](DefaultConfigurationRequest.md) | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDefaultConfiguration

> VoiceV2Response FetchDefaultConfiguration(ctx, Type)

Get the Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDefaultConfigurationParams struct


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

