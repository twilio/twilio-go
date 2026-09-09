# TraitExtractionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable description of this strategy. May be empty. |[optional] 
**IsEnabled** | **bool** | Flag indicating whether the strategy is active. When false, conversation configurations that reference it fall back to the default extraction behaviour. |[optional] [default to true]
**Definitions** | [**[]TraitExtractionDefinition**](TraitExtractionDefinition.md) | The traits this strategy infers. Replaces the existing list on update. |[optional] 
**DisplayName** | **string** | Unique, immutable name identifying a Trait Extraction Strategy within the account. |
**StoreId** | **string** | The Memory Store whose Trait Registry every definition in this strategy resolves against. The store is a property of the strategy, not of an individual definition, and is immutable after creation. |
**Id** | **string** | The unique identifier for the Trait Extraction Strategy. |
**CreatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the strategy was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the strategy was last updated. |
**Version** | **int** | The current version number of the strategy. Incremented on each successful update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


