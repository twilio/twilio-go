# CustomFieldMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | TTID for conversation insights custom field mapping |[optional] 
**Category** | [**MappingCategory**](MappingCategory.md) |  |
**Name** | **string** |  |
**Description** | **string** | Optional, free-text description supplied at registration time. Display-only — not part of any uniqueness check.  |[optional] 
**EntityMetadata** | Pointer to [**IntelligenceOperatorEntityMetadata**](IntelligenceOperatorEntityMetadata.md) |  |
**CreatedAt** | [**time.Time**](time.Time.md) | When the mapping was registered. |[optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | When the mapping was last materially modified in-place. Null until the first update-in-place. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


