# ListConversationByAccountResponseConversationsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A human-readable name for the configuration. Limited to 32 characters. |[optional] 
**Description** | **string** | Human-readable description for the Configuration. |[optional] 
**ConversationGroupingType** | **string** | Type of Conversation grouping strategy: - `GROUP_BY_PROFILE`: Groups Communications by resolved Profile from the Memory Store.   A Profile is looked up or created for `CUSTOMER` Participant types. All Communications from the same Profile are in the same Conversation, regardless of address or channel. - `GROUP_BY_PARTICIPANT_ADDRESSES`: Groups Communications by Participant addresses across all channels.   A customer using +18005550100 will be in the same Conversation whether they contact by SMS, WhatsApp, or RCS. - `GROUP_BY_PARTICIPANT_ADDRESSES_AND_CHANNEL_TYPE`: Groups Communications by both Participant addresses AND channel.   A customer using +18005550100 by SMS will be in a different Conversation than the same customer by Voice.  |[optional] 
**MemoryStoreId** | **string** | Memory Store ID for Profile resolution. |[optional] 
**ChannelSettings** | **map[string]interface{}** | Channel-specific parameters forwarded as-is to the downstream sending service. Allows passing backend-specific fields without requiring API changes.  |[optional] 
**StatusCallbacks** | [**[]ConversationsV2StatusCallbackConfig**](ConversationsV2StatusCallbackConfig.md) | List of default webhook configurations applied to Conversations under this Configuration. |[optional] 
**IntelligenceConfigurationIds** | **[]string** | List of Intelligence Configuration IDs configured for this Configuration. |[optional] 
**MemoryExtractionEnabled** | **bool** | Whether memory extraction is enabled for conversations under this configuration. Defaults to false. |[optional] [default to false]
**ConversationsV1Bridge** | [**ConversationsV2ConversationsV1Bridge**](ConversationsV2ConversationsV1Bridge.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


