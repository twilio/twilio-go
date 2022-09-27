# VerifyV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the verification service |
**CodeLength** | Pointer to **int** | The length of the verification code |
**LookupEnabled** | Pointer to **bool** | Whether to perform a lookup with each verification |
**Psd2Enabled** | Pointer to **bool** | Whether to pass PSD2 transaction parameters when starting a verification |
**SkipSmsToLandlines** | Pointer to **bool** | Whether to skip sending SMS verifications to landlines |
**DtmfInputRequired** | Pointer to **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call |
**TtsName** | Pointer to **string** | The name of an alternative text-to-speech service to use in phone calls |
**DoNotShareWarningEnabled** | Pointer to **bool** | Whether to add a security warning at the end of an SMS. |
**CustomCodeEnabled** | Pointer to **bool** | Whether to allow sending verifications with a custom code. |
**Push** | Pointer to **interface{}** | The service level configuration of factor push type. |
**Totp** | Pointer to **interface{}** | The service level configuration of factor TOTP type. |
**DefaultTemplateSid** | Pointer to **string** |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


