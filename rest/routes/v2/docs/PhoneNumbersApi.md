# PhoneNumbersApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPhoneNumber**](PhoneNumbersApi.md#FetchPhoneNumber) | **Get** /v2/PhoneNumbers/{PhoneNumber} | 
[**UpdatePhoneNumber**](PhoneNumbersApi.md#UpdatePhoneNumber) | **Post** /v2/PhoneNumbers/{PhoneNumber} | 



## FetchPhoneNumber

> RoutesV2PhoneNumber FetchPhoneNumber(ctx, PhoneNumber)



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

[**RoutesV2PhoneNumber**](RoutesV2PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> RoutesV2PhoneNumber UpdatePhoneNumber(ctx, PhoneNumberoptional)



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
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.

### Return type

[**RoutesV2PhoneNumber**](RoutesV2PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

