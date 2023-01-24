# ApiV2010IncomingPhoneNumberAssignedAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that that we created to identify the resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource. |
**ResourceSid** | Pointer to **string** | The SID of the Phone Number to which the Add-on is assigned. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Description** | Pointer to **string** | A short description of the functionality that the Add-on provides. |
**Configuration** | Pointer to **interface{}** | A JSON string that represents the current configuration of this Add-on installation. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


