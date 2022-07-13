# VoiceNumbersApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVoiceNumber**](VoiceNumbersApi.md#FetchVoiceNumber) | **Get** /v1/Voice/Numbers/{Number} | 



## FetchVoiceNumber

> PricingV1VoiceNumber FetchVoiceNumber(ctx, Number)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Number** | **string** | The phone number to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV1VoiceNumber**](PricingV1VoiceNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

