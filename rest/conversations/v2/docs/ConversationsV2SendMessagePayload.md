# ConversationsV2SendMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**ConversationsV2SendMessageParticipant**](ConversationsV2SendMessageParticipant.md) |  |
**To** | [**[]ConversationsV2SendMessageParticipant**](ConversationsV2SendMessageParticipant.md) | The recipients of this action. |
**Content** | [**ConversationsV2SendMessageContent**](ConversationsV2SendMessageContent.md) |  |
**ChannelSettings** | **map[string]interface{}** | Channel-specific parameters forwarded as-is to the downstream sending service. Allows passing backend-specific fields without requiring API changes.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


