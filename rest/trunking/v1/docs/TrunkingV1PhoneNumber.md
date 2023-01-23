# TrunkingV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the PhoneNumber resource. |
**AddressRequirements** | Pointer to [**string**](PhoneNumberEnumAddressRequirement.md) |  |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session. |
**Beta** | Pointer to **bool** | Whether the phone number is new to the Twilio platform. Can be: `true` or `false`. |
**Capabilities** | Pointer to **map[string]interface{}** | The set of Boolean properties that indicate whether a phone number can receive calls or messages.  Capabilities are  `Voice`, `SMS`, and `MMS` and each capability can be: `true` or `false`. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**Sid** | Pointer to **string** | The unique string that we created to identify the PhoneNumber resource. |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application. |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method we use to call `sms_fallback_url`. Can be: `GET` or `POST`. |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML from `sms_url`. |
**SmsMethod** | Pointer to **string** | The HTTP method we use to call `sms_url`. Can be: `GET` or `POST`. |
**SmsUrl** | Pointer to **string** | The URL we call using the `sms_method` when the phone number receives an incoming SMS message. |
**StatusCallback** | Pointer to **string** | The URL we call using the `status_callback_method` to send status information to your application. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `status_callback`. Can be: `GET` or `POST`. |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice URLs and voice applications and use those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice URLs and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether we look up the caller's caller-ID name from the CNAM database ($0.01 per look up). Can be: `true` or `false`. |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method that we use to call `voice_fallback_url`. Can be: `GET` or `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | The URL that we call using the `voice_fallback_method` when an error occurs retrieving or executing the TwiML requested by `url`. |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. |
**VoiceUrl** | Pointer to **string** | The URL we call using the `voice_method` when the phone number receives a call. The `voice_url` is not be used if a `voice_application_sid` or a `trunk_sid` is set. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


