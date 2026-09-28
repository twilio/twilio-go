# CreateConversationWithConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | **string** | The ID of an existing configuration. |
**Name** | **string** | The name of the conversation. |[optional] 
**Configuration** | Pointer to [**CreateConversationWithConfigRequestConfiguration**](CreateConversationWithConfigRequestConfiguration.md) |  |
**Participants** | [**[]CreateConversationWithConfigRequestParticipants**](CreateConversationWithConfigRequestParticipants.md) | Optional list of Participants to create with the Conversation. |[optional] 
**Metadata** | **map[string]string** | Optional customer-managed key-value metadata for this Conversation. Maximum 8 entries; keys up to 128 characters allowing alphanumeric characters, periods, underscores, and dashes; values up to 512 characters. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


