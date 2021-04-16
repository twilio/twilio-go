# DefaultApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVoiceCountry**](DefaultApi.md#FetchVoiceCountry) | **Get** /v2/Voice/Countries/{IsoCountry} | 
[**FetchVoiceNumber**](DefaultApi.md#FetchVoiceNumber) | **Get** /v2/Voice/Numbers/{DestinationNumber} | 
[**ListVoiceCountry**](DefaultApi.md#ListVoiceCountry) | **Get** /v2/Voice/Countries | 



## FetchVoiceCountry

> PricingV2VoiceVoiceCountryInstance FetchVoiceCountry(ctx, IsoCountry)



Fetch a specific Country.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the origin-based voice pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV2VoiceVoiceCountryInstance**](PricingV2VoiceVoiceCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceNumber

> PricingV2VoiceVoiceNumber FetchVoiceNumber(ctx, DestinationNumberoptional)



Fetch pricing information for a specific destination and, optionally, origination phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DestinationNumber** | **string** | The destination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**OriginationNumber** | **string** | The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Return type

[**PricingV2VoiceVoiceNumber**](PricingV2VoiceVoiceNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVoiceCountry

> ListVoiceCountryResponse ListVoiceCountry(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVoiceCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVoiceCountryResponse**](ListVoiceCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

