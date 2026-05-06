# CreateConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |
**Description** | **string** | Human-readable description for the configuration. |
**ConversationGroupingType** | **string** | The strategy Conversation Orchestrator uses to assign communications to conversations. |
**MemoryStoreId** | **string** | The memory store ID that Conversation Orchestrator uses for profile resolution. |
**ChannelSettings** | [**map[string]CreateConfigurationRequestChannelSettingsValue**](CreateConfigurationRequestChannelSettingsValue.md) |  |[optional] 
**StatusCallbacks** | [**[]CreateConfigurationRequestStatusCallbacks**](CreateConfigurationRequestStatusCallbacks.md) | A list of webhook configurations. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | A list of Conversational Intelligence configuration IDs. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for conversations under this configuration. Defaults to false. |[optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


