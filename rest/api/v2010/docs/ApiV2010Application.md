# ApiV2010Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource. |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**MessageStatusCallback** | Pointer to **string** | The URL we call using a POST method to send message status information to your application. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Application resource. |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method we use to call `sms_fallback_url`. Can be: `GET` or `POST`. |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML from `sms_url`. |
**SmsMethod** | Pointer to **string** | The HTTP method we use to call `sms_url`. Can be: `GET` or `POST`. |
**SmsStatusCallback** | Pointer to **string** | The URL we call using a POST method to send status information to your application about SMS messages that refer to the application. |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message. |
**StatusCallback** | Pointer to **string** | The URL we call using the `status_callback_method` to send status information to your application. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `status_callback`. Can be: `GET` or `POST`. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether we look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`. |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs retrieving or executing the TwiML requested by `url`. |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number assigned to this application receives a call. |
**PublicApplicationConnectEnabled** | Pointer to **bool** | Whether to allow other Twilio accounts to dial this applicaton using Dial verb. Can be: `true` or `false`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


