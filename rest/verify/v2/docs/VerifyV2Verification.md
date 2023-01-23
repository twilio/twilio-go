# VerifyV2Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Verification resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Verification resource. |
**To** | Pointer to **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) being verified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). |
**Channel** | Pointer to [**string**](VerificationEnumChannel.md) |  |
**Status** | Pointer to **string** | The status of the verification. One of: `pending`, `approved`, or `canceled` |
**Valid** | Pointer to **bool** | Use \"status\" instead. Legacy property indicating whether the verification was successful. |
**Lookup** | Pointer to **interface{}** | Information about the phone number being verified. |
**Amount** | Pointer to **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. |
**Payee** | Pointer to **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. |
**SendCodeAttempts** | Pointer to **[]interface{}** | An array of verification attempt objects containing the channel attempted and the channel-specific transaction SID. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Sna** | Pointer to **interface{}** | The set of fields used for a silent network auth (`sna`) verification. Contains a single field with the URL to be invoked to verify the phone number. |
**Url** | Pointer to **string** | The absolute URL of the Verification resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


