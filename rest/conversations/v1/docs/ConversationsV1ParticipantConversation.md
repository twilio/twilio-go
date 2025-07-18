# ConversationsV1ParticipantConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |
**ParticipantSid** | Pointer to **string** | The unique ID of the [Participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource). |
**ParticipantUserSid** | Pointer to **string** | The unique string that identifies the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). |
**ParticipantIdentity** | Pointer to **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters. |
**ParticipantMessagingBinding** | Pointer to **interface{}** | Information about how this participant exchanges messages with the conversation. A JSON parameter consisting of type and address fields of the participant. |
**ConversationSid** | Pointer to **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) this Participant belongs to. |
**ConversationUniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Conversation resource. |
**ConversationFriendlyName** | Pointer to **string** | The human-readable name of this conversation, limited to 256 characters. Optional. |
**ConversationAttributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned. |
**ConversationDateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created, given in ISO 8601 format. |
**ConversationDateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated, given in ISO 8601 format. |
**ConversationCreatedBy** | Pointer to **string** | Identity of the creator of this Conversation. |
**ConversationState** | Pointer to [**string**](ParticipantConversationEnumState.md) |  |
**ConversationTimers** | Pointer to **interface{}** | Timer date values representing state update for this conversation. |
**Links** | Pointer to **map[string]interface{}** | Contains absolute URLs to access the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) and [conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) of this conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


