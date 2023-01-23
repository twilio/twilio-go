# VerifyV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the verification service |[optional] 
**CodeLength** | **int** | The length of the verification code |[optional] 
**LookupEnabled** | **bool** | Whether to perform a lookup with each verification |[optional] 
**Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification |[optional] 
**SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines |[optional] 
**DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call |[optional] 
**TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls |[optional] 
**DoNotShareWarningEnabled** | **bool** | Whether to add a security warning at the end of an SMS. |[optional] 
**CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code. |[optional] 
**Push** | Pointer to **interface{}** | The service level configuration of factor push type. |
**Totp** | Pointer to **interface{}** | The service level configuration of factor TOTP type. |
**DefaultTemplateSid** | **string** |  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


