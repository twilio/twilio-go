# \DefaultApi

All URIs are relative to *http://localhost*

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

> PricingV1MessagingMessagingCountryInstance FetchMessagingCountry(ctx, isoCountry)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isoCountry** | **string**| The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch. | 

### Return type

[**PricingV1MessagingMessagingCountryInstance**](pricing.v1.messaging.messaging_country-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumberCountry

> PricingV1PhoneNumberPhoneNumberCountryInstance FetchPhoneNumberCountry(ctx, isoCountry)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isoCountry** | **string**| The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch. | 

### Return type

[**PricingV1PhoneNumberPhoneNumberCountryInstance**](pricing.v1.phone_number.phone_number_country-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceCountry

> PricingV1VoiceVoiceCountryInstance FetchVoiceCountry(ctx, isoCountry)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isoCountry** | **string**| The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch. | 

### Return type

[**PricingV1VoiceVoiceCountryInstance**](pricing.v1.voice.voice_country-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceNumber

> PricingV1VoiceVoiceNumber FetchVoiceNumber(ctx, number)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **string**| The phone number to fetch. | 

### Return type

[**PricingV1VoiceVoiceNumber**](pricing.v1.voice.voice_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingCountry

> PricingV1MessagingMessagingCountryReadResponse ListMessagingCountry(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListMessagingCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessagingCountryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PricingV1MessagingMessagingCountryReadResponse**](pricing_v1_messaging_messaging_countryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberCountry

> PricingV1PhoneNumberPhoneNumberCountryReadResponse ListPhoneNumberCountry(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPhoneNumberCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPhoneNumberCountryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PricingV1PhoneNumberPhoneNumberCountryReadResponse**](pricing_v1_phone_number_phone_number_countryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVoiceCountry

> PricingV1VoiceVoiceCountryReadResponse ListVoiceCountry(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVoiceCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVoiceCountryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PricingV1VoiceVoiceCountryReadResponse**](pricing_v1_voice_voice_countryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

