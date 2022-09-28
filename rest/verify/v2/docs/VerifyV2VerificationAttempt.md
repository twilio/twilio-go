# VerifyV2VerificationAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID that uniquely identifies the verification attempt. |
**AccountSid** | Pointer to **string** | The SID of the Account that created the verification. |
**ServiceSid** | Pointer to **string** | The SID of the verify service that generated this attempt. |
**VerificationSid** | Pointer to **string** | The SID of the verification that generated this attempt. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this Attempt was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date this Attempt was updated |
**ConversionStatus** | Pointer to [**string**](VerificationAttemptEnumConversionStatus.md) |  |
**Channel** | Pointer to [**string**](VerificationAttemptEnumChannels.md) |  |
**Price** | Pointer to **interface{}** | An object containing the charge for this verification attempt. |
**ChannelData** | Pointer to **interface{}** | An object containing the channel specific information for an attempt. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


