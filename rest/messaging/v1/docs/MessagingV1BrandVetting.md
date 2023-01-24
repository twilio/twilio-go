# MessagingV1BrandVetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the vetting record. |
**BrandSid** | Pointer to **string** | The unique string to identify Brand Registration. |
**BrandVettingSid** | Pointer to **string** | The Twilio SID of the third-party vetting record. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**VettingId** | Pointer to **string** | The unique identifier of the vetting from the third-party provider. |
**VettingClass** | Pointer to **string** | The type of vetting that has been conducted. One of “STANDARD” (Aegis) or “POLITICAL” (Campaign Verify). |
**VettingStatus** | Pointer to **string** | The status of the import vetting attempt. One of “PENDING,” “SUCCESS,” or “FAILED”. |
**VettingProvider** | Pointer to [**string**](BrandVettingEnumVettingProvider.md) |  |
**Url** | Pointer to **string** | The absolute URL of the Brand Vetting resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


