# PatchExactMatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | **string** | Human-readable name for the Rule |[optional] 
**Description** | **string** | Optional description of the Rule's purpose and usage |[optional] 
**MatchType** | **string** |  |
**Value** | **string** | New exact URL to match, MUST be unique across all exact match rules and MUST NOT end with `/`. |
**Settings** | [**[]SettingReference**](SettingReference.md) | Settings to apply with traffic distribution (percentages must sum to 100) |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


