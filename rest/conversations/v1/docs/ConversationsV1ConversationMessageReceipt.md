# ConversationsV1ConversationMessageReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this participant. |
**ConversationSid** | Pointer to **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**MessageSid** | Pointer to **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to |
**ChannelMessageSid** | Pointer to **string** | A messaging channel-specific identifier for the message delivered to participant e.g. `SMxx` for SMS, `WAxx` for Whatsapp etc.  |
**ParticipantSid** | Pointer to **string** | The unique ID of the participant the delivery receipt belongs to. |
**Status** | Pointer to [**string**](ConversationMessageReceiptEnumDeliveryStatus.md) |  |
**ErrorCode** | Pointer to **int** | The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a `failed` status,  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. `null` if the delivery receipt has not been updated. |
**Url** | Pointer to **string** | An absolute API resource URL for this delivery receipt. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


