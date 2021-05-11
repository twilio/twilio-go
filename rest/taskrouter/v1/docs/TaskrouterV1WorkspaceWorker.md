# TaskrouterV1WorkspaceWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActivityName** | Pointer to **string** | The friendly_name of the Worker's current Activity |
**ActivitySid** | Pointer to **string** | The SID of the Worker's current Activity |
**Attributes** | Pointer to **string** | The JSON string that describes the Worker |
**Available** | Pointer to **bool** | Whether the Worker is available to perform tasks |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateStatusChanged** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT of the last change to the Worker's activity |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Worker resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Worker |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


