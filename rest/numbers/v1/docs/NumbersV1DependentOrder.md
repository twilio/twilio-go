# NumbersV1DependentOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentOrder resource. |
**AuthorizationDocumentSid** | Pointer to **string** | The SID of the AuthorizationDocument resource associated with the hosted number order. |
**Capabilities** | Pointer to [**NumbersV1AuthorizationDocumentDependentOrderCapabilities**](NumbersV1AuthorizationDocumentDependentOrderCapabilities.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**FailureReason** | Pointer to **string** | The message that explains why the hosted number order `status` became `action-required`. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**IncomingPhoneNumberSid** | Pointer to **string** | The SID of the IncomingPhoneNumber resource created by the hosted number order. |
**PhoneNumber** | Pointer to **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number hosted by the hosted number order. |
**Sid** | Pointer to **string** | The unique string that we created to identify the DependentOrder resource. |
**Status** | Pointer to [**string**](DependentOrderEnumStatus.md) |  |
**VerificationAttempts** | Pointer to **int** | The number of attempts made to verify ownership via a call for the hosted phone number. |
**VerificationCallDelay** | Pointer to **int** | The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive. |
**VerificationCallExtension** | Pointer to **string** | The numerical extension to dial when making the ownership verification call. |
**VerificationCallSids** | Pointer to **[]string** | The Call SIDs that identify the calls placed to verify ownership. |
**VerificationCode** | Pointer to **string** | The digits the user must pass in the ownership verification call. |
**VerificationDocumentSid** | Pointer to **string** | The SID of the identity document resource that represents the document used to verify ownership of the number to be hosted. |
**VerificationType** | Pointer to [**string**](DependentOrderEnumVerificationType.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


