# ServicesApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v2/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v2/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 



## CreateService

> VerifyV2Service CreateService(ctx, optional)



Create a new Verification Service.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 32 characters long. **This value should not contain PII.**
**CodeLength** | **int** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
**LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number.
**SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
**DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call.
**TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
**Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification.
**DoNotShareWarningEnabled** | **bool** | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: `Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code`
**CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
**PushIncludeDate** | **bool** | Optional configuration for the Push factors. If true, include the date in the Challenge's response. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: false. **Deprecated** do not use this parameter. This timestamp value is the same one as the one found in `date_created`, please use that one instead.
**PushApnCredentialSid** | **string** | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**PushFcmCredentialSid** | **string** | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**TotpIssuer** | **string** | Optional configuration for the TOTP factors. Set TOTP Issuer for this service. This will allow to configure the issuer of the TOTP URI. Defaults to the service friendly name if not provided.
**TotpTimeStep** | **int** | Optional configuration for the TOTP factors. Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
**TotpCodeLength** | **int** | Optional configuration for the TOTP factors. Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
**TotpSkew** | **int** | Optional configuration for the TOTP factors. The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
**DefaultTemplateSid** | **string** | The default message [template](https://www.twilio.com/docs/verify/api/templates). Will be used for all SMS verifications unless explicitly overriden. SMS channel only.

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)



Delete a specific Verification Service Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> VerifyV2Service FetchService(ctx, Sid)



Fetch specific Verification Service Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> []VerifyV2Service ListService(ctx, optional)



Retrieve a list of all Verification Services for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> VerifyV2Service UpdateService(ctx, Sidoptional)



Update a specific Verification Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 32 characters long. **This value should not contain PII.**
**CodeLength** | **int** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
**LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number.
**SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
**DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call.
**TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
**Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification.
**DoNotShareWarningEnabled** | **bool** | Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.**
**CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
**PushIncludeDate** | **bool** | Optional configuration for the Push factors. If true, include the date in the Challenge's response. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: false. **Deprecated** do not use this parameter.
**PushApnCredentialSid** | **string** | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**PushFcmCredentialSid** | **string** | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
**TotpIssuer** | **string** | Optional configuration for the TOTP factors. Set TOTP Issuer for this service. This will allow to configure the issuer of the TOTP URI.
**TotpTimeStep** | **int** | Optional configuration for the TOTP factors. Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
**TotpCodeLength** | **int** | Optional configuration for the TOTP factors. Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
**TotpSkew** | **int** | Optional configuration for the TOTP factors. The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
**DefaultTemplateSid** | **string** | The default message [template](https://www.twilio.com/docs/verify/api/templates). Will be used for all SMS verifications unless explicitly overriden. SMS channel only.

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

