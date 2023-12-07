# MessagingV1ChannelSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ChannelSender resource. |
**MessagingServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/messaging/services) the resource is associated with. |
**Sid** | Pointer to **string** | The unique string that we created to identify the ChannelSender resource. |
**Sender** | Pointer to **string** | The unique string that identifies the sender e.g whatsapp:+123456XXXX. |
**SenderType** | Pointer to **string** | A string value that identifies the sender type e.g WhatsApp, Messenger. |
**CountryCode** | Pointer to **string** | The 2-character [ISO Country Code](https://www.iso.org/iso-3166-country-codes.html) of the number. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the ChannelSender resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


