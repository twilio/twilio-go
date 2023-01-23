# TaskrouterV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Event resource. |
**ActorSid** | Pointer to **string** | The SID of the resource that triggered the event. |
**ActorType** | Pointer to **string** | The type of resource that triggered the event. |
**ActorUrl** | Pointer to **string** | The absolute URL of the resource that triggered the event. |
**Description** | Pointer to **string** | A description of the event. |
**EventData** | Pointer to **interface{}** | Data about the event. For more information, see [Event types](https://www.twilio.com/docs/taskrouter/api/event#event-types). |
**EventDate** | Pointer to [**time.Time**](time.Time.md) | The time the event was sent, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EventDateMs** | Pointer to **int64** | The time the event was sent in milliseconds. |
**EventType** | Pointer to **string** | The identifier for the event. |
**ResourceSid** | Pointer to **string** | The SID of the object the event is most relevant to, such as a TaskSid, ReservationSid, or a  WorkerSid. |
**ResourceType** | Pointer to **string** | The type of object the event is most relevant to, such as a Task, Reservation, or a Worker). |
**ResourceUrl** | Pointer to **string** | The URL of the resource the event is most relevant to. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Event resource. |
**Source** | Pointer to **string** | Where the Event originated. |
**SourceIpAddress** | Pointer to **string** | The IP from which the Event originated. |
**Url** | Pointer to **string** | The absolute URL of the Event resource. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Event. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


