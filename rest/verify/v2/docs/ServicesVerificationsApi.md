# ServicesVerificationsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerification**](ServicesVerificationsApi.md#CreateVerification) | **Post** /v2/Services/{ServiceSid}/Verifications | 
[**FetchVerification**](ServicesVerificationsApi.md#FetchVerification) | **Get** /v2/Services/{ServiceSid}/Verifications/{Sid} | 
[**UpdateVerification**](ServicesVerificationsApi.md#UpdateVerification) | **Post** /v2/Services/{ServiceSid}/Verifications/{Sid} | 



## CreateVerification

> VerifyV2Verification CreateVerification(ctx, ServiceSidoptional)



Create a new Verification using a Service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
**Channel** | **string** | The verification method to use. One of: [`email`](https://www.twilio.com/docs/verify/email), `sms`, `whatsapp`, `call`, `sna` or `auto`.
**CustomFriendlyName** | **string** | A custom user defined friendly name that overwrites the existing one in the verification message
**CustomMessage** | **string** | The text of a custom message to use for the verification.
**SendDigits** | **string** | The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits).
**Locale** | **string** | Locale will automatically resolve based on phone number country code for SMS, WhatsApp, and call channel verifications. It will fallback to English or the template’s default translation if the selected translation is not available. This parameter will override the automatic locale resolution. [See supported languages and more information here](https://www.twilio.com/docs/verify/supported-languages).
**CustomCode** | **string** | A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive.
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
**Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
**RateLimits** | [**interface{}**](interface{}.md) | The custom key-value pairs of Programmable Rate Limits. Keys correspond to `unique_name` fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request.
**ChannelConfiguration** | [**interface{}**](interface{}.md) | [`email`](https://www.twilio.com/docs/verify/email) channel configuration in json format. The fields 'from' and 'from_name' are optional but if included the 'from' field must have a valid email address.
**AppHash** | **string** | Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: `<#> Your AppName verification code is: 1234 He42w354ol9`.
**TemplateSid** | **string** | The message [template](https://www.twilio.com/docs/verify/api/templates). If provided, will override the default template for the Service. SMS and Voice channels only.
**TemplateCustomSubstitutions** | **string** | A stringified JSON object in which the keys are the template's special variables and the values are the variables substitutions.
**DeviceIp** | **string** | Strongly encouraged if using the auto channel. The IP address of the client's device. If provided, it has to be a valid IPv4 or IPv6 address.

### Return type

[**VerifyV2Verification**](VerifyV2Verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVerification

> VerifyV2Verification FetchVerification(ctx, ServiceSidSid)



Fetch a specific Verification

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Verification**](VerifyV2Verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVerification

> VerifyV2Verification UpdateVerification(ctx, ServiceSidSidoptional)



Update a Verification status

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 

### Return type

[**VerifyV2Verification**](VerifyV2Verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

