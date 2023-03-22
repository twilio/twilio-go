# LinkShorteningDomainsMessagingServicesApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLinkshorteningMessagingService**](LinkShorteningDomainsMessagingServicesApi.md#CreateLinkshorteningMessagingService) | **Post** /v1/LinkShortening/Domains/{DomainSid}/MessagingServices/{MessagingServiceSid} | 
[**DeleteLinkshorteningMessagingService**](LinkShorteningDomainsMessagingServicesApi.md#DeleteLinkshorteningMessagingService) | **Delete** /v1/LinkShortening/Domains/{DomainSid}/MessagingServices/{MessagingServiceSid} | 



## CreateLinkshorteningMessagingService

> MessagingV1LinkshorteningMessagingService CreateLinkshorteningMessagingService(ctx, DomainSidMessagingServiceSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The domain SID to associate with a messaging service. With URL shortening enabled, links in messages sent with the associated messaging service will be shortened to the provided domain
**MessagingServiceSid** | **string** | A messaging service SID to associate with a domain. With URL shortening enabled, links in messages sent with the provided messaging service will be shortened to the associated domain

### Other Parameters

Other parameters are passed through a pointer to a CreateLinkshorteningMessagingServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1LinkshorteningMessagingService**](MessagingV1LinkshorteningMessagingService.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkshorteningMessagingService

> DeleteLinkshorteningMessagingService(ctx, DomainSidMessagingServiceSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The domain SID to dissociate from a messaging service. With URL shortening enabled, links in messages sent with the associated messaging service will be shortened to the provided domain
**MessagingServiceSid** | **string** | A messaging service SID to dissociate from a domain. With URL shortening enabled, links in messages sent with the provided messaging service will be shortened to the associated domain

### Other Parameters

Other parameters are passed through a pointer to a DeleteLinkshorteningMessagingServiceParams struct


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

