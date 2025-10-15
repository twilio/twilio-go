# NumbersV3HostedNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this HostedNumberOrder. |
**AccountSid** | Pointer to **string** | A 34 character string that uniquely identifies the account. |
**IncomingPhoneNumberSid** | Pointer to **string** | A 34 character string that uniquely identifies the [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the phone number being hosted. |
**AddressSid** | Pointer to **string** | A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. |
**SigningDocumentSid** | Pointer to **string** | A 34 character string that uniquely identifies the [Authorization Document](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource) the user needs to sign. |
**PhoneNumber** | Pointer to **string** | Phone number to be hosted. This must be in [E.164](https://en.wikipedia.org/wiki/E.164) format, e.g., +16175551212 |
**Capabilities** | Pointer to [**NumbersV3HostedNumbersHostedNumberOrderCapabilities**](NumbersV3HostedNumbersHostedNumberOrderCapabilities.md) |  |
**FriendlyName** | Pointer to **string** | A 64 character string that is a human-readable text that describes this resource. |
**UniqueName** | Pointer to **string** | Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. |
**Status** | Pointer to [**string**](HostedNumberOrderEnumStatus.md) |  |
**FailureReason** | Pointer to **string** | A message that explains why a hosted_number_order went to status \"action-required\" |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was updated, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**VerificationAttempts** | **int** | The number of attempts made to verify ownership of the phone number that is being hosted. |[optional] [default to 0]
**Email** | Pointer to **string** | Email of the owner of this phone number that is being hosted. |
**CcEmails** | Pointer to **[]string** | A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. |
**Url** | Pointer to **string** | The URL of this HostedNumberOrder. |
**VerificationType** | Pointer to [**string**](HostedNumberOrderEnumVerificationType.md) |  |
**VerificationDocumentSid** | Pointer to **string** | A 34 character string that uniquely identifies the Identity Document resource that represents the document for verifying ownership of the number to be hosted. |
**Extension** | Pointer to **string** | A numerical extension to be used when making the ownership verification call. |
**CallDelay** | **int** | A value between 0-30 specifying the number of seconds to delay initiating the ownership verification call. |[optional] [default to 0]
**VerificationCode** | Pointer to **string** | A verification code provided in the response for a user to enter when they pick up the phone call. |
**VerificationCallSids** | Pointer to **[]string** | A list of 34 character strings that are unique identifiers for the calls placed as part of ownership verification. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


