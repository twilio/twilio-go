# ConversationsV1ServiceConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this message. |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the resource is associated with. |[optional] 
**ConversationSid** | **string** | The unique ID of the Conversation for this message. |[optional] 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**Index** | **int** | The index of the message within the Conversation. |[optional] 
**Author** | **string** | The channel specific identifier of the message's author. |[optional] 
**Body** | **string** | The content of the message. |[optional] 
**Media** | **[]interface{}** | An array of objects that describe the Message's media if attached, otherwise, null. |[optional] 
**Attributes** | **string** | A string metadata field you can use to store any data you wish. |[optional] 
**ParticipantSid** | **string** | The unique ID of messages's author participant. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 
**Delivery** | Pointer to **interface{}** | An object that contains the summary of delivery statuses for the message to non-chat participants. |
**Url** | **string** | An absolute URL for this message. |[optional] 
**Links** | **map[string]interface{}** | Absolute URL to access the receipts of this message. |[optional] 
**ContentSid** | **string** | The unique ID of the multi-channel Rich Content template. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


