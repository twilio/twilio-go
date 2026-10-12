# CreateCommunicationInConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**CreateCommunicationInConversationRequestAuthor**](CreateCommunicationInConversationRequestAuthor.md) |  |
**Content** | Pointer to [**CreateCommunicationInConversationRequestContent**](CreateCommunicationInConversationRequestContent.md) |  |
**ChannelId** | **string** |  |[optional] 
**ResourceId** | **string** | External resource identifier for this Communication (e.g. MessageSid for SMS/RCS/WhatsApp, TranscriptionSid + MessageIndex for Voice). If a Communication with the same resourceId already exists in the Conversation, it is updated instead of a new one being created. |[optional] 
**Recipients** | [**[]CreateCommunicationInConversationRequestRecipients**](CreateCommunicationInConversationRequestRecipients.md) |  |
**OccurredAt** | [**time.Time**](time.Time.md) | Timestamp when this Communication occurred. If omitted, the server uses the current time. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


