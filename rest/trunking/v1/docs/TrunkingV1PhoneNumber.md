# TrunkingV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AddressRequirements** | Pointer to [**string**](PhoneNumberEnumAddressRequirement.md) |  |
**ApiVersion** | **string** | The API version used to start a new TwiML session |[optional] 
**Beta** | **bool** | Whether the phone number is new to the Twilio platform |[optional] 
**Capabilities** | **map[string]interface{}** | Indicate if a phone can receive calls or messages |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 
**PhoneNumber** | **string** | The phone number in E.164 format |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SmsApplicationSid** | **string** | The SID of the application that handles SMS messages sent to the phone number |[optional] 
**SmsFallbackMethod** | **string** | The HTTP method used with sms_fallback_url |[optional] 
**SmsFallbackUrl** | **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |[optional] 
**SmsMethod** | **string** | The HTTP method to use with sms_url |[optional] 
**SmsUrl** | **string** | The URL we call when the phone number receives an incoming SMS message |[optional] 
**StatusCallback** | **string** | The URL to send status information to your application |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback |[optional] 
**TrunkSid** | **string** | The SID of the Trunk that handles calls to the phone number |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**VoiceApplicationSid** | **string** | The SID of the application that handles calls to the phone number |[optional] 
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller's name |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method that we use to call voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when an error occurs in TwiML |[optional] 
**VoiceMethod** | **string** | The HTTP method used with the voice_url |[optional] 
**VoiceUrl** | **string** | The URL we call when the phone number receives a call |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


