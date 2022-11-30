# LinkShorteningDomainsCertificateApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDomainCertV4**](LinkShorteningDomainsCertificateApi.md#DeleteDomainCertV4) | **Delete** /v1/LinkShortening/Domains/{DomainSid}/Certificate | 
[**FetchDomainCertV4**](LinkShorteningDomainsCertificateApi.md#FetchDomainCertV4) | **Get** /v1/LinkShortening/Domains/{DomainSid}/Certificate | 
[**UpdateDomainCertV4**](LinkShorteningDomainsCertificateApi.md#UpdateDomainCertV4) | **Post** /v1/LinkShortening/Domains/{DomainSid}/Certificate | 



## DeleteDomainCertV4

> DeleteDomainCertV4(ctx, DomainSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this certificate should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a DeleteDomainCertV4Params struct


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


## FetchDomainCertV4

> MessagingV1DomainCertV4 FetchDomainCertV4(ctx, DomainSid)





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

[**MessagingV1DomainCertV4**](MessagingV1DomainCertV4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainCertV4

> MessagingV1DomainCertV4 UpdateDomainCertV4(ctx, DomainSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this certificate should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDomainCertV4Params struct


Name | Type | Description
------------- | ------------- | -------------
**TlsCert** | **string** | Contains the full TLS certificate and private for this domain in PEM format: https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail. Twilio uses this information to process HTTPS traffic sent to your domain.

### Return type

[**MessagingV1DomainCertV4**](MessagingV1DomainCertV4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

