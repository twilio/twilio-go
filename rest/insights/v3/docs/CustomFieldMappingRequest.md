# CustomFieldMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**MappingCategory**](MappingCategory.md) |  |
**Name** | **string** | Customer-facing name for the field; becomes the ClickHouse Map key and the platform field alias. Unique per account across every `sourceType` — the Map key namespace is account-global.  |
**Description** | **string** | Optional, free-text description of this custom field mapping. Display-only — not part of any uniqueness check.  |[optional] 
**EntityMetadata** | Pointer to [**IntelligenceOperatorEntityMetadata**](IntelligenceOperatorEntityMetadata.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


