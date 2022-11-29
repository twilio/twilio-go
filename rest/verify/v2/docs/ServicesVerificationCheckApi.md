# ServicesVerificationCheckApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerificationCheck**](ServicesVerificationCheckApi.md#CreateVerificationCheck) | **Post** /v2/Services/{ServiceSid}/VerificationCheck | 



## CreateVerificationCheck

> VerifyV2VerificationCheck CreateVerificationCheck(ctx, ServiceSidoptional)



challenge a specific Verification Check.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateVerificationCheckParams struct


Name | Type | Description
------------- | ------------- | -------------
**Code** | **string** | The 4-10 character string being verified.
**To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the `verification_sid` must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
**VerificationSid** | **string** | A SID that uniquely identifies the Verification Check. Either this parameter or the `to` phone number/[email](https://www.twilio.com/docs/verify/email) must be specified.
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
**Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.

### Return type

[**VerifyV2VerificationCheck**](VerifyV2VerificationCheck.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

