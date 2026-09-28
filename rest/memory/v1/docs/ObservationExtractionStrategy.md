# ObservationExtractionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable description of this strategy. May be empty. |[optional] 
**IsEnabled** | **bool** | Flag indicating whether the strategy is active. When false, conversation configurations that reference it fall back to the default extraction behaviour. |[optional] [default to true]
**AllowedCategories** | [**[]ObservationCategory**](ObservationCategory.md) | Categories of observation this strategy extracts. Replaces the existing list on update. |[optional] 
**ProhibitedCategories** | [**[]ObservationCategory**](ObservationCategory.md) | Categories of observation this strategy must never extract. Replaces the existing list on update. |[optional] 
**DisplayName** | **string** | Unique, immutable name identifying an Observation Extraction Strategy within the account. |
**Id** | **string** | The unique identifier for the Observation Extraction Strategy. |
**CreatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the strategy was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the strategy was last updated. |
**Version** | **int** | The current version number of the strategy. Incremented on each successful update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


