# FetchConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The `id` of the Conversation attached to the Operator Result. |
**AccountId** | Pointer to **string** | The ID of the account that owns the Conversation. |
**Name** | Pointer to **string** | Display name of the Conversation. |
**Status** | [**ConversationStatus**](ConversationStatus.md) |  |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp for when the Conversation was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp for when the Conversation was last updated. |
**IntelligenceConfigurationIds** | **[]string** | The Intelligence Configuration(s) associated with the Conversation. |
**ConversationConfigurationId** | **string** | The `id` of the Configuration for a Conversation. |
**Channels** | [**[]Channel**](Channel.md) | The communication channel(s) included in the Conversation. |
**ChannelIds** | **[]string** | The underlying channel resource `id`s associated with this Conversation, such as a Call ID or Message ID. |
**Participants** | [**[]Participant**](Participant.md) | Metadata for Participants of the Conversation. |
**Communications** | [**[]Communication**](Communication.md) | Metadata for the Communications that make up the Conversation. |
**OperatorResultIds** | **[]string** | List of Operator Result IDs generated from this Conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


