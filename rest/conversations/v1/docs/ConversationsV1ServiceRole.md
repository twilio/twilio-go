# ConversationsV1ServiceRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Type** | Pointer to [**string**](ServiceRoleEnumRoleType.md) |  |
**Permissions** | Pointer to **[]string** | An array of the permissions the role has been granted |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | An absolute URL for this user role. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


