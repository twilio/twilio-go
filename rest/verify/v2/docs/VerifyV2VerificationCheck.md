# VerifyV2VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**To** | **string** | The phone number or email being verified |[optional] 
**Channel** | Pointer to [**string**](VerificationCheckEnumChannel.md) |  |
**Status** | **string** | The status of the verification resource |[optional] 
**Valid** | **bool** | Whether the verification was successful |[optional] 
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. |[optional] 
**Payee** | **string** | The payee of the associated PSD2 compliant transaction |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Verification Check resource was last updated |[optional] 
**SnaAttemptsErrorCodes** | **[]interface{}** | List of error codes as a result of attempting a verification using the `sna` channel. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


