# ProxyV1Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateDeleted** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date the Participant was removed |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the participant |
**Identifier** | Pointer to **string** | The phone number or channel identifier of the Participant |
**Links** | Pointer to **map[string]interface{}** | The URLs to resources related the participant |
**ProxyIdentifier** | Pointer to **string** | The phone number or short code of the participant's partner |
**ProxyIdentifierSid** | Pointer to **string** | The SID of the Proxy Identifier assigned to the Participant |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**SessionSid** | Pointer to **string** | The SID of the resource's parent Session |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Participant resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


