# RecallObservationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The main content of the observation. |
**OccurredAt** | [**time.Time**](time.Time.md) | The timestamp when the observation originally occurred. |
**Source** | **string** | The source system that generated this observation. Allows letters, numbers, spaces, and URL-safe symbols. Excludes URL-unsafe characters like quotes, angle brackets, and control characters. |
**ConversationIds** | **[]string** | Array of conversation IDs associated with this observation. |[optional] 
**Id** | **string** | A unique identifier for the observation using Twilio Type ID (TTID) format. |
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the observation was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the observation was last updated. |
**Score** | **float64** | The relevance score of the observation in relation to the query. Higher values indicate greater relevance. This field is omitted when results are returned in most-recent order. This may occur when no query is provided and one cannot be inferred from the conversation context, or when the system defaults to chronological retrieval to ensure high availability. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


