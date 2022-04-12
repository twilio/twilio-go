# VerifyV2VerificationAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the verification. |
**Channel** | Pointer to **string** | Communication channel used for the attempt. |
**ChannelData** | Pointer to **interface{}** | An object containing the channel specific information for an attempt. |
**ConversionStatus** | Pointer to **string** | Status of the conversion for the verification. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this Attempt was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date this Attempt was updated |
**Price** | Pointer to **interface{}** | An object containing the charge for this verification attempt. |
**ServiceSid** | Pointer to **string** | The SID of the verify service that generated this attempt. |
**Sid** | Pointer to **string** | The SID that uniquely identifies the verification attempt. |
**Url** | Pointer to **string** |  |
**VerificationSid** | Pointer to **string** | The SID of the verification that generated this attempt. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


