# CallerIdsVoiceVerificationsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoiceVerification**](CallerIdsVoiceVerificationsApi.md#CreateVoiceVerification) | **Post** /v1/CallerIds/VoiceVerifications | Initiate voice verification for a caller ID



## CreateVoiceVerification

> NumbersV1VoiceVerification CreateVoiceVerification(ctx, optional)

Initiate voice verification for a caller ID

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateVoiceVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**NumbersV1VoiceVerificationRequest** | [**NumbersV1VoiceVerificationRequest**](NumbersV1VoiceVerificationRequest.md) | 

### Return type

[**NumbersV1VoiceVerification**](NumbersV1VoiceVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

