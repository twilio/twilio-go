# SipDomainsApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSipDomain**](SipDomainsApi.md#FetchSipDomain) | **Get** /v2/SipDomains/{SipDomain} | 
[**UpdateSipDomain**](SipDomainsApi.md#UpdateSipDomain) | **Post** /v2/SipDomains/{SipDomain} | 



## FetchSipDomain

> RoutesV2SipDomain FetchSipDomain(ctx, SipDomain)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SipDomain** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV2SipDomain**](RoutesV2SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipDomain

> RoutesV2SipDomain UpdateSipDomain(ctx, SipDomainoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SipDomain** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**VoiceRegion** | **string** | 
**FriendlyName** | **string** | 

### Return type

[**RoutesV2SipDomain**](RoutesV2SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

