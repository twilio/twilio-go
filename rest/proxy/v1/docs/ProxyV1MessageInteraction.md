# ProxyV1MessageInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Data** | Pointer to **string** | A JSON string that includes the message body sent to the participant |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**InboundParticipantSid** | Pointer to **string** | Always empty for Message Interactions |
**InboundResourceSid** | Pointer to **string** | Always empty for Message Interactions |
**InboundResourceStatus** | Pointer to **string** | Always empty for Message Interactions |
**InboundResourceType** | Pointer to **string** | Always empty for Message Interactions |
**InboundResourceUrl** | Pointer to **string** | Always empty for Message Interactions |
**OutboundParticipantSid** | Pointer to **string** | The SID of the outbound Participant resource |
**OutboundResourceSid** | Pointer to **string** | The SID of the outbound Message resource |
**OutboundResourceStatus** | Pointer to **string** | The outbound resource status |
**OutboundResourceType** | Pointer to **string** | The outbound resource type |
**OutboundResourceUrl** | Pointer to **string** | The URL of the Twilio message resource |
**ParticipantSid** | Pointer to **string** | The SID of the Participant resource |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**SessionSid** | Pointer to **string** | The SID of the resource's parent Session |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Type** | Pointer to **string** | The Type of Message Interaction |
**Url** | Pointer to **string** | The absolute URL of the MessageInteraction resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


