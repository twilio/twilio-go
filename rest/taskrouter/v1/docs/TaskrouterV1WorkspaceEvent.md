# TaskrouterV1WorkspaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActorSid** | Pointer to **string** | The SID of the resource that triggered the event |
**ActorType** | Pointer to **string** | The type of resource that triggered the event |
**ActorUrl** | Pointer to **string** | The absolute URL of the resource that triggered the event |
**Description** | Pointer to **string** | A description of the event |
**EventData** | Pointer to **map[string]interface{}** | Data about the event |
**EventDate** | Pointer to [**time.Time**](time.Time.md) | The time the event was sent |
**EventDateMs** | Pointer to **int32** | The time the event was sent in milliseconds |
**EventType** | Pointer to **string** | The identifier for the event |
**ResourceSid** | Pointer to **string** | The SID of the object the event is most relevant to |
**ResourceType** | Pointer to **string** | The type of object the event is most relevant to |
**ResourceUrl** | Pointer to **string** | The URL of the resource the event is most relevant to |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Source** | Pointer to **string** | Where the Event originated |
**SourceIpAddress** | Pointer to **string** | The IP from which the Event originated |
**Url** | Pointer to **string** | The absolute URL of the Event resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Event |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


