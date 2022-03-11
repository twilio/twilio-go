# AccountsSIPDomainsIpAccessControlListMappingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipIpAccessControlListMapping**](AccountsSIPDomainsIpAccessControlListMappingsApi.md#CreateSipIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 
[**DeleteSipIpAccessControlListMapping**](AccountsSIPDomainsIpAccessControlListMappingsApi.md#DeleteSipIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipIpAccessControlListMapping**](AccountsSIPDomainsIpAccessControlListMappingsApi.md#FetchSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**ListSipIpAccessControlListMapping**](AccountsSIPDomainsIpAccessControlListMappingsApi.md#ListSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 



## CreateSipIpAccessControlListMapping

> ApiV2010SipIpAccessControlListMapping CreateSipIpAccessControlListMapping(ctx, DomainSidoptional)



Create a new IpAccessControlListMapping resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**IpAccessControlListSid** | **string** | The unique id of the IP access control list to map to the SIP domain.

### Return type

[**ApiV2010SipIpAccessControlListMapping**](ApiV2010SipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlListMapping

> DeleteSipIpAccessControlListMapping(ctx, DomainSidSidoptional)



Delete an IpAccessControlListMapping resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

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


## FetchSipIpAccessControlListMapping

> ApiV2010SipIpAccessControlListMapping FetchSipIpAccessControlListMapping(ctx, DomainSidSidoptional)



Fetch an IpAccessControlListMapping resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

### Return type

[**ApiV2010SipIpAccessControlListMapping**](ApiV2010SipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlListMapping

> []ApiV2010SipIpAccessControlListMapping ListSipIpAccessControlListMapping(ctx, DomainSidoptional)



Retrieve a list of IpAccessControlListMapping resources.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAccessControlListMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipIpAccessControlListMapping**](ApiV2010SipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

