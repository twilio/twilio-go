# ProxyV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent Service |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**ShortCode** | Pointer to **string** | The short code's number |
**IsoCountry** | Pointer to **string** | The ISO Country Code |
**Capabilities** | Pointer to [**ProxyV1ServiceShortCodeCapabilities**](ProxyV1ServiceShortCodeCapabilities.md) |  |
**Url** | Pointer to **string** | The absolute URL of the ShortCode resource |
**IsReserved** | Pointer to **bool** | Whether the short code should be reserved for manual assignment to participants only |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


