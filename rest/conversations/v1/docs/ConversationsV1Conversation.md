# ConversationsV1Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**MessagingServiceSid** | Pointer to **string** | The unique ID of the Messaging Service this conversation belongs to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**State** | Pointer to [**string**](ConversationEnumState.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Timers** | Pointer to **interface{}** | Timer date values for this conversation. |
**Url** | Pointer to **string** | An absolute URL for this conversation. |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participants, messages and webhooks of this conversation. |
**Bindings** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


