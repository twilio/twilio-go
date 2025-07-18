# ConversationsV1ServiceConversationScopedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with. |
**ConversationSid** | Pointer to **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook. |
**Target** | Pointer to **string** | The target of this webhook: `webhook`, `studio`, `trigger` |
**Url** | Pointer to **string** | An absolute API resource URL for this webhook. |
**Configuration** | Pointer to **interface{}** | The configuration of this webhook. Is defined based on target. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


