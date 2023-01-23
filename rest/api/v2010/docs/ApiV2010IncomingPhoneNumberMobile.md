# ApiV2010IncomingPhoneNumberMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AddressSid** | **string** | The SID of the Address resource associated with the phone number |[optional] 
**AddressRequirements** | Pointer to [**string**](IncomingPhoneNumberMobileEnumAddressRequirement.md) |  |
**ApiVersion** | **string** | The API version used to start a new TwiML session |[optional] 
**Beta** | **bool** | Whether the phone number is new to the Twilio platform |[optional] 
**Capabilities** | [**ApiV2010AccountIncomingPhoneNumberCapabilities**](ApiV2010AccountIncomingPhoneNumberCapabilities.md) |  |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**IdentitySid** | **string** | The SID of the Identity resource associated with number |[optional] 
**PhoneNumber** | **string** | The phone number in E.164 format |[optional] 
**Origin** | **string** | The phone number's origin. Can be twilio or hosted. |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SmsApplicationSid** | **string** | The SID of the application that handles SMS messages sent to the phone number |[optional] 
**SmsFallbackMethod** | **string** | The HTTP method used with sms_fallback_url |[optional] 
**SmsFallbackUrl** | **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |[optional] 
**SmsMethod** | **string** | The HTTP method to use with sms_url |[optional] 
**SmsUrl** | **string** | The URL we call when the phone number receives an incoming SMS message |[optional] 
**StatusCallback** | **string** | The URL to send status information to your application |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback |[optional] 
**TrunkSid** | **string** | The SID of the Trunk that handles calls to the phone number |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**VoiceReceiveMode** | Pointer to [**string**](IncomingPhoneNumberMobileEnumVoiceReceiveMode.md) |  |
**VoiceApplicationSid** | **string** | The SID of the application that handles calls to the phone number |[optional] 
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller's name |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method used with voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when an error occurs in TwiML |[optional] 
**VoiceMethod** | **string** | The HTTP method used with the voice_url |[optional] 
**VoiceUrl** | **string** | The URL we call when the phone number receives a call |[optional] 
**EmergencyStatus** | Pointer to [**string**](IncomingPhoneNumberMobileEnumEmergencyStatus.md) |  |
**EmergencyAddressSid** | **string** | The emergency address configuration to use for emergency calling |[optional] 
**EmergencyAddressStatus** | Pointer to [**string**](IncomingPhoneNumberMobileEnumEmergencyAddressStatus.md) |  |
**BundleSid** | **string** | The SID of the Bundle resource associated with number |[optional] 
**Status** | **string** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


