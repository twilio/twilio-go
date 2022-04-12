# VerifyV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CodeLength** | Pointer to **int** | The length of the verification code |
**CustomCodeEnabled** | Pointer to **bool** | Whether to allow sending verifications with a custom code. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**DefaultTemplateSid** | Pointer to **string** |  |
**DoNotShareWarningEnabled** | Pointer to **bool** | Whether to add a security warning at the end of an SMS. |
**DtmfInputRequired** | Pointer to **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the verification service |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**LookupEnabled** | Pointer to **bool** | Whether to perform a lookup with each verification |
**Psd2Enabled** | Pointer to **bool** | Whether to pass PSD2 transaction parameters when starting a verification |
**Push** | Pointer to **interface{}** | The service level configuration of factor push type. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SkipSmsToLandlines** | Pointer to **bool** | Whether to skip sending SMS verifications to landlines |
**Totp** | Pointer to **interface{}** | The service level configuration of factor TOTP type. |
**TtsName** | Pointer to **string** | The name of an alternative text-to-speech service to use in phone calls |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


