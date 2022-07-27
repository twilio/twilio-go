# ConversationsV1ConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this message. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this message. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Index** | Pointer to **int** | The index of the message within the Conversation. |
**Author** | Pointer to **string** | The channel specific identifier of the message's author. |
**Body** | Pointer to **string** | The content of the message. |
**Media** | Pointer to **[]interface{}** | An array of objects that describe the Message's media if attached, otherwise, null. |
**Attributes** | Pointer to **string** | A string metadata field you can use to store any data you wish. |
**ParticipantSid** | Pointer to **string** | The unique ID of messages's author participant. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Url** | Pointer to **string** | An absolute API URL for this message. |
**Delivery** | Pointer to **interface{}** | An object that contains the summary of delivery statuses for the message to non-chat participants. |
**Links** | Pointer to **map[string]interface{}** | Absolute URL to access the receipts of this message. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


