# VerifyV2VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**To** | Pointer to **string** | The phone number or email being verified |
**Channel** | Pointer to [**string**](VerificationCheckEnumChannel.md) |  |
**Status** | Pointer to **string** | The status of the verification resource |
**Valid** | Pointer to **bool** | Whether the verification was successful |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was last updated |
**SnaAttemptsErrorCodes** | Pointer to **[]interface{}** | List of error codes as a result of attempting a verification using the `sna` channel. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


