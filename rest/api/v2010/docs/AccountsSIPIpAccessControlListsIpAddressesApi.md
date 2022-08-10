# AccountsSIPIpAccessControlListsIpAddressesApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipIpAddress**](AccountsSIPIpAccessControlListsIpAddressesApi.md#CreateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**DeleteSipIpAddress**](AccountsSIPIpAccessControlListsIpAddressesApi.md#DeleteSipIpAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**FetchSipIpAddress**](AccountsSIPIpAccessControlListsIpAddressesApi.md#FetchSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**ListSipIpAddress**](AccountsSIPIpAccessControlListsIpAddressesApi.md#ListSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**UpdateSipIpAddress**](AccountsSIPIpAccessControlListsIpAddressesApi.md#UpdateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 



## CreateSipIpAddress

> ApiV2010SipIpAddress CreateSipIpAddress(ctx, IpAccessControlListSidoptional)



Create a new IpAddress resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid with which to associate the created IpAddress resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**FriendlyName** | **string** | A human readable descriptive text for this resource, up to 255 characters long.
**IpAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
**CidrPrefixLength** | **int** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.

### Return type

[**ApiV2010SipIpAddress**](ApiV2010SipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAddress

> DeleteSipIpAddress(ctx, IpAccessControlListSidSidoptional)



Delete an IpAddress resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to delete.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAddressParams struct


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


## FetchSipIpAddress

> ApiV2010SipIpAddress FetchSipIpAddress(ctx, IpAccessControlListSidSidoptional)



Read one IpAddress resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to fetch.
**Sid** | **string** | A 34 character string that uniquely identifies the IpAddress resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Return type

[**ApiV2010SipIpAddress**](ApiV2010SipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAddress

> []ApiV2010SipIpAddress ListSipIpAddress(ctx, IpAccessControlListSidoptional)



Read multiple IpAddress resources.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipIpAddress**](ApiV2010SipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAddress

> ApiV2010SipIpAddress UpdateSipIpAddress(ctx, IpAccessControlListSidSidoptional)



Update an IpAddress resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to update.
**Sid** | **string** | A 34 character string that identifies the IpAddress resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipIpAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
**FriendlyName** | **string** | A human readable descriptive text for this resource, up to 255 characters long.
**CidrPrefixLength** | **int** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.

### Return type

[**ApiV2010SipIpAddress**](ApiV2010SipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

