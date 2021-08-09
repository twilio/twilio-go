# VerifyV2Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. |
**Channel** | Pointer to **string** | The verification method used. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Lookup** | Pointer to **map[string]interface{}** | Information about the phone number being verified |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction |
**SendCodeAttempts** | Pointer to **[]map[string]interface{}** | An array of verification attempt objects. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the verification resource |
**To** | Pointer to **string** | The phone number or email being verified |
**Url** | Pointer to **string** | The absolute URL of the Verification resource |
**Valid** | Pointer to **bool** | Whether the verification was successful |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


