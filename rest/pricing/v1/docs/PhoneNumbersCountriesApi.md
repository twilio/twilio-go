# PhoneNumbersCountriesApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPhoneNumberCountry**](PhoneNumbersCountriesApi.md#FetchPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries/{IsoCountry} | 
[**ListPhoneNumberCountry**](PhoneNumbersCountriesApi.md#ListPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries | 



## FetchPhoneNumberCountry

> PricingV1PhoneNumberCountryInstance FetchPhoneNumberCountry(ctx, IsoCountry)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV1PhoneNumberCountryInstance**](PricingV1PhoneNumberCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberCountry

> []PricingV1PhoneNumberCountry ListPhoneNumberCountry(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]PricingV1PhoneNumberCountry**](PricingV1PhoneNumberCountry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

