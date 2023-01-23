# ConversationsV1ConversationMessageReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique ID of the Account responsible for this participant. |[optional] 
**ConversationSid** | **string** | The unique ID of the Conversation for this message. |[optional] 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**MessageSid** | **string** | The SID of the message the delivery receipt belongs to |[optional] 
**ChannelMessageSid** | **string** | A messaging channel-specific identifier for the message delivered to participant |[optional] 
**ParticipantSid** | **string** | The unique ID of the participant the delivery receipt belongs to. |[optional] 
**Status** | Pointer to [**string**](ConversationMessageReceiptEnumDeliveryStatus.md) |  |
**ErrorCode** | **int** | The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a `failed` status |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 
**Url** | **string** | An absolute URL for this delivery receipt. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


