# ConversationsV1ServiceParticipantConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this conversation. |[optional] 
**ChatServiceSid** | **string** | The unique ID of the Conversation Service this conversation belongs to. |[optional] 
**ParticipantSid** | **string** | The unique ID of the Participant. |[optional] 
**ParticipantUserSid** | **string** | The unique ID for the conversation participant as Conversation User. |[optional] 
**ParticipantIdentity** | **string** | A unique string identifier for the conversation participant as Conversation User. |[optional] 
**ParticipantMessagingBinding** | Pointer to **interface{}** | Information about how this participant exchanges messages with the conversation. |
**ConversationSid** | **string** | The unique ID of the Conversation this Participant belongs to. |[optional] 
**ConversationUniqueName** | **string** | An application-defined string that uniquely identifies the Conversation resource. |[optional] 
**ConversationFriendlyName** | **string** | The human-readable name of this conversation. |[optional] 
**ConversationAttributes** | **string** | An optional string metadata field you can use to store any data you wish. |[optional] 
**ConversationDateCreated** | [**time.Time**](time.Time.md) | The date that this conversation was created. |[optional] 
**ConversationDateUpdated** | [**time.Time**](time.Time.md) | The date that this conversation was last updated. |[optional] 
**ConversationCreatedBy** | **string** | Creator of this conversation. |[optional] 
**ConversationState** | Pointer to [**string**](ServiceParticipantConversationEnumState.md) |  |
**ConversationTimers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Links** | **map[string]interface{}** | Absolute URLs to access the participant and conversation of this Participant Conversation. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


