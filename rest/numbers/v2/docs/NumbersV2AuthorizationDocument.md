# NumbersV2AuthorizationDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this AuthorizationDocument. |
**AddressSid** | Pointer to **string** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. |
**Status** | Pointer to [**string**](AuthorizationDocumentEnumStatus.md) |  |
**Email** | Pointer to **string** | Email that this AuthorizationDocument will be sent to for signing. |
**CcEmails** | Pointer to **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was updated, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** |  |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


