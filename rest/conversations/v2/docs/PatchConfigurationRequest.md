# PatchConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A human-readable name for the configuration. Limited to 32 characters. |
**Description** | Pointer to **string** | Human-readable description for the configuration. |
**ConversationGroupingType** | **string** | Type of Conversation grouping strategy: - `GROUP_BY_PROFILE`: Groups Communications by resolved Profile from the Memory Store.   A Profile is looked up or created for `CUSTOMER` Participant types. All Communications from the same Profile are in the same Conversation, regardless of address or channel. - `GROUP_BY_PARTICIPANT_ADDRESSES`: Groups Communications by Participant addresses across all channels.   A customer using +18005550100 will be in the same Conversation whether they contact by SMS, WhatsApp, or RCS. - `GROUP_BY_PARTICIPANT_ADDRESSES_AND_CHANNEL_TYPE`: Groups Communications by both Participant addresses AND channel.   A customer using +18005550100 by SMS will be in a different Conversation than the same customer by Voice.  |[optional] 
**MemoryStoreId** | Pointer to **string** | The Memory Store ID for profile resolution. |
**ChannelSettings** | [**map[string]PatchConfigurationRequestChannelSettingsValue**](PatchConfigurationRequestChannelSettingsValue.md) | Channel-specific settings to merge onto the existing channelSettings map. A channel key mapped to a value replaces that channel's settings; a channel key explicitly mapped to null removes it; an omitted channel key is left untouched. |[optional] 
**StatusCallbacks** | Pointer to [**[]UpdateConfigurationRequestStatusCallbacks**](UpdateConfigurationRequestStatusCallbacks.md) |  |
**IntelligenceConfigurationIds** | Pointer to **[]string** | A list of Conversational Intelligence configuration IDs. |
**MemoryExtractionEnabled** | Pointer to **bool** | Whether memory extraction is enabled for conversations under this configuration. |
**ConversationsV1Bridge** | Pointer to [**PatchConfigurationRequestConversationsV1Bridge**](PatchConfigurationRequestConversationsV1Bridge.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


