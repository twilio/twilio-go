# MessagingV1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource. |
**MessagingServiceSid** | Pointer to **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) the message belongs to. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Message resource. |
**Url** | Pointer to **string** | The absolute URL of the Message resource. |
**Recipients** | Pointer to **[]interface{}** | The list of message recipients, formatted as `channel:number`. |
**From** | Pointer to **string** | A single colon-delimited address. |
**Body** | Pointer to **string** | The message body. |
**BroadcastSid** | Pointer to **string** | The SID of the Broadcast resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateOfExpiry** | Pointer to [**time.Time**](time.Time.md) | The date when the resource expires. |
**Template** | Pointer to **string** | Template for message body. |
**TemplateSid** | Pointer to **string** | Template sid to get template for message body. |
**TemplateLanguage** | Pointer to **string** | Template language which complements template sid to get template for message body. |
**TemplateArgs** | Pointer to **interface{}** | The dictionary of arguments to construct message body from template. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


