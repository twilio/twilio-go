# TrunkingCountriesApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTrunkingCountry**](TrunkingCountriesApi.md#FetchTrunkingCountry) | **Get** /v2/Trunking/Countries/{IsoCountry} | 
[**ListTrunkingCountry**](TrunkingCountriesApi.md#ListTrunkingCountry) | **Get** /v2/Trunking/Countries | 



## FetchTrunkingCountry

> PricingV2TrunkingCountryInstance FetchTrunkingCountry(ctx, IsoCountry)



Fetch a specific Country.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the origin-based voice pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunkingCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV2TrunkingCountryInstance**](PricingV2TrunkingCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrunkingCountry

> []PricingV2TrunkingCountry ListTrunkingCountry(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTrunkingCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]PricingV2TrunkingCountry**](PricingV2TrunkingCountry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

