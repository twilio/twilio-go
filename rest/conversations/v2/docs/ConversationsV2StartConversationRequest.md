# ConversationsV2StartConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | [**ConversationsV2StartConversationChannel**](ConversationsV2StartConversationChannel.md) |  |
**From** | [**ConversationsV2StartConversationParticipant**](ConversationsV2StartConversationParticipant.md) |  |
**To** | [**[]ConversationsV2StartConversationParticipant**](ConversationsV2StartConversationParticipant.md) | The recipient of the Conversation. Exactly one recipient is supported (1:1). |
**Content** | [**ConversationsV2SendContent**](ConversationsV2SendContent.md) |  |[optional] 
**ConfigurationId** | **string** | Configuration governing grouping, timeouts, and channel configuration. |
**ChannelSettings** | **map[string]interface{}** | Channel-specific parameters forwarded as-is to the downstream sending service. Allows passing backend-specific fields without requiring API changes.  |[optional] 
**OrchestratorPolicy** | [**ConversationsV2OrchestratorPolicy**](ConversationsV2OrchestratorPolicy.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


