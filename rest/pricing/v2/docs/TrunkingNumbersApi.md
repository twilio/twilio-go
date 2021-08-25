# TrunkingNumbersApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTrunkingNumber**](TrunkingNumbersApi.md#FetchTrunkingNumber) | **Get** /v2/Trunking/Numbers/{DestinationNumber} | 



## FetchTrunkingNumber

> PricingV2TrunkingNumber FetchTrunkingNumber(ctx, DestinationNumberoptional)



Fetch pricing information for a specific destination and, optionally, origination phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DestinationNumber** | **string** | The destination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunkingNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**OriginationNumber** | **string** | The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Return type

[**PricingV2TrunkingNumber**](PricingV2TrunkingNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

