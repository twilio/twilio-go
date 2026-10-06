# ShortCodesApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchShortCode**](ShortCodesApi.md#FetchShortCode) | **Get** /v3/ShortCodes/{isoCountryCode}/{shortCode} | Fetch the Processing Region assigned to a short code.
[**UpdateShortCode**](ShortCodesApi.md#UpdateShortCode) | **Post** /v3/ShortCodes/{isoCountryCode}/{shortCode} | Assign an Processing Region to a short code.



## FetchShortCode

> RoutesV3ShortCode FetchShortCode(ctx, IsoCountryCodeShortCode)

Fetch the Processing Region assigned to a short code.

Fetch the Processing Region assigned to a short code. Short codes only support a Messaging region; Voice is not supported.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountryCode** | **string** | The ISO country code of the short code
**ShortCode** | **string** | The short code

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV3ShortCode**](RoutesV3ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> RoutesV3ShortCode UpdateShortCode(ctx, IsoCountryCodeShortCodeoptional)

Assign an Processing Region to a short code.

Assign an Processing Region to a short code. Short codes only support a Messaging region; Voice is not supported.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountryCode** | **string** | The ISO country code of the short code
**ShortCode** | **string** | The short code

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessagingRegion** | **string** | The Processing Region used for this short code for messaging
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.

### Return type

[**RoutesV3ShortCode**](RoutesV3ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

