# DefaultRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Rule identifier (TypeID/TTID format) |[readonly] 
**FriendlyName** | **string** | Human-readable name for identification |[optional] 
**Description** | **string** | Optional description of the Rule's purpose and usage |[optional] 
**MatchType** | **string** |  |
**Settings** | [**[]SettingReference**](SettingReference.md) | Settings to apply with traffic distribution (percentages must sum to 100) |
**CreatedAt** | [**time.Time**](time.Time.md) |  |[optional] [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  |[optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


