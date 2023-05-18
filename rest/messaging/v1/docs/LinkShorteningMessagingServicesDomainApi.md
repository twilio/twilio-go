# LinkShorteningMessagingServicesDomainApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchLinkshorteningMessagingServiceDomainAssociation**](LinkShorteningMessagingServicesDomainApi.md#FetchLinkshorteningMessagingServiceDomainAssociation) | **Get** /v1/LinkShortening/MessagingServices/{MessagingServiceSid}/Domain | 



## FetchLinkshorteningMessagingServiceDomainAssociation

> MessagingV1LinkshorteningMessagingServiceDomainAssociation FetchLinkshorteningMessagingServiceDomainAssociation(ctx, MessagingServiceSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | Unique string used to identify the Messaging service that this domain should be associated with.

### Other Parameters

Other parameters are passed through a pointer to a FetchLinkshorteningMessagingServiceDomainAssociationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1LinkshorteningMessagingServiceDomainAssociation**](MessagingV1LinkshorteningMessagingServiceDomainAssociation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

