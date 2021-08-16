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
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
**AppHash** | **string** | Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: &#x60;&lt;#&gt; Your AppName verification code is: 1234 He42w354ol9&#x60;.
**Channel** | **string** | The verification method to use. Can be: [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email), &#x60;sms&#x60; or &#x60;call&#x60;.
**ChannelConfiguration** | [**map[string]interface{}**](map[string]interface{}.md) | [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include &#39;from&#39; and &#39;from_name&#39;.
**CustomCode** | **string** | A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive.
**CustomFriendlyName** | **string** | A custom user defined friendly name that overwrites the existing one in the verification message
**CustomMessage** | **string** | The text of a custom message to use for the verification.
**Locale** | **string** | The locale to use for the verification SMS or call. Can be: &#x60;af&#x60;, &#x60;ar&#x60;, &#x60;ca&#x60;, &#x60;cs&#x60;, &#x60;da&#x60;, &#x60;de&#x60;, &#x60;el&#x60;, &#x60;en&#x60;, &#x60;en-GB&#x60;, &#x60;es&#x60;, &#x60;fi&#x60;, &#x60;fr&#x60;, &#x60;he&#x60;, &#x60;hi&#x60;, &#x60;hr&#x60;, &#x60;hu&#x60;, &#x60;id&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;ko&#x60;, &#x60;ms&#x60;, &#x60;nb&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;, &#x60;pr-BR&#x60;, &#x60;ro&#x60;, &#x60;ru&#x60;, &#x60;sv&#x60;, &#x60;th&#x60;, &#x60;tl&#x60;, &#x60;tr&#x60;, &#x60;vi&#x60;, &#x60;zh&#x60;, &#x60;zh-CN&#x60;, or &#x60;zh-HK.&#x60;
**Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
**RateLimits** | [**map[string]interface{}**](map[string]interface{}.md) | The custom key-value pairs of Programmable Rate Limits. Keys correspond to &#x60;unique_name&#x60; fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request.
**SendDigits** | **string** | The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits).
**TemplateSid** | **string** | The SID of the custom template to be used.
**To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

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
**Status** | **string** | The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;approved&#x60;.

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

