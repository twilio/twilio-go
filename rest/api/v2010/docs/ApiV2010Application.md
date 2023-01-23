# ApiV2010Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ApiVersion** | **string** | The API version used to start a new TwiML session |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**MessageStatusCallback** | **string** | The URL to send message status information to your application |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SmsFallbackMethod** | **string** | The HTTP method used with sms_fallback_url |[optional] 
**SmsFallbackUrl** | **string** | The URL that we call when an error occurs while retrieving or executing the TwiML |[optional] 
**SmsMethod** | **string** | The HTTP method to use with sms_url |[optional] 
**SmsStatusCallback** | **string** | The URL to send status information to your application |[optional] 
**SmsUrl** | **string** | The URL we call when the phone number receives an incoming SMS message |[optional] 
**StatusCallback** | **string** | The URL to send status information to your application |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller's name |[optional] 
**VoiceFallbackMethod** | **string** | The HTTP method used with voice_fallback_url |[optional] 
**VoiceFallbackUrl** | **string** | The URL we call when a TwiML error occurs |[optional] 
**VoiceMethod** | **string** | The HTTP method used with the voice_url |[optional] 
**VoiceUrl** | **string** | The URL we call when the phone number receives a call |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


