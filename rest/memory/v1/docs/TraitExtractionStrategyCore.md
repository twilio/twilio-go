# TraitExtractionStrategyCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A human readable description of this strategy. May be empty. |[optional] 
**IsEnabled** | **bool** | Flag indicating whether the strategy is active. When false, conversation configurations that reference it fall back to the default extraction behaviour. |[optional] [default to true]
**Definitions** | [**[]TraitExtractionDefinition**](TraitExtractionDefinition.md) | The traits this strategy infers. Replaces the existing list on update. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


