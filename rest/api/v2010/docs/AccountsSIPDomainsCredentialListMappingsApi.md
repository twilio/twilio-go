# AccountsSIPDomainsCredentialListMappingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipCredentialListMapping**](AccountsSIPDomainsCredentialListMappingsApi.md#CreateSipCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 
[**DeleteSipCredentialListMapping**](AccountsSIPDomainsCredentialListMappingsApi.md#DeleteSipCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**FetchSipCredentialListMapping**](AccountsSIPDomainsCredentialListMappingsApi.md#FetchSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**ListSipCredentialListMapping**](AccountsSIPDomainsCredentialListMappingsApi.md#ListSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 



## CreateSipCredentialListMapping

> ApiV2010SipCredentialListMapping CreateSipCredentialListMapping(ctx, DomainSidoptional)



Create a CredentialListMapping resource for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**CredentialListSid** | **string** | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010SipCredentialListMapping**](ApiV2010SipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialListMapping

> DeleteSipCredentialListMapping(ctx, DomainSidSidoptional)



Delete a CredentialListMapping resource from an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

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


## FetchSipCredentialListMapping

> ApiV2010SipCredentialListMapping FetchSipCredentialListMapping(ctx, DomainSidSidoptional)



Fetch a single CredentialListMapping resource from an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Return type

[**ApiV2010SipCredentialListMapping**](ApiV2010SipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialListMapping

> []ApiV2010SipCredentialListMapping ListSipCredentialListMapping(ctx, DomainSidoptional)



Read multiple CredentialListMapping resources from an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipCredentialListMapping**](ApiV2010SipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

