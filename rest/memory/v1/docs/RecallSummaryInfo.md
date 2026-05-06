# RecallSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source system that generated the summary. Allows letters, numbers, spaces, and URL-safe symbols. Excludes URL-unsafe characters like quotes, angle brackets, and control characters. |[optional] 
**Content** | **string** | The main content of the summary. |
**OccurredAt** | [**time.Time**](time.Time.md) | The timestamp when the summary was originally created. |
**ConversationId** | **string** | A unique identifier for the conversation using Twilio Type ID (TTID) format. |
**Id** | **string** | A unique identifier for the summary using Twilio Type ID (TTID) format. |
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the summary was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the summary was last updated. |
**Score** | **float64** | The relevance score of the summary in relation to the query. Higher values indicate greater relevance. This field is omitted when results are returned in most-recent order. This may occur when no query is provided and one cannot be inferred from the conversation context, or when the system defaults to chronological retrieval to ensure high availability. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


