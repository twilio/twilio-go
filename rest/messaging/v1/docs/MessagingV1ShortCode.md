# MessagingV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the ShortCode resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the resource is associated with. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ShortCode** | Pointer to **string** | The [E.164](https://www.twilio.com/docs/glossary/what-e164) format of the short code. |
**CountryCode** | Pointer to **string** | The 2-character [ISO Country Code](https://www.iso.org/iso-3166-country-codes.html) of the number. |
**Capabilities** | Pointer to **[]string** | An array of values that describe whether the number can receive calls or messages. Can be: `SMS` and `MMS`. |
**Url** | Pointer to **string** | The absolute URL of the ShortCode resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


