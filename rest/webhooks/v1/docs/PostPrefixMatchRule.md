# PostPrefixMatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** | Human-readable name for the Rule (optional, defaults to blank) |
**Description** | **string** | Optional description of the Rule's purpose and usage |[optional] 
**MatchType** | **string** |  |
**Settings** | [**[]SettingReference**](SettingReference.md) | Settings to apply with traffic distribution (percentages must sum to 100) |
**Prefix** | **string** | URL prefix to match, MUST be unique across all prefix match rules and MUST end with `/`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


