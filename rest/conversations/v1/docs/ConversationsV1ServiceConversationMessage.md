# ConversationsV1ServiceConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this message. |
**ChatServiceSid** | Pointer to **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with. |
**ConversationSid** | Pointer to **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Index** | Pointer to **int** | The index of the message within the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource). |
**Author** | Pointer to **string** | The channel specific identifier of the message's author. Defaults to `system`. |
**Body** | Pointer to **string** | The content of the message, can be up to 1,600 characters long. |
**Media** | Pointer to **[]interface{}** | An array of objects that describe the Message's media, if the message contains media. Each object contains these fields: `content_type` with the MIME type of the media, `filename` with the name of the media, `sid` with the SID of the Media resource, and `size` with the media object's file size in bytes. If the Message has no media, this value is `null`. |
**Attributes** | Pointer to **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned. |
**ParticipantSid** | Pointer to **string** | The unique ID of messages's author participant. Null in case of `system` sent message. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. `null` if the message has not been edited. |
**Delivery** | Pointer to **interface{}** | An object that contains the summary of delivery statuses for the message to non-chat participants. |
**Url** | Pointer to **string** | An absolute API resource URL for this message. |
**Links** | Pointer to **map[string]interface{}** | Contains an absolute API resource URL to access the delivery & read receipts of this message. |
**ContentSid** | Pointer to **string** | The unique ID of the multi-channel [Rich Content](https://www.twilio.com/docs/content-api) template. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


