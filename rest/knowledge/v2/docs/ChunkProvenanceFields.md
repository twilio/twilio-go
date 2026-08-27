# ChunkProvenanceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChunkIndex** | **int** | 0-based position of this chunk within its source document for a single ingestion run. |[optional] 
**DocumentTitle** | **string** | Human-readable title of the source document. Web: HTML <title> from the crawled page. File: filename from Unstructured metadata. Text: knowledge name from the knowledge source. |[optional] 
**DocumentUrl** | **string** | Specific page URL this chunk was crawled from. Web sources only; null for File and Text sources. |[optional] 
**DocumentNumber** | **int** | Physical page number (1-based). PDF sources only; omitted for all other source types. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


