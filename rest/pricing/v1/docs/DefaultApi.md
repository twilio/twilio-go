# DefaultApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMessagingCountry**](DefaultApi.md#FetchMessagingCountry) | **Get** /v1/Messaging/Countries/{IsoCountry} | 
[**FetchPhoneNumberCountry**](DefaultApi.md#FetchPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries/{IsoCountry} | 
[**FetchVoiceCountry**](DefaultApi.md#FetchVoiceCountry) | **Get** /v1/Voice/Countries/{IsoCountry} | 
[**FetchVoiceNumber**](DefaultApi.md#FetchVoiceNumber) | **Get** /v1/Voice/Numbers/{Number} | 
[**ListMessagingCountry**](DefaultApi.md#ListMessagingCountry) | **Get** /v1/Messaging/Countries | 
[**ListPhoneNumberCountry**](DefaultApi.md#ListPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries | 
[**ListVoiceCountry**](DefaultApi.md#ListVoiceCountry) | **Get** /v1/Voice/Countries | 



## FetchMessagingCountry

> PricingV1MessagingMessagingCountryInstance FetchMessagingCountry(ctx, IsoCountry)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessagingCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV1MessagingMessagingCountryInstance**](PricingV1MessagingMessagingCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumberCountry

> PricingV1PhoneNumberPhoneNumberCountryInstance FetchPhoneNumberCountry(ctx, IsoCountry)



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

[**PricingV1PhoneNumberPhoneNumberCountryInstance**](PricingV1PhoneNumberPhoneNumberCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceCountry

> PricingV1VoiceVoiceCountryInstance FetchVoiceCountry(ctx, IsoCountry)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceCountryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**PricingV1VoiceVoiceCountryInstance**](PricingV1VoiceVoiceCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceNumber

> PricingV1VoiceVoiceNumber FetchVoiceNumber(ctx, Number)



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

[**PricingV1VoiceVoiceNumber**](PricingV1VoiceVoiceNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingCountry

> ListMessagingCountryResponse ListMessagingCountry(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMessagingCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessagingCountryResponse**](ListMessagingCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberCountry

> ListPhoneNumberCountryResponse ListPhoneNumberCountry(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberCountryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListPhoneNumberCountryResponse**](ListPhoneNumberCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

