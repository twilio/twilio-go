# AccountsSIPDomainsAuthCallsIpAccessControlListMappingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipAuthCallsIpAccessControlListMapping**](AccountsSIPDomainsAuthCallsIpAccessControlListMappingsApi.md#CreateSipAuthCallsIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 
[**DeleteSipAuthCallsIpAccessControlListMapping**](AccountsSIPDomainsAuthCallsIpAccessControlListMappingsApi.md#DeleteSipAuthCallsIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipAuthCallsIpAccessControlListMapping**](AccountsSIPDomainsAuthCallsIpAccessControlListMappingsApi.md#FetchSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**ListSipAuthCallsIpAccessControlListMapping**](AccountsSIPDomainsAuthCallsIpAccessControlListMappingsApi.md#ListSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 



## CreateSipAuthCallsIpAccessControlListMapping

> ApiV2010SipAuthCallsIpAccessControlListMapping CreateSipAuthCallsIpAccessControlListMapping(ctx, DomainSidoptional)



Create a new IP Access Control List mapping

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthCallsIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**IpAccessControlListSid** | **string** | The SID of the IpAccessControlList resource to map to the SIP domain.

### Return type

[**ApiV2010SipAuthCallsIpAccessControlListMapping**](ApiV2010SipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsIpAccessControlListMapping

> DeleteSipAuthCallsIpAccessControlListMapping(ctx, DomainSidSidoptional)



Delete an IP Access Control List mapping from the requested domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthCallsIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete.

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


## FetchSipAuthCallsIpAccessControlListMapping

> ApiV2010SipAuthCallsIpAccessControlListMapping FetchSipAuthCallsIpAccessControlListMapping(ctx, DomainSidSidoptional)



Fetch a specific instance of an IP Access Control List mapping

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthCallsIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch.

### Return type

[**ApiV2010SipAuthCallsIpAccessControlListMapping**](ApiV2010SipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsIpAccessControlListMapping

> []ApiV2010SipAuthCallsIpAccessControlListMapping ListSipAuthCallsIpAccessControlListMapping(ctx, DomainSidoptional)



Retrieve a list of IP Access Control List mappings belonging to the domain used in the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthCallsIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipAuthCallsIpAccessControlListMapping**](ApiV2010SipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

