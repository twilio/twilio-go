# NumbersV2DependentHostedNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Authorization Document |
**BulkHostingRequestSid** | Pointer to **string** | A 34 character string that uniquely identifies the bulk hosting request associated with this HostedNumberOrder. |
**NextStep** | Pointer to **string** | The next step you need to take to complete the hosted number order and request it successfully. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**IncomingPhoneNumberSid** | Pointer to **string** | A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. |
**AddressSid** | Pointer to **string** | A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. |
**SigningDocumentSid** | Pointer to **string** | A 34 character string that uniquely identifies the LOA document associated with this HostedNumberOrder. |
**PhoneNumber** | Pointer to **string** | An E164 formatted phone number hosted by this HostedNumberOrder. |
**Capabilities** | Pointer to [**NumbersV2AuthorizationDocumentDependentHostedNumberOrderCapabilities**](NumbersV2AuthorizationDocumentDependentHostedNumberOrderCapabilities.md) |  |
**FriendlyName** | Pointer to **string** | A human readable description of this resource, up to 128 characters. |
**Status** | Pointer to [**string**](DependentHostedNumberOrderEnumStatus.md) |  |
**FailureReason** | Pointer to **string** | A message that explains why a hosted_number_order went to status \"action-required\" |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was updated, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Email** | Pointer to **string** | Email of the owner of this phone number that is being hosted. |
**CcEmails** | Pointer to **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed |
**ContactTitle** | Pointer to **string** | The title of the person authorized to sign the Authorization Document for this phone number. |
**ContactPhoneNumber** | Pointer to **string** | The contact phone number of the person authorized to sign the Authorization Document. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


