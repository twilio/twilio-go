# ApiV2010SmsVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number being verified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format. |
**VerificationSid** | Pointer to **string** | The SID that uniquely identifies the verification. |
**SendCodeAttempts** | Pointer to [**[]ApiV2010AccountOutgoingCallerIdSmsVerificationSendCodeAttempts**](ApiV2010AccountOutgoingCallerIdSmsVerificationSendCodeAttempts.md) | An array of verification attempt objects containing the channel attempted and the channel-specific transaction SID. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


