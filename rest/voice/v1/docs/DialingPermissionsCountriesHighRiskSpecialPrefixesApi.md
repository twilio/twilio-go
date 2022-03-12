# DialingPermissionsCountriesHighRiskSpecialPrefixesApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDialingPermissionsHrsPrefixes**](DialingPermissionsCountriesHighRiskSpecialPrefixesApi.md#ListDialingPermissionsHrsPrefixes) | **Get** /v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes | 



## ListDialingPermissionsHrsPrefixes

> []VoiceV1DialingPermissionsHrsPrefixes ListDialingPermissionsHrsPrefixes(ctx, IsoCodeoptional)



Fetch the high-risk special services prefixes from the country resource corresponding to the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string** | The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched

### Other Parameters

Other parameters are passed through a pointer to a ListDialingPermissionsHrsPrefixesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1DialingPermissionsHrsPrefixes**](VoiceV1DialingPermissionsHrsPrefixes.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

