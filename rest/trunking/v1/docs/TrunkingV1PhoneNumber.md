# TrunkingV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AddressRequirements** | Pointer to **string** | Whether the phone number requires an Address registered with Twilio |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**Beta** | Pointer to **bool** | Whether the phone number is new to the Twilio platform |
**Capabilities** | Pointer to **map[string]interface{}** | Indicate if a phone can receive calls or messages |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method used with sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |
**SmsMethod** | Pointer to **string** | The HTTP method to use with sms_url |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether to lookup the caller's name |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method that we use to call voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when an error occurs in TwiML |
**VoiceMethod** | Pointer to **string** | The HTTP method used with the voice_url |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


