# ConversationsV1Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this conversation. |[optional] 
**ChatServiceSid** | **string** | The unique ID of the Conversation Service this conversation belongs to. |[optional] 
**MessagingServiceSid** | **string** | The unique ID of the Messaging Service this conversation belongs to. |[optional] 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**FriendlyName** | **string** | The human-readable name of this conversation. |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. |[optional] 
**State** | Pointer to [**string**](ConversationEnumState.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 
**Timers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Url** | **string** | An absolute URL for this conversation. |[optional] 
**Links** | **map[string]interface{}** | Absolute URLs to access the participants, messages and webhooks of this conversation. |[optional] 
**Bindings** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


