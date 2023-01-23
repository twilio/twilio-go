# ConversationsV1ServiceConversationScopedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**AccountSid** | **string** | The unique ID of the Account responsible for this conversation. |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the resource is associated with. |[optional] 
**ConversationSid** | **string** | The unique ID of the Conversation for this webhook. |[optional] 
**Target** | **string** | The target of this webhook. |[optional] 
**Url** | **string** | An absolute URL for this webhook. |[optional] 
**Configuration** | Pointer to **interface{}** | The configuration of this webhook. |
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


