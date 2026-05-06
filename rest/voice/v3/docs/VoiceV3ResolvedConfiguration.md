# VoiceV3ResolvedConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscriptionEngine** | **string** | The engine used for transcription (Deepgram, Google, or auto) |[optional] 
**SpeechModel** | **string** | The speech model used for transcription (e.g., nova-2, nova-3, chirp_2) |[optional] 
**Language** | **string** | The language code for transcription |[optional] 
**TranscriptionStatusCallback** | [**VoiceV3TranscriptionStatusCallback**](VoiceV3TranscriptionStatusCallback.md) |  |[optional] 
**ConversationConfigurationId** | Pointer to **string** | Maestro conversation configuration ID |
**ParticipantDefaults** | [**[]VoiceV3TranscriptionResolvedConfigurationParticipantDefaults**](VoiceV3TranscriptionResolvedConfigurationParticipantDefaults.md) | Default participant configurations for the transcription |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


