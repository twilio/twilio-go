# ServicesBindingsApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBinding**](ServicesBindingsApi.md#DeleteBinding) | **Delete** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**FetchBinding**](ServicesBindingsApi.md#FetchBinding) | **Get** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**ListBinding**](ServicesBindingsApi.md#ListBinding) | **Get** /v2/Services/{ServiceSid}/Bindings | 



## DeleteBinding

> DeleteBinding(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteBindingParams struct


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


## FetchBinding

> IpMessagingV2Binding FetchBinding(ctx, ServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV2Binding**](IpMessagingV2Binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> []IpMessagingV2Binding ListBinding(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | [**[]BindingEnumBindingType**](BindingEnumBindingType.md) | 
**Identity** | **[]string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IpMessagingV2Binding**](IpMessagingV2Binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

