# DevicesApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDevice**](DevicesApi.md#FetchDevice) | **Get** /v1/Devices/{Sid} | 
[**ListDevice**](DevicesApi.md#ListDevice) | **Get** /v1/Devices | 
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Post** /v1/Devices/{Sid} | 



## FetchDevice

> MicrovisorV1Device FetchDevice(ctx, Sid)



Fetch a specific Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34-character string that uniquely identifies this Device.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeviceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1Device**](MicrovisorV1Device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevice

> []MicrovisorV1Device ListDevice(ctx, optional)



Retrieve a list of all Devices registered with the Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDeviceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MicrovisorV1Device**](MicrovisorV1Device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> MicrovisorV1Device UpdateDevice(ctx, Sidoptional)



Update a specific Device.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34-character string that uniquely identifies this Device.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDeviceParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | A unique and addressable name to be assigned to this Device by the developer. It may be used in place of the Device SID.
**TargetApp** | **string** | The SID or unique name of the App to be targeted to the Device.
**LoggingEnabled** | **bool** | A Boolean flag specifying whether to enable application logging. Logs will be enabled or extended for 24 hours.

### Return type

[**MicrovisorV1Device**](MicrovisorV1Device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

