# ChatV2Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Message resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource. |
**Attributes** | Pointer to **string** | The JSON string that stores application-specific data. If attributes have not been set, `{}` is returned. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Message resource is associated with. |
**To** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) that the message was sent to. |
**ChannelSid** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Message resource belongs to. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**LastUpdatedBy** | Pointer to **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable. |
**WasEdited** | Pointer to **bool** | Whether the message has been edited since it was created. |
**From** | Pointer to **string** | The [Identity](https://www.twilio.com/docs/chat/identity) of the message's author. The default value is `system`. |
**Body** | Pointer to **string** | The content of the message. |
**Index** | Pointer to **int** | The index of the message within the [Channel](https://www.twilio.com/docs/chat/channels). Indices may skip numbers, but will always be in order of when the message was received. |
**Type** | Pointer to **string** | The Message type. Can be: `text` or `media`. |
**Media** | Pointer to **interface{}** | An object that describes the Message's media, if the message contains media. The object contains these fields: `content_type` with the MIME type of the media, `filename` with the name of the media, `sid` with the SID of the Media resource, and `size` with the media object's file size in bytes. If the Message has no media, this value is `null`. |
**Url** | Pointer to **string** | The absolute URL of the Message resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


