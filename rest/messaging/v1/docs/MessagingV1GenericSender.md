# MessagingV1GenericSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The SID to identify the number or channel sender resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the number or channel sender resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the resource is associated with. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Sender** | Pointer to **string** | The unique string that identifies the number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the channel sender e.g whatsapp:+123456XXXX. |
**SenderType** | Pointer to **string** | A string value that identifies the number or channel sender type e.g AlphaSenderId, LongCode, ShortCode, Whatsapp, RCS. |
**CountryCode** | Pointer to **string** | The 2-character [ISO Country Code](https://www.iso.org/iso-3166-country-codes.html) of the number or channel sender. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


