# CustomFieldMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the custom field mapping. |[optional] 
**Category** | [**MappingCategory**](MappingCategory.md) |  |
**Name** | **string** | Name of the custom field, unique within the account. |
**Description** | **string** | Optional, free-text description supplied at registration time. Display-only — not part of any uniqueness check.  |[optional] 
**EntityMetadata** | Pointer to [**EntityMetadata**](EntityMetadata.md) |  |
**CreatedAt** | [**time.Time**](time.Time.md) | When the mapping was registered. |[optional] 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | When the mapping was last modified via `PATCH`. Null until the first update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


