# ProxyV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the ShortCode resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource. |
**ServiceSid** | Pointer to **string** | The SID of the ShortCode resource's parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**ShortCode** | Pointer to **string** | The short code's number. |
**IsoCountry** | Pointer to **string** | The ISO Country Code for the short code. |
**Capabilities** | Pointer to [**ProxyV1ServiceShortCodeCapabilities**](ProxyV1ServiceShortCodeCapabilities.md) |  |
**Url** | Pointer to **string** | The absolute URL of the ShortCode resource. |
**IsReserved** | Pointer to **bool** | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


