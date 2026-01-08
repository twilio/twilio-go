# LinkShorteningDomainsCertificateApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDomainCertV4**](LinkShorteningDomainsCertificateApi.md#FetchDomainCertV4) | **Get** /v2/LinkShortening/Domains/{DomainSid}/Certificate | 



## FetchDomainCertV4

> MessagingV2DomainCertV4 FetchDomainCertV4(ctx, DomainSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this certificate should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a FetchDomainCertV4Params struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV2DomainCertV4**](MessagingV2DomainCertV4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

