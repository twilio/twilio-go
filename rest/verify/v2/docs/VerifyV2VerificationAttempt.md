# VerifyV2VerificationAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The SID that uniquely identifies the verification attempt. |[optional] 
**AccountSid** | **string** | The SID of the Account that created the verification. |[optional] 
**ServiceSid** | **string** | The SID of the verify service that generated this attempt. |[optional] 
**VerificationSid** | **string** | The SID of the verification that generated this attempt. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date this Attempt was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date this Attempt was updated |[optional] 
**ConversionStatus** | Pointer to [**string**](VerificationAttemptEnumConversionStatus.md) |  |
**Channel** | Pointer to [**string**](VerificationAttemptEnumChannels.md) |  |
**Price** | Pointer to **interface{}** | An object containing the charge for this verification attempt. |
**ChannelData** | Pointer to **interface{}** | An object containing the channel specific information for an attempt. |
**Url** | **string** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


