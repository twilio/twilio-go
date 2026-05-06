# MemoryRetrievalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** | A unique identifier for the conversation using Twilio Type ID (TTID) format. |[optional] 
**Query** | **string** | Hybrid search query for finding relevant memories. Omit to use query expansion to generate query from previous 10 communications in conversation. |[optional] 
**BeginDate** | [**time.Time**](time.Time.md) | Start date for filtering memories (inclusive). |[optional] 
**EndDate** | [**time.Time**](time.Time.md) | End date for filtering memories (exclusive). |[optional] 
**CommunicationsLimit** | **int** | Maximum number of conversational session memories to return. If omitted or set to 0, no session memories will be fetched. |[optional] [default to 0]
**ObservationsLimit** | **int** | Maximum number of observation memories to return. If omitted, defaults to 20. If set to 0, no observation memories will be fetched. |[optional] [default to 20]
**SummariesLimit** | **int** | Maximum number of summary memories to return. If omitted, defaults to 5. If set to 0, no summary memories will be fetched. |[optional] [default to 5]
**RelevanceThreshold** | **float64** | Minimum relevance score threshold for observations and summaries to be returned. Only memories with a relevance score greater than or equal to this threshold will be included in the response. This threshold only applies when results are ranked by relevance. When results are returned in most-recent order, this field has no effect. |[optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


