# ConversationsV1ServiceParticipantConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**ConversationAttributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**ConversationCreatedBy** | Pointer to **string** | Creator of this conversation. |
**ConversationDateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created. |
**ConversationDateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated. |
**ConversationFriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation this Participant belongs to. |
**ConversationState** | Pointer to **string** | The current state of this User Conversation |
**ConversationTimers** | Pointer to **interface{}** | Timer date values for this conversation. |
**ConversationUniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Conversation resource. |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participant and conversation of this Participant Conversation. |
**ParticipantIdentity** | Pointer to **string** | A unique string identifier for the conversation participant as Conversation User. |
**ParticipantMessagingBinding** | Pointer to **interface{}** | Information about how this participant exchanges messages with the conversation. |
**ParticipantSid** | Pointer to **string** | The unique ID of the Participant. |
**ParticipantUserSid** | Pointer to **string** | The unique ID for the conversation participant as Conversation User. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


