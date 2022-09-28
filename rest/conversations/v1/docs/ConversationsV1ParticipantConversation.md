# ConversationsV1ParticipantConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**ParticipantSid** | Pointer to **string** | The unique ID of the Participant. |
**ParticipantUserSid** | Pointer to **string** | The unique ID for the conversation participant as Conversation User. |
**ParticipantIdentity** | Pointer to **string** | A unique string identifier for the conversation participant as Conversation User. |
**ParticipantMessagingBinding** | Pointer to **interface{}** | Information about how this participant exchanges messages with the conversation. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation this Participant belongs to. |
**ConversationUniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Conversation resource |
**ConversationFriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**ConversationAttributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**ConversationDateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created. |
**ConversationDateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated. |
**ConversationCreatedBy** | Pointer to **string** | Creator of this conversation. |
**ConversationState** | Pointer to [**string**](ParticipantConversationEnumState.md) |  |
**ConversationTimers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participant and conversation of this Participant Conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


