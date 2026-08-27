# ListConfigurationResponseConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Configuration ID. |
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |
**Description** | **string** | Human-readable description for the Configuration. Allows spaces and special characters, typically limited to a paragraph of text. This serves as a descriptive field rather than just a name. |
**ConversationGroupingType** | [**ConversationsV2ConversationGroupingType**](ConversationsV2ConversationGroupingType.md) |  |
**MemoryStoreId** | **string** | Memory Store ID for Profile resolution. |
**ChannelSettings** | [**map[string]ConversationsV2ChannelSetting**](ConversationsV2ChannelSetting.md) | Channel-specific configuration settings by channel type. Keys should be valid channel types (`VOICE`, `SMS`, `RCS`, `WHATSAPP`, `CHAT`). |[optional] 
**StatusCallbacks** | [**[]ConversationsV2StatusCallbackConfig**](ConversationsV2StatusCallbackConfig.md) | List of default webhook configurations applied to Conversations under this Configuration. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | A list of Conversational Intelligence configuration IDs. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for conversations under this configuration. Defaults to false. |[optional] [default to false]
**ConversationsV1Bridge** | [**ConversationsV2ConversationsV1Bridge**](ConversationsV2ConversationsV1Bridge.md) |  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Configuration was created. |[optional] [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Configuration was last updated. |[optional] [readonly] 
**Version** | **int64** | Version number used for optimistic locking. |[optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


