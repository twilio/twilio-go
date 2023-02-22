# DevicesSecretsApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSecret**](DevicesSecretsApi.md#CreateDeviceSecret) | **Post** /v1/Devices/{DeviceSid}/Secrets | 
[**DeleteDeviceSecret**](DevicesSecretsApi.md#DeleteDeviceSecret) | **Delete** /v1/Devices/{DeviceSid}/Secrets/{Key} | 
[**FetchDeviceSecret**](DevicesSecretsApi.md#FetchDeviceSecret) | **Get** /v1/Devices/{DeviceSid}/Secrets/{Key} | 
[**ListDeviceSecret**](DevicesSecretsApi.md#ListDeviceSecret) | **Get** /v1/Devices/{DeviceSid}/Secrets | 
[**UpdateDeviceSecret**](DevicesSecretsApi.md#UpdateDeviceSecret) | **Post** /v1/Devices/{DeviceSid}/Secrets/{Key} | 



## CreateDeviceSecret

> MicrovisorV1DeviceSecret CreateDeviceSecret(ctx, DeviceSidoptional)



Create a secret for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.

### Other Parameters

Other parameters are passed through a pointer to a CreateDeviceSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | The secret key; up to 100 characters.
**Value** | **string** | The secret value; up to 4096 characters.

### Return type

[**MicrovisorV1DeviceSecret**](MicrovisorV1DeviceSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceSecret

> DeleteDeviceSecret(ctx, DeviceSidKey)



Delete a secret for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDeviceSecretParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeviceSecret

> MicrovisorV1DeviceSecret FetchDeviceSecret(ctx, DeviceSidKey)



Retrieve a Secret for a Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeviceSecretParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1DeviceSecret**](MicrovisorV1DeviceSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceSecret

> []MicrovisorV1DeviceSecret ListDeviceSecret(ctx, DeviceSidoptional)



Retrieve a list of all Secrets for a Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.

### Other Parameters

Other parameters are passed through a pointer to a ListDeviceSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1DeviceSecret**](MicrovisorV1DeviceSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSecret

> MicrovisorV1DeviceSecret UpdateDeviceSecret(ctx, DeviceSidKeyoptional)



Update a secret for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The secret key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDeviceSecretParams struct


Name | Type | Description
------------- | ------------- | -------------
**Value** | **string** | The secret value; up to 4096 characters.

### Return type

[**MicrovisorV1DeviceSecret**](MicrovisorV1DeviceSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

