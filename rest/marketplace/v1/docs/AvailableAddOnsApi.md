# AvailableAddOnsApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMarketplaceAvailableAddOn**](AvailableAddOnsApi.md#FetchMarketplaceAvailableAddOn) | **Get** /v1/AvailableAddOns/{Sid} | 
[**ListMarketplaceAvailableAddOn**](AvailableAddOnsApi.md#ListMarketplaceAvailableAddOn) | **Get** /v1/AvailableAddOns | 



## FetchMarketplaceAvailableAddOn

> MarketplaceAvailableAddOn FetchMarketplaceAvailableAddOn(ctx, Sid)



Fetch an instance of an Add-on currently available to be installed.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the AvailableAddOn resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMarketplaceAvailableAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceAvailableAddOn**](MarketplaceAvailableAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketplaceAvailableAddOn

> []MarketplaceAvailableAddOn ListMarketplaceAvailableAddOn(ctx, optional)



Retrieve a list of Add-ons currently available to be installed.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMarketplaceAvailableAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MarketplaceAvailableAddOn**](MarketplaceAvailableAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

