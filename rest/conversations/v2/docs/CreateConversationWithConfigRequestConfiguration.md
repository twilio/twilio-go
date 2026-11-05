# CreateConversationWithConfigRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationGroupingType** | **string** | Type of Conversation grouping strategy: - `GROUP_BY_PROFILE`: Groups Communications by resolved Profile from the Memory Store.   A Profile is looked up or created for `CUSTOMER` Participant types. All Communications from the same Profile are in the same Conversation, regardless of address or channel. - `GROUP_BY_PARTICIPANT_ADDRESSES`: Groups Communications by Participant addresses across all channels.   A customer using +18005550100 will be in the same Conversation whether they contact by SMS, WhatsApp, or RCS. - `GROUP_BY_PARTICIPANT_ADDRESSES_AND_CHANNEL_TYPE`: Groups Communications by both Participant addresses AND channel.   A customer using +18005550100 by SMS will be in a different Conversation than the same customer by Voice.  |[optional] 
**MemoryStoreId** | **string** | The memory store ID that Conversation Orchestrator uses for profile resolution. |[optional] 
**ChannelSettings** | [**map[string]CreateConversationWithConfigRequestConfigurationChannelSettingsValue**](CreateConversationWithConfigRequestConfigurationChannelSettingsValue.md) | Status timeouts by channel type. Capture rules can only be set on a Configuration. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for this Conversation. Without `configurationId`, defaults to false; with it, defaults to the Configuration's value. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | A list of Conversational Intelligence configuration IDs. |[optional] 
**TraitExtractionStrategyIds** | **[]string** | A list of Trait Extraction Strategy IDs. |[optional] 
**ObservationExtractionStrategyIds** | **[]string** | A list of Observation Extraction Strategy IDs. |[optional] 
**StatusCallbacks** | [**[]CreateConversationWithConfigRequestConfigurationStatusCallbacks**](CreateConversationWithConfigRequestConfigurationStatusCallbacks.md) | The webhook configurations for this Conversation. Overrides the Configuration's own. |[optional] 
**Workflows** | [**[]ConversationWorkflow**](ConversationWorkflow.md) | The Workflows to associate with this Conversation. Overrides the Configuration's own. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


