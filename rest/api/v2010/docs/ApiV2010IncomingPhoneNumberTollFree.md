# ApiV2010IncomingPhoneNumberTollFree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AddressSid** | Pointer to **string** | The SID of the Address resource associated with the phone number |
**AddressRequirements** | Pointer to [**string**](IncomingPhoneNumberTollFreeEnumAddressRequirement.md) |  |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**Beta** | Pointer to **bool** | Whether the phone number is new to the Twilio platform |
**Capabilities** | Pointer to [**ApiV2010AccountIncomingPhoneNumberCapabilities**](ApiV2010AccountIncomingPhoneNumberCapabilities.md) |  |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**IdentitySid** | Pointer to **string** | The SID of the Identity resource associated with number |
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**Origin** | Pointer to **string** | The phone number's origin. Can be twilio or hosted. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method used with sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |
**SmsMethod** | Pointer to **string** | The HTTP method to use with sms_url |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**VoiceReceiveMode** | Pointer to [**string**](IncomingPhoneNumberTollFreeEnumVoiceReceiveMode.md) |  |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether to lookup the caller's name |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs in TwiML |
**VoiceMethod** | Pointer to **string** | The HTTP method used with the voice_url |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call |
**EmergencyStatus** | Pointer to [**string**](IncomingPhoneNumberTollFreeEnumEmergencyStatus.md) |  |
**EmergencyAddressSid** | Pointer to **string** | The emergency address configuration to use for emergency calling |
**EmergencyAddressStatus** | Pointer to [**string**](IncomingPhoneNumberTollFreeEnumEmergencyAddressStatus.md) |  |
**BundleSid** | Pointer to **string** | The SID of the Bundle resource associated with number |
**Status** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


