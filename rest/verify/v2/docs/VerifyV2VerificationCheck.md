# VerifyV2VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. |
**Channel** | Pointer to **string** | The verification method to use |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was last updated |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the verification resource |
**To** | Pointer to **string** | The phone number or email being verified |
**Valid** | Pointer to **bool** | Whether the verification was successful |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


