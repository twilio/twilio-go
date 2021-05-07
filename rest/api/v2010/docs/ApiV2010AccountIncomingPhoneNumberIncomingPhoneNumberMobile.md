# ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AddressRequirements** | Pointer to **string** | Whether the phone number requires an Address registered with Twilio. |
**AddressSid** | Pointer to **string** | The SID of the Address resource associated with the phone number |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**Beta** | Pointer to **bool** | Whether the phone number is new to the Twilio platform |
**BundleSid** | Pointer to **string** | The SID of the Bundle resource associated with number |
**Capabilities** | Pointer to [**ApiV2010AccountIncomingPhoneNumberCapabilities**](api_v2010_account_incoming_phone_number_capabilities.md) |  |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**EmergencyAddressSid** | Pointer to **string** | The emergency address configuration to use for emergency calling |
**EmergencyStatus** | Pointer to **string** | Whether the phone number is enabled for emergency calling |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**IdentitySid** | Pointer to **string** | The SID of the Identity resource associated with number |
**Origin** | Pointer to **string** | The phone number's origin. Can be twilio or hosted. |
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method used with sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |
**SmsMethod** | Pointer to **string** | The HTTP method to use with sms_url |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message |
**Status** | Pointer to **string** |  |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether to lookup the caller's name |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs in TwiML |
**VoiceMethod** | Pointer to **string** | The HTTP method used with the voice_url |
**VoiceReceiveMode** | Pointer to **string** |  |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


