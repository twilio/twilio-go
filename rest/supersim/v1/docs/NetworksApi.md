# NetworksApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchNetwork**](NetworksApi.md#FetchNetwork) | **Get** /v1/Networks/{Sid} | 
[**ListNetwork**](NetworksApi.md#ListNetwork) | **Get** /v1/Networks | 



## FetchNetwork

> SupersimV1Network FetchNetwork(ctx, Sid)



Fetch a Network resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1Network**](SupersimV1Network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork

> []SupersimV1Network ListNetwork(ctx, optional)



Retrieve a list of Network resources.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
**Mcc** | **string** | The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read.
**Mnc** | **string** | The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1Network**](SupersimV1Network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

