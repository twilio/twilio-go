# KnowledgeSourceTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Raw text knowledge sources |
**Content** | **string** | The raw text content to be processed |
**Url** | **string** | The URL to crawl for web content |
**CrawlDepth** | **int** | The maximum depth to crawl from the source URL |[optional] [default to 2]
**CrawlPeriod** | **string** | Frequency of re-crawling the website for updated content |[optional] [default to "NEVER"]
**Errors** | [**[]KnowledgeErrorGroup**](KnowledgeErrorGroup.md) | Processing errors encountered during web crawling, grouped by title. Array of error groups, where each group has a title and list of error instances. Only present when crawl errors occurred. |[optional] 
**FileName** | **string** | Name of the file to be uploaded |
**FileSize** | **int** | Expected size of the file in bytes |
**MimeType** | [**SupportedFileMimeType**](SupportedFileMimeType.md) |  |
**ImportUrl** | **string** | Presigned S3 URL for file upload (when status is SCHEDULED).  Use PUT method to upload the file to this URL when status is SCHEDULED. |[optional] [readonly] 
**UploadExpiration** | [**time.Time**](time.Time.md) | Expiration time of the presigned upload URL in ISO 8601 format (only present when status is SCHEDULED) |[optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


