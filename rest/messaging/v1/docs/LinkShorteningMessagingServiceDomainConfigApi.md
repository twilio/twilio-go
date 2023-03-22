# LinkShorteningMessagingServiceDomainConfigApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDomainConfigMessagingService**](LinkShorteningMessagingServiceDomainConfigApi.md#FetchDomainConfigMessagingService) | **Get** /v1/LinkShortening/MessagingService/{MessagingServiceSid}/DomainConfig | 



## FetchDomainConfigMessagingService

> MessagingV1DomainConfigMessagingService FetchDomainConfigMessagingService(ctx, MessagingServiceSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | Unique string used to identify the Messaging service that this domain should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a FetchDomainConfigMessagingServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1DomainConfigMessagingService**](MessagingV1DomainConfigMessagingService.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

