# ConversationsV2StartConversationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this Action. |
**Status** | [**ConversationsV2ActionStatus**](ConversationsV2ActionStatus.md) |  |
**Related** | **map[string]string** | Named identifiers, populated as the asynchronous work completes: - conversationId: The created Conversation (present once the Conversation is created) - channelId: The downstream channel identifier (present once dispatch succeeds) - executionSid: VOICE only. The Studio Flow execution that placed the call - callId: VOICE only. The Call SID of the outbound call  Absent from the 202 response (returned as an empty map).  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Action was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Action was last updated. |[optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the Action reached a terminal status. Null until the Action reaches one. |
**FailureReason** | Pointer to **string** | Human-readable failure reason. Null unless status is FAILED. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


