# LinkShorteningDomainsRequestManagedCertApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRequestManagedCert**](LinkShorteningDomainsRequestManagedCertApi.md#UpdateRequestManagedCert) | **Post** /v1/LinkShortening/Domains/{DomainSid}/RequestManagedCert | 



## UpdateRequestManagedCert

> MessagingV1RequestManagedCert UpdateRequestManagedCert(ctx, DomainSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | Unique string used to identify the domain that this certificate should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRequestManagedCertParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1RequestManagedCert**](MessagingV1RequestManagedCert.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

