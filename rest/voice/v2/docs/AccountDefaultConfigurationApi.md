# AccountDefaultConfigurationApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDefaultConfiguration**](AccountDefaultConfigurationApi.md#CreateAccountDefaultConfiguration) | **Post** /v2/AccountDefaultConfiguration/{Type} | Create an Account Default Configuration
[**DeleteAccountDefaultConfiguration**](AccountDefaultConfigurationApi.md#DeleteAccountDefaultConfiguration) | **Delete** /v2/AccountDefaultConfiguration/{Type} | Delete the Account Default Configuration
[**FetchAccountDefaultConfiguration**](AccountDefaultConfigurationApi.md#FetchAccountDefaultConfiguration) | **Get** /v2/AccountDefaultConfiguration/{Type} | Get the Account Default Configuration
[**UpdateAccountDefaultConfiguration**](AccountDefaultConfigurationApi.md#UpdateAccountDefaultConfiguration) | **Put** /v2/AccountDefaultConfiguration/{Type} | Update the Account Default Configuration



## CreateAccountDefaultConfiguration

> VoiceV2Response CreateAccountDefaultConfiguration(ctx, Typeoptional)

Create an Account Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountDefaultConfigurationParams struct


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


## DeleteAccountDefaultConfiguration

> DeleteAccountDefaultConfiguration(ctx, Type)

Delete the Account Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteAccountDefaultConfigurationParams struct


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


## FetchAccountDefaultConfiguration

> VoiceV2Response FetchAccountDefaultConfiguration(ctx, Type)

Get the Account Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountDefaultConfigurationParams struct


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


## UpdateAccountDefaultConfiguration

> VoiceV2Response UpdateAccountDefaultConfiguration(ctx, Typeoptional)

Update the Account Default Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountDefaultConfigurationParams struct


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

