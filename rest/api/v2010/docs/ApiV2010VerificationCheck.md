# ApiV2010VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number being verified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format. |
**VerificationSid** | Pointer to **string** | The SID that uniquely identifies the verification. |
**Status** | Pointer to **string** | The status of the verification. Can be: `pending`, `approved`, or `failed`. |
**CallerIdSid** | Pointer to **string** | The SID of the OutgoingCallerId resource created when verification is approved. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


