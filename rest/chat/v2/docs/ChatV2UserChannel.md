# ChatV2UserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the User Channel resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the User Channel resource is associated with. |
**ChannelSid** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the User Channel resource belongs to. |
**UserSid** | Pointer to **string** | The SID of the [User](https://www.twilio.com/docs/chat/rest/user-resource) the User Channel belongs to. |
**MemberSid** | Pointer to **string** | The SID of a [Member](https://www.twilio.com/docs/chat/rest/member-resource) that represents the User on the Channel. |
**Status** | Pointer to [**string**](UserChannelEnumChannelStatus.md) |  |
**LastConsumedMessageIndex** | Pointer to **int** | The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. |
**UnreadMessagesCount** | Pointer to **int** | The number of unread Messages in the Channel for the User. Note that retrieving messages on a client endpoint does not mean that messages are consumed or read. See [Consumption Horizon feature](https://www.twilio.com/docs/chat/consumption-horizon) to learn how to mark messages as consumed. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the [Members](https://www.twilio.com/docs/chat/rest/member-resource), [Messages](https://www.twilio.com/docs/chat/rest/message-resource) , [Invites](https://www.twilio.com/docs/chat/rest/invite-resource) and, if it exists, the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) for the Channel. |
**Url** | Pointer to **string** | The absolute URL of the User Channel resource. |
**NotificationLevel** | Pointer to [**string**](UserChannelEnumNotificationLevel.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


