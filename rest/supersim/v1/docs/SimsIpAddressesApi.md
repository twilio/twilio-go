# SimsIpAddressesApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSimIpAddress**](SimsIpAddressesApi.md#ListSimIpAddress) | **Get** /v1/Sims/{SimSid}/IpAddresses | 



## ListSimIpAddress

> []SupersimV1SimIpAddress ListSimIpAddress(ctx, SimSidoptional)



Retrieve a list of IP Addresses for the given Super SIM.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SimSid** | **string** | The SID of the Super SIM to list IP Addresses for.

### Other Parameters

Other parameters are passed through a pointer to a ListSimIpAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1SimIpAddress**](SupersimV1SimIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

