# ChatV2Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Channel resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Channel resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Channel resource is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data. If attributes have not been set, `{}` is returned. |
**Type** | Pointer to [**string**](ChannelEnumChannelType.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**CreatedBy** | Pointer to **string** | The `identity` of the User that created the channel. If the Channel was created by using the API, the value is `system`. |
**MembersCount** | Pointer to **int** | The number of Members in the Channel. |
**MessagesCount** | Pointer to **int** | The number of Messages that have been passed in the Channel. |
**Url** | Pointer to **string** | The absolute URL of the Channel resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the [Members](https://www.twilio.com/docs/chat/rest/member-resource), [Messages](https://www.twilio.com/docs/chat/rest/message-resource), [Invites](https://www.twilio.com/docs/chat/rest/invite-resource), Webhooks and, if it exists, the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) for the Channel. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


