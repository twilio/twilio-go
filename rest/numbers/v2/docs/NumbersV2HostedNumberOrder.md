# NumbersV2HostedNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this HostedNumberOrder. |
**AccountSid** | Pointer to **string** | A 34 character string that uniquely identifies the account. |
**IncomingPhoneNumberSid** | Pointer to **string** | A 34 character string that uniquely identifies the [IncomingPhoneNumber](https://www.twilio.com/docs/api/rest/incoming-phone-numbers) resource that represents the phone number being hosted. |
**AddressSid** | Pointer to **string** | A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. |
**SigningDocumentSid** | Pointer to **string** | A 34 character string that uniquely identifies the [Authorization Document](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents) the user needs to sign. |
**PhoneNumber** | Pointer to **string** | Phone number to be hosted. This must be in [E.164](https://en.wikipedia.org/wiki/E.164) format, e.g., +16175551212 |
**Capabilities** | Pointer to [**NumbersV2HostedNumberOrderCapabilities**](NumbersV2HostedNumberOrderCapabilities.md) |  |
**FriendlyName** | Pointer to **string** | A 64 character string that is a human-readable text that describes this resource. |
**Status** | Pointer to [**string**](HostedNumberOrderEnumStatus.md) |  |
**FailureReason** | Pointer to **string** | A message that explains why a hosted_number_order went to status \"action-required\" |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**SmsCapability** | Pointer to **bool** | Whether the SMS capability will be hosted on our platform. Can be `true` of `false`. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was updated, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Email** | Pointer to **string** | Email of the owner of this phone number that is being hosted. |
**CcEmails** | Pointer to **[]string** | A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. |
**Url** | Pointer to **string** | The URL of this HostedNumberOrder. |
**ContactTitle** | Pointer to **string** | The title of the person authorized to sign the Authorization Document for this phone number. |
**ContactPhoneNumber** | Pointer to **string** | The contact phone number of the person authorized to sign the Authorization Document. |
**BulkHostingRequestSid** | Pointer to **string** | A 34 character string that uniquely identifies the bulk hosting request associated with this HostedNumberOrder. |
**NextStep** | Pointer to **string** | The next step you need to take to complete the hosted number order and request it successfully. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


