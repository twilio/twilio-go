# MonitorV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActorSid** | Pointer to **string** | The SID of the actor that caused the event, if available |
**ActorType** | Pointer to **string** | The type of actor that caused the event |
**Description** | Pointer to **string** | A description of the event |
**EventData** | Pointer to **interface{}** | A JSON string that represents an object with additional data about the event |
**EventDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the event was recorded |
**EventType** | Pointer to **string** | The event's type |
**ResourceSid** | Pointer to **string** | The SID of the resource that was affected |
**ResourceType** | Pointer to **string** | The type of resource that was affected |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Source** | Pointer to **string** | The originating system or interface that caused the event |
**SourceIpAddress** | Pointer to **string** | The IP address of the source |
**Url** | Pointer to **string** | The absolute URL of the resource that was affected |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


