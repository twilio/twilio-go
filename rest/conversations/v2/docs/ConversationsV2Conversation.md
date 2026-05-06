# ConversationsV2Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Conversation ID. |
**AccountId** | **string** | Account ID. |
**ConfigurationId** | **string** | Configuration ID. |
**Status** | **string** | Conversation status. |[optional] 
**Name** | Pointer to **string** | Conversation name. |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Conversation was created. |[optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Conversation was last updated. |[optional] 
**Configuration** | [**ListConversationByAccountResponseConversationsConfiguration**](ListConversationByAccountResponseConversationsConfiguration.md) |  |[optional] 
**Participants** | [**[]ConversationsV2Participant**](ConversationsV2Participant.md) | Participants in this Conversation. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


