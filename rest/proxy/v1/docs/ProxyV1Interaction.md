# ProxyV1Interaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SessionSid** | **string** | The SID of the resource's parent Session |[optional] 
**ServiceSid** | **string** | The SID of the resource's parent Service |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Data** | **string** | A JSON string that includes the message body of message interactions |[optional] 
**Type** | Pointer to [**string**](InteractionEnumType.md) |  |
**InboundParticipantSid** | **string** | The SID of the inbound Participant resource |[optional] 
**InboundResourceSid** | **string** | The SID of the inbound resource |[optional] 
**InboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**InboundResourceType** | **string** | The inbound resource type |[optional] 
**InboundResourceUrl** | **string** | The URL of the Twilio inbound resource |[optional] 
**OutboundParticipantSid** | **string** | The SID of the outbound Participant |[optional] 
**OutboundResourceSid** | **string** | The SID of the outbound resource |[optional] 
**OutboundResourceStatus** | Pointer to [**string**](InteractionEnumResourceStatus.md) |  |
**OutboundResourceType** | **string** | The outbound resource type |[optional] 
**OutboundResourceUrl** | **string** | The URL of the Twilio outbound resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Interaction was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Interaction resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


