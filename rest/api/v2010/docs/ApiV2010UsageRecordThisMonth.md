# ApiV2010UsageRecordThisMonth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account accrued the usage |[optional] 
**ApiVersion** | **string** | The API version used to create the resource |[optional] 
**AsOf** | **string** | Usage records up to date as of this timestamp |[optional] 
**Category** | Pointer to [**string**](UsageRecordThisMonthEnumCategory.md) |  |
**Count** | **string** | The number of usage events |[optional] 
**CountUnit** | **string** | The units in which count is measured |[optional] 
**Description** | **string** | A plain-language description of the usage category |[optional] 
**EndDate** | **string** | The last date for which usage is included in the UsageRecord |[optional] 
**Price** | **float32** | The total price of the usage |[optional] 
**PriceUnit** | **string** | The currency in which `price` is measured |[optional] 
**StartDate** | **string** | The first date for which usage is included in this UsageRecord |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their relative URIs |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**Usage** | **string** | The amount of usage |[optional] 
**UsageUnit** | **string** | The units in which usage is measured |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


