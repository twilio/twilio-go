# ApiV2010DependentPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**PhoneNumber** | **string** | The phone number in E.164 format |[optional] 
**VoiceUrl** | **string** | The URL we call when the phone number receives a call |[optional] 
**VoiceMethod** | **string** | The HTTP method used with the voice_url |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method used with voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when an error occurs in TwiML |[optional] 
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller's name |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**SmsFallbackMethod** | **string** | The HTTP method used with sms_fallback_url |[optional] 
**SmsFallbackUrl** | **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |[optional] 
**SmsMethod** | **string** | The HTTP method to use with sms_url |[optional] 
**SmsUrl** | **string** | The URL we call when the phone number receives an incoming SMS message |[optional] 
**AddressRequirements** | Pointer to [**string**](DependentPhoneNumberEnumAddressRequirement.md) |  |
**Capabilities** | Pointer to **interface{}** | Indicate if a phone can receive calls or messages |
**StatusCallback** | **string** | The URL to send status information to your application |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback |[optional] 
**ApiVersion** | **string** | The API version used to start a new TwiML session |[optional] 
**SmsApplicationSid** | **string** | The SID of the application that handles SMS messages sent to the phone number |[optional] 
**VoiceApplicationSid** | **string** | The SID of the application that handles calls to the phone number |[optional] 
**TrunkSid** | **string** | The SID of the Trunk that handles calls to the phone number |[optional] 
**EmergencyStatus** | Pointer to [**string**](DependentPhoneNumberEnumEmergencyStatus.md) |  |
**EmergencyAddressSid** | **string** | The emergency address configuration to use for emergency calling |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


