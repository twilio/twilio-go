# NumbersV1AuthorizationDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressSid** | Pointer to **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the AuthorizationDocument resource. |
**CcEmails** | Pointer to **[]string** | The email addresses that the authorization document will be copied to. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Email** | Pointer to **string** | The email address that the authorization document will be sent to for signing. |
**Links** | Pointer to **map[string]interface{}** | A list of the URLs to the HostedNumberOrder resources that the authorization document will authorize for hosting phone number capabilities on our platform. |
**Sid** | Pointer to **string** | The unique string that we created to identify the AuthorizationDocument resource. |
**Status** | Pointer to [**string**](AuthorizationDocumentEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the AuthorizationDocument resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


