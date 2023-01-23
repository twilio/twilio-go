# ApiV2010UsageRecordMonthly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that accrued the usage. |
**ApiVersion** | Pointer to **string** | The API version used to create the resource. |
**AsOf** | Pointer to **string** | Usage records up to date as of this timestamp, formatted as YYYY-MM-DDTHH:MM:SS+00:00. All timestamps are in GMT |
**Category** | Pointer to [**string**](UsageRecordMonthlyEnumCategory.md) |  |
**Count** | Pointer to **string** | The number of usage events, such as the number of calls. |
**CountUnit** | Pointer to **string** | The units in which `count` is measured, such as `calls` for calls or `messages` for SMS. |
**Description** | Pointer to **string** | A plain-language description of the usage category. |
**EndDate** | Pointer to **string** | The last date for which usage is included in the UsageRecord. The date is specified in GMT and formatted as `YYYY-MM-DD`. |
**Price** | Pointer to **float32** | The total price of the usage in the currency specified in `price_unit` and associated with the account. |
**PriceUnit** | Pointer to **string** | The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format, such as `usd`, `eur`, and `jpy`. |
**StartDate** | Pointer to **string** | The first date for which usage is included in this UsageRecord. The date is specified in GMT and formatted as `YYYY-MM-DD`. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their URIs. For more information, see [List Subresources](https://www.twilio.com/docs/usage/api/usage-record#list-subresources). |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**Usage** | Pointer to **string** | The amount used to bill usage and measured in units described in `usage_unit`. |
**UsageUnit** | Pointer to **string** | The units in which `usage` is measured, such as `minutes` for calls or `messages` for SMS. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


