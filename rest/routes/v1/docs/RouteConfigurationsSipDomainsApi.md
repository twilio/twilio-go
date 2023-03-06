# RouteConfigurationsSipDomainsApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipDomain**](RouteConfigurationsSipDomainsApi.md#CreateSipDomain) | **Post** /v1/RouteConfigurations/{RouteConfigurationSid}/SipDomains | 
[**DeleteSipDomain**](RouteConfigurationsSipDomainsApi.md#DeleteSipDomain) | **Delete** /v1/RouteConfigurations/{RouteConfigurationSid}/SipDomains | 
[**FetchSipDomain**](RouteConfigurationsSipDomainsApi.md#FetchSipDomain) | **Get** /v1/RouteConfigurations/{RouteConfigurationSid}/SipDomains | 



## CreateSipDomain

> RoutesV1SipDomain CreateSipDomain(ctx, RouteConfigurationSidoptional)



Link a Resource to an existing Route Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**SipDomain** | **string** | TBD

### Return type

[**RoutesV1SipDomain**](RoutesV1SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipDomain

> DeleteSipDomain(ctx, RouteConfigurationSid)



Delete a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipDomainParams struct


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


## FetchSipDomain

> RoutesV1SipDomain FetchSipDomain(ctx, RouteConfigurationSid)



Fetch a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV1SipDomain**](RoutesV1SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

