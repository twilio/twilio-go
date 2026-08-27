# CreateConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |
**Description** | **string** | Human-readable description for the configuration. |
**ConversationGroupingType** | **string** | Type of Conversation grouping strategy: - `GROUP_BY_PROFILE`: Groups Communications by resolved Profile from the Memory Store.   A Profile is looked up or created for `CUSTOMER` Participant types. All Communications from the same Profile are in the same Conversation, regardless of address or channel. - `GROUP_BY_PARTICIPANT_ADDRESSES`: Groups Communications by Participant addresses across all channels.   A customer using +18005550100 will be in the same Conversation whether they contact by SMS, WhatsApp, or RCS. - `GROUP_BY_PARTICIPANT_ADDRESSES_AND_CHANNEL_TYPE`: Groups Communications by both Participant addresses AND channel.   A customer using +18005550100 by SMS will be in a different Conversation than the same customer by Voice.  |
**MemoryStoreId** | **string** | The memory store ID that Conversation Orchestrator uses for profile resolution. |
**ChannelSettings** | [**map[string]CreateConfigurationRequestChannelSettingsValue**](CreateConfigurationRequestChannelSettingsValue.md) |  |[optional] 
**StatusCallbacks** | [**[]CreateConfigurationRequestStatusCallbacks**](CreateConfigurationRequestStatusCallbacks.md) | A list of webhook configurations. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | A list of Conversational Intelligence configuration IDs. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for conversations under this configuration. Defaults to false. |[optional] [default to false]
**ConversationsV1Bridge** | [**CreateConfigurationRequestConversationsV1Bridge**](CreateConfigurationRequestConversationsV1Bridge.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


