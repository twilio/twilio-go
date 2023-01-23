# ConversationsV1ConversationParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this participant. |[optional] 
**ConversationSid** | **string** | The unique ID of the Conversation for this participant. |[optional] 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**Identity** | **string** | A unique string identifier for the conversation participant as Conversation User. |[optional] 
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. |[optional] 
**MessagingBinding** | Pointer to **interface{}** | Information about how this participant exchanges messages with the conversation. |
**RoleSid** | **string** | The SID of a conversation-level Role to assign to the participant |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 
**Url** | **string** | An absolute URL for this participant. |[optional] 
**LastReadMessageIndex** | Pointer to **int** | Index of last “read” message in the Conversation for the Participant. |
**LastReadTimestamp** | **string** | Timestamp of last “read” message in the Conversation for the Participant. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


