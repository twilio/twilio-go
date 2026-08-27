# PhoneNumbersApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPhoneNumber**](PhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v3/PhoneNumbers/{phoneNumber} | Fetch the Inbound Processing Region assigned to a phone number.
[**UpdatePhoneNumber**](PhoneNumbersApi.md#UpdatePhoneNumber) | **Post** /v3/PhoneNumbers/{phoneNumber} | Assign an Inbound Processing Region to a phone number.



## FetchPhoneNumber

> RoutesV3PhoneNumber FetchPhoneNumber(ctx, PhoneNumber)

Fetch the Inbound Processing Region assigned to a phone number.

Fetch the Inbound Processing Region assigned to a phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number in E.164 format

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV3PhoneNumber**](RoutesV3PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> RoutesV3PhoneNumber UpdatePhoneNumber(ctx, PhoneNumberoptional)

Assign an Inbound Processing Region to a phone number.

Assign an Inbound Processing Region to a phone number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number in E.164 format

### Other Parameters

Other parameters are passed through a pointer to a UpdatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**VoiceRegion** | **string** | The Inbound Processing Region used for this phone number for voice
**MessagingRegion** | **string** | The Inbound Processing Region used for this phone number for messaging
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.

### Return type

[**RoutesV3PhoneNumber**](RoutesV3PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

