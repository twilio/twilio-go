# CreateVerificationCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | [optional] 
**Code** | **string** | The 4-10 character string being verified. | 
**Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | [optional] 
**To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the &#x60;verification_sid&#x60; must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**VerificationSid** | **string** | A SID that uniquely identifies the Verification Check. Either this parameter or the &#x60;to&#x60; phone number/[email](https://www.twilio.com/docs/verify/email) must be specified. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


