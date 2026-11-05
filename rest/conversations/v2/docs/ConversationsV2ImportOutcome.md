# ConversationsV2ImportOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**ConversationsV2ImportSourceType**](ConversationsV2ImportSourceType.md) |  |
**ResourceId** | **string** | The recording SID or Classic conversation SID. |
**Status** | [**ConversationsV2ImportOutcomeStatus**](ConversationsV2ImportOutcomeStatus.md) |  |
**FailureReason** | [**ConversationsV2ImportOutcomeFailureReason**](ConversationsV2ImportOutcomeFailureReason.md) |  |[optional] 
**ConversationId** | **string** | The conversation the resource was imported into, when there is one. |[optional] 
**ImportedCommunicationCount** | **int** | Communications created from the Classic conversation. Only present for a completed Classic conversation. |[optional] 
**ConversationClosed** | **bool** | Whether the import closed the conversation. Only present for a completed Classic conversation. |[optional] 
**AudioSeconds** | **int** | Seconds of audio sent for transcription; 0 when none was sent. Only present for a recording. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


