# RouteConfigurationsPhoneNumbersApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumber**](RouteConfigurationsPhoneNumbersApi.md#CreatePhoneNumber) | **Post** /v1/RouteConfigurations/{RouteConfigurationSid}/PhoneNumbers | 
[**DeletePhoneNumber**](RouteConfigurationsPhoneNumbersApi.md#DeletePhoneNumber) | **Delete** /v1/RouteConfigurations/{RouteConfigurationSid}/PhoneNumbers | 
[**FetchPhoneNumber**](RouteConfigurationsPhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v1/RouteConfigurations/{RouteConfigurationSid}/PhoneNumbers | 



## CreatePhoneNumber

> RoutesV1PhoneNumber CreatePhoneNumber(ctx, RouteConfigurationSidoptional)



Link a Resource to an existing Route Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | TBD

### Return type

[**RoutesV1PhoneNumber**](RoutesV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneNumber

> DeletePhoneNumber(ctx, RouteConfigurationSid)



Delete a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct


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


## FetchPhoneNumber

> RoutesV1PhoneNumber FetchPhoneNumber(ctx, RouteConfigurationSid)



Fetch a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV1PhoneNumber**](RoutesV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

