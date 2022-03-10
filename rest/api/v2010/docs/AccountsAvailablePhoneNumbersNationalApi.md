# AccountsAvailablePhoneNumbersNationalApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumberNational**](AccountsAvailablePhoneNumbersNationalApi.md#ListAvailablePhoneNumberNational) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/National.json | 



## ListAvailablePhoneNumberNational

> []ApiV2010AvailablePhoneNumberNational ListAvailablePhoneNumberNational(ctx, CountryCodeoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberNationalParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**AreaCode** | **int** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
**Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
**SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
**NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
**NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
**Distance** | **int** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
**InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
**InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
**InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
**InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
**InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
**FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010AvailablePhoneNumberNational**](ApiV2010AvailablePhoneNumberNational.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

