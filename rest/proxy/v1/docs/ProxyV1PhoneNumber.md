# ProxyV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the PhoneNumber resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the PhoneNumber resource. |
**ServiceSid** | Pointer to **string** | The SID of the PhoneNumber resource's parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**IsoCountry** | Pointer to **string** | The ISO Country Code for the phone number. |
**Capabilities** | Pointer to [**ProxyV1ServicePhoneNumberCapabilities**](ProxyV1ServicePhoneNumberCapabilities.md) |  |
**Url** | Pointer to **string** | The absolute URL of the PhoneNumber resource. |
**IsReserved** | Pointer to **bool** | Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. |
**InUse** | Pointer to **int** | The number of open session assigned to the number. See the [How many Phone Numbers do I need?](https://www.twilio.com/docs/proxy/phone-numbers-needed) guide for more information. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


