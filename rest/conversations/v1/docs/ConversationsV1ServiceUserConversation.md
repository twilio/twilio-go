# ConversationsV1ServiceUserConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this User Conversation. |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Conversation. |
**LastReadMessageIndex** | Pointer to **int** | The index of the last read Message . |
**ParticipantSid** | Pointer to **string** | Participant Sid. |
**UserSid** | Pointer to **string** | The unique ID for the User. |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**ConversationState** | Pointer to [**string**](ServiceUserConversationEnumState.md) |  |
**Timers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated. |
**CreatedBy** | Pointer to **string** | Creator of this conversation. |
**NotificationLevel** | Pointer to [**string**](ServiceUserConversationEnumNotificationLevel.md) |  |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Conversation resource. |
**Url** | Pointer to **string** |  |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participant and conversation of this user conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


