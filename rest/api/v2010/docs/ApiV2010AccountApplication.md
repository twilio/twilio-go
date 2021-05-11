# ApiV2010AccountApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**MessageStatusCallback** | Pointer to **string** | The URL to send message status information to your application |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method used with sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |
**SmsMethod** | Pointer to **string** | The HTTP method to use with sms_url |
**SmsStatusCallback** | Pointer to **string** | The URL to send status information to your application |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether to lookup the caller's name |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method used with voice_fallback_url |
**VoiceFallbackUrl** | Pointer to **string** | The URL we call when a TwiML error occurs |
**VoiceMethod** | Pointer to **string** | The HTTP method used with the voice_url |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


