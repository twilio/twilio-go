# ProxyV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Capabilities** | Pointer to [**ProxyV1ServicePhoneNumberCapabilities**](ProxyV1ServicePhoneNumberCapabilities.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**InUse** | Pointer to **int** | The number of open session assigned to the number. |
**IsReserved** | Pointer to **bool** | Reserve the phone number for manual assignment to participants only |
**IsoCountry** | Pointer to **string** | The ISO Country Code |
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**ServiceSid** | Pointer to **string** | The SID of the PhoneNumber resource's parent Service resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the PhoneNumber resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


