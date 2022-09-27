# TaskrouterV1TaskChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Task Channel |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Task Channel |
**ChannelOptimizedRouting** | Pointer to **bool** | Whether the Task Channel will prioritize Workers that have been idle |
**Url** | Pointer to **string** | The absolute URL of the Task Channel resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


