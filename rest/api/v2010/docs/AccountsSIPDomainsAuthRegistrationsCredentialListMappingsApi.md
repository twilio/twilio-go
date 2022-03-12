# AccountsSIPDomainsAuthRegistrationsCredentialListMappingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipAuthRegistrationsCredentialListMapping**](AccountsSIPDomainsAuthRegistrationsCredentialListMappingsApi.md#CreateSipAuthRegistrationsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 
[**DeleteSipAuthRegistrationsCredentialListMapping**](AccountsSIPDomainsAuthRegistrationsCredentialListMappingsApi.md#DeleteSipAuthRegistrationsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**FetchSipAuthRegistrationsCredentialListMapping**](AccountsSIPDomainsAuthRegistrationsCredentialListMappingsApi.md#FetchSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**ListSipAuthRegistrationsCredentialListMapping**](AccountsSIPDomainsAuthRegistrationsCredentialListMappingsApi.md#ListSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 



## CreateSipAuthRegistrationsCredentialListMapping

> ApiV2010SipAuthRegistrationsCredentialListMapping CreateSipAuthRegistrationsCredentialListMapping(ctx, DomainSidoptional)



Create a new credential list mapping resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthRegistrationsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**CredentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010SipAuthRegistrationsCredentialListMapping**](ApiV2010SipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthRegistrationsCredentialListMapping

> DeleteSipAuthRegistrationsCredentialListMapping(ctx, DomainSidSidoptional)



Delete a credential list mapping from the requested domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthRegistrationsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.

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


## FetchSipAuthRegistrationsCredentialListMapping

> ApiV2010SipAuthRegistrationsCredentialListMapping FetchSipAuthRegistrationsCredentialListMapping(ctx, DomainSidSidoptional)



Fetch a specific instance of a credential list mapping

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthRegistrationsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.

### Return type

[**ApiV2010SipAuthRegistrationsCredentialListMapping**](ApiV2010SipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthRegistrationsCredentialListMapping

> []ApiV2010SipAuthRegistrationsCredentialListMapping ListSipAuthRegistrationsCredentialListMapping(ctx, DomainSidoptional)



Retrieve a list of credential list mappings belonging to the domain used in the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthRegistrationsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipAuthRegistrationsCredentialListMapping**](ApiV2010SipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

