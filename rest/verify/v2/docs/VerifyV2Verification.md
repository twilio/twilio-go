# VerifyV2Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**To** | **string** | The phone number or email being verified |[optional] 
**Channel** | Pointer to [**string**](VerificationEnumChannel.md) |  |
**Status** | **string** | The status of the verification resource |[optional] 
**Valid** | **bool** | Whether the verification was successful |[optional] 
**Lookup** | Pointer to **interface{}** | Information about the phone number being verified |
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. |[optional] 
**Payee** | **string** | The payee of the associated PSD2 compliant transaction |[optional] 
**SendCodeAttempts** | **[]interface{}** | An array of verification attempt objects. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Sna** | Pointer to **interface{}** | The set of fields used for a silent network auth (`sna`) verification |
**Url** | **string** | The absolute URL of the Verification resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


