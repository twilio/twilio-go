# TaskrouterV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ActorSid** | **string** | The SID of the resource that triggered the event |[optional] 
**ActorType** | **string** | The type of resource that triggered the event |[optional] 
**ActorUrl** | **string** | The absolute URL of the resource that triggered the event |[optional] 
**Description** | **string** | A description of the event |[optional] 
**EventData** | Pointer to **interface{}** | Data about the event |
**EventDate** | [**time.Time**](time.Time.md) | The time the event was sent |[optional] 
**EventDateMs** | **int64** | The time the event was sent in milliseconds |[optional] 
**EventType** | **string** | The identifier for the event |[optional] 
**ResourceSid** | **string** | The SID of the object the event is most relevant to |[optional] 
**ResourceType** | **string** | The type of object the event is most relevant to |[optional] 
**ResourceUrl** | **string** | The URL of the resource the event is most relevant to |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Source** | **string** | Where the Event originated |[optional] 
**SourceIpAddress** | **string** | The IP from which the Event originated |[optional] 
**Url** | **string** | The absolute URL of the Event resource |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Event |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


