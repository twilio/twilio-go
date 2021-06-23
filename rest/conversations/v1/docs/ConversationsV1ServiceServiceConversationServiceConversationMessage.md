# ConversationsV1ServiceServiceConversationServiceConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this message. |
**Attributes** | Pointer to **string** | A string metadata field you can use to store any data you wish. |
**Author** | Pointer to **string** | The channel specific identifier of the message's author. |
**Body** | Pointer to **string** | The content of the message. |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this message. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Delivery** | Pointer to **map[string]interface{}** | An object that contains the summary of delivery statuses for the message to non-chat participants. |
**Index** | Pointer to **int** | The index of the message within the Conversation. |
**Links** | Pointer to **map[string]interface{}** | Absolute URL to access the receipts of this message. |
**Media** | Pointer to **[]map[string]interface{}** | An array of objects that describe the Message's media if attached, otherwise, null. |
**ParticipantSid** | Pointer to **string** | The unique ID of messages's author participant. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Url** | Pointer to **string** | An absolute URL for this message. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


