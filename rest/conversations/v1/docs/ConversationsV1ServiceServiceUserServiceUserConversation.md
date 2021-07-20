# ConversationsV1ServiceServiceUserServiceUserConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this message. |
**ConversationState** | Pointer to **string** | The current state of this User Conversation |
**CreatedBy** | Pointer to **string** | Creator of this conversation. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this conversation was last updated. |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**LastReadMessageIndex** | Pointer to **int** | The index of the last read Message . |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participant and conversation of this user conversation. |
**NotificationLevel** | Pointer to **string** | The Notification Level of this User Conversation. |
**ParticipantSid** | Pointer to **string** | Participant Sid. |
**Timers** | Pointer to **map[string]interface{}** | Timer date values for this conversation. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Conversation. |
**Url** | Pointer to **string** |  |
**UserSid** | Pointer to **string** | The unique ID for the User. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


