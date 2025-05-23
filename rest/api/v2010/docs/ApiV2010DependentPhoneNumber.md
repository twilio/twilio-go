# ApiV2010DependentPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that that we created to identify the DependentPhoneNumber resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**VoiceUrl** | Pointer to **string** | The URL we call when the phone number receives a call. The `voice_url` will not be used if a `voice_application_sid` or a `trunk_sid` is set. |
**VoiceMethod** | Pointer to **string** | The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`. |
**VoiceFallbackMethod** | Pointer to **string** | The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`. |
**VoiceFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs retrieving or executing the TwiML requested by `url`. |
**VoiceCallerIdLookup** | Pointer to **bool** | Whether we look up the caller's caller-ID name from the CNAM database. Can be: `true` or `false`. Caller ID lookups can cost $0.01 each. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method we use to call `sms_fallback_url`. Can be: `GET` or `POST`. |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call when an error occurs while retrieving or executing the TwiML from `sms_url`. |
**SmsMethod** | Pointer to **string** | The HTTP method we use to call `sms_url`. Can be: `GET` or `POST`. |
**SmsUrl** | Pointer to **string** | The URL we call when the phone number receives an incoming SMS message. |
**AddressRequirements** | Pointer to [**string**](DependentPhoneNumberEnumAddressRequirement.md) |  |
**Capabilities** | Pointer to [**ApiV2010AccountAddressDependentPhoneNumberCapabilities**](ApiV2010AccountAddressDependentPhoneNumberCapabilities.md) |  |
**StatusCallback** | Pointer to **string** | The URL we call using the `status_callback_method` to send status information to your application. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `status_callback`. Can be: `GET` or `POST`. |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session. |
**SmsApplicationSid** | Pointer to **string** | The SID of the application that handles SMS messages sent to the phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application. |
**VoiceApplicationSid** | Pointer to **string** | The SID of the application that handles calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that handles calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. |
**EmergencyStatus** | Pointer to [**string**](DependentPhoneNumberEnumEmergencyStatus.md) |  |
**EmergencyAddressSid** | Pointer to **string** | The SID of the emergency address configuration that we use for emergency calling from the phone number. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


