# ConversationsV1ServiceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this configuration. |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the Configuration applies to. |[optional] 
**NewMessage** | Pointer to **interface{}** | The Push Notification configuration for New Messages. |
**AddedToConversation** | Pointer to **interface{}** | The Push Notification configuration for being added to a Conversation. |
**RemovedFromConversation** | Pointer to **interface{}** | The Push Notification configuration for being removed from a Conversation. |
**LogEnabled** | **bool** | Weather the notification logging is enabled. |[optional] 
**Url** | **string** | An absolute URL for this configuration. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


