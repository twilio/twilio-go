# WebSourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Web based knowledge sources |
**Url** | **string** | The URL to crawl for web content |
**CrawlDepth** | **int** | The maximum depth to crawl from the source URL |[optional] [default to 2]
**CrawlPeriod** | **string** | Frequency of re-crawling the website for updated content |[optional] [default to "NEVER"]
**Errors** | [**[]KnowledgeErrorGroup**](KnowledgeErrorGroup.md) | Processing errors encountered during web crawling, grouped by title. Array of error groups, where each group has a title and list of error instances. Only present when crawl errors occurred. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


