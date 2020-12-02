# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLength** | **int32** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | [optional] 
**CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | [optional] 
**DoNotShareWarningEnabled** | **bool** | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: &#x60;Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code&#x60; | [optional] 
**DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 64 characters long. **This value should not contain PII.** | 
**LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number. | [optional] 
**Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification. | [optional] 
**Push** | **map[string]interface{}** | Configurations for the Push factors (channel) created under this Service. If present, it must be a json string with the following format: {\&quot;notify_service_sid\&quot;: \&quot;ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\&quot;, \&quot;include_date\&quot;: true}. If &#x60;include_date&#x60; is set to &#x60;true&#x60;, which is the default, that means that the push challenge’s response will include the date created value. If &#x60;include_date&#x60; is set to &#x60;false&#x60;, then the date created value will not be included. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info | [optional] 
**SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | [optional] 
**TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


