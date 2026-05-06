# SummaryCorePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source system that generated the summary. Allows letters, numbers, spaces, and URL-safe symbols. Excludes URL-unsafe characters like quotes, angle brackets, and control characters. |[optional] 
**Content** | **string** | The main content of the summary. |[optional] 
**OccurredAt** | [**time.Time**](time.Time.md) | The timestamp when the summary was originally created. If not provided, defaults to the time the summary was received. |[optional] 
**ConversationId** | **string** | A unique identifier for the conversation using Twilio Type ID (TTID) format. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


