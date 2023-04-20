# VerifyV2VerificationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies a Verification Template. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**FriendlyName** | Pointer to **string** | A descriptive string that you create to describe a Template. It can be up to 32 characters long. |
**Channels** | Pointer to **[]string** | A list of channels that support the Template. Can include: sms, voice. |
**Translations** | Pointer to **interface{}** | An object that contains the different translations of the template. Every translation is identified by the language short name and contains its respective information as the approval status, text and created/modified date. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


