# PhoneNumbersApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPhoneNumber**](PhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v2/PhoneNumbers/{PhoneNumber} | 



## FetchPhoneNumber

> LookupsV2PhoneNumber FetchPhoneNumber(ctx, PhoneNumberoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number to lookup in E.164 or national format. Default country code is +1 (North America).

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**Fields** | **string** | A comma-separated list of fields to return. Possible values are caller_name, sim_swap, call_forwarding, live_activity, line_type_intelligence, identity_match.
**CountryCode** | **string** | The [country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) used if the phone number provided is in national format.
**FirstName** | **string** | User’s first name. This query parameter is only used (optionally) for identity_match package requests.
**LastName** | **string** | User’s last name. This query parameter is only used (optionally) for identity_match package requests.
**AddressLine1** | **string** | User’s first address line. This query parameter is only used (optionally) for identity_match package requests.
**AddressLine2** | **string** | User’s second address line. This query parameter is only used (optionally) for identity_match package requests.
**City** | **string** | User’s city. This query parameter is only used (optionally) for identity_match package requests.
**State** | **string** | User’s country subdivision, such as state, province, or locality. This query parameter is only used (optionally) for identity_match package requests.
**PostalCode** | **string** | User’s postal zip code. This query parameter is only used (optionally) for identity_match package requests.
**AddressCountryCode** | **string** | User’s country, up to two characters. This query parameter is only used (optionally) for identity_match package requests.
**NationalId** | **string** | User’s national ID, such as SSN or Passport ID. This query parameter is only used (optionally) for identity_match package requests.
**DateOfBirth** | **string** | User’s date of birth, in YYYYMMDD format. This query parameter is only used (optionally) for identity_match package requests.

### Return type

[**LookupsV2PhoneNumber**](LookupsV2PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

