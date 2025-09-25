# ServicesPasskeysVerifyFactorApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePasskeysFactor**](ServicesPasskeysVerifyFactorApi.md#UpdatePasskeysFactor) | **Post** /v2/Services/{ServiceSid}/Passkeys/VerifyFactor | Verify a Passkeys Factor



## UpdatePasskeysFactor

> UpdatePasskeysFactorResponse UpdatePasskeysFactor(ctx, ServiceSidoptional)

Verify a Passkeys Factor

Verify a Passkeys Factor

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePasskeysFactorParams struct


Name | Type | Description
------------- | ------------- | -------------
**VerifyPasskeysFactorRequest** | [**VerifyPasskeysFactorRequest**](VerifyPasskeysFactorRequest.md) | 

### Return type

[**UpdatePasskeysFactorResponse**](UpdatePasskeysFactor200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

