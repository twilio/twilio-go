# VerifyV2VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the VerificationCheck resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the VerificationCheck resource. |
**To** | Pointer to **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) being verified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). |
**Channel** | Pointer to [**string**](VerificationCheckEnumChannel.md) |  |
**Status** | Pointer to **string** | The status of the verification. Can be: `pending`, `approved`, or `canceled`. |
**Valid** | Pointer to **bool** | Use \"status\" instead. Legacy property indicating whether the verification was successful. |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the Verification Check resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the Verification Check resource was last updated. |
**SnaAttemptsErrorCodes** | Pointer to **[]interface{}** | List of error codes as a result of attempting a verification using the `sna` channel. The error codes are chronologically ordered, from the first attempt to the latest attempt. This will be an empty list if no errors occured or `null` if the last channel used wasn't `sna`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


