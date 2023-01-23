# TaskrouterV1Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**DefaultActivityName** | **string** | The name of the default activity |[optional] 
**DefaultActivitySid** | **string** | The SID of the Activity that will be used when new Workers are created in the Workspace |[optional] 
**EventCallbackUrl** | **string** | The URL we call when an event occurs |[optional] 
**EventsFilter** | **string** | The list of Workspace events for which to call event_callback_url |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the Workspace resource |[optional] 
**MultiTaskEnabled** | **bool** | Whether multi-tasking is enabled |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TimeoutActivityName** | **string** | The name of the timeout activity |[optional] 
**TimeoutActivitySid** | **string** | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response |[optional] 
**PrioritizeQueueOrder** | Pointer to [**string**](WorkspaceEnumQueueOrder.md) |  |
**Url** | **string** | The absolute URL of the Workspace resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


