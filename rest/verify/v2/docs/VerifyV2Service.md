# VerifyV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the verification service. It can be up to 32 characters long. **This value should not contain PII.** |
**CodeLength** | Pointer to **int** | The length of the verification code to generate. |
**LookupEnabled** | Pointer to **bool** | Whether to perform a lookup with each verification started and return info about the phone number. |
**Psd2Enabled** | Pointer to **bool** | Whether to pass PSD2 transaction parameters when starting a verification. |
**SkipSmsToLandlines** | Pointer to **bool** | Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`. |
**DtmfInputRequired** | Pointer to **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call. |
**TtsName** | Pointer to **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. |
**DoNotShareWarningEnabled** | Pointer to **bool** | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: `Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code` |
**CustomCodeEnabled** | Pointer to **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. |
**Push** | Pointer to **interface{}** | Configurations for the Push factors (channel) created under this Service. |
**Totp** | Pointer to **interface{}** | Configurations for the TOTP factors (channel) created under this Service. |
**DefaultTemplateSid** | Pointer to **string** |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


