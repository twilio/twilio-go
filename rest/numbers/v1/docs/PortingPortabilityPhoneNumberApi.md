# PortingPortabilityPhoneNumberApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPortingPortability**](PortingPortabilityPhoneNumberApi.md#FetchPortingPortability) | **Get** /v1/Porting/Portability/PhoneNumber/{PhoneNumber} | 



## FetchPortingPortability

> NumbersV1PortingPortability FetchPortingPortability(ctx, PhoneNumberoptional)



Allows to check if a single phone number can be ported to Twilio or not.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number which portability is to be checked. Phone numbers are in E.164 format (e.g. +16175551212).

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingPortabilityParams struct


Name | Type | Description
------------- | ------------- | -------------
**TargetAccountSid** | **string** | The SID of the account where the phone number(s) will be ported.

### Return type

[**NumbersV1PortingPortability**](NumbersV1PortingPortability.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

