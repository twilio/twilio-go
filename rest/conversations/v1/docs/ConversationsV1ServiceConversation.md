# ConversationsV1ServiceConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to. |
**MessagingServiceSid** | Pointer to **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation, limited to 256 characters. Optional. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned. |
**State** | Pointer to [**string**](ServiceConversationEnumState.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Timers** | Pointer to **interface{}** | Timer date values representing state update for this conversation. |
**Url** | Pointer to **string** | An absolute API resource URL for this conversation. |
**Links** | Pointer to **map[string]interface{}** | Contains absolute URLs to access the [participants](https://www.twilio.com/docs/conversations/api/conversation-participant-resource), [messages](https://www.twilio.com/docs/conversations/api/conversation-message-resource) and [webhooks](https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource) of this conversation. |
**Bindings** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


