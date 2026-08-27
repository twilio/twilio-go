# ConfigurationsRecordingApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecordingConfiguration**](ConfigurationsRecordingApi.md#CreateRecordingConfiguration) | **Post** /v2/Configurations/Recording | Create a Recording Configuration
[**DeleteRecordingConfiguration**](ConfigurationsRecordingApi.md#DeleteRecordingConfiguration) | **Delete** /v2/Configurations/Recording/{IdOrUniqueName} | Delete a Recording Configuration by Id
[**FetchRecordingConfiguration**](ConfigurationsRecordingApi.md#FetchRecordingConfiguration) | **Get** /v2/Configurations/Recording/{IdOrUniqueName} | Get a Recording Configuration by Id or UniqueName
[**UpdateRecordingConfiguration**](ConfigurationsRecordingApi.md#UpdateRecordingConfiguration) | **Put** /v2/Configurations/Recording/{IdOrUniqueName} | Update a Recording Configuration by Id



## CreateRecordingConfiguration

> VoiceV2Response CreateRecordingConfiguration(ctx, optional)

Create a Recording Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRecordingConfigurationParams struct


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


## DeleteRecordingConfiguration

> DeleteRecordingConfiguration(ctx, IdOrUniqueName)

Delete a Recording Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingConfigurationParams struct


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


## FetchRecordingConfiguration

> VoiceV2Response FetchRecordingConfiguration(ctx, IdOrUniqueName)

Get a Recording Configuration by Id or UniqueName

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | Config id or unique name

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingConfigurationParams struct


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


## UpdateRecordingConfiguration

> VoiceV2Response UpdateRecordingConfiguration(ctx, IdOrUniqueNameoptional)

Update a Recording Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateRecordingConfigurationParams struct


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

