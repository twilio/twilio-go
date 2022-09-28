# ProxyV1Interaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SessionSid** | Pointer to **string** | The SID of the resource's parent Session |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Data** | Pointer to **string** | A JSON string that includes the message body of message interactions |
**Type** | Pointer to [**string**](InteractionEnumType.md) |  |
**InboundParticipantSid** | Pointer to **string** | The SID of the inbound Participant resource |
**InboundResourceSid** | Pointer to **string** | The SID of the inbound resource |
**InboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**InboundResourceType** | Pointer to **string** | The inbound resource type |
**InboundResourceUrl** | Pointer to **string** | The URL of the Twilio inbound resource |
**OutboundParticipantSid** | Pointer to **string** | The SID of the outbound Participant |
**OutboundResourceSid** | Pointer to **string** | The SID of the outbound resource |
**OutboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**OutboundResourceType** | Pointer to **string** | The outbound resource type |
**OutboundResourceUrl** | Pointer to **string** | The URL of the Twilio outbound resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Interaction was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Interaction resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


