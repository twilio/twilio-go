# VerifyV2Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**To** | Pointer to **string** | The phone number or email being verified |
**Channel** | Pointer to [**string**](VerificationEnumChannel.md) |  |
**Status** | Pointer to **string** | The status of the verification resource |
**Valid** | Pointer to **bool** | Whether the verification was successful |
**Lookup** | Pointer to **interface{}** | Information about the phone number being verified |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction |
**SendCodeAttempts** | Pointer to **[]interface{}** | An array of verification attempt objects. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Sna** | Pointer to **interface{}** | The set of fields used for a silent network auth (`sna`) verification |
**Url** | Pointer to **string** | The absolute URL of the Verification resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


