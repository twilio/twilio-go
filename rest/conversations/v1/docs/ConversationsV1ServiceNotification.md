# ConversationsV1ServiceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this configuration. |
**ChatServiceSid** | Pointer to **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to. |
**NewMessage** | Pointer to **interface{}** | The Push Notification configuration for New Messages. |
**AddedToConversation** | Pointer to **interface{}** | The Push Notification configuration for being added to a Conversation. |
**RemovedFromConversation** | Pointer to **interface{}** | The Push Notification configuration for being removed from a Conversation. |
**LogEnabled** | Pointer to **bool** | Weather the notification logging is enabled. |
**Url** | Pointer to **string** | An absolute API resource URL for this configuration. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


