# ConversationsV1ServiceUserConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |
**ConversationSid** | Pointer to **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this User Conversation. |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Conversation for the Participant. |
**LastReadMessageIndex** | Pointer to **int** | The index of the last Message in the Conversation that the Participant has read. |
**ParticipantSid** | Pointer to **string** | The unique ID of the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) the user conversation belongs to. |
**UserSid** | Pointer to **string** | The unique string that identifies the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation, limited to 256 characters. Optional. |
**ConversationState** | Pointer to [**string**](ServiceUserConversationEnumState.md) |  |
**Timers** | Pointer to **interface{}** | Timer date values representing state update for this conversation. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated, given in ISO 8601 format. |
**CreatedBy** | Pointer to **string** | Identity of the creator of this Conversation. |
**NotificationLevel** | Pointer to [**string**](ServiceUserConversationEnumNotificationLevel.md) |  |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Conversation resource. It can be used to address the resource in place of the resource's `conversation_sid` in the URL. |
**Url** | Pointer to **string** |  |
**Links** | Pointer to **map[string]interface{}** | Contains absolute URLs to access the [participant](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) and [conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) of this conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


