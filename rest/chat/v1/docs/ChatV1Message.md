# ChatV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Message resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Message resource. |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data. **Note** If this property has been assigned a value, it's only  displayed in a FETCH action that returns a single resource; otherwise, it's null. If the attributes have not been set, `{}` is returned. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) the resource is associated with. |
**To** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/api/channels) that the message was sent to. |
**ChannelSid** | Pointer to **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the Message resource belongs to. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**WasEdited** | Pointer to **bool** | Whether the message has been edited since it was created. |
**From** | Pointer to **string** | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the message's author. The default value is `system`. |
**Body** | Pointer to **string** | The content of the message. |
**Index** | Pointer to **int** | The index of the message within the [Channel](https://www.twilio.com/docs/chat/api/channels). |
**Url** | Pointer to **string** | The absolute URL of the Message resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


