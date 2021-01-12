# InlineObject18

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLength** | **int32** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | [optional] 
**CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | [optional] 
**DoNotShareWarningEnabled** | **bool** | Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.** | [optional] 
**DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.** | [optional] 
**LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number. | [optional] 
**Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification. | [optional] 
**PushApnCredentialSid** | **string** | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | [optional] 
**PushFcmCredentialSid** | **string** | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | [optional] 
**PushIncludeDate** | **bool** | Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true | [optional] 
**SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | [optional] 
**TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


