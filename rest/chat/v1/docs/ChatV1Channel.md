# ChatV1Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Channel resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Channel resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) the resource is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data. **Note** If this property has been assigned a value, it's only  displayed in a FETCH action that returns a single resource; otherwise, it's null. If the attributes have not been set, `{}` is returned. |
**Type** | Pointer to [**string**](ChannelEnumChannelType.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**CreatedBy** | Pointer to **string** | The `identity` of the User that created the channel. If the Channel was created by using the API, the value is `system`. |
**MembersCount** | Pointer to **int** | The number of Members in the Channel. |
**MessagesCount** | Pointer to **int** | The number of Messages in the Channel. |
**Url** | Pointer to **string** | The absolute URL of the Channel resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the [Members](https://www.twilio.com/docs/chat/api/members), [Messages](https://www.twilio.com/docs/chat/api/messages) , [Invites](https://www.twilio.com/docs/chat/api/invites) and, if it exists, the last [Message](https://www.twilio.com/docs/chat/api/messages) for the Channel. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


