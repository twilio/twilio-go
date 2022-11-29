# AccountsAvailablePhoneNumbersSharedCostApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumberSharedCost**](AccountsAvailablePhoneNumbersSharedCostApi.md#ListAvailablePhoneNumberSharedCost) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/SharedCost.json | 



## ListAvailablePhoneNumberSharedCost

> []ApiV2010AvailablePhoneNumberSharedCost ListAvailablePhoneNumberSharedCost(ctx, CountryCodeoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberSharedCostParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**AreaCode** | **int** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
**Contains** | **string** | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
**SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: `true` or `false`.
**MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`.
**VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: `true` or `false`.
**ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
**ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
**ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
**Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
**NearNumber** | **string** | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
**NearLatLong** | **string** | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada.
**Distance** | **int** | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada.
**InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
**InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
**InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada.
**InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
**InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
**FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: `true` or `false`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010AvailablePhoneNumberSharedCost**](ApiV2010AvailablePhoneNumberSharedCost.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

