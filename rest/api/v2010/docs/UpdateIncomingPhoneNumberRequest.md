# UpdateIncomingPhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). | [optional] 
**AddressSid** | **string** | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations. | [optional] 
**ApiVersion** | **string** | The API version to use for incoming calls made to the phone number. The default is &#x60;2010-04-01&#x60;. | [optional] 
**BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | [optional] 
**EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from this phone number. | [optional] 
**EmergencyStatus** | **string** | The configuration status parameter that determines whether the phone number is enabled for emergency calling. | [optional] 
**FriendlyName** | **string** | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | [optional] 
**IdentitySid** | **string** | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations. | [optional] 
**SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | [optional] 
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | [optional] 
**SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message. | [optional] 
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | [optional] 
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**TrunkSid** | **string** | The SID of the Trunk we should use to handle phone calls to the phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | [optional] 
**VoiceApplicationSid** | **string** | The SID of the application we should use to handle phone calls to the phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | [optional] 
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | [optional] 
**VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | [optional] 
**VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**VoiceReceiveMode** | **string** | The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | [optional] 
**VoiceUrl** | **string** | The URL that we should call to answer a call to the phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


