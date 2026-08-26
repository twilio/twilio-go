# ConfigurationsTranscriptionApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscriptionConfiguration**](ConfigurationsTranscriptionApi.md#CreateTranscriptionConfiguration) | **Post** /v2/Configurations/Transcription | Create a Transcription Configuration
[**DeleteTranscriptionConfiguration**](ConfigurationsTranscriptionApi.md#DeleteTranscriptionConfiguration) | **Delete** /v2/Configurations/Transcription/{IdOrUniqueName} | Delete a Transcription Configuration by Id
[**FetchTranscriptionConfiguration**](ConfigurationsTranscriptionApi.md#FetchTranscriptionConfiguration) | **Get** /v2/Configurations/Transcription/{IdOrUniqueName} | Get a Transcription Configuration by Id or UniqueName
[**UpdateTranscriptionConfiguration**](ConfigurationsTranscriptionApi.md#UpdateTranscriptionConfiguration) | **Put** /v2/Configurations/Transcription/{IdOrUniqueName} | Update a Transcription Configuration by Id



## CreateTranscriptionConfiguration

> VoiceV2Response CreateTranscriptionConfiguration(ctx, optional)

Create a Transcription Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTranscriptionConfigurationParams struct


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


## DeleteTranscriptionConfiguration

> DeleteTranscriptionConfiguration(ctx, IdOrUniqueName)

Delete a Transcription Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteTranscriptionConfigurationParams struct


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


## FetchTranscriptionConfiguration

> VoiceV2Response FetchTranscriptionConfiguration(ctx, IdOrUniqueName)

Get a Transcription Configuration by Id or UniqueName

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | Config id or unique name

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptionConfigurationParams struct


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


## UpdateTranscriptionConfiguration

> VoiceV2Response UpdateTranscriptionConfiguration(ctx, IdOrUniqueNameoptional)

Update a Transcription Configuration by Id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IdOrUniqueName** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateTranscriptionConfigurationParams struct


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

