# TaskrouterV1Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**DefaultActivityName** | Pointer to **string** | The name of the default activity |
**DefaultActivitySid** | Pointer to **string** | The SID of the Activity that will be used when new Workers are created in the Workspace |
**EventCallbackUrl** | Pointer to **string** | The URL we call when an event occurs |
**EventsFilter** | Pointer to **string** | The list of Workspace events for which to call event_callback_url |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Workspace resource |
**MultiTaskEnabled** | Pointer to **bool** | Whether multi-tasking is enabled |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TimeoutActivityName** | Pointer to **string** | The name of the timeout activity |
**TimeoutActivitySid** | Pointer to **string** | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response |
**PrioritizeQueueOrder** | Pointer to [**string**](WorkspaceEnumQueueOrder.md) |  |
**Url** | Pointer to **string** | The absolute URL of the Workspace resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


