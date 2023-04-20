# LinkShorteningDomainsConfigApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDomainConfig**](LinkShorteningDomainsConfigApi.md#FetchDomainConfig) | **Get** /v1/LinkShortening/Domains/{DomainSid}/Config | 
[**UpdateDomainConfig**](LinkShorteningDomainsConfigApi.md#UpdateDomainConfig) | **Post** /v1/LinkShortening/Domains/{DomainSid}/Config | 



## FetchDomainConfig

> MessagingV1DomainConfig FetchDomainConfig(ctx, DomainSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this config should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a FetchDomainConfigParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1DomainConfig**](MessagingV1DomainConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainConfig

> MessagingV1DomainConfig UpdateDomainConfig(ctx, DomainSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this config should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDomainConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**FallbackUrl** | **string** | Any requests we receive to this domain that do not match an existing shortened message will be redirected to the fallback url. These will likely be either expired messages, random misdirected traffic, or intentional scraping.
**CallbackUrl** | **string** | URL to receive click events to your webhook whenever the recipients click on the shortened links

### Return type

[**MessagingV1DomainConfig**](MessagingV1DomainConfig.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

