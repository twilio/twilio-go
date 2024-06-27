# MarketplaceV1InstalledAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the InstalledAddOn resource. This Sid can also be found in the Console on that specific Add-ons page as the 'Available Add-on Sid'. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the InstalledAddOn resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Description** | Pointer to **string** | A short description of the Add-on's functionality. |
**Configuration** | Pointer to **interface{}** | The JSON object that represents the current configuration of installed Add-on. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


