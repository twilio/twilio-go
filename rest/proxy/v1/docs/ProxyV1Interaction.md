# ProxyV1Interaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Interaction resource. |
**SessionSid** | Pointer to **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource. |
**ServiceSid** | Pointer to **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Interaction resource. |
**Data** | Pointer to **string** | A JSON string that includes the message body of message interactions (e.g. `{\"body\": \"hello\"}`) or the call duration (when available) of a call (e.g. `{\"duration\": \"5\"}`). |
**Type** | Pointer to [**string**](InteractionEnumType.md) |  |
**InboundParticipantSid** | Pointer to **string** | The SID of the inbound [Participant](https://www.twilio.com/docs/proxy/api/participant) resource. |
**InboundResourceSid** | Pointer to **string** | The SID of the inbound resource; either the [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message. |
**InboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**InboundResourceType** | Pointer to **string** | The inbound resource type. Can be [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource). |
**InboundResourceUrl** | Pointer to **string** | The URL of the Twilio inbound resource |
**OutboundParticipantSid** | Pointer to **string** | The SID of the outbound [Participant](https://www.twilio.com/docs/proxy/api/participant)). |
**OutboundResourceSid** | Pointer to **string** | The SID of the outbound resource; either the [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource). |
**OutboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**OutboundResourceType** | Pointer to **string** | The outbound resource type. Can be: [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource). |
**OutboundResourceUrl** | Pointer to **string** | The URL of the Twilio outbound resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the Interaction was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**Url** | Pointer to **string** | The absolute URL of the Interaction resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


