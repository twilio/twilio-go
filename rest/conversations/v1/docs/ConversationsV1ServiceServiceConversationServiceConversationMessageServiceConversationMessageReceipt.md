# ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this participant. |
**ChannelMessageSid** | Pointer to **string** | A messaging channel-specific identifier for the message delivered to participant |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with. |
**ConversationSid** | Pointer to **string** | The unique ID of the Conversation for this message. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**ErrorCode** | Pointer to **int32** | The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a `failed` status |
**MessageSid** | Pointer to **string** | The SID of the message the delivery receipt belongs to |
**ParticipantSid** | Pointer to **string** | The unique ID of the participant the delivery receipt belongs to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Status** | Pointer to **string** | The message delivery status |
**Url** | Pointer to **string** | An absolute URL for this delivery receipt. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


