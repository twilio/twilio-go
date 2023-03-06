# RouteConfigurationsApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteConfiguration**](RouteConfigurationsApi.md#CreateRouteConfiguration) | **Post** /v1/RouteConfigurations | 
[**DeleteRouteConfiguration**](RouteConfigurationsApi.md#DeleteRouteConfiguration) | **Delete** /v1/RouteConfigurations/{Sid} | 
[**FetchRouteConfiguration**](RouteConfigurationsApi.md#FetchRouteConfiguration) | **Get** /v1/RouteConfigurations/{Sid} | 
[**ListRouteConfiguration**](RouteConfigurationsApi.md#ListRouteConfiguration) | **Get** /v1/RouteConfigurations | 
[**UpdateRouteConfiguration**](RouteConfigurationsApi.md#UpdateRouteConfiguration) | **Post** /v1/RouteConfigurations/{Sid} | 



## CreateRouteConfiguration

> RoutesV1RouteConfiguration CreateRouteConfiguration(ctx, optional)



Create a new Route Configuration.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRouteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**VoiceRegion** | **string** | TBD
**MessagingRegion** | **string** | TBD
**DataRegion** | **string** | TBD

### Return type

[**RoutesV1RouteConfiguration**](RoutesV1RouteConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteConfiguration

> DeleteRouteConfiguration(ctx, Sid)



Delete a specific Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRouteConfigurationParams struct


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


## FetchRouteConfiguration

> RoutesV1RouteConfiguration FetchRouteConfiguration(ctx, Sid)



Fetch a specific Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchRouteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV1RouteConfiguration**](RoutesV1RouteConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteConfiguration

> []RoutesV1RouteConfiguration ListRouteConfiguration(ctx, optional)



Read a list of Route Configuration instances.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRouteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumbers** | **string** | The unique string that identifies the phone number
**DomainNames** | **string** | The unique string that identifies the domain name
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]RoutesV1RouteConfiguration**](RoutesV1RouteConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteConfiguration

> RoutesV1RouteConfiguration UpdateRouteConfiguration(ctx, Sidoptional)



Update a specific Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRouteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**VoiceRegion** | **string** | TBD
**MessagingRegion** | **string** | TBD
**DataRegion** | **string** | TBD

### Return type

[**RoutesV1RouteConfiguration**](RoutesV1RouteConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

