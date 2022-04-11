# ConversationsV1ServiceConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this conversation. |
**Attributes** | Pointer to **string** | An optional string metadata field you can use to store any data you wish. |
**Bindings** | Pointer to **interface{}** |  |
**ChatServiceSid** | Pointer to **string** | The unique ID of the Conversation Service this conversation belongs to. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**FriendlyName** | Pointer to **string** | The human-readable name of this conversation. |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participants, messages and webhooks of this conversation. |
**MessagingServiceSid** | Pointer to **string** | The unique ID of the Messaging Service this conversation belongs to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**State** | Pointer to **string** | Current state of this conversation. |
**Timers** | Pointer to **interface{}** | Timer date values for this conversation. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | An absolute URL for this conversation. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


