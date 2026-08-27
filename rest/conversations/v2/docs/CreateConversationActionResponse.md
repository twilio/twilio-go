# CreateConversationActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this Action. |
**Type** | **string** | The type of action. Accepted values: SEND_MESSAGE. |
**Status** | [**ConversationsV2ActionStatus**](ConversationsV2ActionStatus.md) |  |
**ConversationId** | **string** | The conversation this action belongs to. |
**Related** | **map[string]string** | Named identifiers from downstream. For SEND_MESSAGE: - messageSid: The downstream message SID (present when PENDING or COMPLETED) - communicationId: The Communication ID (present when COMPLETED)  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the action was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the action was last updated. |[optional] 
**CompletedAt** | [**time.Time**](time.Time.md) | Timestamp when the action reached a terminal status. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


