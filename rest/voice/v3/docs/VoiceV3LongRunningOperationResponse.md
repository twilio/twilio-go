# VoiceV3LongRunningOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | Unique identifier for the transcription operation. |
**Status** | **string** | Current status of the transcription operation. PENDING: accepted but not yet started. RUNNING: currently in progress. COMPLETED: successfully completed. FAILED: failed and cannot be completed.  |
**StatusUrl** | **string** | URL to poll for the latest operation status. |
**Transcription** | [**VoiceV3Transcription**](VoiceV3Transcription.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


