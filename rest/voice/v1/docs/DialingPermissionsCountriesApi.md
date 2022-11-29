# DialingPermissionsCountriesApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDialingPermissionsCountry**](DialingPermissionsCountriesApi.md#FetchDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries/{IsoCode} | 
[**ListDialingPermissionsCountry**](DialingPermissionsCountriesApi.md#ListDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries | 



## FetchDialingPermissionsCountry

> VoiceV1DialingPermissionsCountryInstance FetchDialingPermissionsCountry(ctx, IsoCode)



Retrieve voice dialing country permissions identified by the given ISO country code

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchDialingPermissionsCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1DialingPermissionsCountryInstance**](VoiceV1DialingPermissionsCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsCountry

> []VoiceV1DialingPermissionsCountry ListDialingPermissionsCountry(ctx, optional)



Retrieve all voice dialing country permissions for this account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDialingPermissionsCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**IsoCode** | **string** | Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
**Continent** | **string** | Filter to retrieve the country permissions by specifying the continent
**CountryCode** | **string** | Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html)
**LowRiskNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: `true` or `false`.
**HighRiskSpecialNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: `true` or `false`
**HighRiskTollfraudNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: `true` or `false`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1DialingPermissionsCountry**](VoiceV1DialingPermissionsCountry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

