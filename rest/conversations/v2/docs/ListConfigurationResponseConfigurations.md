# ListConfigurationResponseConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Configuration ID. |
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |
**Description** | **string** | Human-readable description for the Configuration. Allows spaces and special characters, typically limited to a paragraph of text. This serves as a descriptive field rather than just a name. |
**ConversationGroupingType** | **string** | Type of Conversation grouping strategy: - `GROUP_BY_PROFILE`: Groups Communications by resolved Profile from the Memory Store.   A Profile is looked up or created for `CUSTOMER` Participant types. All Communications from the same Profile are in the same Conversation, regardless of address or channel. - `GROUP_BY_PARTICIPANT_ADDRESSES`: Groups Communications by Participant addresses across all channels.   A customer using +18005550100 will be in the same Conversation whether they contact by SMS, WhatsApp, or RCS. - `GROUP_BY_PARTICIPANT_ADDRESSES_AND_CHANNEL_TYPE`: Groups Communications by both Participant addresses AND channel.   A customer using +18005550100 by SMS will be in a different Conversation than the same customer by Voice.  |
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


