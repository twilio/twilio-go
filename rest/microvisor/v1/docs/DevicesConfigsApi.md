# DevicesConfigsApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceConfig**](DevicesConfigsApi.md#CreateDeviceConfig) | **Post** /v1/Devices/{DeviceSid}/Configs | 
[**DeleteDeviceConfig**](DevicesConfigsApi.md#DeleteDeviceConfig) | **Delete** /v1/Devices/{DeviceSid}/Configs/{Key} | 
[**FetchDeviceConfig**](DevicesConfigsApi.md#FetchDeviceConfig) | **Get** /v1/Devices/{DeviceSid}/Configs/{Key} | 
[**ListDeviceConfig**](DevicesConfigsApi.md#ListDeviceConfig) | **Get** /v1/Devices/{DeviceSid}/Configs | 
[**UpdateDeviceConfig**](DevicesConfigsApi.md#UpdateDeviceConfig) | **Post** /v1/Devices/{DeviceSid}/Configs/{Key} | 



## CreateDeviceConfig

> MicrovisorV1DeviceConfig CreateDeviceConfig(ctx, DeviceSidoptional)



Create a config for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.

### Other Parameters

Other parameters are passed through a pointer to a CreateDeviceConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | The config key; up to 100 characters.
**Value** | **string** | The config value; up to 4096 characters.

### Return type

[**MicrovisorV1DeviceConfig**](MicrovisorV1DeviceConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceConfig

> DeleteDeviceConfig(ctx, DeviceSidKey)



Delete a config for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDeviceConfigParams struct


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


## FetchDeviceConfig

> MicrovisorV1DeviceConfig FetchDeviceConfig(ctx, DeviceSidKey)



Retrieve a Config for a Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeviceConfigParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1DeviceConfig**](MicrovisorV1DeviceConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceConfig

> []MicrovisorV1DeviceConfig ListDeviceConfig(ctx, DeviceSidoptional)



Retrieve a list of all Configs for a Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.

### Other Parameters

Other parameters are passed through a pointer to a ListDeviceConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1DeviceConfig**](MicrovisorV1DeviceConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceConfig

> MicrovisorV1DeviceConfig UpdateDeviceConfig(ctx, DeviceSidKeyoptional)



Update a config for a Microvisor Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DeviceSid** | **string** | A 34-character string that uniquely identifies the Device.
**Key** | **string** | The config key; up to 100 characters.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDeviceConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**Value** | **string** | The config value; up to 4096 characters.

### Return type

[**MicrovisorV1DeviceConfig**](MicrovisorV1DeviceConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

