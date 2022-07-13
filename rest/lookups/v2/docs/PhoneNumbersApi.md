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
**Fields** | **string** | A comma-separated list of fields to return. Possible values are caller_name, sim_swap, call_forwarding, live_activity, line_type_intelligence.
**CountryCode** | **string** | The [country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) used if the phone number provided is in national format.

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

