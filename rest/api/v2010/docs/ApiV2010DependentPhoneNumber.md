# ApiV2010DependentPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call |
**VoiceMethod** | Pointer to **string** | The HTTP method used with the voice_url |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs in TwiML |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether to lookup the caller's name |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method used with sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |
**SmsMethod** | Pointer to **string** | The HTTP method to use with sms_url |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message |
**AddressRequirements** | Pointer to [**string**](DependentPhoneNumberEnumAddressRequirement.md) |  |
**Capabilities** | Pointer to **interface{}** | Indicate if a phone can receive calls or messages |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number |
**EmergencyStatus** | Pointer to [**string**](DependentPhoneNumberEnumEmergencyStatus.md) |  |
**EmergencyAddressSid** | Pointer to **string** | The emergency address configuration to use for emergency calling |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


