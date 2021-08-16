# ConversationsV1ServiceConversationParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this participant. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this participant. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Identity** | Pointer to **string** | A unique string identifier for the conversation participant as Conversation User. |
**LastReadMessageIndex** | Pointer to **int** | Index of last “read” message in the Conversation for the Participant. |
**LastReadTimestamp** | Pointer to **string** | Timestamp of last “read” message in the Conversation for the Participant. |
**MessagingBinding** | Pointer to **map[string]interface{}** | Information about how this participant exchanges messages with the conversation. |
**RoleSid** | Pointer to **string** | The SID of a conversation-level Role to assign to the participant |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Url** | Pointer to **string** | An absolute URL for this participant. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


