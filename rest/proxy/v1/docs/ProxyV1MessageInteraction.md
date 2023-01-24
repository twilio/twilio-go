# ProxyV1MessageInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the MessageInteraction resource. |
**SessionSid** | Pointer to **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource. |
**ServiceSid** | Pointer to **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the MessageInteraction resource. |
**Data** | Pointer to **string** | A JSON string that includes the message body sent to the participant. (e.g. `{\"body\": \"hello\"}`) |
**Type** | Pointer to [**string**](MessageInteractionEnumType.md) |  |
**ParticipantSid** | Pointer to **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource. |
**InboundParticipantSid** | Pointer to **string** | Always empty for created Message Interactions. |
**InboundResourceSid** | Pointer to **string** | Always empty for created Message Interactions. |
**InboundResourceStatus** | Pointer to [**string**](MessageInteractionEnumResourceStatus.md) |  |
**InboundResourceType** | Pointer to **string** | Always empty for created Message Interactions. |
**InboundResourceUrl** | Pointer to **string** | Always empty for created Message Interactions. |
**OutboundParticipantSid** | Pointer to **string** | The SID of the outbound [Participant](https://www.twilio.com/docs/proxy/api/participant) resource. |
**OutboundResourceSid** | Pointer to **string** | The SID of the outbound [Message](https://www.twilio.com/docs/sms/api/message-resource) resource. |
**OutboundResourceStatus** | Pointer to [**string**](MessageInteractionEnumResourceStatus.md) |  |
**OutboundResourceType** | Pointer to **string** | The outbound resource type. This value is always `Message`. |
**OutboundResourceUrl** | Pointer to **string** | The URL of the Twilio message resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**Url** | Pointer to **string** | The absolute URL of the MessageInteraction resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


