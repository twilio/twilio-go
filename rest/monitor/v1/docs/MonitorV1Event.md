# MonitorV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ActorSid** | **string** | The SID of the actor that caused the event, if available |[optional] 
**ActorType** | **string** | The type of actor that caused the event |[optional] 
**Description** | **string** | A description of the event |[optional] 
**EventData** | Pointer to **interface{}** | A JSON string that represents an object with additional data about the event |
**EventDate** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the event was recorded |[optional] 
**EventType** | **string** | The event's type |[optional] 
**ResourceSid** | **string** | The SID of the resource that was affected |[optional] 
**ResourceType** | **string** | The type of resource that was affected |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Source** | **string** | The originating system or interface that caused the event |[optional] 
**SourceIpAddress** | **string** | The IP address of the source |[optional] 
**Url** | **string** | The absolute URL of the resource that was affected |[optional] 
**Links** | **map[string]interface{}** | The absolute URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


