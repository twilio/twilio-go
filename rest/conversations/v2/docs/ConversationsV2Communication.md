# ConversationsV2Communication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Communication ID. |
**ConversationId** | **string** | Conversation ID. |
**AccountId** | **string** | Account ID. |
**Author** | [**ConversationsV2ParticipantAddress**](ConversationsV2ParticipantAddress.md) |  |
**Content** | [**ListCommunicationByConversationResponseCommunicationsContent**](ListCommunicationByConversationResponseCommunicationsContent.md) |  |
**ChannelId** | **string** | Channel-specific reference ID. |[optional] 
**ResourceId** | **string** | External resource identifier for this Communication (e.g. MessageSid for SMS/RCS/WhatsApp, TranscriptionSid + MessageIndex for Voice). When set, used for Communication deduplication/uniqueness within a Conversation. |[optional] 
**Recipients** | [**[]ListCommunicationByConversationResponseCommunicationsRecipients**](ListCommunicationByConversationResponseCommunicationsRecipients.md) | Communication recipients. |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Communication was created. |[optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Communication was last updated. |[optional] 
**OccurredAt** | [**time.Time**](time.Time.md) | ISO 8601 timestamp when the communication occurred. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


