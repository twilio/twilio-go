# ConversationsV2StartConversationSendMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**ConversationsV2StartConversationParticipantSelector**](ConversationsV2StartConversationParticipantSelector.md) |  |[optional] 
**To** | [**[]ConversationsV2StartConversationParticipantSelector**](ConversationsV2StartConversationParticipantSelector.md) | The single recipient (1:1). Defaults to the roster's only `CUSTOMER` Participant. |[optional] 
**Content** | [**ConversationsV2SendContent**](ConversationsV2SendContent.md) |  |
**ChannelSettings** | **map[string]interface{}** | Channel-specific parameters forwarded as-is to the downstream sending service. Allows passing backend-specific fields without requiring API changes.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


