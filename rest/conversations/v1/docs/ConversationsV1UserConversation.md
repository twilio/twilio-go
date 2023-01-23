# ConversationsV1UserConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this conversation. |[optional] 
**ChatServiceSid** | **string** | The unique ID of the Conversation Service this conversation belongs to. |[optional] 
**ConversationSid** | **string** | The unique ID of the Conversation for this User Conversation. |[optional] 
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Conversation. |
**LastReadMessageIndex** | Pointer to **int** | The index of the last read Message . |
**ParticipantSid** | **string** | Participant Sid. |[optional] 
**UserSid** | **string** | The unique ID for the User. |[optional] 
**FriendlyName** | **string** | The human-readable name of this conversation. |[optional] 
**ConversationState** | Pointer to [**string**](UserConversationEnumState.md) |  |
**Timers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this conversation was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this conversation was last updated. |[optional] 
**CreatedBy** | **string** | Creator of this conversation. |[optional] 
**NotificationLevel** | Pointer to [**string**](UserConversationEnumNotificationLevel.md) |  |
**UniqueName** | **string** | An application-defined string that uniquely identifies the Conversation resource. |[optional] 
**Url** | **string** |  |[optional] 
**Links** | **map[string]interface{}** | Absolute URLs to access the participant and conversation of this user conversation. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


