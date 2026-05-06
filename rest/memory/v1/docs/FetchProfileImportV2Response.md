# FetchProfileImportV0Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Current processing status of the import task |
**Filename** | **string** | Original filename of the uploaded CSV |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the import was created |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the import was last updated |[optional] 
**FileSize** | **int** | Size of the uploaded file in bytes (1 byte to 100 MiB) |[optional] 
**ColumnMappings** | [**[]ColumnMappingItem**](ColumnMappingItem.md) | Mappings of CSV header columns to traits' fields |[optional] 
**Summary** | [**FetchProfileImportV0ResponseSummary**](FetchProfileImportV0ResponseSummary.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


