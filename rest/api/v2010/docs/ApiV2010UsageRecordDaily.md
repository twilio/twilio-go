# ApiV2010UsageRecordDaily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account accrued the usage |
**ApiVersion** | Pointer to **string** | The API version used to create the resource |
**AsOf** | Pointer to **string** | Usage records up to date as of this timestamp |
**Category** | Pointer to **string** | The category of usage |
**Count** | Pointer to **string** | The number of usage events |
**CountUnit** | Pointer to **string** | The units in which count is measured |
**Description** | Pointer to **string** | A plain-language description of the usage category |
**EndDate** | Pointer to **string** | The last date for which usage is included in the UsageRecord |
**Price** | Pointer to **float32** | The total price of the usage |
**PriceUnit** | Pointer to **string** | The currency in which `price` is measured |
**StartDate** | Pointer to **string** | The first date for which usage is included in this UsageRecord |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |
**Usage** | Pointer to **string** | The amount of usage |
**UsageUnit** | Pointer to **string** | The units in which usage is measured |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


