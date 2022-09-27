# ApiV2010IncomingPhoneNumberAssignedAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ResourceSid** | Pointer to **string** | The SID of the Phone Number that installed this Add-on |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Description** | Pointer to **string** | A short description of the Add-on functionality |
**Configuration** | Pointer to **interface{}** | A JSON string that represents the current configuration |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


