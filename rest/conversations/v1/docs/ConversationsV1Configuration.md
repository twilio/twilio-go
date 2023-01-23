# ConversationsV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account responsible for this configuration. |[optional] 
**DefaultChatServiceSid** | **string** | The SID of the default Conversation Service that every new conversation is associated with. |[optional] 
**DefaultMessagingServiceSid** | **string** | The SID of the default Messaging Service that every new conversation is associated with. |[optional] 
**DefaultInactiveTimer** | **string** | Default ISO8601 duration when conversation will be switched to `inactive` state. |[optional] 
**DefaultClosedTimer** | **string** | Default ISO8601 duration when conversation will be switched to `closed` state. |[optional] 
**Url** | **string** | An absolute URL for this global configuration. |[optional] 
**Links** | **map[string]interface{}** | Absolute URLs to access the webhook and default service configurations. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


