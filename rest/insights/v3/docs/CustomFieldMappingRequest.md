# CustomFieldMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**MappingCategory**](MappingCategory.md) |  |
**Name** | **string** | Name of the custom field, used to reference it in Insights queries. Contains only letters, digits and `_`, and must not be empty. Must be unique within the account across every `sourceType`; registering an existing name returns `409`. Cannot be changed after registration.  |
**Description** | **string** | Optional, free-text description of this custom field mapping, up to 150 characters. Display-only — not part of any uniqueness check.  |[optional] 
**EntityMetadata** | Pointer to [**EntityMetadata**](EntityMetadata.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


