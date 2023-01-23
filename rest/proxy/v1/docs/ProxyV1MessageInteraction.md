# ProxyV1MessageInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SessionSid** | **string** | The SID of the resource's parent Session |[optional] 
**ServiceSid** | **string** | The SID of the resource's parent Service |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Data** | **string** | A JSON string that includes the message body sent to the participant |[optional] 
**Type** | Pointer to [**string**](MessageInteractionEnumType.md) |  |
**ParticipantSid** | **string** | The SID of the Participant resource |[optional] 
**InboundParticipantSid** | **string** | Always empty for Message Interactions |[optional] 
**InboundResourceSid** | **string** | Always empty for Message Interactions |[optional] 
**InboundResourceStatus** | Pointer to [**string**](MessageInteractionEnumResourceStatus.md) |  |
**InboundResourceType** | **string** | Always empty for Message Interactions |[optional] 
**InboundResourceUrl** | **string** | Always empty for Message Interactions |[optional] 
**OutboundParticipantSid** | **string** | The SID of the outbound Participant resource |[optional] 
**OutboundResourceSid** | **string** | The SID of the outbound Message resource |[optional] 
**OutboundResourceStatus** | Pointer to [**string**](MessageInteractionEnumResourceStatus.md) |  |
**OutboundResourceType** | **string** | The outbound resource type |[optional] 
**OutboundResourceUrl** | **string** | The URL of the Twilio message resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the MessageInteraction resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


