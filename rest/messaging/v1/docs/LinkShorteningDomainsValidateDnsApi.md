# LinkShorteningDomainsValidateDnsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDomainDnsValidation**](LinkShorteningDomainsValidateDnsApi.md#FetchDomainDnsValidation) | **Get** /v1/LinkShortening/Domains/{DomainSid}/ValidateDns | 



## FetchDomainDnsValidation

> MessagingV1DomainDnsValidation FetchDomainDnsValidation(ctx, DomainSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain.

### Other Parameters

Other parameters are passed through a pointer to a FetchDomainDnsValidationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1DomainDnsValidation**](MessagingV1DomainDnsValidation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

