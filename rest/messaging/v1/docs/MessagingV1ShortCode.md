# MessagingV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**ShortCode** | Pointer to **string** | The E.164 format of the short code |
**CountryCode** | Pointer to **string** | The 2-character ISO Country Code of the number |
**Capabilities** | Pointer to **[]string** | An array of values that describe whether the number can receive calls or messages |
**Url** | Pointer to **string** | The absolute URL of the ShortCode resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


