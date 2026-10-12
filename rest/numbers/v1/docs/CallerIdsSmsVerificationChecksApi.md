# CallerIdsSmsVerificationChecksApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsVerificationCheck**](CallerIdsSmsVerificationChecksApi.md#CreateSmsVerificationCheck) | **Post** /v1/CallerIds/SmsVerificationChecks | Check SMS verification code



## CreateSmsVerificationCheck

> NumbersV1SmsVerificationCheck CreateSmsVerificationCheck(ctx, optional)

Check SMS verification code

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSmsVerificationCheckParams struct


Name | Type | Description
------------- | ------------- | -------------
**NumbersV1SmsVerificationCheckRequest** | [**NumbersV1SmsVerificationCheckRequest**](NumbersV1SmsVerificationCheckRequest.md) | 

### Return type

[**NumbersV1SmsVerificationCheck**](NumbersV1SmsVerificationCheck.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

