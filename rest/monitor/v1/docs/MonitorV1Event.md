# MonitorV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Event resource. |
**ActorSid** | Pointer to **string** | The SID of the actor that caused the event, if available. Can be `null`. |
**ActorType** | Pointer to **string** | The type of actor that caused the event. Can be: `user` for a change made by a logged-in user in the Twilio Console, `account` for an event caused by an API request by an authenticating Account, `twilio-admin` for an event caused by a Twilio employee, and so on. |
**Description** | Pointer to **string** | A description of the event. Can be `null`. |
**EventData** | Pointer to **interface{}** | An object with additional data about the event. The  contents depend on `event_type`. For example, event-types of the form `RESOURCE.updated`, this value contains a `resource_properties` dictionary that describes the previous and updated properties of the resource. |
**EventDate** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the event was recorded specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EventType** | Pointer to **string** | The event's type. Event-types are typically in the form: `RESOURCE_TYPE.ACTION`, where `RESOURCE_TYPE` is the type of resource that was affected and `ACTION` is what happened to it. For example, `phone-number.created`. For a full list of all event-types, see the [Monitor Event Types](https://www.twilio.com/docs/usage/monitor-events#event-types). |
**ResourceSid** | Pointer to **string** | The SID of the resource that was affected. |
**ResourceType** | Pointer to **string** | The type of resource that was affected. For a full list of all resource-types, see the [Monitor Event Types](https://www.twilio.com/docs/usage/monitor-events#event-types). |
**Sid** | Pointer to **string** | The unique string that we created to identify the Event resource. |
**Source** | Pointer to **string** | The originating system or interface that caused the event.  Can be: `web` for events caused by user action in the Twilio Console, `api` for events caused by a request to our API, or   `twilio` for events caused by an automated or internal Twilio system. |
**SourceIpAddress** | Pointer to **string** | The IP address of the source, if the source is outside the Twilio cloud. This value is `null` for events with `source` of `twilio` |
**Url** | Pointer to **string** | The absolute URL of the resource that was affected. Can be `null`. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


