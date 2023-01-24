# VerifyV2VerificationAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID that uniquely identifies the verification attempt resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Verification resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) used to generate the attempt. |
**VerificationSid** | Pointer to **string** | The SID of the [Verification](https://www.twilio.com/docs/verify/api/verification) that generated the attempt. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Attempt was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Attempt was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ConversionStatus** | Pointer to [**string**](VerificationAttemptEnumConversionStatus.md) |  |
**Channel** | Pointer to [**string**](VerificationAttemptEnumChannels.md) |  |
**Price** | Pointer to **interface{}** | An object containing the charge for this verification attempt related to the channel costs and the currency used. The costs related to the succeeded verifications are not included. May not be immediately available. More information on pricing is available [here](https://www.twilio.com/verify/pricing). |
**ChannelData** | Pointer to **interface{}** | An object containing the channel specific information for an attempt. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


