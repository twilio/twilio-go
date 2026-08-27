# ConversationsV2StartConversationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this Action. |
**Status** | [**ConversationsV2ActionStatus**](ConversationsV2ActionStatus.md) |  |
**Related** | **map[string]string** | Named identifiers, populated as the asynchronous work completes and varying by action type: - conversationId: The created Conversation - channelId: The downstream channel identifier, for SEND_MESSAGE - executionSid: The Studio Flow execution started for the Conversation, for START_FLOW - callId: The Call SID of the outbound call, for CALL and for START_FLOW once its Flow reports back  Absent from the 202 response (returned as an empty map).  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Action was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Action was last updated. |[optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the Action reached a terminal status. Null until the Action reaches one. |
**FailureReason** | Pointer to **string** | Human-readable failure reason. Null unless status is FAILED. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


