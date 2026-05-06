# VoiceV3Transcription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for a Transcription. This is also the transcriptionId returned in the LRO 202 response. |
**AccountId** | **string** | Twilio Account SID |
**Status** | **string** | The current status of the transcription operation |
**TranscriptionConfigurationId** | **string** | Unique identifier for a Transcription configuration. |
**MediaUrl** | Pointer to **string** | The third party media URL |
**SourceId** | Pointer to **string** | The source ID (recording ID) - used for tracking only |
**AudioStartedAt** | [**time.Time**](time.Time.md) | The call/recording start time. When the transcription was created using a sourceId, this value is inferred from the recording resource's start time. When created using a mediaUrl, this reflects the value supplied by the caller.  |[optional] 
**ConversationId** | Pointer to **string** | Maestro conversation ID, populated once the transcription has been stored in Maestro. |
**Participants** | [**[]VoiceV3Participant**](VoiceV3Participant.md) | Array of participants in the conversation |[optional] 
**Duration** | Pointer to **int** | Audio duration in seconds |
**ResolvedConfiguration** | [**VoiceV3ResolvedConfiguration**](VoiceV3ResolvedConfiguration.md) |  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | When this transcript was created |
**UpdatedAt** | [**time.Time**](time.Time.md) | When this transcript was last updated |
**Url** | **string** | The URL of this resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


