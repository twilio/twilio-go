# AccountsSIPDomainsAuthCallsCredentialListMappingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipAuthCallsCredentialListMapping**](AccountsSIPDomainsAuthCallsCredentialListMappingsApi.md#CreateSipAuthCallsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 
[**DeleteSipAuthCallsCredentialListMapping**](AccountsSIPDomainsAuthCallsCredentialListMappingsApi.md#DeleteSipAuthCallsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**FetchSipAuthCallsCredentialListMapping**](AccountsSIPDomainsAuthCallsCredentialListMappingsApi.md#FetchSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**ListSipAuthCallsCredentialListMapping**](AccountsSIPDomainsAuthCallsCredentialListMappingsApi.md#ListSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 



## CreateSipAuthCallsCredentialListMapping

> ApiV2010SipAuthCallsCredentialListMapping CreateSipAuthCallsCredentialListMapping(ctx, DomainSidoptional)



Create a new credential list mapping resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthCallsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**CredentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010SipAuthCallsCredentialListMapping**](ApiV2010SipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsCredentialListMapping

> DeleteSipAuthCallsCredentialListMapping(ctx, DomainSidSidoptional)



Delete a credential list mapping from the requested domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthCallsCredentialListMappingParams struct


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


## FetchSipAuthCallsCredentialListMapping

> ApiV2010SipAuthCallsCredentialListMapping FetchSipAuthCallsCredentialListMapping(ctx, DomainSidSidoptional)



Fetch a specific instance of a credential list mapping

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthCallsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.

### Return type

[**ApiV2010SipAuthCallsCredentialListMapping**](ApiV2010SipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsCredentialListMapping

> []ApiV2010SipAuthCallsCredentialListMapping ListSipAuthCallsCredentialListMapping(ctx, DomainSidoptional)



Retrieve a list of credential list mappings belonging to the domain used in the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthCallsCredentialListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipAuthCallsCredentialListMapping**](ApiV2010SipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

