# ConversationsV1ServiceWebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this service. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |
**PreWebhookUrl** | Pointer to **string** | The absolute url the pre-event webhook request should be sent to. |
**PostWebhookUrl** | Pointer to **string** | The absolute url the post-event webhook request should be sent to. |
**Filters** | Pointer to **[]string** | The list of events that your configured webhook targets will receive. Events not configured here will not fire. Possible values are `onParticipantAdd`, `onParticipantAdded`, `onDeliveryUpdated`, `onConversationUpdated`, `onConversationRemove`, `onParticipantRemove`, `onConversationUpdate`, `onMessageAdd`, `onMessageRemoved`, `onParticipantUpdated`, `onConversationAdded`, `onMessageAdded`, `onConversationAdd`, `onConversationRemoved`, `onParticipantUpdate`, `onMessageRemove`, `onMessageUpdated`, `onParticipantRemoved`, `onMessageUpdate` or `onConversationStateUpdated`. |
**Method** | Pointer to [**string**](ServiceWebhookConfigurationEnumMethod.md) |  |
**Url** | Pointer to **string** | An absolute API resource URL for this webhook. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


