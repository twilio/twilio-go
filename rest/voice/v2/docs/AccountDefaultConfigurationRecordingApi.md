# AccountDefaultConfigurationRecordingApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDefaultRecordingConfiguration**](AccountDefaultConfigurationRecordingApi.md#CreateAccountDefaultRecordingConfiguration) | **Post** /v2/AccountDefaultConfiguration/Recording | Create an Account Default Recording Configuration
[**DeleteAccountDefaultRecordingConfiguration**](AccountDefaultConfigurationRecordingApi.md#DeleteAccountDefaultRecordingConfiguration) | **Delete** /v2/AccountDefaultConfiguration/Recording | Delete the Account Default Recording Configuration
[**FetchAccountDefaultRecordingConfiguration**](AccountDefaultConfigurationRecordingApi.md#FetchAccountDefaultRecordingConfiguration) | **Get** /v2/AccountDefaultConfiguration/Recording | Get the Account Default Recording Configuration
[**UpdateAccountDefaultRecordingConfiguration**](AccountDefaultConfigurationRecordingApi.md#UpdateAccountDefaultRecordingConfiguration) | **Put** /v2/AccountDefaultConfiguration/Recording | Update the Account Default Recording Configuration



## CreateAccountDefaultRecordingConfiguration

> VoiceV2Response CreateAccountDefaultRecordingConfiguration(ctx, optional)

Create an Account Default Recording Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountDefaultRecordingConfigurationParams struct


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


## DeleteAccountDefaultRecordingConfiguration

> DeleteAccountDefaultRecordingConfiguration(ctx, )

Delete the Account Default Recording Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAccountDefaultRecordingConfigurationParams struct


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


## FetchAccountDefaultRecordingConfiguration

> VoiceV2Response FetchAccountDefaultRecordingConfiguration(ctx, )

Get the Account Default Recording Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountDefaultRecordingConfigurationParams struct


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


## UpdateAccountDefaultRecordingConfiguration

> VoiceV2Response UpdateAccountDefaultRecordingConfiguration(ctx, optional)

Update the Account Default Recording Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountDefaultRecordingConfigurationParams struct


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

