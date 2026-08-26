# ListConversationByAccountResponseConversationsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |[optional] 
**Description** | **string** | Human-readable description for the Configuration. |[optional] 
**ConversationGroupingType** | [**ConversationsV2ConversationGroupingType**](ConversationsV2ConversationGroupingType.md) |  |[optional] 
**MemoryStoreId** | **string** | Memory Store ID for Profile resolution. |[optional] 
**ChannelSettings** | **map[string]interface{}** | Channel-specific parameters forwarded as-is to the downstream sending service. Allows passing backend-specific fields without requiring API changes.  |[optional] 
**StatusCallbacks** | [**[]ConversationsV2StatusCallbackConfig**](ConversationsV2StatusCallbackConfig.md) | List of default webhook configurations applied to Conversations under this Configuration. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | List of Intelligence Configuration IDs configured for this Configuration. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for conversations under this configuration. Defaults to false. |[optional] [default to false]
**ConversationsV1Bridge** | [**ConversationsV2ConversationsV1Bridge**](ConversationsV2ConversationsV1Bridge.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


